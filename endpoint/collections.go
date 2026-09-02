package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Collections struct {
	BaseEndpoint[pb.Collection]
}

func NewCollections(request RequestFunc) *Collections {
	a := &Collections{
		BaseEndpoint[pb.Collection]{
			endpointName: EPCollections,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CollectionResult) []*pb.Collection { return r.Collections })
	return a
}
