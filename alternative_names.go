package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetAlternativeNames(query string) ([]*pb.AlternativeName, error) {
	resp, err := g.Request("https://api.igdb.com/v4/alternative_names.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AlternativeNameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Alternativenames) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Alternativenames, nil
}

func (g *igdb) GetAlternativeNameByID(id uint64) (*pb.AlternativeName, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	alternativeNames, err := g.GetAlternativeNames(query)
	if err != nil {
		return nil, err
	}
	return alternativeNames[0], nil
}

func (g *igdb) GetAlternativeNamesByIDs(ids []uint64) ([]*pb.AlternativeName, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAlternativeNames(idStr)
}

func (g *igdb) GetAlternativeNamesByGameID(id uint64) ([]*pb.AlternativeName, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetAlternativeNames(query)
}

func (g *igdb) GetAlternativeNamesByGameIDs(ids []uint64) ([]*pb.AlternativeName, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetAlternativeNames(idStr)
}

func (g *igdb) GetAlternativeNamesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	alternativeNames, err := g.GetAlternativeNames(query)
	if err != nil {
		return 0, err
	}
	return int(alternativeNames[0].Id), nil
}
