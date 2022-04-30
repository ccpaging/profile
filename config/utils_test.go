package config

import (
	"encoding/json"
	"testing"
)

func TestFlatSubMapStringInterface(t *testing.T) {
	var subsub map[string]interface{} = map[string]interface{}{
		"c": "sea",
	}
	var sub map[string]interface{} = map[string]interface{}{
		"b":  subsub,
		"b1": "beach1",
	}
	var in map[string]interface{} = map[string]interface{}{
		"a": sub,
		"b": "beach",
	}

	b, _ := json.MarshalIndent(in, "", "\t")
	println(string(b))

	out := flattenMap(in, "")

	b1, _ := json.MarshalIndent(out, "", "\t")
	println(string(b1))

	if want := 53; want != len(b1) {
		t.Fatalf("want:%d got:%d\n%s", want, len(b1), string(b1))
	}
}

func TestFlatSubMapStringInt(t *testing.T) {
	var in map[string]interface{} = map[string]interface{}{
		"a": map[string]int{
			"b": 1,
		},
		"c": "sea",
	}

	out := flattenMap(in, "")

	b, _ := json.MarshalIndent(out, "", "\t")
	if want := 36; want != len(b) {
		t.Fatalf("want:\n%d\ngot:\n%s", want, string(b))
	}
}

func TestFlatSubMapIntString(t *testing.T) {
	var in map[string]interface{} = map[string]interface{}{
		"a": map[int]string{
			1: "b",
		},
		"c": "sea",
	}

	out := flattenMap(in, "")

	b, _ := json.MarshalIndent(out, "", "\t")
	if want := 38; want != len(b) {
		t.Fatalf("want:\n%d\ngot:\n%s", want, string(b))
	}
}
