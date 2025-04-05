package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetNetworkTypes(query string) ([]*pb.NetworkType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/network_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.NetworkTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Networktypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Networktypes, nil
}

func (g *igdb) GetNetworkTypeByID(id uint64) (*pb.NetworkType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	networkTypes, err := g.GetNetworkTypes(query)
	if err != nil {
		return nil, err
	}
	return networkTypes[0], nil
}

func (g *igdb) GetNetworkTypesByIDs(ids []uint64) ([]*pb.NetworkType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetNetworkTypes(idStr)
}

func (g *igdb) GetNetworkTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	networkTypes, err := g.GetNetworkTypes(query)
	if err != nil {
		return 0, err
	}
	return int(networkTypes[0].Id), nil
}
