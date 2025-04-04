package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

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
