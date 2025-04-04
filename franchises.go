package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetFranchises(query string) ([]*pb.Franchise, error) {
	resp, err := g.Request("https://api.igdb.com/v4/franchises.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.FranchiseResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Franchises) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Franchises, nil
}

func (g *igdb) GetFranchiseByID(id uint64) (*pb.Franchise, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	franchises, err := g.GetFranchises(query)
	if err != nil {
		return nil, err
	}
	return franchises[0], nil
}

func (g *igdb) GetFranchisesByIDs(ids []uint64) ([]*pb.Franchise, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetFranchises(idStr)
}

func (g *igdb) GetFranchisesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	franchises, err := g.GetFranchises(query)
	if err != nil {
		return 0, err
	}
	return int(franchises[0].Id), nil
}
