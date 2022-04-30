package config

import (
	"reflect"
	"testing"
)

func TestProfile(t *testing.T) {
	c := New("")

	t.Run("GetSection", func(t *testing.T) {
		name := "test1"
		c.GetSection(name)
		if _, ok := c.HasSection(name); !ok {
			t.Errorf("want %v, got %v", name, "NOTHING")
		}
	})

	t.Run("WriteSection", func(t *testing.T) {
		name := "test2"
		sm := c.GetSection(name)
		sm["key"] = "value"
		c.WriteSection(name, sm)
		sm1 := c.GetSection(name)
		if !reflect.DeepEqual(sm, sm1) {
			t.Errorf("want %v, got %v", sm, sm1)
		}
	})

	t.Run("GetRootValue", func(t *testing.T) {
		name := ""
		key := "Version"
		v1 := "1.0.0"
		v2 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("want %v, got %v", v1, v2)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		name := "test3"
		key := "key"
		v1 := 12345
		v2 := c.GetValue(name, key, v1)
		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("want %v, got %v", v1, v2)
		}
		v3 := c.GetValue("Test3", "Key", 12346)
		if !reflect.DeepEqual(v1, v3) {
			t.Errorf("want %v, got %v", v1, v3)
		}
	})

	t.Run("WriteValue", func(t *testing.T) {
		name := "test4"
		key := "key"
		v1 := 12345
		c.WriteValue(name, key, v1)
		v2 := c.GetValue("TESt4", "kEY", 12346)
		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("want %v, got %v", v1, v2)
		}
	})

	t.Run("HasSection", func(t *testing.T) {
		name := "test5"
		key := "key"
		if _, ok := c.HasSection(name); ok {
			t.Error("Should not has section")
		}
		v1 := 12345
		c.WriteValue(name, key, v1)
		if _, ok := c.HasSection("test5"); !ok {
			t.Error("Should has section")
		}
	})

	t.Run("HasKey", func(t *testing.T) {
		name := "test6"
		key := "key"
		if _, ok := c.HasKey(name, key); ok {
			t.Error("Should not has key")
		}
		v1 := 12345
		c.WriteValue(name, key, v1)
		if _, ok := c.HasKey("Test6", "Key"); !ok {
			t.Error("Should has key")
		}
	})
}
