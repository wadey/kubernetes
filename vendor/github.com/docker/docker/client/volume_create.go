package client

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	volumetypes "github.com/docker/docker/api/types/volume"
)

// VolumeCreate creates a volume in the docker host.
func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumeCreateBody) (types.Volume, error) {
	var volume types.Volume
	resp, err := cli.post(ctx, "/volumes/create", nil, options, nil)
	if err != nil {
		return volume, err
	}
	err = json.NewDecoder(resp.body).Decode(&volume)
	ensureReaderClosed(resp)
	return volume, err
}
