package main

import (
	"reflect"
	"testing"
)

var packsarr = []int{250, 500, 1000, 2000, 5000}

func TestCalcOrder(t *testing.T) {
	//create test structure

	type test struct {
		name  string
		input int
		want  []int
	}

	//populate tests
	tests := []test{
		{name: "test 1", input: 1, want: []int{250}},
		{name: "test 250", input: 250, want: []int{250}},
		{name: "test 251", input: 251, want: []int{500}},
		{name: "test 501", input: 501, want: []int{500, 250}},
		{name: "test 12001", input: 12001, want: []int{5000, 5000, 2000, 250}},
	}

	//loop through tests
	for _, tc := range tests {
		got := calcOrder(packsarr, tc.input)
		// if expected result matches the result then tests have passed else log error
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		} else {
			t.Log(tc.name + ": PASSED")
		}
	}

}
