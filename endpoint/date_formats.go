package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type DateFormats struct {
	BaseEndpoint[pb.DateFormat]
}

func NewDateFormats(request RequestFunc) *DateFormats {
	a := &DateFormats{
		BaseEndpoint[pb.DateFormat]{
			endpointName: EPDateFormats,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *DateFormats) Query(query string) ([]*pb.DateFormat, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
