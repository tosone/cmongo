package bson

/*
#include <bson.h>
*/
import "C"

// Bson The bson_t structure represents a BSON document.
// This structure manages the underlying BSON encoded buffer.
// For mutable documents, it can append new data to the document.
// http://mongoc.org/libbson/current/bson_t.html
type Bson struct {
	bson *C.bson_t
}

// New The bson_new() function shall create a new bson_t structure on the heap.
// It should be freed with bson_destroy() when it is no longer in use.
// http://mongoc.org/libbson/current/bson_new.html
func New() (b *Bson) {
	b = new(Bson)
	b.bson = C.bson_new()
	return
}

// AppendInt32 The bson_append_int32() function shall append a new element to bson containing a 32-bit signed integer.
// http://mongoc.org/libbson/current/bson_append_int32.html
func (b *Bson) AppendInt32(key string, val int32) bool {
	return bool(C.bson_append_int32(b.bson, C.CString(key), C.int(len(key)), C.int32_t(val)))
}

// AppendInt64 The bson_append_int64() function shall append a new element to bson containing a 64-bit signed integer.
// http://mongoc.org/libbson/current/bson_append_int64.html
func (b *Bson) AppendInt64(key string, val int64) bool {
	return bool(C.bson_append_int64(b.bson, C.CString(key), C.int(len(key)), C.int64_t(val)))
}

// Destroy The bson_destroy() function shall free an allocated bson_t structure. Does nothing if bson is NULL.
// This function should always be called when you are done with a bson_t unless otherwise specified.
// http://mongoc.org/libbson/current/bson_destroy.html
func (b *Bson) Destroy() {
	C.bson_destroy(b.bson)
}

// AppendMaxKey The bson_append_maxkey() function shall append an element of type BSON_TYPE_MAXKEY to a bson document.
// This is primarily used in queries and unlikely to be used when storing a document to MongoDB.
// http://mongoc.org/libbson/current/bson_append_maxkey.html
func (b *Bson) AppendMaxKey(key string) bool {
	return bool(C.bson_append_maxkey(b.bson, C.CString(key), C.int(len(key))))
}

// AppendMinKey The bson_append_minkey() function shall append an element of type BSON_TYPE_MINKEY to a bson document.
// This is primarily used in queries and unlikely to be used when storing a document to MongoDB.
// http://mongoc.org/libbson/current/bson_append_minkey.html
func (b *Bson) AppendMinKey(key string) bool {
	return bool(C.bson_append_minkey(b.bson, C.CString(key), C.int(len(key))))
}

// AppendNowUTC The bson_append_now_utc() function is a helper to get the current date
// and time in UTC and append it to bson as a BSON_TYPE_DATE_TIME element.
// This function calls bson_append_date_time() internally.
// http://mongoc.org/libbson/current/bson_append_now_utc.html
func (b *Bson) AppendNowUTC(key string) bool {
	return bool(C.bson_append_now_utc(b.bson, C.CString(key), C.int(len(key))))
}

// AppendDateTime The bson_append_date_time() function shall append a new element to a bson document
// containing a date and time with no timezone information.
// value is assumed to be in UTC format of milliseconds since the UNIX epoch. value MAY be negative.
// http://mongoc.org/libbson/current/bson_append_date_time.html
func (b *Bson) AppendDateTime(key string, timestamp int64) bool {
	return bool(C.bson_append_date_time(b.bson, C.CString(key), C.int(len(key)), C.int64_t(timestamp)))
}
