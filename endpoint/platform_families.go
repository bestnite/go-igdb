package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformFamilies struct {
	BaseEndpoint[pb.PlatformFamily]
}

func NewPlatformFamilies(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformFamilies {
	a := &PlatformFamilies{
		BaseEndpoint[pb.PlatformFamily]{
			endpointName: EPPlatformFamilies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformFamilies) Query(query string) ([]*pb.PlatformFamily, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformFamilyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformfamilies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformfamilies, nil
}
