package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameLocalizations(query string) ([]*pb.GameLocalization, error) {
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

func (g *Client) GetGameLocalizationByID(id uint64) (*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameLocalizations, err := g.GetGameLocalizations(query)
	if err != nil {
		return nil, err
	}
	return gameLocalizations[0], nil
}

func (g *Client) GetGameLocalizationsByIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *Client) GetGameLocalizationsByGameID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *Client) GetGameLocalizationsByGameIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *Client) GetGameLocalizationsByCoverID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where cover = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *Client) GetGameLocalizationsByCoverIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where cover = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *Client) GetGameLocalizationsByRegionID(id uint64) ([]*pb.GameLocalization, error) {
	query := fmt.Sprintf(`where region = %d; fields *;`, id)
	return g.GetGameLocalizations(query)
}

func (g *Client) GetGameLocalizationsByRegionIDs(ids []uint64) ([]*pb.GameLocalization, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where region = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameLocalizations(idStr)
}

func (g *Client) GetGameLocalizationsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameLocalizations, err := g.GetGameLocalizations(query)
	if err != nil {
		return 0, err
	}
	return int(gameLocalizations[0].Id), nil
}
