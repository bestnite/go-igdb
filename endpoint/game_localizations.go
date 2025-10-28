package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameLocalizations struct {
	BaseEndpoint[pb.GameLocalization]
}

func NewGameLocalizations(request RequestFunc) *GameLocalizations {
	a := &GameLocalizations{
		BaseEndpoint: BaseEndpoint[pb.GameLocalization]{
			endpointName: EPGameLocalizations,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameLocalizations) Query(ctx context.Context, query string) ([]*pb.GameLocalization, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameLocalizationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Gamelocalizations, nil
}
