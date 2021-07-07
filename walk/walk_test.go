package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age     int
	Address string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name      string
		Input     interface{}
		Expecteds []string
	}{
		{"输入一个string字段的struct",
			struct {
				S string
			}{"sicong"},
			[]string{"sicong"}},
		{"输入二个string字段的struct",
			struct {
				S string
				I int
			}{"sicong", 33},
			[]string{"sicong"}},
		{"输入嵌套string的struct",
			&Person{Name: "sicong", Profile: Profile{33, "shaoxing"}},
			[]string{"sicong", "shaoxing"}},
		{"输入string",
			"sicong",
			[]string{"sicong"}},
		{"输入切片",
			[]Profile{
				{33, "shaoxing"},
				{44, "shanghai"},
			},
			[]string{"shaoxing", "shanghai"}},
		{"输入数组",
			[2]Profile{
				{33, "shaoxing"},
				{44, "shanghai"},
			},
			[]string{"shaoxing", "shanghai"}},
		{"输入Map",
			map[string]string{
				"22": "aa",
				"33": "bb",
			},
			[]string{"aa", "bb"}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var got []string
			Walk(c.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, c.Expecteds) {
				t.Errorf("got %#v, want %#v", got, c.Expecteds)
			}
		})

	}
	t.Run("Test Maps", func(t *testing.T) {
		input := map[string]string{
			"22": "aa",
			"33": "bb",
		}
		var got []string
		Walk(input, func(input string) {
			got = append(got, input)
		})
		assertContain(t, got, "aa")
		assertContain(t, got, "bb")
	})
}

func assertContain(t *testing.T, got []string, expected string) {
	contains := false
	for _, word := range got {
		if word == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("got %#v,don't have %s", got, expected)
	}
}
