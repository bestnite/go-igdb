package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type LanguageSupportTypes struct {
	BaseEndpoint[pb.LanguageSupportType]
}

func NewLanguageSupportTypes(request RequestFunc) *LanguageSupportTypes {
	a := &LanguageSupportTypes{
		BaseEndpoint[pb.LanguageSupportType]{
			endpointName: EPLanguageSupportTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *LanguageSupportTypes) Query(query string) ([]*pb.LanguageSupportType, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
