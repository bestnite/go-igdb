package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetAgeRatingContentDescriptionsV2(query string) ([]*pb.AgeRatingContentDescriptionV2, error) {
	resp, err := g.Request("https://api.igdb.com/v4/age_rating_content_descriptions_v2.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingContentDescriptionV2Result{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingcontentdescriptionsv2) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingcontentdescriptionsv2, nil
}

func (g *Client) GetAgeRatingContentDescriptionV2ByID(id uint64) (*pb.AgeRatingContentDescriptionV2, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatingContentDescriptions, err := g.GetAgeRatingContentDescriptionsV2(query)
	if err != nil {
		return nil, err
	}
	return ageRatingContentDescriptions[0], nil
}

func (g *Client) GetAgeRatingContentDescriptionsV2ByIDs(ids []uint64) ([]*pb.AgeRatingContentDescriptionV2, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingContentDescriptionsV2(idStr)
}

func (g *Client) GetAgeRatingContentDescriptionsV2ByOrganizationID(id uint64) ([]*pb.AgeRatingContentDescriptionV2, error) {
	query := fmt.Sprintf(`where organization = %d; fields *;`, id)
	return g.GetAgeRatingContentDescriptionsV2(query)
}

func (g *Client) GetAgeRatingContentDescriptionsV2ByOrganizationIDs(ids []uint64) ([]*pb.AgeRatingContentDescriptionV2, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where organization = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingContentDescriptionsV2(idStr)
}

func (g *Client) GetAgeRatingContentDescriptionsV2Length() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	ageRatingContentDescriptions, err := g.GetAgeRatingContentDescriptionsV2(query)
	if err != nil {
		return 0, err
	}
	return int(ageRatingContentDescriptions[0].Id), nil
}
