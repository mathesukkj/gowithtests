package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"nested struct",
			Person{"Matheus", Profile{18, "BH"}},
			[]string{"Matheus", "BH"},
		},
		{
			"pointers to things",
			&Person{
				"Matheus",
				Profile{18, "BH"},
			},
			[]string{"Matheus", "BH"},
		},
		{
			"slices",
			[]Profile{
				{18, "BH"},
				{19, "SP"},
			},
			[]string{"BH", "SP"},
		},
		{
			"arrays",
			[2]Profile{
				{18, "BH"},
				{19, "SP"},
			},
			[]string{"BH", "SP"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
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
