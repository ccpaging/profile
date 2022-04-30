package config

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var testReplacer = strings.NewReplacer(".", DELIMITER)

func TestOverrieProfile(t *testing.T) {
	c := New("")

	t.Run("arguments", func(t *testing.T) {
		name := "test1"
		key := "key"
		v1 := 12345
		v2 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("want %v, got %v", v1, v2)
		}

		want := "123456"
		c.SetArg(map[string]string{
			"Test1.Key": want,
		}, testReplacer)

		v3 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(want, v3) {
			t.Errorf("want %v, got %v", want, v3)
		}
	})

	t.Run("environment", func(t *testing.T) {
		name := "test2"
		key := "key"
		v1 := 12345
		v2 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("want %v, got %v", v1, v2)
		}

		want := "123456"
		os.Setenv("CFG_Test2.Key", want)
		c.SetEnv([]string{"CFG_"}, testReplacer)

		v3 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(want, v3) {
			t.Errorf("want %v, got %v", want, v3)
		}
	})
}
