package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameVideos struct {
	BaseEndpoint[pb.GameVideo]
}

func NewGameVideos(request RequestFunc) *GameVideos {
	a := &GameVideos{
		BaseEndpoint: BaseEndpoint[pb.GameVideo]{
			endpointName: EPGameVideos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameVideos) Query(ctx context.Context, query string) ([]*pb.GameVideo, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVideoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Gamevideos, nil
}
