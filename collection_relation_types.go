package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCollectionRelationTypes(query string) ([]*pb.CollectionRelationType, error) {
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

func (g *Client) GetCollectionRelationTypeByID(id uint64) (*pb.CollectionRelationType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionRelationTypes, err := g.GetCollectionRelationTypes(query)
	if err != nil {
		return nil, err
	}
	return collectionRelationTypes[0], nil
}

func (g *Client) GetCollectionRelationTypesByIDs(ids []uint64) ([]*pb.CollectionRelationType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelationTypes(idStr)
}

func (g *Client) GetCollectionRelationTypesByAllowedChildTypeID(id uint64) ([]*pb.CollectionRelationType, error) {
	query := fmt.Sprintf(`where allowed_child_type = %d; fields *;`, id)
	return g.GetCollectionRelationTypes(query)
}

func (g *Client) GetCollectionRelationTypesByAllowedChildTypeIDs(ids []uint64) ([]*pb.CollectionRelationType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where allowed_child_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelationTypes(idStr)
}

func (g *Client) GetCollectionRelationTypesByAllowedParentTypeID(id uint64) ([]*pb.CollectionRelationType, error) {
	query := fmt.Sprintf(`where allowed_parent_type = %d; fields *;`, id)
	return g.GetCollectionRelationTypes(query)
}

func (g *Client) GetCollectionRelationTypesByAllowedParentTypeIDs(ids []uint64) ([]*pb.CollectionRelationType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where allowed_parent_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionRelationTypes(idStr)
}

func (g *Client) GetCollectionRelationTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	collectionRelationTypes, err := g.GetCollectionRelationTypes(query)
	if err != nil {
		return 0, err
	}
	return int(collectionRelationTypes[0].Id), nil
}
