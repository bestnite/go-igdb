package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetAgeRatingContentDescriptionsV2(query string) ([]*pb.AgeRatingContentDescriptionV2, error) {
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

func (g *igdb) GetAgeRatingContentDescriptionV2ByID(id uint64) (*pb.AgeRatingContentDescriptionV2, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	ageRatingContentDescriptions, err := g.GetAgeRatingContentDescriptionsV2(query)
	if err != nil {
		return nil, err
	}
	return ageRatingContentDescriptions[0], nil
}

func (g *igdb) GetAgeRatingContentDescriptionsV2ByIDs(ids []uint64) ([]*pb.AgeRatingContentDescriptionV2, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAgeRatingContentDescriptionsV2(idStr)
}
