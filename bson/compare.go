package bson

/*
#include <bson.h>
*/
import "C"
import (
	"unsafe"
)

// Equal ..
func Equal() {

}

// IsValid ..
func IsValid(oid string) bool {
	var cOid = C.CString(oid)
	defer C.free(unsafe.Pointer(cOid))
	return bool(C.bson_oid_is_valid(cOid, C.size_t(len(oid))))
}
