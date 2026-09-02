package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Events struct {
	BaseEndpoint[pb.Event]
}

func NewEvents(request RequestFunc) *Events {
	a := &Events{
		BaseEndpoint[pb.Event]{
			endpointName: EPEvents,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.EventResult) []*pb.Event { return r.Events })
	return a
}
