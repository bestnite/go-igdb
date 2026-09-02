package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type EventNetworks struct {
	BaseEndpoint[pb.EventNetwork]
}

func NewEventNetworks(request RequestFunc) *EventNetworks {
	a := &EventNetworks{
		BaseEndpoint[pb.EventNetwork]{
			endpointName: EPEventNetworks,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.EventNetworkResult) []*pb.EventNetwork { return r.Eventnetworks })
	return a
}
