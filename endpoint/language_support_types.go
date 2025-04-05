package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type LanguageSupportTypes struct{ BaseEndpoint }

func (a *LanguageSupportTypes) Query(query string) ([]*pb.LanguageSupportType, error) {
	resp, err := a.request("https://api.igdb.com/v4/language_support_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageSupportTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languagesupporttypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languagesupporttypes, nil
}
