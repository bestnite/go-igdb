package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGames(query string) ([]*pb.Game, error) {
	resp, err := g.Request("https://api.igdb.com/v4/games.pb", query)
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

func (g *Client) GetGameByID(id uint64) (*pb.Game, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	games, err := g.GetGames(query)
	if err != nil {
		return nil, err
	}
	return games[0], nil
}

func (g *Client) GetGameByIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByCollectionID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where collection = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByCollectionIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where collection = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByCoverID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where cover = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByCoverIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where cover = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByFranchiseID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where franchise = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByFranchiseIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where franchise = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByGameStatusID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where game_status = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByGameStatusIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_status = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByGameTypeID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where game_type = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByGameTypeIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByParentGameID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where parent_game = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByParentGameIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where parent_game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGameByVersionParentGameID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where version_parent = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *Client) GetGamesByVersionParentGameIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where version_parent = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *Client) GetGamesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	games, err := g.GetGames(query)
	if err != nil {
		return 0, err
	}
	return int(games[0].Id), nil
}
