package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetReleaseDates(query string) ([]*pb.ReleaseDate, error) {
	resp, err := g.Request("https://api.igdb.com/v4/release_dates.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedates) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedates, nil
}

func (g *igdb) GetReleaseDateByID(id uint64) (*pb.ReleaseDate, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	releaseDates, err := g.GetReleaseDates(query)
	if err != nil {
		return nil, err
	}
	return releaseDates[0], nil
}

func (g *igdb) GetReleaseDatesByIDs(ids []uint64) ([]*pb.ReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDates(idStr)
}
