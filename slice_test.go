package jgunify

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 2, 3, 4, 4, 4, 5}, []int{1, 2, 3, 4, 5}},          // 测试包含重复元素的整数切片
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, []int{1}},                      // 测试所有元素都相同的整数切片
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}}, // 测试没有重复元素的整数切片
	}
	var tests2 = []struct {
		input []string
		want  []string
	}{
		{[]string{"apple", "banana", "cherry", "banana", "apple"}, []string{"apple", "banana", "cherry"}}, // 测试包含重复元素的字符串切片
		{[]string{"apple", "apple", "apple"}, []string{"apple"}},                                          // 测试所有元素都相同的字符串切片
		{[]string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},                    // 测试没有重复元素的字符串切片
	}

	for _, tt := range tests {
		if got := Unique(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Unique(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
	for _, tt := range tests2 {
		if got := Unique(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Unique(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
