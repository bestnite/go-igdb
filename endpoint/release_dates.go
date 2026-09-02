package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ReleaseDates struct {
	BaseEndpoint[pb.ReleaseDate]
}

func NewReleaseDates(request RequestFunc) *ReleaseDates {
	a := &ReleaseDates{
		BaseEndpoint[pb.ReleaseDate]{
			endpointName: EPReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ReleaseDateResult) []*pb.ReleaseDate { return r.Releasedates })
	return a
}
