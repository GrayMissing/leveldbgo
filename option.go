package leveldb

// #cgo LDFLAGS: ${SRCDIR}/leveldb/out-static/libleveldb.a -lstdc++
// #include "leveldb/include/leveldb/c.h"
import "C"

type Option struct {
	internal *C.leveldb_options_t
}

func NewOption(createIfMissing bool) *Option {
	opt := new(Option)
	opt.internal = C.leveldb_options_create()
	opt.SetCreateIfMissing(createIfMissing)
	return opt
}

func (opt *Option) Close() {
	C.leveldb_options_destroy(opt.internal)
}

func (opt *Option) SetCreateIfMissing(value bool) {
	C.leveldb_options_set_create_if_missing(opt.internal, bool2Uchar(value))
}
