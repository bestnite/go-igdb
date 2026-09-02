package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type EventLogos struct {
	BaseEndpoint[pb.EventLogo]
}

func NewEventLogos(request RequestFunc) *EventLogos {
	a := &EventLogos{
		BaseEndpoint[pb.EventLogo]{
			endpointName: EPEventLogos,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.EventLogoResult) []*pb.EventLogo { return r.Eventlogos })
	return a
}
