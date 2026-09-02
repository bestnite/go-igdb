package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CollectionTypes struct {
	BaseEndpoint[pb.CollectionType]
}

func NewCollectionTypes(request RequestFunc) *CollectionTypes {
	a := &CollectionTypes{
		BaseEndpoint[pb.CollectionType]{
			endpointName: EPCollectionTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionTypeResult) []*pb.CollectionType { return r.Collectiontypes })
	return a
}
