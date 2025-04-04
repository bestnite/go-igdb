package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCollectionMemberships(query string) ([]*pb.CollectionMembership, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collection_memberships.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionMembershipResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionmemberships) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionmemberships, nil
}

func (g *igdb) GetCollectionMembershipByID(id uint64) (*pb.CollectionMembership, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionMemberships, err := g.GetCollectionMemberships(query)
	if err != nil {
		return nil, err
	}
	return collectionMemberships[0], nil
}

func (g *igdb) GetCollectionMembershipsByIDs(ids []uint64) ([]*pb.CollectionMembership, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMemberships(idStr)
}
