package igdb

import (
	"errors"
	"fmt"
	"github/bestnite/go-igdb/constant"
	pb "github/bestnite/go-igdb/proto"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameIDBySteamAppID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game_type.id = 0 & uid = "%d"; fields game;`, id)
	resp, err := g.Request(constant.IGDBExternalGameURL, query)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch IGDB ID by Steam App ID %d: %w", id, err)
	}

	res := pb.ExternalGameResult{}
	if err = proto.Unmarshal(resp.Body(), &res); err != nil {
		return 0, fmt.Errorf("failed to unmarshal IGDB response for Steam App ID %d: %w", id, err)
	}

	if len(res.Externalgames) == 0 || res.Externalgames[0].Game.Id == 0 {
		return 0, errors.New("no matching IGDB game found")
	}

	return res.Externalgames[0].Game.Id, nil
}

func (g *igdb) GetSteamIDByGameID(id uint64) (uint64, error) {
	query := fmt.Sprintf(`where game = %v & game_type.id = 0; fields *;`, id)
	resp, err := g.Request(constant.IGDBExternalGameURL, query)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch IGDB websites for IGDB ID %d: %w", id, err)
	}

	res := pb.ExternalGameResult{}
	if err := proto.Unmarshal(resp.Body(), &res); err != nil {
		return 0, fmt.Errorf("failed to unmarshal IGDB websites response for IGDB ID %d: %w", id, err)
	}

	if len(res.Externalgames) == 0 {
		return 0, errors.New("steam ID not found")
	}

	return strconv.ParseUint(res.Externalgames[0].Uid, 10, 64)
}
