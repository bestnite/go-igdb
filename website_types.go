package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetWebsiteTypes(query string) ([]*pb.WebsiteType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/website_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.WebsiteTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Websitetypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Websitetypes, nil
}

func (g *igdb) GetWebsiteTypeByID(id uint64) (*pb.WebsiteType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	websiteTypes, err := g.GetWebsiteTypes(query)
	if err != nil {
		return nil, err
	}
	return websiteTypes[0], nil
}

func (g *igdb) GetWebsiteTypesByIDs(ids []uint64) ([]*pb.WebsiteType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetWebsiteTypes(idStr)
}

func (g *igdb) GetWebsiteTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	websiteTypes, err := g.GetWebsiteTypes(query)
	if err != nil {
		return 0, err
	}
	return int(websiteTypes[0].Id), nil
}
