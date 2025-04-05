package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCollectionMemberships(query string) ([]*pb.CollectionMembership, error) {
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

func (g *Client) GetCollectionMembershipByID(id uint64) (*pb.CollectionMembership, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	collectionMemberships, err := g.GetCollectionMemberships(query)
	if err != nil {
		return nil, err
	}
	return collectionMemberships[0], nil
}

func (g *Client) GetCollectionMembershipsByIDs(ids []uint64) ([]*pb.CollectionMembership, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMemberships(idStr)
}

func (g *Client) GetCollectionMembershipsByGameID(id uint64) ([]*pb.CollectionMembership, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetCollectionMemberships(query)
}

func (g *Client) GetCollectionMembershipsByGameIDs(ids []uint64) ([]*pb.CollectionMembership, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMemberships(idStr)
}

func (g *Client) GetCollectionMembershipsByCollectionID(id uint64) ([]*pb.CollectionMembership, error) {
	query := fmt.Sprintf(`where collection = %d; fields *;`, id)
	return g.GetCollectionMemberships(query)
}

func (g *Client) GetCollectionMembershipsByCollectionMembershipTypeID(id uint64) ([]*pb.CollectionMembership, error) {
	query := fmt.Sprintf(`where type = %d; fields *;`, id)
	return g.GetCollectionMemberships(query)
}

func (g *Client) GetCollectionMembershipsByCollectionMembershipTypeIDs(ids []uint64) ([]*pb.CollectionMembership, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCollectionMemberships(idStr)
}

func (g *Client) GetCollectionMembershipsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	collectionMemberships, err := g.GetCollectionMemberships(query)
	if err != nil {
		return 0, err
	}
	return int(collectionMemberships[0].Id), nil
}
