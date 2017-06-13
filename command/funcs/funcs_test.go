package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"log"
	"os"
	"testing"
)

type dummyFlagContext struct{}

func (d *dummyFlagContext) IsSet(name string) bool {
	return true
}

var dummyContext command.Context

func TestMain(m *testing.M) {
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
		os.Exit(0) // exit normal
	}
	zone := os.Getenv("SAKURACLOUD_ZONE")
	if zone == "" {
		zone = "tk1v"
	}

	command.GlobalOption.AccessToken = accessToken
	command.GlobalOption.AccessTokenSecret = accessTokenSecret
	command.GlobalOption.Zone = zone

	dummyContext = command.NewContext(&dummyFlagContext{}, []string{}, nil)
	dummyContext.GetAPIClient().UserAgent = fmt.Sprintf("usacloud-unit-test")

	ret := m.Run()
	os.Exit(ret)
}
