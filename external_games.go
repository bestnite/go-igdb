package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetExternalGames(query string) ([]*pb.ExternalGame, error) {
	resp, err := g.Request("https://api.igdb.com/v4/external_games.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Externalgames) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Externalgames, nil
}

func (g *igdb) GetExternalGameByID(id uint64) (*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return nil, err
	}
	return externalGames[0], nil
}

func (g *igdb) GetExternalGamesByIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *igdb) GetGameIDBySteamAppID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game_type.id = 0 & uid = "%d"; fields game;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return 0, err
	}
	return externalGames[0].Game.Id, nil
}

func (g *igdb) GetSteamIDByGameID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game = %v & game_type.id = 0; fields *;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(externalGames[0].Uid, 10, 64)
}

func (g *igdb) GetExternalGamesByGameID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *igdb) GetExternalGamesByExternalGameSourceID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where external_game_source = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *igdb) GetExternalGamesByExternalGameSourceIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where external_game_source = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *igdb) GetExternalGamesByGameReleaseFormatID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where game_release_format = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *igdb) GetExternalGamesByGameReleaseFormatIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_release_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *igdb) GetExternalGamesByPlatformVersionID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where platform_version = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *igdb) GetExternalGamesByPlatformVersionIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_version = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}
