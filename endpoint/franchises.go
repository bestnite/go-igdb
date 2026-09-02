package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Franchises struct {
	BaseEndpoint[pb.Franchise]
}

func NewFranchises(request RequestFunc) *Franchises {
	a := &Franchises{
		BaseEndpoint[pb.Franchise]{
			endpointName: EPFranchises,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.FranchiseResult) []*pb.Franchise { return r.Franchises })
	return a
}
