package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Matheus"},
			[]string{"Matheus"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Matheus", "BH"},
			[]string{"Matheus", "BH"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Matheus", 18},
			[]string{"Matheus"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
