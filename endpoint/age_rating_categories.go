package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AgeRatingCategories struct {
	BaseEndpoint
}

func (a *AgeRatingCategories) Query(query string) ([]*pb.AgeRatingCategory, error) {
	resp, err := a.request("https://api.igdb.com/v4/age_rating_categories.pb", query)
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
