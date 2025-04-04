package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetEventNetworks(query string) ([]*pb.EventNetwork, error) {
	resp, err := g.Request("https://api.igdb.com/v4/event_networks.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventNetworkResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Eventnetworks) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Eventnetworks, nil
}

func (g *igdb) GetEventNetworkByID(id uint64) (*pb.EventNetwork, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	eventNetworks, err := g.GetEventNetworks(query)
	if err != nil {
		return nil, err
	}
	return eventNetworks[0], nil
}

func (g *igdb) GetEventNetworksByIDs(ids []uint64) ([]*pb.EventNetwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEventNetworks(idStr)
}

func (g *igdb) GetEventNetworksByEventID(id uint64) ([]*pb.EventNetwork, error) {
	query := fmt.Sprintf(`where event = %d; fields *;`, id)
	return g.GetEventNetworks(query)
}

func (g *igdb) GetEventNetworksByEventIDs(ids []uint64) ([]*pb.EventNetwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where event = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEventNetworks(idStr)
}

func (g *igdb) GetEventNetworksByNetworkTypeID(id uint64) ([]*pb.EventNetwork, error) {
	query := fmt.Sprintf(`where network_type = %d; fields *;`, id)
	return g.GetEventNetworks(query)
}

func (g *igdb) GetEventNetworksByNetworkTypeIDs(ids []uint64) ([]*pb.EventNetwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where network_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEventNetworks(idStr)
}

func (g *igdb) GetEventNetworksLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	eventNetworks, err := g.GetEventNetworks(query)
	if err != nil {
		return 0, err
	}
	return int(eventNetworks[0].Id), nil
}
