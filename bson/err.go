package bson

/*
#include <bson.h>

void bson_set_error_customer(bson_error_t *error, uint32_t domain, uint32_t code, const char *msg) {
  bson_set_error(error, domain, code, "%s", msg);
}
*/
import "C"

// Error The bson_error_t structure is used as an out-parameter to pass error information to the caller.
// It should be stack-allocated and does not requiring freeing.
// http://mongoc.org/libbson/current/bson_error_t.html
type Error C.bson_error_t

// SetError This is a helper function to set the parameters of a bson_error_t.
// It handles the case where error is NULL by doing nothing.
// http://mongoc.org/libbson/current/bson_set_error.html
func (e *Error) SetError(domain, code uint32, str string) {
	C.bson_set_error_customer((*C.bson_error_t)(e), C.uint32_t(domain), C.uint32_t(code), C.CString(str))
}

// StrError This is a portability wrapper around strerror().
// http://mongoc.org/libbson/current/bson_strerror_r.html
func StrError(code int, err string) string {
	return C.GoString(C.bson_strerror_r(C.int(code), C.CString(err), C.ulong(len(err))))
}
