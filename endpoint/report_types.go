package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ReportTypes struct {
	BaseEndpoint[pb.ReportType]
}

func NewReportTypes(request RequestFunc) *ReportTypes {
	a := &ReportTypes{
		BaseEndpoint[pb.ReportType]{
			endpointName: EPReportTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ReportTypeResult) []*pb.ReportType { return r.Reporttypes })
	return a
}
