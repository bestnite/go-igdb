package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetEventLogos(query string) ([]*pb.EventLogo, error) {
	resp, err := g.Request("https://api.igdb.com/v4/event_logos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.EventLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Eventlogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Eventlogos, nil
}

func (g *igdb) GetEventLogoByID(id uint64) (*pb.EventLogo, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	eventLogos, err := g.GetEventLogos(query)
	if err != nil {
		return nil, err
	}
	return eventLogos[0], nil
}

func (g *igdb) GetEventLogosByIDs(ids []uint64) ([]*pb.EventLogo, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetEventLogos(idStr)
}
