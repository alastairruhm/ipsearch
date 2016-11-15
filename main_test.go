package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_trimSlice(t *testing.T) {
	type args struct {
		src    []string
		cutset string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test trim quota rune",
			args{
				[]string{"\"中国\"", "\"天津\"", "\"天津\"", "\"\"", "\"鹏博士\""},
				"\"",
			},
			[]string{"中国", "天津", "天津", "鹏博士"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimSlice(tt.args.src, tt.args.cutset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractData(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test extractData",
			args{
				`["中国","天津","天津","","鹏博士"]`,
			},
			[]string{"中国", "天津", "天津", "鹏博士"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractData(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractData() = %v, want %v", got, tt.want)
			}
		})
	}
}
