package main

/*
#cgo CFLAGS: -I./dist/include
#cgo LDFLAGS: -L./dist/lib -lleveldb

#include "leveldb/c.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

type Options struct {
	options *C.leveldb_options_t
}

type CompressionType uint

const (
	NO_COMPRESSION     CompressionType = C.leveldb_no_compression
	SNAPPY_COMPRESSION CompressionType = C.leveldb_snappy_compression
)

type OptionsBuilder struct {
	createIfMissing      bool
	errorIfExists        bool
	paranoidChecks       bool
	writeBufferSize      uint64
	maxOpenFiles         int
	blockSize            uint64
	blockRestartInterval int
	maxFileSize          int
	compression          CompressionType
}

func NewOptionsBuilder() *OptionsBuilder {
	return &OptionsBuilder{}
}

func (b *OptionsBuilder) SetCreateIfMissing(value bool) *OptionsBuilder {
	b.createIfMissing = value
	return b
}

func (b *OptionsBuilder) SetErrorIfExists(value bool) *OptionsBuilder {
	b.errorIfExists = value
	return b
}

func (b *OptionsBuilder) SetParanoidChecks(value bool) *OptionsBuilder {
	b.paranoidChecks = value
	return b
}

func (b *OptionsBuilder) build() *Options {
	o := &Options{}
	o.options = C.leveldb_options_create()
	if b.createIfMissing {
		C.leveldb_options_set_create_if_missing(o.options, 1)
	} else {
		C.leveldb_options_set_create_if_missing(o.options, 0)
	}

	if b.errorIfExists {
		C.leveldb_options_set_error_if_exists(o.options, 1)
	} else {
		C.leveldb_options_set_error_if_exists(o.options, 0)
	}

	if b.paranoidChecks {
		C.leveldb_options_set_paranoid_checks(o.options, 1)
	} else {
		C.leveldb_options_set_paranoid_checks(o.options, 0)
	}

	C.leveldb_options_set_write_buffer_size(o.options, C.ulong(b.writeBufferSize))
	C.leveldb_options_set_max_open_files(o.options, C.int(b.maxOpenFiles))
	C.leveldb_options_set_block_size(o.options, C.ulong(b.blockSize))
	C.leveldb_options_set_block_restart_interval(o.options, C.int(b.blockRestartInterval))
	C.leveldb_options_set_max_file_size(o.options, C.ulong(b.maxFileSize))
	C.leveldb_options_set_compression(o.options, C.int(b.compression))

	return o
}

type DB struct {
	db *C.leveldb_t
}

func OpenDB(options *Options, name string) (*DB, error) {
	db := &DB{}

	errptr := C.CString("")
	db.db = C.leveldb_open(options.options, C.CString(name), &errptr)

	goErrPtr := C.GoString(errptr)
	if len(goErrPtr) != 0 {
		return nil, errors.New(goErrPtr)
	}
	return db, nil
}

func (db *DB) Close() {
	C.leveldb_close(db.db)
}

func (db *DB) Put(key, value string) {

}

func main() {
	builder := NewOptionsBuilder().SetCreateIfMissing(true).SetErrorIfExists(false)
	options := builder.build()

	db, err := OpenDB(options, "test")
	defer db.Close()
	fmt.Printf("db: %+v, err: %+v\n", db, err)
}
