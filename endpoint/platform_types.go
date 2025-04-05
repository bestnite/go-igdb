package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformTypes struct {
	BaseEndpoint[pb.PlatformType]
}

func NewPlatformTypes(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformTypes {
	a := &PlatformTypes{
		BaseEndpoint[pb.PlatformType]{
			endpointName: EPPlatformTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformTypes) Query(query string) ([]*pb.PlatformType, error) {
	resp, err := a.request("https://api.igdb.com/v4/platform_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformtypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformtypes, nil
}
