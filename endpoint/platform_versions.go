package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformVersions struct {
	BaseEndpoint[pb.PlatformVersion]
}

func NewPlatformVersions(request RequestFunc) *PlatformVersions {
	a := &PlatformVersions{
		BaseEndpoint[pb.PlatformVersion]{
			endpointName: EPPlatformVersions,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformVersions) Query(query string) ([]*pb.PlatformVersion, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversions, nil
}
