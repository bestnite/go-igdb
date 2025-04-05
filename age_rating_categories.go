package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

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

func (g *igdb) GetAgeRatingCategoriesByOrganizationID(id uint64) ([]*pb.AgeRatingCategory, error) {
	query := fmt.Sprintf(`where organization = %d; fields *;`, id)
	return g.GetAgeRatingCategories(query)
}

func (g *igdb) GetAgeRatingCategoriesByOrganizationIDs(ids []uint64) ([]*pb.AgeRatingCategory, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where organization = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingCategories(idStr)
}

func (g *igdb) GetAgeRatingCategoriesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	ageRatingCategories, err := g.GetAgeRatingCategories(query)
	if err != nil {
		return 0, err
	}
	return int(ageRatingCategories[0].Id), nil
}
