package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type EventLogos struct{ BaseEndpoint }

func (a *EventLogos) Query(query string) ([]*pb.EventLogo, error) {
	resp, err := a.request("https://api.igdb.com/v4/event_logos.pb", query)
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
