package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformTypes struct {
	BaseEndpoint[pb.PlatformType]
}

func NewPlatformTypes(request RequestFunc) *PlatformTypes {
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
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
