package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CharacterGenders struct {
	BaseEndpoint[pb.CharacterGender]
}

func NewCharacterGenders(request RequestFunc) *CharacterGenders {
	a := &CharacterGenders{
		BaseEndpoint[pb.CharacterGender]{
			endpointName: EPCharacterGenders,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CharacterGenderResult) []*pb.CharacterGender { return r.Charactergenders })
	return a
}
