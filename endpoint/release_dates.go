package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ReleaseDates struct{ BaseEndpoint }

func (a *ReleaseDates) Query(query string) ([]*pb.ReleaseDate, error) {
	resp, err := a.request("https://api.igdb.com/v4/release_dates.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedates) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedates, nil
}
