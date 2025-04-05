package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetScreenshots(query string) ([]*pb.Screenshot, error) {
	resp, err := g.Request("https://api.igdb.com/v4/screenshots.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ScreenshotResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Screenshots) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Screenshots, nil
}

func (g *igdb) GetScreenshotByID(id uint64) (*pb.Screenshot, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	screenshots, err := g.GetScreenshots(query)
	if err != nil {
		return nil, err
	}
	return screenshots[0], nil
}

func (g *igdb) GetScreenshotsByIDs(ids []uint64) ([]*pb.Screenshot, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetScreenshots(idStr)
}

func (g *igdb) GetScreenshotsByGameID(id uint64) ([]*pb.Screenshot, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetScreenshots(query)
}

func (g *igdb) GetScreenshotsByGameIDs(ids []uint64) ([]*pb.Screenshot, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetScreenshots(idStr)
}

func (g *igdb) GetScreenshotsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	screenshots, err := g.GetScreenshots(query)
	if err != nil {
		return 0, err
	}
	return int(screenshots[0].Id), nil
}
