package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Screenshots struct {
	BaseEndpoint[pb.Screenshot]
}

func NewScreenshots(request RequestFunc) *Screenshots {
	a := &Screenshots{
		BaseEndpoint[pb.Screenshot]{
			endpointName: EPScreenshots,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ScreenshotResult) []*pb.Screenshot { return r.Screenshots })
	return a
}
