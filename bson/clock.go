// +build darwin

package bson

/*
#include <bson.h>
#include <sys/time.h>
*/
import "C"
import "unsafe"

// GetMonotonicTime return the system startup till now microsecond
func GetMonotonicTime() uint64 {
	return uint64(C.bson_get_monotonic_time())
}

// GetTimeOfDay get the time and timezone
func GetTimeOfDay() (tv *C.struct_timeval, tz *C.struct_timezone, err error) {
	tv = new(C.struct_timeval)
	tz = new(C.struct_timezone)
	var n C.int
	if n, err = C.gettimeofday(tv, unsafe.Pointer(tz)); int(n) != 0 {
		return
	}
	return
}
