package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AgeRatingCategories struct {
	BaseEndpoint[pb.AgeRatingCategory]
}

func NewAgeRatingCategories(request RequestFunc) *AgeRatingCategories {
	a := &AgeRatingCategories{
		BaseEndpoint: BaseEndpoint[pb.AgeRatingCategory]{
			endpointName: EPAgeRatingCategories,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AgeRatingCategories) Query(ctx context.Context, query string) ([]*pb.AgeRatingCategory, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingCategoryResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Ageratingcategories, nil
}
