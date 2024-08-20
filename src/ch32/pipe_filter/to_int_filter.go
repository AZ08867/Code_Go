package pipe_filter

import (
	"errors"
	"strconv"
)

var ErrToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ErrToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		atoi, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, atoi)
	}
	return ret, nil
}
