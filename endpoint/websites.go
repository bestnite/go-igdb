package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Websites struct {
	BaseEndpoint[pb.Website]
}

func NewWebsites(request func(URL string, dataBody any) (*resty.Response, error)) *Websites {
	a := &Websites{
		BaseEndpoint[pb.Website]{
			endpointName: EPWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Websites) Query(query string) ([]*pb.Website, error) {
	resp, err := a.request("https://api.igdb.com/v4/websites.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.WebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Websites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Websites, nil
}
