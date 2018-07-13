package funcs

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/helper/mutexkv"
)

const SummaryAPIThrottleSize = 5

var summaryBuildLock = mutexkv.NewMutexKV()

type resourceSummaryItem struct {
	key    string
	zone   string
	global bool
	count  int
}

type resourceCounterContext struct {
	client     *api.Client
	counterDef *resourceCounter
	results    chan *resourceSummaryItem
	errs       chan error
}

func SummaryShow(ctx command.Context, params *params.ShowSummaryParam) error {

	// prepare
	var summary []map[string]string
	results := make(chan *resourceSummaryItem)
	errs := make(chan error)

	counters := buildCounters(ctx, params.PaidResourcesOnly, results, errs)

	err := internal.ExecWithProgress(
		"Still calculating...",
		"Calculate resource count",
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {

			// do count & build result
			s, err := buildSummary(counters, params.PaidResourcesOnly, results, errs)
			if err != nil {
				errChan <- fmt.Errorf("Building summary is failed: %s", err)
				return
			}
			summary = s
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("SummaryShow is failed: %s", err)
	}

	// output
	list := []interface{}{}
	for i := range summary {

		list = append(list, &summary[i])
	}
	return ctx.GetOutput().Print(list...)
}

func buildCounters(ctx command.Context, paidOnly bool, results chan *resourceSummaryItem, errs chan error) []func() {
	client := ctx.GetAPIClient()
	counters := []func(){}

	for _, def := range resourceCounters {
		if paidOnly && !def.paid {
			continue
		}
	zoneLoop:
		for i := range define.AllowZones {

			c := client.Clone()
			c.Zone = define.AllowZones[i]

			ctx := &resourceCounterContext{
				client:     c,
				counterDef: def,
				results:    results,
				errs:       errs,
			}

			counters = append(counters, buildCounter(ctx))
			if def.global {
				break zoneLoop
			}
		}
	}
	return counters
}

func buildCounter(ctx *resourceCounterContext) func() {

	return func() {

		def := ctx.counterDef

		count, err := def.finder(ctx.client)
		if err != nil {
			ctx.errs <- err
		}

		ctx.results <- &resourceSummaryItem{
			key:    def.displayName,
			count:  count,
			global: def.global,
			zone:   ctx.client.Zone,
		}

	}

}

func buildSummary(counters []func(), paidOnly bool, results chan *resourceSummaryItem, errs chan error) ([]map[string]string, error) {
	//prepare summary
	summary := []map[string]string{}
	for _, def := range resourceCounters {
		if paidOnly && !def.paid {
			continue
		}
		v := map[string]string{
			"Name":  def.displayName,
			"Total": "0",
		}
		for _, zone := range define.AllowZones {
			v[zone] = "0"
		}
		summary = append(summary, v)
	}
	errors := []error{}

	wg := sync.WaitGroup{}
	wg.Add(len(counters))

	apiThrottle := make(chan bool, SummaryAPIThrottleSize)

	go func() {
		for {
			select {
			case res := <-results:
				buildSummaryItem(summary, res)
			case err := <-errs:
				errors = append(errors, err)
			}
			wg.Done()
			<-apiThrottle
		}
	}()

	for _, counter := range counters {
		apiThrottle <- true
		go counter()
	}

	wg.Wait()
	return summary, nil
}

func buildSummaryItem(summary []map[string]string, item *resourceSummaryItem) {

	var target map[string]string
	for i := range summary {
		if summary[i]["Name"] == item.key {
			target = summary[i]
		}
	}

	if target == nil {
		panic(fmt.Sprintf("summary[%q] is not found", item.key))
	}

	summaryBuildLock.Lock(item.key)
	defer summaryBuildLock.Unlock(item.key)

	// calc total
	calcSummaryItem(target, "Total", item.count)
	if item.global {
		for _, zone := range define.AllowZones {
			target[zone] = "-"
		}

		return
	}

	// calc per zone
	calcSummaryItem(target, item.zone, item.count)
}

func calcSummaryItem(target map[string]string, key string, value int) {
	v := target[key]
	n, _ := strconv.ParseInt(v, 10, 64)
	target[key] = fmt.Sprintf("%d", n+int64(value))
}
