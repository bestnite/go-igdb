package igdb

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetExternalGames(query string) ([]*pb.ExternalGame, error) {
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

func (g *Client) GetExternalGameByID(id uint64) (*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return nil, err
	}
	return externalGames[0], nil
}

func (g *Client) GetExternalGamesByIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *Client) GetGameIDBySteamAppID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game_type.id = 0 & uid = "%d"; fields game;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return 0, err
	}
	return externalGames[0].Game.Id, nil
}

func (g *Client) GetSteamIDByGameID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game = %v & game_type.id = 0; fields *;`, id)
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(externalGames[0].Uid, 10, 64)
}

func (g *Client) GetExternalGamesByGameID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *Client) GetExternalGamesByExternalGameSourceID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where external_game_source = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *Client) GetExternalGamesByExternalGameSourceIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where external_game_source = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *Client) GetExternalGamesByGameReleaseFormatID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where game_release_format = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *Client) GetExternalGamesByGameReleaseFormatIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_release_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *Client) GetExternalGamesByPlatformVersionID(id uint64) ([]*pb.ExternalGame, error) {
	query := fmt.Sprintf(`where platform_version = %d; fields *;`, id)
	return g.GetExternalGames(query)
}

func (g *Client) GetExternalGamesByPlatformVersionIDs(ids []uint64) ([]*pb.ExternalGame, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_version = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetExternalGames(idStr)
}

func (g *Client) GetExternalGamesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	externalGames, err := g.GetExternalGames(query)
	if err != nil {
		return 0, err
	}
	return int(externalGames[0].Id), nil
}
