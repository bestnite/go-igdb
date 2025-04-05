package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCollectionRelations(query string) ([]*pb.CollectionRelation, error) {
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

func (g *Client) GetCollectionRelationByID(id uint64) (*pb.CollectionRelation, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionRelations, err := g.GetCollectionRelations(query)
	if err != nil {
		return nil, err
	}
	return collectionRelations[0], nil
}

func (g *Client) GetCollectionRelationsByIDs(ids []uint64) ([]*pb.CollectionRelation, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelations(idStr)
}

func (g *Client) GetCollectionRelationsByChildCollectionID(id uint64) ([]*pb.CollectionRelation, error) {
	query := fmt.Sprintf(`where child_collection = %d; fields *;`, id)
	return g.GetCollectionRelations(query)
}

func (g *Client) GetCollectionRelationsByChildCollectionIDs(ids []uint64) ([]*pb.CollectionRelation, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where child_collection = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelations(idStr)
}

func (g *Client) GetCollectionRelationsByParentCollectionID(id uint64) ([]*pb.CollectionRelation, error) {
	query := fmt.Sprintf(`where parent_collection = %d; fields *;`, id)
	return g.GetCollectionRelations(query)
}

func (g *Client) GetCollectionRelationsByParentCollectionIDs(ids []uint64) ([]*pb.CollectionRelation, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where parent_collection = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelations(idStr)
}

func (g *Client) GetCollectionRelationsByCollectionRelationTypeID(id uint64) ([]*pb.CollectionRelation, error) {
	query := fmt.Sprintf(`where type = %d; fields *;`, id)
	return g.GetCollectionRelations(query)
}

func (g *Client) GetCollectionRelationsByCollectionRelationTypeIDs(ids []uint64) ([]*pb.CollectionRelation, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelations(idStr)
}

func (g *Client) GetCollectionRelationsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	collectionRelations, err := g.GetCollectionRelations(query)
	if err != nil {
		return 0, err
	}
	return int(collectionRelations[0].Id), nil
}
