package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetThemes(query string) ([]*pb.Theme, error) {
	resp, err := g.Request("https://api.igdb.com/v4/themes.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ThemeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Themes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Themes, nil
}

func (g *igdb) GetThemeByID(id uint64) (*pb.Theme, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	themes, err := g.GetThemes(query)
	if err != nil {
		return nil, err
	}
	return themes[0], nil
}

func (g *igdb) GetThemesByIDs(ids []uint64) ([]*pb.Theme, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetThemes(idStr)
}

func (g *igdb) GetThemesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	themes, err := g.GetThemes(query)
	if err != nil {
		return 0, err
	}
	return int(themes[0].Id), nil
}
