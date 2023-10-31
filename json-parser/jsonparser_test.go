package jsonparser_test

import (
	"jsonparser"
	"reflect"
	"testing"
)

func TestParseNull(t *testing.T) {
	var want interface{}

	got, err := jsonparser.Parse("null")
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestParseBool(t *testing.T) {
	var want bool

	got, err := jsonparser.Parse("false")
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}

	want = true
	got, err = jsonparser.Parse("true")
	if err != nil {
		t.Error(err)
	}

	if want != got {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestParseNumber(t *testing.T) {
	var numberTests = []struct {
		in  string
		out int
	}{
		{"1", 1},
		{"20", 20},
		{"14526", 14526},
		{"69", 69},
	}

	for i := range numberTests {
		t.Run(numberTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(numberTests[i].in)
			if err != nil {
				t.Error(err)
			}
			if got != numberTests[i].out {
				t.Errorf("got: %v, want: %v", got, numberTests[i].out)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	var numberTests = []struct {
		in  string
		out string
	}{
		{" \"one\" ", "one"},
		{" \" two  \"", " two  "},
		{"\" \\\"three\\\" \"", " \"three\" "},
	}

	for i := range numberTests {
		t.Run(numberTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(numberTests[i].in)
			if err != nil {
				t.Error(err)
			}
			if got != numberTests[i].out {
				t.Errorf("got: %v, want: %v", got, numberTests[i].out)
			}
		})
	}
}

func TestParseBoolArray(t *testing.T) {
	var boolArrayTests = []struct {
		in  string
		out []bool
	}{
		{"[]", []bool{}},
		{"  [false,   false, true]", []bool{
			false,
			false,
			true,
		}},
	}

	for i := range boolArrayTests {
		t.Run(boolArrayTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(boolArrayTests[i].in)
			if err != nil {
				t.Error(err)
			}

			interfaceSlice, ok := got.([]interface{})
			if !ok {
				t.Error("Cannot convert to slice")
			}

			boolSlice := make([]bool, len(interfaceSlice))

			for i, v := range interfaceSlice {
				boolValue, ok := v.(bool)
				if !ok {
					t.Errorf("element at index %d is not a bool", i)
				}
				boolSlice[i] = boolValue
			}

			if !reflect.DeepEqual(boolSlice, boolArrayTests[i].out) {
				t.Errorf("got: %v, want: %v", got, boolArrayTests[i].out)
			}
		})
	}
}

func TestParseNumberArray(t *testing.T) {
	var numberArrayTests = []struct {
		in  string
		out []int
	}{
		{"[1, 2]", []int{
			1,
			2,
		}},
		{"  [1001,   581531, 235]", []int{
			1001,
			581531,
			235,
		}},
	}

	for i := range numberArrayTests {
		t.Run(numberArrayTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(numberArrayTests[i].in)
			if err != nil {
				t.Error(err)
			}

			interfaceSlice, ok := got.([]interface{})
			if !ok {
				t.Error("Cannot convert to slice")
			}

			numberSlice := make([]int, len(interfaceSlice))

			for i, v := range interfaceSlice {
				numberValue, ok := v.(int)
				if !ok {
					t.Errorf("element at index %d is not a bool", i)
				}
				numberSlice[i] = numberValue
			}

			if !reflect.DeepEqual(numberSlice, numberArrayTests[i].out) {
				t.Errorf("got: %v, want: %v", got, numberArrayTests[i].out)
			}
		})
	}
}

func TestParseStringArray(t *testing.T) {
	var stringArrayTests = []struct {
		in  string
		out []string
	}{
		{"[\"   hello\", \"world\"]", []string{
			"   hello",
			"world",
		}},
		{"  [\"\\\"three\\\"\", \"235\"]", []string{
			"\"three\"",
			"235",
		}},
	}

	for i := range stringArrayTests {
		t.Run(stringArrayTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(stringArrayTests[i].in)
			if err != nil {
				t.Error(err)
			}

			interfaceSlice, ok := got.([]interface{})
			if !ok {
				t.Error("Cannot convert to slice")
			}

			stringSlice := make([]string, len(interfaceSlice))

			for i, v := range interfaceSlice {
				stringValue, ok := v.(string)
				if !ok {
					t.Errorf("element at index %d is not a bool", i)
				}
				stringSlice[i] = stringValue
			}

			if !reflect.DeepEqual(stringSlice, stringArrayTests[i].out) {
				t.Errorf("got: %v, want: %v", got, stringArrayTests[i].out)
			}
		})
	}
}

func TestParseDynamicArray(t *testing.T) {
	var dynamicArrayTests = []struct {
		in  string
		out []interface{}
	}{
		{"[123, null, \"hello\", false]", []interface{}{
			123,
			nil,
			"hello",
			false,
		}},
	}

	for i := range dynamicArrayTests {
		t.Run(dynamicArrayTests[i].in, func(t *testing.T) {
			got, err := jsonparser.Parse(dynamicArrayTests[i].in)
			if err != nil {
				t.Error(err)
			}

			interfaceSlice, ok := got.([]interface{})
			if !ok {
				t.Error("Cannot convert to slice")
			}

			if !reflect.DeepEqual(interfaceSlice, dynamicArrayTests[i].out) {
				t.Errorf("got: %v, want: %v", got, dynamicArrayTests[i].out)
			}
		})
	}
}

func TestParseNestedArray(t *testing.T) {

	testInput := "[[123, \"hello world\"], null, true]"

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		interfaceSlice, ok := got.([]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		nestedSlice, ok := interfaceSlice[0].([]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}
		wantSlice := []interface{}{
			123,
			"hello world",
		}

		if !reflect.DeepEqual(nestedSlice, wantSlice) {
			t.Errorf("got: %v, want: %v", got, wantSlice)
		}

		if interfaceSlice[1] != nil {
			t.Errorf("got: %v, want: %v", got, nil)
		}

		if interfaceSlice[2] != true {
			t.Errorf("got: %v, want: %v", got, true)
		}
	})
}

