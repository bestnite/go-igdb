package igdb

import "fmt"

func AssertSingle[T any](data any, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	datas, ok := data.([]*T)
	if !ok {
		return nil, fmt.Errorf("failed to convert to []*T")
	}

	if len(datas) == 0 {
		return nil, fmt.Errorf("no results")
	}

	return datas[0], nil
}

func AssertSlice[T any](data any, err error) ([]*T, error) {
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	datas, ok := data.([]*T)
	if !ok {
		return nil, fmt.Errorf("failed to convert to []*T")
	}

	return datas, nil
}
