package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type WebsiteTypes struct {
	BaseEndpoint[pb.WebsiteType]
}

func NewWebsiteTypes(request func(URL string, dataBody any) (*resty.Response, error)) *WebsiteTypes {
	a := &WebsiteTypes{
		BaseEndpoint[pb.WebsiteType]{
			endpointName: EPWebsiteTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *WebsiteTypes) Query(query string) ([]*pb.WebsiteType, error) {
	resp, err := a.request("https://api.igdb.com/v4/website_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.WebsiteTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Websitetypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Websitetypes, nil
}
