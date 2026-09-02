package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CollectionMemberships struct {
	BaseEndpoint[pb.CollectionMembership]
}

func NewCollectionMemberships(request RequestFunc) *CollectionMemberships {
	a := &CollectionMemberships{
		BaseEndpoint[pb.CollectionMembership]{
			endpointName: EPCollectionMemberships,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionMembershipResult) []*pb.CollectionMembership { return r.Collectionmemberships })
	return a
}
