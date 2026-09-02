package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CharacterMugShots struct {
	BaseEndpoint[pb.CharacterMugShot]
}

func NewCharacterMugShots(request RequestFunc) *CharacterMugShots {
	a := &CharacterMugShots{
		BaseEndpoint[pb.CharacterMugShot]{
			endpointName: EPCharacterMugShots,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CharacterMugShotResult) []*pb.CharacterMugShot { return r.Charactermugshots })
	return a
}
