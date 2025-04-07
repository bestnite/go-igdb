package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CharacterGenders struct {
	BaseEndpoint[pb.CharacterGender]
}

func NewCharacterGenders(request func(URL string, dataBody any) (*resty.Response, error)) *CharacterGenders {
	a := &CharacterGenders{
		BaseEndpoint[pb.CharacterGender]{
			endpointName: EPCharacterGenders,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CharacterGenders) Query(query string) ([]*pb.CharacterGender, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterGenderResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Charactergenders) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Charactergenders, nil
}
