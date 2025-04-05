package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Languages struct{ BaseEndpoint }

func (a *Languages) Query(query string) ([]*pb.Language, error) {
	resp, err := a.request("https://api.igdb.com/v4/languages.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languages) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languages, nil
}
