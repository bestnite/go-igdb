package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCovers(query string) ([]*pb.Cover, error) {
	resp, err := g.Request("https://api.igdb.com/v4/covers.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CoverResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Covers) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Covers, nil
}

func (g *igdb) GetCoverByID(id uint64) (*pb.Cover, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	covers, err := g.GetCovers(query)
	if err != nil {
		return nil, err
	}
	return covers[0], nil
}

func (g *igdb) GetCoversByIDs(ids []uint64) ([]*pb.Cover, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCovers(idStr)
}
