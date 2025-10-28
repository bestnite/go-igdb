package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Characters struct {
	BaseEndpoint[pb.Character]
}

func NewCharacters(request RequestFunc) *Characters {
	a := &Characters{
		BaseEndpoint[pb.Character]{
			endpointName: EPCharacters,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Characters) Query(ctx context.Context, query string) ([]*pb.Character, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Characters, nil
}
