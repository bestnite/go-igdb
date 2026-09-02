package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CollectionMembershipTypes struct {
	BaseEndpoint[pb.CollectionMembershipType]
}

func NewCollectionMembershipTypes(request RequestFunc) *CollectionMembershipTypes {
	a := &CollectionMembershipTypes{
		BaseEndpoint[pb.CollectionMembershipType]{
			endpointName: EPCollectionMembershipTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionMembershipTypeResult) []*pb.CollectionMembershipType {
		return r.Collectionmembershiptypes
	})
	return a
}
