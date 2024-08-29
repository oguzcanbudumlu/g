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
		{
			"arrays",
			[2]struct {
				Year     int
				Location string
			}{{
				1907,
				"Kadikoy",
			}, {
				1903, "Besiktas",
			},
			},
			[]string{"Kadikoy", "Besiktas"},
		},
		{
			"maps",
			map[string]string{
				"Fenerbahce": "Kadikoy",
				"Besiktas":   "Besiktas",
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

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
