package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameLocalizations struct{ BaseEndpoint }

func (a *GameLocalizations) Query(query string) ([]*pb.GameLocalization, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_localizations.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameLocalizationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamelocalizations) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamelocalizations, nil
}
