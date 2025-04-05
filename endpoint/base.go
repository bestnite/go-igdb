package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type BaseEndpoint[T any] struct {
	request      func(URL string, dataBody any) (*resty.Response, error)
	endpointName EndpointName
	queryFunc    func(string) ([]*T, error)
}

func (b *BaseEndpoint[T]) GetEndpointName() EndpointName {
	return b.endpointName
}

func (b *BaseEndpoint[T]) Query(query string) ([]*T, error) {
	if b.queryFunc == nil {
		return nil, fmt.Errorf("Query method must be implemented by specific endpoint")
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
	builder := strings.Builder{}
	for i, v := range ids {
		if i > 0 {
			builder.WriteByte(',')
		}
		builder.WriteString(strconv.FormatUint(v, 10))
	}
	return b.Query(fmt.Sprintf("where id = (%s); fields *;", builder.String()))
}

func (b *BaseEndpoint[T]) GetLastOneId() (uint64, error) {
	res, err := b.Query("fields *; sort id desc; limit 1;")
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, fmt.Errorf("no results")
	}
	type IdGetter interface {
		GetId() uint64
	}
	item, ok := any(res[0]).(IdGetter)
	if !ok {
		return 0, fmt.Errorf("invalid type")
	}
	return item.GetId(), nil
}

func (b *BaseEndpoint[T]) Paginated(offset, limit uint64) ([]*T, error) {
	return b.Query(fmt.Sprintf("offset %d; limit %d; fields *; sort id asc;", offset, limit))
}

type EntityEndpoint[T any] interface {
	GetEndpointName() EndpointName
	Query(string) ([]*T, error)
	GetByID(uint64) (*T, error)
	GetByIDs([]uint64) ([]*T, error)
	GetLastOneId() (uint64, error)
	Paginated(uint64, uint64) ([]*T, error)
}