func TestParseEmptyObject(t *testing.T) {

	testInput := "{}"

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		mapGot, ok := got.(map[string]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		if len(mapGot) != 0 {
			t.Errorf("got: %v, want: %v", got, mapGot)
		}
	})
}

func TestParseObject(t *testing.T) {

	testInput := `
	{
		"a": 123,
		"b": "hello",
		"c": null,
		"d": true
	}
	`

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		mapGot, ok := got.(map[string]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		numberVal := mapGot["a"].(int)
		if numberVal != 123 {
			t.Errorf("got: %v, want: %v", got, 123)
		}

		stringVal := mapGot["b"].(string)
		if stringVal != "hello" {
			t.Errorf("got: %v, want: %v", got, "hello")
		}

		if mapGot["c"] != nil {
			t.Errorf("got: %v, want: %v", got, nil)
		}

		if mapGot["d"] != true {
			t.Errorf("got: %v, want: %v", got, true)
		}
	})
}

func TestParseNestedObject(t *testing.T) {

	testInput := `
	{
		"a": true,
		"b": {
			"c": 123,
			"d": null
		}
	}
	`

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		mapGot, ok := got.(map[string]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		if mapGot["a"] != true {
			t.Errorf("got: %v, want: %v", got, true)
		}

		nestedMap := mapGot["b"].(map[string]interface{})

		numberVal := nestedMap["c"].(int)
		if numberVal != 123 {
			t.Errorf("got: %v, want: %v", got, 123)
		}

		if nestedMap["d"] != nil {
			t.Errorf("got: %v, want: %v", got, nil)
		}
	})
}

func TestParseNestedArrayInObject(t *testing.T) {

	testInput := `
	{
		"a": 123,
		"b": "hello",
		"c": null,
		"d": [123, "hello world"]
	}
	`

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		mapGot, ok := got.(map[string]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		numberVal := mapGot["a"].(int)
		if numberVal != 123 {
			t.Errorf("got: %v, want: %v", got, 123)
		}

		stringVal := mapGot["b"].(string)
		if stringVal != "hello" {
			t.Errorf("got: %v, want: %v", got, "hello")
		}

		if mapGot["c"] != nil {
			t.Errorf("got: %v, want: %v", got, nil)
		}

		nestedSlice, ok := mapGot["d"].([]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}
		wantSlice := []interface{}{
			123,
			"hello world",
		}

		if !reflect.DeepEqual(nestedSlice, wantSlice) {
			t.Errorf("got: %v, want: %v", got, wantSlice)
		}
	})
}

func TestParseComplexObject(t *testing.T) {

	testInput := `
	{
		"a": null,
		"b": {
			"b": [
				{
					"b": true,
					"c": 360,
					"d": "no scope"
				}
			],
			"c": 360,
			"d": "no scope"
		}
	}
	`

	t.Run(testInput, func(t *testing.T) {
		got, err := jsonparser.Parse(testInput)
		if err != nil {
			t.Error(err)
		}

		mapGot, ok := got.(map[string]interface{})
		if !ok {
			t.Error("Cannot convert to slice")
		}

		t.Log(mapGot)

		if mapGot["a"] != nil {
			t.Errorf("got: %v, want: %v", got, nil)
		}

		nestedMap := mapGot["b"].(map[string]interface{})
		nestedSlice := nestedMap["b"].([]interface{})

		innerMap := nestedSlice[0].(map[string]interface{})

		if innerMap["b"] != true {
			t.Errorf("got: %v, want: %v", got, true)
		}

		numberVal := innerMap["c"].(int)
		if numberVal != 360 {
			t.Errorf("got: %v, want: %v", got, 360)
		}

		stringVal := innerMap["d"].(string)
		if stringVal != "no scope" {
			t.Errorf("got: %v, want: %v", got, "no scope")
		}

	})
}
