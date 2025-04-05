package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Platforms struct{ BaseEndpoint }

func (a *Platforms) Query(query string) ([]*pb.Platform, error) {
	resp, err := a.request("https://api.igdb.com/v4/platforms.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platforms) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}
	return data.Platforms, nil
}
