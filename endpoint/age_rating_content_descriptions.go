package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AgeRatingContentDescriptions struct {
	BaseEndpoint[pb.AgeRatingContentDescription]
}

func NewAgeRatingContentDescriptions(request RequestFunc) *AgeRatingContentDescriptions {
	a := &AgeRatingContentDescriptions{
		BaseEndpoint[pb.AgeRatingContentDescription]{
			endpointName: EPAgeRatingContentDescriptions,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AgeRatingContentDescriptionResult) []*pb.AgeRatingContentDescription {
		return r.Ageratingcontentdescriptions
	})
	return a
}
