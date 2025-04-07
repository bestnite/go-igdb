package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformWebsites struct {
	BaseEndpoint[pb.PlatformWebsite]
}

func NewPlatformWebsites(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformWebsites {
	a := &PlatformWebsites{
		BaseEndpoint[pb.PlatformWebsite]{
			endpointName: EPPlatformWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformWebsites) Query(query string) ([]*pb.PlatformWebsite, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformwebsites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformwebsites, nil
}
