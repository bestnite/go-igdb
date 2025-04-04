package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetReleaseDateStatuses(query string) ([]*pb.ReleaseDateStatus, error) {
	resp, err := g.Request("https://api.igdb.com/v4/release_date_statuses.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedatestatuses) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedatestatuses, nil
}

func (g *igdb) GetReleaseDateStatusByID(id uint64) (*pb.ReleaseDateStatus, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	releaseDateStatuses, err := g.GetReleaseDateStatuses(query)
	if err != nil {
		return nil, err
	}
	return releaseDateStatuses[0], nil
}

func (g *igdb) GetReleaseDateStatusesByIDs(ids []uint64) ([]*pb.ReleaseDateStatus, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDateStatuses(idStr)
}
