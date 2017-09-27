package leveldb

// #cgo LDFLAGS: ${SRCDIR}/leveldb/out-static/libleveldb.a -lstdc++
// #include "leveldb/include/leveldb/c.h"
import "C"

func emptyCString() *C.char {
	var str *C.char = nil
	return str
}

func bool2Uchar(value bool) C.uchar {
	if value {
		return C.uchar(1)
	} else {
		return C.uchar(0)
	}
}
