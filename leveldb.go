package main

// #cgo LDFLAGS: ${SRCDIR}/out-static/libleveldb.a -lstdc++
// #include "include/leveldb/c.h"
import "C"
import (
    "errors"
    "log"
)

func EmptyCString() *C.char {
    var str *C.char = nil
    return str
}

type Option struct {
    internal *C.leveldb_options_t
}

func NewOption() *Option {
    opt := new(Option)
    opt.internal = C.leveldb_options_create()
    return opt
}

func (opt *Option) Close() {
    C.leveldb_options_destroy(opt.internal)
}

func (opt *Option) SetCreateIfMissing(v bool) {
    if v {
        C.leveldb_options_set_create_if_missing(opt.internal, 0)
    } else {
        C.leveldb_options_set_create_if_missing(opt.internal, 1)
    }

}

type LevelDB struct {
    internal *C.leveldb_t
}

func NewLevelDB(opt *Option, name string) (*LevelDB, error) {
    db := new(LevelDB)
    err := EmptyCString()
    db.internal = C.leveldb_open(opt.internal, C.CString(name), &err)
    if err != nil {
        return nil, errors.New(C.GoString(err))
    }
    return db, nil
}

func (db *LevelDB) Close() {
    C.leveldb_close(db.internal)
}

func (db *LevelDB) Put(key, value string, sync bool) error {
    var opt = C.leveldb_writeoptions_create()
    defer C.leveldb_writeoptions_destroy(opt)
    if sync {
        C.leveldb_writeoptions_set_sync(opt, 0)
    } else {
        C.leveldb_writeoptions_set_sync(opt, 1)
    }
    err := EmptyCString()
    Ckey := C.CString(key)
    Cvalue := C.CString(value)
    C.leveldb_put(db.internal, opt, Ckey, C.size_t(len(key)), Cvalue, C.size_t(len(value)), &err)
    if err != nil {
        return errors.New(C.GoString(err))
    }
    return nil
}

func (db *LevelDB) Get(key string, verify_checksums, fill_cache bool) (string, error) {
    var opt *C.leveldb_readoptions_t = C.leveldb_readoptions_create()
    defer C.leveldb_readoptions_destroy(opt)
    if verify_checksums {
        C.leveldb_readoptions_set_verify_checksums(opt, 0)
    } else {
        C.leveldb_readoptions_set_verify_checksums(opt, 1)
    }
    if fill_cache {
        C.leveldb_readoptions_set_fill_cache(opt, 0)
    } else {
        C.leveldb_readoptions_set_fill_cache(opt, 1)
    }
    var str_len C.size_t
    err := EmptyCString()
    var Cvalue = C.leveldb_get(db.internal, opt, C.CString(key), C.size_t(len(key)), &str_len, &err)
    if err != nil {
        return "", errors.New(C.GoString(err))
    }
    return C.GoStringN(Cvalue, C.int(str_len)), nil
}

func main() {
    opt := NewOption()
    defer opt.Close()
    opt.SetCreateIfMissing(true)
    db, err := NewLevelDB(opt, "test")
    defer db.Close()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(db)
    err = db.Put("key", "test", false)
    if err != nil {
        log.Fatal(err)
    }
    v, err := db.Get("key", true, true)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(v)
}

