package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetEvents(query string) ([]*pb.Event, error) {
	resp, err := g.Request("https://api.igdb.com/v4/events.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Events) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Events, nil
}

func (g *Client) GetEventByID(id uint64) (*pb.Event, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	events, err := g.GetEvents(query)
	if err != nil {
		return nil, err
	}
	return events[0], nil
}

func (g *Client) GetEventsByIDs(ids []uint64) ([]*pb.Event, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEvents(idStr)
}

func (g *Client) GetEventsByEventLogoID(id uint64) ([]*pb.Event, error) {
	query := fmt.Sprintf(`where event_logo = %d; fields *;`, id)
	return g.GetEvents(query)
}

func (g *Client) GetEventsByEventLogoIDs(ids []uint64) ([]*pb.Event, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where event_logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEvents(idStr)
}

func (g *Client) GetEventsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	events, err := g.GetEvents(query)
	if err != nil {
		return 0, err
	}
	return int(events[0].Id), nil
}
