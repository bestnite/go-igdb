package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollectionRelationTypes(query string) ([]*pb.CollectionRelationType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collection_relation_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionRelationTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionrelationtypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionrelationtypes, nil
}

func (g *igdb) GetCollectionRelationTypeByID(id uint64) (*pb.CollectionRelationType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionRelationTypes, err := g.GetCollectionRelationTypes(query)
	if err != nil {
		return nil, err
	}
	return collectionRelationTypes[0], nil
}

func (g *igdb) GetCollectionRelationTypesByIDs(ids []uint64) ([]*pb.CollectionRelationType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelationTypes(idStr)
}
