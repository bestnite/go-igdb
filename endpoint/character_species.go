package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
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
	a.queryFunc = a.queryPB(func(r *pb.CharacterSpecieResult) []*pb.CharacterSpecie { return r.Characterspecies })
	return a
}
