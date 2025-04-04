package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameReleaseFormats(query string) ([]*pb.GameReleaseFormat, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_release_formats.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameReleaseFormatResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamereleaseformats) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamereleaseformats, nil
}

func (g *igdb) GetGameReleaseFormatByID(id uint64) (*pb.GameReleaseFormat, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameReleaseFormats, err := g.GetGameReleaseFormats(query)
	if err != nil {
		return nil, err
	}
	return gameReleaseFormats[0], nil
}

func (g *igdb) GetGameReleaseFormatsByIDs(ids []uint64) ([]*pb.GameReleaseFormat, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameReleaseFormats(idStr)
}
