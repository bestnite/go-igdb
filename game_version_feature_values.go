package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameVersionFeatureValues(query string) ([]*pb.GameVersionFeatureValue, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_version_feature_values.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionFeatureValueResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameversionfeaturevalues) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameversionfeaturevalues, nil
}

func (g *igdb) GetGameVersionFeatureValueByID(id uint64) (*pb.GameVersionFeatureValue, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameVersionFeatureValues, err := g.GetGameVersionFeatureValues(query)
	if err != nil {
		return nil, err
	}
	return gameVersionFeatureValues[0], nil
}

func (g *igdb) GetGameVersionFeatureValuesByIDs(ids []uint64) ([]*pb.GameVersionFeatureValue, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersionFeatureValues(idStr)
}

func (g *igdb) GetGameVersionFeatureValuesByGameID(id uint64) ([]*pb.GameVersionFeatureValue, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetGameVersionFeatureValues(query)
}

func (g *igdb) GetGameVersionFeatureValuesByGameIDs(ids []uint64) ([]*pb.GameVersionFeatureValue, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersionFeatureValues(idStr)
}

func (g *igdb) GetGameVersionFeatureValuesByGameVersionFeatureID(id uint64) ([]*pb.GameVersionFeatureValue, error) {
	query := fmt.Sprintf(`where game_version_feature = %d; fields *;`, id)
	return g.GetGameVersionFeatureValues(query)
}

func (g *igdb) GetGameVersionFeatureValuesByGameVersionFeatureIDs(ids []uint64) ([]*pb.GameVersionFeatureValue, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_version_feature = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersionFeatureValues(idStr)
}

func (g *igdb) GetGameVersionFeatureValuesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameVersionFeatureValues, err := g.GetGameVersionFeatureValues(query)
	if err != nil {
		return 0, err
	}
	return int(gameVersionFeatureValues[0].Id), nil
}
