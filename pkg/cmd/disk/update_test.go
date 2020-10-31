package disk

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/stretchr/testify/require"
)

func TestUpdate_ConvertToServiceRequest(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		in := &UpdateParameter{
			ExecContext:     nil,
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
			Description:     pointer.NewString("desc"),
			Tags:            pointer.NewStringSlice([]string{"tag1", "tag2"}),
			IconID:          pointer.NewID(types.ID(2)),
			Connection:      pointer.NewString(types.DiskConnections.VirtIO.String()),
		}

		out := &disk.UpdateRequest{}
		if err := mapconv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: pointer.NewString("desc"),
			Tags:        pointer.NewTags(types.Tags{"tag1", "tag2"}),
			IconID:      pointer.NewID(types.ID(2)),
			Connection:  &types.DiskConnections.VirtIO,
		}, out)
	})

	t.Run("nil", func(t *testing.T) {
		in := &UpdateParameter{
			ExecContext:     nil,
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
		}

		out := &disk.UpdateRequest{}
		if err := mapconv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: nil,
			Tags:        nil,
			IconID:      nil,
			Connection:  nil,
		}, out)
	})
	t.Run("empty", func(t *testing.T) {
		in := &UpdateParameter{
			ExecContext:     nil,
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
			Description:     pointer.NewString(""),
			Tags:            pointer.NewStringSlice([]string{}),
			IconID:          pointer.NewID(types.ID(0)),
		}

		out := &disk.UpdateRequest{}
		if err := mapconv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: pointer.NewString(""),
			Tags:        pointer.NewTags(types.Tags{}),
			IconID:      pointer.NewID(0),
			Connection:  nil,
		}, out)
	})
}
