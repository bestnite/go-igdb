package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameEngineLogos(query string) ([]*pb.GameEngineLogo, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_engine_logos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameenginelogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameenginelogos, nil
}

func (g *igdb) GetGameEngineLogoByID(id uint64) (*pb.GameEngineLogo, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameEngineLogos, err := g.GetGameEngineLogos(query)
	if err != nil {
		return nil, err
	}
	return gameEngineLogos[0], nil
}

func (g *igdb) GetGameEngineLogosByIDs(ids []uint64) ([]*pb.GameEngineLogo, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameEngineLogos(idStr)
}
