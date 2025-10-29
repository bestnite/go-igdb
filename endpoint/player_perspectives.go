package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlayerPerspectives struct {
	BaseEndpoint[pb.PlayerPerspective]
}

func NewPlayerPerspectives(request RequestFunc) *PlayerPerspectives {
	a := &PlayerPerspectives{
		BaseEndpoint: BaseEndpoint[pb.PlayerPerspective]{
			endpointName: EPPlayerPerspectives,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlayerPerspectives) Query(ctx context.Context, query string) ([]*pb.PlayerPerspective, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlayerPerspectiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Playerperspectives, nil
}
