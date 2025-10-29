package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CharacterSpecies struct {
	BaseEndpoint[pb.CharacterSpecie]
}

func NewCharacterSpecies(request RequestFunc) *CharacterSpecies {
	a := &CharacterSpecies{
		BaseEndpoint[pb.CharacterSpecie]{
			endpointName: EPCharacterSpecies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CharacterSpecies) Query(ctx context.Context, query string) ([]*pb.CharacterSpecie, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterSpecieResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Characterspecies, nil
}
