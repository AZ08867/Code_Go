package pipe_filter

import (
	"reflect"
	"testing"
)

func TestToIntFilter_Process(t *testing.T) {
	tests := []struct {
		name    string
		input   Request
		want    Response
		wantErr bool
	}{
		{
			name:    "正常转换",
			input:   []string{"1", "2", "3"},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "空切片",
			input:   []string{},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "包含非数字字符串",
			input:   []string{"1", "a", "3"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "错误输入类型",
			input:   "不是切片",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tif := NewToIntFilter()
			got, err := tif.Process(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntFilter.Process() 错误 = %v, 期望错误 %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ToIntFilter.Process() = %v, 期望 %v", got, tt.want)
				}
			}
		})
	}
}