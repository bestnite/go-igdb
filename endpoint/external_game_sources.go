package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ExternalGameSources struct{ BaseEndpoint }

func (a *ExternalGameSources) Query(query string) ([]*pb.ExternalGameSource, error) {
	resp, err := a.request("https://api.igdb.com/v4/external_game_sources.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameSourceResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Externalgamesources) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Externalgamesources, nil
}
