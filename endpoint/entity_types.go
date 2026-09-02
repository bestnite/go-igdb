package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type EntityTypes struct {
	BaseEndpoint[pb.EntityType]
}

func NewEntityTypes(request RequestFunc) *EntityTypes {
	a := &EntityTypes{
		BaseEndpoint[pb.EntityType]{
			endpointName: EPEntityTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.EntityTypeResult) []*pb.EntityType { return r.Entitytypes })
	return a
}
