package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type LanguageSupports struct {
	BaseEndpoint[pb.LanguageSupport]
}

func NewLanguageSupports(request func(URL string, dataBody any) (*resty.Response, error)) *LanguageSupports {
	a := &LanguageSupports{
		BaseEndpoint[pb.LanguageSupport]{
			endpointName: EPLanguageSupports,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *LanguageSupports) Query(query string) ([]*pb.LanguageSupport, error) {
	resp, err := a.request("https://api.igdb.com/v4/language_supports.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.LanguageSupportResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Languagesupports) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Languagesupports, nil
}
