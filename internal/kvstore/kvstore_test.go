package kvstore

import "testing"

func TestUnitNoSuchKey(t *testing.T) {
	kvs := NewMemoryKeyValueStore()
	value, err := kvs.Get("nonexistent")
	if err == nil {
		t.Fatal("expected an error here")
	}
	if value != nil {
		t.Fatal("expected empty string here")
	}
}

func TestUnitExistingKey(t *testing.T) {
	kvs := NewMemoryKeyValueStore()
	if err := kvs.Set("antani", []byte("mascetti")); err != nil {
		t.Fatal(err)
	}
	value, err := kvs.Get("antani")
	if err != nil {
		t.Fatal(err)
	}
	if string(value) != "mascetti" {
		t.Fatal("not the result we expected")
	}
}
