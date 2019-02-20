package funcs

import "github.com/sacloud/libsacloud/sacloud"

type dnsRecordValueType struct {
	*sacloud.DNSRecordSet
	Index int
}

type dnsRecordsType []*dnsRecordValueType
