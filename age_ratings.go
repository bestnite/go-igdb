package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetAgeRatings(query string) ([]*pb.AgeRating, error) {
	resp, err := g.Request("https://api.igdb.com/v4/age_ratings.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratings) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratings, nil
}

func (g *igdb) GetAgeRatingByID(id uint64) (*pb.AgeRating, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatings, err := g.GetAgeRatings(query)
	if err != nil {
		return nil, err
	}
	return ageRatings[0], nil
}

func (g *igdb) GetAgeRatingsByIDs(ids []uint64) ([]*pb.AgeRating, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatings(idStr)
}

func (g *igdb) GetAgeRatingsByOrganizationID(id uint64) ([]*pb.AgeRating, error) {
	query := fmt.Sprintf(`where organization = %d; fields *;`, id)
	return g.GetAgeRatings(query)
}

func (g *igdb) GetAgeRatingsByOrganizationIDs(ids []uint64) ([]*pb.AgeRating, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where organization = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatings(idStr)
}

func (g *igdb) GetAgeRatingsByAgeRatingCategoryID(id uint64) ([]*pb.AgeRating, error) {
	query := fmt.Sprintf(`where rating_category = %d; fields *;`, id)
	return g.GetAgeRatings(query)
}

func (g *igdb) GetAgeRatingsByAgeRatingCategoryIDs(ids []uint64) ([]*pb.AgeRating, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where rating_category = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatings(idStr)
}
