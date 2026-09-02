package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type NetworkTypes struct {
	BaseEndpoint[pb.NetworkType]
}

func NewNetworkTypes(request RequestFunc) *NetworkTypes {
	a := &NetworkTypes{
		BaseEndpoint[pb.NetworkType]{
			endpointName: EPNetworkTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.NetworkTypeResult) []*pb.NetworkType { return r.Networktypes })
	return a
}
