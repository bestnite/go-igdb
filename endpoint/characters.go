package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
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
	a.queryFunc = a.queryPB(func(r *pb.CharacterResult) []*pb.Character { return r.Characters })
	return a
}
