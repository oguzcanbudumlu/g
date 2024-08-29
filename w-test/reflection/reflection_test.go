package reflection

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
			}{"Fenerbahce"},
			[]string{"Fenerbahce"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Fener", "bahce"},
			[]string{"Fener", "bahce"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Year int
			}{"Fenerbahce", 1907},
			[]string{"Fenerbahce"},
		},
		{
			"nested fields",
			struct {
				Name    string
				Profile struct {
					Year     int
					Location string
				}
			}{
				"Fenerbahce", struct {
					Year     int
					Location string
				}{
					1907,
					"Istanbul",
				},
			},
			[]string{"Fenerbahce", "Istanbul"},
		},
		{
			"pointer to things",
			&struct {
				Name    string
				Profile struct {
					Year     int
					Location string
				}
			}{
				"Fenerbahce", struct {
					Year     int
					Location string
				}{
					1907,
					"Istanbul",
				},
			},
			[]string{"Fenerbahce", "Istanbul"},
		},
		{
			"slices",
			[]struct {
				Year     int
				Location string
			}{
				{
					1907,
					"Kadikoy",
				},
				{
					1903,
					"Besiktas",
				},
			},
			[]string{"Kadikoy", "Besiktas"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
