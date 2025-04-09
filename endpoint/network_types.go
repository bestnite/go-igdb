package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type NetworkTypes struct {
	BaseEndpoint[pb.NetworkType]
}

func NewNetworkTypes(request RequestFunc) *NetworkTypes {
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
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
