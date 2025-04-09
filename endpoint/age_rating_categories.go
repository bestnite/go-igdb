package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

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

func (a *AgeRatingCategories) Query(query string) ([]*pb.AgeRatingCategory, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingCategoryResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingcategories) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingcategories, nil
}
