package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollectionMembershipTypes(query string) ([]*pb.CollectionMembershipType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collection_membership_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionMembershipTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionmembershiptypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionmembershiptypes, nil
}

func (g *igdb) GetCollectionMembershipTypeByID(id uint64) (*pb.CollectionMembershipType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionMembershipTypes, err := g.GetCollectionMembershipTypes(query)
	if err != nil {
		return nil, err
	}
	return collectionMembershipTypes[0], nil
}

func (g *igdb) GetCollectionMembershipTypesByIDs(ids []uint64) ([]*pb.CollectionMembershipType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMembershipTypes(idStr)
}

func (g *igdb) GetCollectionMembershipTypesByAllowedCollectionTypeID(id uint64) ([]*pb.CollectionMembershipType, error) {
	query := fmt.Sprintf(`where allowed_collection_type = %d; fields *;`, id)
	return g.GetCollectionMembershipTypes(query)
}

func (g *igdb) GetCollectionMembershipTypesByAllowedCollectionTypeIDs(ids []uint64) ([]*pb.CollectionMembershipType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where allowed_collection_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMembershipTypes(idStr)
}

func (g *igdb) GetCollectionMembershipTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	collectionMembershipTypes, err := g.GetCollectionMembershipTypes(query)
	if err != nil {
		return 0, err
	}
	return int(collectionMembershipTypes[0].Id), nil
}
