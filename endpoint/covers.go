package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Covers struct{ BaseEndpoint }

func (a *Covers) Query(query string) ([]*pb.Cover, error) {
	resp, err := a.request("https://api.igdb.com/v4/covers.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CoverResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Covers) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Covers, nil
}
