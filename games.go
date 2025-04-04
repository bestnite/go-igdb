package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGames(query string) ([]*pb.Game, error) {
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

func (g *igdb) GetGameByID(id uint64) (*pb.Game, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	games, err := g.GetGames(query)
	if err != nil {
		return nil, err
	}
	return games[0], nil
}

func (g *igdb) GetGameByIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByCollectionID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where collection = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByCollectionIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where collection = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByCoverID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where cover = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByCoverIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where cover = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByFranchiseID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where franchise = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByFranchiseIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where franchise = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByGameStatusID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where game_status = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByGameStatusIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_status = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByGameTypeID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where game_type = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByGameTypeIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByParentGameID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where parent_game = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByParentGameIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where parent_game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}

func (g *igdb) GetGameByVersionParentGameID(id uint64) ([]*pb.Game, error) {
	query := fmt.Sprintf(`where version_parent = %d; fields *;`, id)
	return g.GetGames(query)
}

func (g *igdb) GetGamesByVersionParentGameIDs(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where version_parent = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGames(idStr)
}
