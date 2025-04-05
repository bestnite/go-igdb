package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type DateFormats struct{ BaseEndpoint }

func (a *DateFormats) Query(query string) ([]*pb.DateFormat, error) {
	resp, err := a.request("https://api.igdb.com/v4/date_formats.pb", query)
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
