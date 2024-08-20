package pipe_filter

import (
	"reflect"
	"testing"
)

func TestSplitFilter_Process(t *testing.T) {
	tests := []struct {
		name      string
		delimiter string
		input     interface{}
		want      interface{}
		wantErr   bool
	}{
		{
			name:      "正常分割",
			delimiter: ",",
			input:     "a,b,c",
			want:      []string{"a", "b", "c"},
			wantErr:   false,
		},
		{
			name:      "空字符串分割",
			delimiter: ",",
			input:     "",
			want:      []string{""},
			wantErr:   false,
		},
		{
			name:      "单个字符分割",
			delimiter: "|",
			input:     "abc",
			want:      []string{"abc"},
			wantErr:   false,
		},
		{
			name:      "错误输入类型",
			delimiter: ",",
			input:     123, // 非字符串类型
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := NewSplitFilter(tt.delimiter)
			got, err := sf.Process(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("SplitFilter.Process() 错误 = %v, 期望错误 %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("SplitFilter.Process() = %v, 期望 %v", got, tt.want)
				}
			}
		})
	}
}
