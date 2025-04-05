package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetAgeRatingOrganizations(query string) ([]*pb.AgeRatingOrganization, error) {
	resp, err := g.Request("https://api.igdb.com/v4/age_rating_organizations.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingOrganizationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingorganizations) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingorganizations, nil
}

func (g *igdb) GetAgeRatingOrganizationByID(id uint64) (*pb.AgeRatingOrganization, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatingOrganizations, err := g.GetAgeRatingOrganizations(query)
	if err != nil {
		return nil, err
	}
	return ageRatingOrganizations[0], nil
}

func (g *igdb) GetAgeRatingOrganizationsByIDs(ids []uint64) ([]*pb.AgeRatingOrganization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingOrganizations(idStr)
}

func (g *igdb) GetAgeRatingOrganizationsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	ageRatingOrganizations, err := g.GetAgeRatingOrganizations(query)
	if err != nil {
		return 0, err
	}
	return int(ageRatingOrganizations[0].Id), nil
}
