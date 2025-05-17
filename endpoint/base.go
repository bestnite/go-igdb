package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"
	"google.golang.org/protobuf/proto"

	"github.com/go-resty/resty/v2"
)

type RequestFunc func(method string, URL string, dataBody any) (*resty.Response, error)

type BaseEndpoint[T any] struct {
	request      RequestFunc
	endpointName Name
	queryFunc    func(string) ([]*T, error)
}

func (b *BaseEndpoint[T]) GetEndpointName() Name {
	return b.endpointName
}

func (b *BaseEndpoint[T]) Query(query string) ([]*T, error) {
	if b.queryFunc == nil {
		return nil, fmt.Errorf("query method must be implemented by specific endpoint")
	}
	return b.queryFunc(query)
}

func (b *BaseEndpoint[T]) GetByID(id uint64) (*T, error) {
	res, err := b.Query(fmt.Sprintf("where id = %d; fields *;", id))
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no results")
	}
	return res[0], nil
}

func (b *BaseEndpoint[T]) GetByIDs(ids []uint64) ([]*T, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("ids cant be empty")
	}
	batches := make([][]uint64, 0)
	for i := 0; i < len(ids); i += 500 {
		end := min(i+500, len(ids))
		batches = append(batches, ids[i:end])
	}
	res := []*T{}
	for _, batch := range batches {
		builder := strings.Builder{}
		for i, v := range batch {
			if i > 0 {
				builder.WriteByte(',')
			}
			builder.WriteString(strconv.FormatUint(v, 10))
		}
		batchRes, err := b.Query(fmt.Sprintf("where id = (%s); fields *; limit 500;", builder.String()))
		if err != nil {
			return nil, err
		}
		res = append(res, batchRes...)
	}
	return res, nil
}

func (b *BaseEndpoint[T]) Count() (uint64, error) {
	resp, err := b.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s/count.pb", b.endpointName), "")
	if err != nil {
		return 0, fmt.Errorf("failed to request: %w", err)
	}

	var res pb.Count
	if err = proto.Unmarshal(resp.Body(), &res); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return uint64(res.Count), nil
}

func (b *BaseEndpoint[T]) Paginated(offset, limit uint64) ([]*T, error) {
	return b.Query(fmt.Sprintf("offset %d; limit %d; fields *; sort id asc;", offset, limit))
}

type EntityEndpoint[T any] interface {
	GetEndpointName() Name
	Query(string) ([]*T, error)
	GetByID(uint64) (*T, error)
	GetByIDs([]uint64) ([]*T, error)
	Count() (uint64, error)
	Paginated(uint64, uint64) ([]*T, error)
}
