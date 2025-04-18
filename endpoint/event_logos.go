package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type EventLogos struct {
	BaseEndpoint[pb.EventLogo]
}

func NewEventLogos(request RequestFunc) *EventLogos {
	a := &EventLogos{
		BaseEndpoint[pb.EventLogo]{
			endpointName: EPEventLogos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *EventLogos) Query(query string) ([]*pb.EventLogo, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Eventlogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Eventlogos, nil
}
