package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollectionRelations(query string) ([]*pb.CollectionRelation, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collection_relations.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionRelationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionrelations) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionrelations, nil
}

func (g *igdb) GetCollectionRelationByID(id uint64) (*pb.CollectionRelation, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionRelations, err := g.GetCollectionRelations(query)
	if err != nil {
		return nil, err
	}
	return collectionRelations[0], nil
}

func (g *igdb) GetCollectionRelationsByIDs(ids []uint64) ([]*pb.CollectionRelation, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelations(idStr)
}
