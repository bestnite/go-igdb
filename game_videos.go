package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameVideos(query string) ([]*pb.GameVideo, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_videos.pb", query)
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

func (g *igdb) GetGameVideoByID(id uint64) (*pb.GameVideo, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameVideos, err := g.GetGameVideos(query)
	if err != nil {
		return nil, err
	}
	return gameVideos[0], nil
}

func (g *igdb) GetGameVideosByIDs(ids []uint64) ([]*pb.GameVideo, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVideos(idStr)
}
