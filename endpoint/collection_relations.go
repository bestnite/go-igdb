package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CollectionRelations struct {
	BaseEndpoint[pb.CollectionRelation]
}

func NewCollectionRelations(request RequestFunc) *CollectionRelations {
	a := &CollectionRelations{
		BaseEndpoint[pb.CollectionRelation]{
			endpointName: EPCollectionRelations,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionRelationResult) []*pb.CollectionRelation { return r.Collectionrelations })
	return a
}
