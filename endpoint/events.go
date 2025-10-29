package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Events struct {
	BaseEndpoint[pb.Event]
}

func NewEvents(request RequestFunc) *Events {
	a := &Events{
		BaseEndpoint[pb.Event]{
			endpointName: EPEvents,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Events) Query(ctx context.Context, query string) ([]*pb.Event, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Events, nil
}
