package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ReleaseDateStatuses struct {
	BaseEndpoint[pb.ReleaseDateStatus]
}

func NewReleaseDateStatuses(request RequestFunc) *ReleaseDateStatuses {
	a := &ReleaseDateStatuses{
		BaseEndpoint[pb.ReleaseDateStatus]{
			endpointName: EPReleaseDateStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ReleaseDateStatusResult) []*pb.ReleaseDateStatus { return r.Releasedatestatuses })
	return a
}
