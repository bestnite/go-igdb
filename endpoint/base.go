package endpoint

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	pb "git.nite07.com/nite/go-igdb/proto"
	"google.golang.org/protobuf/proto"

	"github.com/go-resty/resty/v2"
)

type RequestFunc func(ctx context.Context, method string, URL string, dataBody any) (*resty.Response, error)

type BaseEndpoint[T any] struct {
	request      RequestFunc
	endpointName Name
	queryFunc    func(context.Context, string) ([]*T, error)
}

func (b *BaseEndpoint[T]) GetEndpointName() Name {
	return b.endpointName
}

func (b *BaseEndpoint[T]) Query(ctx context.Context, query string) ([]*T, error) {
	if b.queryFunc == nil {
		return nil, fmt.Errorf("query method must be implemented by specific endpoint")
	}
	return b.queryFunc(ctx, query)
}

func (b *BaseEndpoint[T]) GetByID(ctx context.Context, id uint64) (*T, error) {
	if id == 0 {
		return nil, errors.New("id cant be 0")
	}
	res, err := b.Query(ctx, fmt.Sprintf("where id = %d; fields *;", id))
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("no results")
	}
	return res[0], nil
}

func (b *BaseEndpoint[T]) GetByIDs(ctx context.Context, ids []uint64) ([]*T, error) {
	if len(ids) == 0 {
		return []*T{}, nil
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
		batchRes, err := b.Query(ctx, fmt.Sprintf("where id = (%s); fields *; limit 500;", builder.String()))
		if err != nil {
			return nil, err
		}
		res = append(res, batchRes...)
	}
	return res, nil
}

func (b *BaseEndpoint[T]) Count(ctx context.Context) (uint64, error) {
	resp, err := b.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s/count.pb", b.endpointName), "")
	if err != nil {
		return 0, fmt.Errorf("failed to request: %w", err)
	}

	var res pb.Count
	if err = proto.Unmarshal(resp.Body(), &res); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if res.Count > 0 {
		return uint64(res.Count), nil
	} else {
		return 0, fmt.Errorf("failed to count, count should larger than 0, but got %v", res.Count)
	}
}

func (b *BaseEndpoint[T]) Paginated(ctx context.Context, offset, limit uint64) ([]*T, error) {
	return b.Query(ctx, fmt.Sprintf("offset %d; limit %d; fields *; sort id asc;", offset, limit))
}

// queryPB 是 Go 1.27 泛型方法：发起 protobuf 查询并把各 endpoint 的
// XxxResult 解包为实体切片。R 为 Result 结构体类型，P 为其指针（须实现
// proto.Message），两者均可从 extract 闭包推断；extract 负责取出实体
// 字段（如 r.Companies）。
func (b *BaseEndpoint[T]) queryPB[R any, P interface {
	*R
	proto.Message
}](extract func(P) []*T) func(context.Context, string) ([]*T, error) {
	return func(ctx context.Context, query string) ([]*T, error) {
		resp, err := b.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", b.endpointName), query)
		if err != nil {
			return nil, fmt.Errorf("failed to request: %w", err)
		}

		data := P(new(R))
		if err = proto.Unmarshal(resp.Body(), data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal: %w", err)
		}

		return extract(data), nil
	}
}

type EntityEndpoint[T any] interface {
	GetEndpointName() Name
	Query(context.Context, string) ([]*T, error)
	GetByID(context.Context, uint64) (*T, error)
	GetByIDs(context.Context, []uint64) ([]*T, error)
	Count(context.Context) (uint64, error)
	Paginated(context.Context, uint64, uint64) ([]*T, error)
}
