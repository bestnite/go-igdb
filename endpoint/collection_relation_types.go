package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CollectionRelationTypes struct {
	BaseEndpoint[pb.CollectionRelationType]
}

func NewCollectionRelationTypes(request RequestFunc) *CollectionRelationTypes {
	a := &CollectionRelationTypes{
		BaseEndpoint[pb.CollectionRelationType]{
			endpointName: EPCollectionRelationTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionRelationTypeResult) []*pb.CollectionRelationType {
		return r.Collectionrelationtypes
	})
	return a
}
