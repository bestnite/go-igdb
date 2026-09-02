package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AgeRatingOrganizations struct {
	BaseEndpoint[pb.AgeRatingOrganization]
}

func NewAgeRatingOrganizations(request RequestFunc) *AgeRatingOrganizations {
	a := &AgeRatingOrganizations{
		BaseEndpoint[pb.AgeRatingOrganization]{
			endpointName: EPAgeRatingOrganizations,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AgeRatingOrganizationResult) []*pb.AgeRatingOrganization { return r.Ageratingorganizations })
	return a
}
