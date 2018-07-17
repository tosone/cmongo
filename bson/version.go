package bson

/*
#include <bson.h>
*/
import "C"

// GetVersion version
func GetVersion() string {
	return C.GoString(C.bson_get_version())
}

// GetMicroVersion micro version
func GetMicroVersion() int {
	return int(C.bson_get_micro_version())
}

// GetMinorVersion Minor version
func GetMinorVersion() int {
	return int(C.bson_get_minor_version())
}

// GetMajorVersion major version
func GetMajorVersion() int {
	return int(C.bson_get_major_version())
}
