package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollections(query string) ([]*pb.Collection, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collections.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collections) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collections, nil
}

func (g *igdb) GetCollectionByID(id uint64) (*pb.Collection, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collections, err := g.GetCollections(query)
	if err != nil {
		return nil, err
	}
	return collections[0], nil
}

func (g *igdb) GetCollectionsByIDs(ids []uint64) ([]*pb.Collection, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollections(idStr)
}
