package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CharacterSpecies struct {
	BaseEndpoint[pb.CharacterSpecie]
}

func NewCharacterSpecies(request func(URL string, dataBody any) (*resty.Response, error)) *CharacterSpecies {
	a := &CharacterSpecies{
		BaseEndpoint[pb.CharacterSpecie]{
			endpointName: EPCharacterSpecies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CharacterSpecies) Query(query string) ([]*pb.CharacterSpecie, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterSpecieResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Characterspecies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Characterspecies, nil
}
