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

func (g *igdb) GetReleaseDatesByGameID(id uint64) ([]*pb.ReleaseDate, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetReleaseDates(query)
}

func (g *igdb) GetReleaseDatesByGameIDs(ids []uint64) ([]*pb.ReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDates(idStr)
}

func (g *igdb) GetReleaseDatesByPlatformID(id uint64) ([]*pb.ReleaseDate, error) {
	query := fmt.Sprintf(`where platform = %d; fields *;`, id)
	return g.GetReleaseDates(query)
}

func (g *igdb) GetReleaseDatesByPlatformIDs(ids []uint64) ([]*pb.ReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDates(idStr)
}

func (g *igdb) GetReleaseDatesByReleaseRegionID(id uint64) ([]*pb.ReleaseDate, error) {
	query := fmt.Sprintf(`where release_region = %d; fields *;`, id)
	return g.GetReleaseDates(query)
}

func (g *igdb) GetReleaseDatesByReleaseRegionIDs(ids []uint64) ([]*pb.ReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where release_region = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDates(idStr)
}

func (g *igdb) GetReleaseDatesByStatusID(id uint64) ([]*pb.ReleaseDate, error) {
	query := fmt.Sprintf(`where status = %d; fields *;`, id)
	return g.GetReleaseDates(query)
}

func (g *igdb) GetReleaseDatesByStatusIDs(ids []uint64) ([]*pb.ReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where status = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDates(idStr)
}

func (g *igdb) GetReleaseDatesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	releaseDates, err := g.GetReleaseDates(query)
	if err != nil {
		return 0, err
	}
	return int(releaseDates[0].Id), nil
}
