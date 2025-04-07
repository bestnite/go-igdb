package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Platforms struct {
	BaseEndpoint[pb.Platform]
}

func NewPlatforms(request func(URL string, dataBody any) (*resty.Response, error)) *Platforms {
	a := &Platforms{
		BaseEndpoint[pb.Platform]{
			endpointName: EPPlatforms,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Platforms) Query(query string) ([]*pb.Platform, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
