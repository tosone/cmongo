package bson

/*
#include <bson.h>
*/
import "C"
import (
	"encoding/hex"
	"unsafe"
)

// MD5 ..
type MD5 struct {
	md5 *C.bson_md5_t
}

// MD5New ..
func MD5New() (m *MD5) {
	var cm = new(C.bson_md5_t)
	C.bson_md5_init(cm)
	m = new(MD5)
	m.md5 = cm
	return
}

// Append ..
func (m *MD5) Append(data []byte) {
	C.bson_md5_append(m.md5, (*C.uint8_t)(C.CBytes(data)), C.uint32_t(len(data)))
}

// Finish ..
func (m *MD5) Finish() string {
	var result = new(C.uint8_t)
	C.bson_md5_finish(m.md5, result)
	return hex.EncodeToString(C.GoBytes(unsafe.Pointer(result), 16))
}
