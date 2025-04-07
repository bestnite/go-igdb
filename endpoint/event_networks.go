package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type EventNetworks struct {
	BaseEndpoint[pb.EventNetwork]
}

func NewEventNetworks(request func(URL string, dataBody any) (*resty.Response, error)) *EventNetworks {
	a := &EventNetworks{
		BaseEndpoint[pb.EventNetwork]{
			endpointName: EPEventNetworks,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *EventNetworks) Query(query string) ([]*pb.EventNetwork, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventNetworkResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Eventnetworks) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Eventnetworks, nil
}
