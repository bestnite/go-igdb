package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetDateFormats(query string) ([]*pb.DateFormat, error) {
	resp, err := g.Request("https://api.igdb.com/v4/date_formats.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.DateFormatResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Dateformats) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Dateformats, nil
}

func (g *Client) GetDateFormatByID(id uint64) (*pb.DateFormat, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	dateFormats, err := g.GetDateFormats(query)
	if err != nil {
		return nil, err
	}
	return dateFormats[0], nil
}

func (g *Client) GetDateFormatsByIDs(ids []uint64) ([]*pb.DateFormat, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetDateFormats(idStr)
}

func (g *Client) GetDateFormatsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	dateFormats, err := g.GetDateFormats(query)
	if err != nil {
		return 0, err
	}
	return int(dateFormats[0].Id), nil
}
