package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

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

func (g *igdb) GetCollectionsByCollectionTypeID(id uint64) ([]*pb.Collection, error) {
	query := fmt.Sprintf(`where collection_type = %d; fields *;`, id)
	return g.GetCollections(query)
}

func (g *igdb) GetCollectionsByCollectionTypeIDs(ids []uint64) ([]*pb.Collection, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where collection_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollections(idStr)
}

func (g *igdb) GetCollectionsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	collections, err := g.GetCollections(query)
	if err != nil {
		return 0, err
	}
	return int(collections[0].Id), nil
}
