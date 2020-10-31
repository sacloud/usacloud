package disk

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/usacloud/pkg/cmd/base"
	"github.com/stretchr/testify/require"
)

func TestList_ConvertToServiceRequest(t *testing.T) {
	in := &ListParameter{
		ExecContext:     nil,
		OutputParameter: nil,
		Zone:            "is1a",
		Names:           []string{"name1", "name2"},
		Tags:            []string{"tag1", "tag2"},
		FindParameter: &base.FindParameter{
			Count: 1,
			From:  2,
		},
	}

	out := &disk.FindRequest{}
	if err := mapconv.ConvertTo(in, out); err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, &disk.FindRequest{
		Zone:  "is1a",
		Names: []string{"name1", "name2"},
		Tags:  []string{"tag1", "tag2"},
		Count: 1,
		From:  2,
	}, out)
}
