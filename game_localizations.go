package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameLocalizations(query string) ([]*pb.GameLocalization, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_localizations.pb", query)
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

func (g *igdb) GetGameLocalizationByID(id uint64) (*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameLocalizations, err := g.GetGameLocalizations(query)
	if err != nil {
		return nil, err
	}
	return gameLocalizations[0], nil
}

func (g *igdb) GetGameLocalizationsByIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *igdb) GetGameLocalizationsByGameID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *igdb) GetGameLocalizationsByGameIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *igdb) GetGameLocalizationsByCoverID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where cover = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *igdb) GetGameLocalizationsByCoverIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where cover = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *igdb) GetGameLocalizationsByRegionID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where region = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *igdb) GetGameLocalizationsByRegionIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where region = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *igdb) GetGameLocalizationsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameLocalizations, err := g.GetGameLocalizations(query)
	if err != nil {
		return 0, err
	}
	return int(gameLocalizations[0].Id), nil
}
