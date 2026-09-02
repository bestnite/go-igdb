package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AgeRatingContentDescriptionsV2 struct {
	BaseEndpoint[pb.AgeRatingContentDescriptionV2]
}

func NewAgeRatingContentDescriptionsV2(request RequestFunc) *AgeRatingContentDescriptionsV2 {
	a := &AgeRatingContentDescriptionsV2{
		BaseEndpoint[pb.AgeRatingContentDescriptionV2]{
			endpointName: EPAgeRatingContentDescriptionsV2,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AgeRatingContentDescriptionV2Result) []*pb.AgeRatingContentDescriptionV2 {
		return r.Ageratingcontentdescriptionsv2
	})
	return a
}
