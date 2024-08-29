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

	t.Run("with maps because ordering is not guarenteed", func(t *testing.T) {
		aMap := map[string]string{
			"Fenerbahce": "Kadikoy",
			"Besiktas":   "Besiktas",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Kadikoy")
		assertContains(t, got, "Besiktas")
	})

	t.Run("with channels", func(t *testing.T) {
		type Profile struct {
			Year     int
			Location string
		}
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{1907, "Kadikoy"}
			aChannel <- Profile{1903, "Besiktas"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Kadikoy", "Besiktas"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
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
		t.Errorf("expected %v to contain %q but it did not", haystack, needle)
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
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
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
