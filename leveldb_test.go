package leveldb

import (
	"testing"
)

func TestLeveldb(t *testing.T) {
	opt := NewOption(true)
	defer opt.Close()
	db, err := NewLevelDB(opt, "test")
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}
	key, value := "key", "value"
	err = db.Put(key, value, true)
	if err != nil {
		t.Fatal(err)
	}
	v, err := db.Get(key, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if v != value {
		t.Fatal("value is not matching")
	}
	err = db.Delete(key, true)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkLeveldbPut(b *testing.B) {
	opt := NewOption(true)
	defer opt.Close()
	db, err := NewLevelDB(opt, "test")
	defer db.Close()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		key, value := "key", "value"
		err = db.Put(key, value, true)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLeveldbGet(b *testing.B) {
	opt := NewOption(true)
	defer opt.Close()
	db, err := NewLevelDB(opt, "test")
	defer db.Close()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		key := "key"
		_, err := db.Get(key, true, false)
		if err != nil {
			b.Fatal(err)
		}
	}
}
