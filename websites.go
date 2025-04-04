package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

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
