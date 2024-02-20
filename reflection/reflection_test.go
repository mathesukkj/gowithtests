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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "mu",
			"Sheep": "bé",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "mu")
		assertContains(t, got, "bé")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didnt", haystack, needle)
	}
}
