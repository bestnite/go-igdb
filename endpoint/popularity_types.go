package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PopularityTypes struct {
	BaseEndpoint[pb.PopularityType]
}

func NewPopularityTypes(request func(URL string, dataBody any) (*resty.Response, error)) *PopularityTypes {
	a := &PopularityTypes{
		BaseEndpoint[pb.PopularityType]{
			endpointName: EPPopularityTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PopularityTypes) Query(query string) ([]*pb.PopularityType, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PopularityTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Popularitytypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Popularitytypes, nil
}
