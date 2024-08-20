package pipe_filter

import (
	"reflect"
	"testing"
)

func TestSumFilter_Process(t *testing.T) {
	sf := NewSumFilter()

	tests := []struct {
		name    string
		input   Request
		want    Response
		wantErr error
	}{
		{
			name:    "正常情况",
			input:   []int{1, 2, 3, 4, 5},
			want:    15,
			wantErr: nil,
		},
		{
			name:    "空切片",
			input:   []int{},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "负数",
			input:   []int{-1, -2, 3},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "错误输入类型",
			input:   []string{"1", "2", "3"},
			want:    nil,
			wantErr: ErrSumFilterWrongFormatError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sf.Process(tt.input)
			if err != tt.wantErr {
				t.Errorf("SumFilter.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumFilter.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}