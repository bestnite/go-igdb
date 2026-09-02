package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Reports struct {
	BaseEndpoint[pb.Report]
}

func NewReports(request RequestFunc) *Reports {
	a := &Reports{
		BaseEndpoint[pb.Report]{
			endpointName: EPReports,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ReportResult) []*pb.Report { return r.Reports })
	return a
}
