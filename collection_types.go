package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollectionTypes(query string) ([]*pb.CollectionType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collection_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectiontypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectiontypes, nil
}

func (g *igdb) GetCollectionTypeByID(id uint64) (*pb.CollectionType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionTypes, err := g.GetCollectionTypes(query)
	if err != nil {
		return nil, err
	}
	return collectionTypes[0], nil
}

func (g *igdb) GetCollectionTypesByIDs(ids []uint64) ([]*pb.CollectionType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionTypes(idStr)
}
