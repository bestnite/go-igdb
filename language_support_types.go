package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetLanguageSupportTypes(query string) ([]*pb.LanguageSupportType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/language_support_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageSupportTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languagesupporttypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languagesupporttypes, nil
}

func (g *igdb) GetLanguageSupportTypeByID(id uint64) (*pb.LanguageSupportType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	languageSupportTypes, err := g.GetLanguageSupportTypes(query)
	if err != nil {
		return nil, err
	}
	return languageSupportTypes[0], nil
}

func (g *igdb) GetLanguageSupportTypesByIDs(ids []uint64) ([]*pb.LanguageSupportType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetLanguageSupportTypes(idStr)
}

func (g *igdb) GetLanguageSupportTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	languageSupportTypes, err := g.GetLanguageSupportTypes(query)
	if err != nil {
		return 0, err
	}
	return int(languageSupportTypes[0].Id), nil
}
