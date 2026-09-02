package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AgeRatings struct {
	BaseEndpoint[pb.AgeRating]
}

func NewAgeRatings(request RequestFunc) *AgeRatings {
	a := &AgeRatings{
		BaseEndpoint[pb.AgeRating]{
			endpointName: EPAgeRatings,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AgeRatingResult) []*pb.AgeRating { return r.Ageratings })
	return a
}
