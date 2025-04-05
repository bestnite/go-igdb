package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AlternativeNames struct{ BaseEndpoint }

func (a *AlternativeNames) Query(query string) ([]*pb.AlternativeName, error) {
	resp, err := a.request("https://api.igdb.com/v4/alternative_names.pb", query)
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
