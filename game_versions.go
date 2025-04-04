package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameVersions(query string) ([]*pb.GameVersion, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_versions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameversions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameversions, nil
}

func (g *igdb) GetGameVersionByID(id uint64) (*pb.GameVersion, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameVersions, err := g.GetGameVersions(query)
	if err != nil {
		return nil, err
	}
	return gameVersions[0], nil
}

func (g *igdb) GetGameVersionsByIDs(ids []uint64) ([]*pb.GameVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersions(idStr)
}

func (g *igdb) GetGameVersionsByGameID(id uint64) ([]*pb.GameVersion, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameVersions(query)
}

func (g *igdb) GetGameVersionsByGameIDs(ids []uint64) ([]*pb.GameVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersions(idStr)
}
