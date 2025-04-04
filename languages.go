package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetLanguages(query string) ([]*pb.Language, error) {
	resp, err := g.Request("https://api.igdb.com/v4/languages.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languages) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languages, nil
}

func (g *igdb) GetLanguageByID(id uint64) (*pb.Language, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	languages, err := g.GetLanguages(query)
	if err != nil {
		return nil, err
	}
	return languages[0], nil
}

func (g *igdb) GetLanguagesByIDs(ids []uint64) ([]*pb.Language, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguages(idStr)
}
