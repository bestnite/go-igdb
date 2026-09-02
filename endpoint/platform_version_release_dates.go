package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformVersionReleaseDates struct {
	BaseEndpoint[pb.PlatformVersionReleaseDate]
}

func NewPlatformVersionReleaseDates(request RequestFunc) *PlatformVersionReleaseDates {
	a := &PlatformVersionReleaseDates{
		BaseEndpoint[pb.PlatformVersionReleaseDate]{
			endpointName: EPPlatformVersionReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformVersionReleaseDateResult) []*pb.PlatformVersionReleaseDate {
		return r.Platformversionreleasedates
	})
	return a
}
