package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetWebsites(query string) ([]*pb.Website, error) {
	resp, err := g.Request("https://api.igdb.com/v4/websites.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.WebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Websites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Websites, nil
}

func (g *igdb) GetWebsiteByID(id uint64) (*pb.Website, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	websites, err := g.GetWebsites(query)
	if err != nil {
		return nil, err
	}
	return websites[0], nil
}

func (g *igdb) GetWebsitesByIDs(ids []uint64) ([]*pb.Website, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetWebsites(idStr)
}

func (g *igdb) GetWebsitesByGameID(id uint64) ([]*pb.Website, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetWebsites(query)
}

func (g *igdb) GetWebsitesByGameIDs(ids []uint64) ([]*pb.Website, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetWebsites(idStr)
}

func (g *igdb) GetWebsitesByTypeID(id uint64) ([]*pb.Website, error) {
	query := fmt.Sprintf(`where type = %d; fields *;`, id)
	return g.GetWebsites(query)
}

func (g *igdb) GetWebsitesByTypeIDs(ids []uint64) ([]*pb.Website, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetWebsites(idStr)
}

func (g *igdb) GetWebsitesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	websites, err := g.GetWebsites(query)
	if err != nil {
		return 0, err
	}
	return int(websites[0].Id), nil
}
