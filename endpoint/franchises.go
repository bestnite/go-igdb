package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Franchises struct {
	BaseEndpoint[pb.Franchise]
}

func NewFranchises(request func(URL string, dataBody any) (*resty.Response, error)) *Franchises {
	a := &Franchises{
		BaseEndpoint[pb.Franchise]{
			endpointName: EPFranchises,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Franchises) Query(query string) ([]*pb.Franchise, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.FranchiseResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Franchises) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Franchises, nil
}
