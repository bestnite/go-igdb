package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetAgeRatingCategories(query string) ([]*pb.AgeRatingCategory, error) {
	resp, err := g.Request("https://api.igdb.com/v4/age_rating_categories.pb", query)
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

func (g *igdb) GetAgeRatingCategoryByID(id uint64) (*pb.AgeRatingCategory, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatingCategories, err := g.GetAgeRatingCategories(query)
	if err != nil {
		return nil, err
	}
	return ageRatingCategories[0], nil
}

func (g *igdb) GetAgeRatingCategoriesByIDs(ids []uint64) ([]*pb.AgeRatingCategory, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingCategories(idStr)
}
