package leveldb

// #cgo LDFLAGS: ${SRCDIR}/leveldb/out-static/libleveldb.a -lstdc++
// #include "leveldb/include/leveldb/c.h"
import "C"
import (
	"errors"
)

type LevelDB struct {
	internal *C.leveldb_t
	option   *Option
}

func NewLevelDB(opt *Option, name string) (*LevelDB, error) {
	db := new(LevelDB)
	err := emptyCString()
	db.internal = C.leveldb_open(opt.internal, C.CString(name), &err)
	db.option = opt
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
	C.leveldb_writeoptions_set_sync(opt, bool2Uchar(sync))
	err := emptyCString()
	Ckey := C.CString(key)
	Cvalue := C.CString(value)
	C.leveldb_put(db.internal, opt, Ckey, C.size_t(len(key)), Cvalue, C.size_t(len(value)), &err)
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

func (db *LevelDB) Get(key string, verifyChecksums, fillCache bool) (string, error) {
	var opt *C.leveldb_readoptions_t = C.leveldb_readoptions_create()
	defer C.leveldb_readoptions_destroy(opt)
	C.leveldb_readoptions_set_verify_checksums(opt, bool2Uchar(verifyChecksums))
	if fillCache {
		C.leveldb_readoptions_set_fill_cache(opt, 0)
	} else {
		C.leveldb_readoptions_set_fill_cache(opt, 1)
	}
	var str_len C.size_t
	err := emptyCString()
	var Cvalue = C.leveldb_get(db.internal, opt, C.CString(key), C.size_t(len(key)), &str_len, &err)
	if err != nil {
		return "", errors.New(C.GoString(err))
	}
	return C.GoStringN(Cvalue, C.int(str_len)), nil
}

func (db *LevelDB) Delete(key string, sync bool) error {
	var opt = C.leveldb_writeoptions_create()
	defer C.leveldb_writeoptions_destroy(opt)
	C.leveldb_writeoptions_set_sync(opt, bool2Uchar(sync))
	err := emptyCString()
	C.leveldb_delete(db.internal, opt, C.CString(key), C.size_t(len(key)), &err)
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}
