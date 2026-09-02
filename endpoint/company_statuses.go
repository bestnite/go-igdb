package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanyStatuses struct {
	BaseEndpoint[pb.CompanyStatus]
}

func NewCompanyStatuses(request RequestFunc) *CompanyStatuses {
	a := &CompanyStatuses{
		BaseEndpoint[pb.CompanyStatus]{
			endpointName: EPCompanyStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyStatusResult) []*pb.CompanyStatus { return r.Companystatuses })
	return a
}
