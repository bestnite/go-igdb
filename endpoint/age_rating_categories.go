package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AgeRatingCategories struct {
	BaseEndpoint[pb.AgeRatingCategory]
}

func NewAgeRatingCategories(request RequestFunc) *AgeRatingCategories {
	a := &AgeRatingCategories{
		BaseEndpoint[pb.AgeRatingCategory]{
			endpointName: EPAgeRatingCategories,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AgeRatingCategoryResult) []*pb.AgeRatingCategory { return r.Ageratingcategories })
	return a
}
