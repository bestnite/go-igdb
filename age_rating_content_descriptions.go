package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetAgeRatingContentDescriptions(query string) ([]*pb.AgeRatingContentDescription, error) {
	resp, err := g.Request("https://api.igdb.com/v4/age_rating_content_descriptions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingContentDescriptionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingcontentdescriptions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingcontentdescriptions, nil
}

func (g *Client) GetAgeRatingContentDescriptionByID(id uint64) (*pb.AgeRatingContentDescription, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatingContentDescriptions, err := g.GetAgeRatingContentDescriptions(query)
	if err != nil {
		return nil, err
	}
	return ageRatingContentDescriptions[0], nil
}

func (g *Client) GetAgeRatingContentDescriptionsByIDs(ids []uint64) ([]*pb.AgeRatingContentDescription, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingContentDescriptions(idStr)
}

func (g *Client) GetAgeRatingContentDescriptionsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	ageRatingContentDescriptions, err := g.GetAgeRatingContentDescriptions(query)
	if err != nil {
		return 0, err
	}
	return int(ageRatingContentDescriptions[0].Id), nil
}
