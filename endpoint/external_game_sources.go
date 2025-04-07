package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type ExternalGameSources struct {
	BaseEndpoint[pb.ExternalGameSource]
}

func NewExternalGameSources(request func(URL string, dataBody any) (*resty.Response, error)) *ExternalGameSources {
	a := &ExternalGameSources{
		BaseEndpoint[pb.ExternalGameSource]{
			endpointName: EPExternalGameSources,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ExternalGameSources) Query(query string) ([]*pb.ExternalGameSource, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameSourceResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Externalgamesources) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Externalgamesources, nil
}
