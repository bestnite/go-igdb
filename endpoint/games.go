package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Games struct{ BaseEndpoint }

func (a *Games) Query(query string) ([]*pb.Game, error) {
	resp, err := a.request("https://api.igdb.com/v4/games.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Games) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Games, nil
}
