package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type DateFormats struct {
	BaseEndpoint[pb.DateFormat]
}

func NewDateFormats(request RequestFunc) *DateFormats {
	a := &DateFormats{
		BaseEndpoint[pb.DateFormat]{
			endpointName: EPDateFormats,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.DateFormatResult) []*pb.DateFormat { return r.Dateformats })
	return a
}
