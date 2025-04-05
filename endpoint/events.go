package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Events struct {
	BaseEndpoint[pb.Event]
}

func NewEvents(request func(URL string, dataBody any) (*resty.Response, error)) *Events {
	a := &Events{
		BaseEndpoint[pb.Event]{
			endpointName: EPEvents,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Events) Query(query string) ([]*pb.Event, error) {
	resp, err := a.request("https://api.igdb.com/v4/events.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Events) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Events, nil
}
