package endpoint

import (
	"github.com/go-resty/resty/v2"
)

type BaseEndpoint struct {
	request      func(URL string, dataBody any) (*resty.Response, error)
	endpointName EndpointName
}

func NewBaseEndpoint(request func(URL string, dataBody any) (*resty.Response, error), endpointName EndpointName) *BaseEndpoint {
	return &BaseEndpoint{
		request:      request,
		endpointName: endpointName,
	}
}

func (b *BaseEndpoint) GetEndpointName() EndpointName {
	return b.endpointName
}

func (b *BaseEndpoint) Query(query string) (any, error) {
	return nil, nil
}

func (b *BaseEndpoint) QueryAny(query string) (any, error) {
	return b.Query(query)
}

type Endpoint interface {
	GetEndpointName() EndpointName
}

type EntityEndpoint interface {
	QueryAny(query string) (any, error)
	GetEndpointName() EndpointName
}
