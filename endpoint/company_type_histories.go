package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanyTypeHistories struct {
	BaseEndpoint[pb.CompanyTypeHistory]
}

func NewCompanyTypeHistories(request RequestFunc) *CompanyTypeHistories {
	a := &CompanyTypeHistories{
		BaseEndpoint[pb.CompanyTypeHistory]{
			endpointName: EPCompanyTypeHistories,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyTypeHistoryResult) []*pb.CompanyTypeHistory { return r.Companytypehistories })
	return a
}
