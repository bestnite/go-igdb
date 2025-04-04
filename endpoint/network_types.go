package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type NetworkTypes struct {
	BaseEndpoint[pb.NetworkType]
}

func NewNetworkTypes(request func(URL string, dataBody any) (*resty.Response, error)) *NetworkTypes {
	a := &NetworkTypes{
		BaseEndpoint[pb.NetworkType]{
			endpointName: EPNetworkTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *NetworkTypes) Query(query string) ([]*pb.NetworkType, error) {
	resp, err := a.request("https://api.igdb.com/v4/network_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.NetworkTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Networktypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Networktypes, nil
}
