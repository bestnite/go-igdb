package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameVideos struct{ BaseEndpoint }

func (a *GameVideos) Query(query string) ([]*pb.GameVideo, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_videos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVideoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamevideos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamevideos, nil
}
