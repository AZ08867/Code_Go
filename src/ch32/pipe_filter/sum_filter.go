package pipe_filter

import "errors"

var ErrSumFilterWrongFormatError = errors.New("输入数据应为整数切片")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, ErrSumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
