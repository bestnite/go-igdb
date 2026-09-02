package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type InvolvedCompanies struct {
	BaseEndpoint[pb.InvolvedCompany]
}

func NewInvolvedCompanies(request RequestFunc) *InvolvedCompanies {
	a := &InvolvedCompanies{
		BaseEndpoint[pb.InvolvedCompany]{
			endpointName: EPInvolvedCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.InvolvedCompanyResult) []*pb.InvolvedCompany { return r.Involvedcompanies })
	return a
}
