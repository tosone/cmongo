package bson

/*
#include <bson.h>
*/
import "C"

// Bson The bson_t structure represents a BSON document.
// This structure manages the underlying BSON encoded buffer.
// For mutable documents, it can append new data to the document.
// http://mongoc.org/libbson/current/bson_t.html
type Bson C.bson_t

// New The bson_new() function shall create a new bson_t structure on the heap.
// It should be freed with bson_destroy() when it is no longer in use.
// http://mongoc.org/libbson/current/bson_new.html
func New() (b *Bson) {
	b = (*Bson)(C.bson_new())
	return
}

// AppendInt32 The bson_append_int32() function shall append a new element to bson containing a 32-bit signed integer.
// http://mongoc.org/libbson/current/bson_append_int32.html
func (b *Bson) AppendInt32(key string, val int32) bool {
	return bool(C.bson_append_int32(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.int32_t(val),
	))
}

// AppendInt64 The bson_append_int64() function shall append a new element to bson containing a 64-bit signed integer.
// http://mongoc.org/libbson/current/bson_append_int64.html
func (b *Bson) AppendInt64(key string, val int64) bool {
	return bool(C.bson_append_int64(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.int64_t(val),
	))
}

// Destroy The bson_destroy() function shall free an allocated bson_t structure. Does nothing if bson is NULL.
// This function should always be called when you are done with a bson_t unless otherwise specified.
// http://mongoc.org/libbson/current/bson_destroy.html
func (b *Bson) Destroy() {
	C.bson_destroy((*C.bson_t)(b))
}

// AppendMaxKey The bson_append_maxkey() function shall append an element of type BSON_TYPE_MAXKEY to a bson document.
// This is primarily used in queries and unlikely to be used when storing a document to MongoDB.
// http://mongoc.org/libbson/current/bson_append_maxkey.html
func (b *Bson) AppendMaxKey(key string) bool {
	return bool(C.bson_append_maxkey(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
	))
}

// AppendMinKey The bson_append_minkey() function shall append an element of type BSON_TYPE_MINKEY to a bson document.
// This is primarily used in queries and unlikely to be used when storing a document to MongoDB.
// http://mongoc.org/libbson/current/bson_append_minkey.html
func (b *Bson) AppendMinKey(key string) bool {
	return bool(C.bson_append_minkey(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
	))
}

// AppendNowUTC The bson_append_now_utc() function is a helper to get the current date
// and time in UTC and append it to bson as a BSON_TYPE_DATE_TIME element.
// This function calls bson_append_date_time() internally.
// http://mongoc.org/libbson/current/bson_append_now_utc.html
func (b *Bson) AppendNowUTC(key string) bool {
	return bool(C.bson_append_now_utc(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
	))
}

// AppendDateTime The bson_append_date_time() function shall append a new element to a bson document
// containing a date and time with no timezone information.
// value is assumed to be in UTC format of milliseconds since the UNIX epoch. value MAY be negative.
// http://mongoc.org/libbson/current/bson_append_date_time.html
func (b *Bson) AppendDateTime(key string, timestamp int64) bool {
	return bool(C.bson_append_date_time(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.int64_t(timestamp),
	))
}

// AppendDocument The bson_append_document() function shall append child to bson using the specified key.
// The type of the field will be a document.
// http://mongoc.org/libbson/current/bson_append_document.html
func (b *Bson) AppendDocument(key string, doc *Bson) bool {
	return bool(C.bson_append_document(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		(*C.bson_t)(doc),
	))
}

// AppendOid The bson_append_oid() function shall append a new element to bson of type BSON_TYPE_OID.
// oid MUST be a pointer to a bson_oid_t.
// http://mongoc.org/libbson/current/bson_append_oid.html
func (b *Bson) AppendOid(key string, oid *Oid) bool {
	return bool(C.bson_append_oid(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		(*C.bson_oid_t)(oid),
	))
}

// AppendTimestamp This function is not similar in functionality to bson_append_date_time().
// Timestamp elements are different in that they include only second precision and an increment field.
// They are primarily used for intra-MongoDB server communication.
// The bson_append_timestamp() function shall append a new element of type BSON_TYPE_TIMESTAMP.
// http://mongoc.org/libbson/current/bson_append_timestamp.html
func (b *Bson) AppendTimestamp(key string, timestamp, increment uint32) bool {
	return bool(C.bson_append_timestamp(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.uint32_t(timestamp),
		C.uint32_t(increment),
	))
}

// AppendRegex Appends a new field to bson of type BSON_TYPE_REGEX.
// regex should be the regex string.
// options should contain the options for the regex.
// Valid characters for options include:
// 'i' for case-insensitive.
// 'm' for multiple matching.
// 'x' for verbose mode.
// 'l' to make w and W locale dependent.
// 's' for dotall mode (‘.’ matches everything)
// 'u' to make w and W match unicode.
// http://mongoc.org/libbson/current/bson_append_regex.html
func (b *Bson) AppendRegex(key, regex, options string) bool {
	return bool(C.bson_append_regex(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.CString(regex),
		C.CString(options),
	))
}

// AppendRegexWLen Appends a new field to bson of type BSON_TYPE_REGEX.
// regex should be the regex string.
// options should contain the options for the regex.
// Valid characters for options include:
// 'i' for case-insensitive.
// 'm' for multiple matching.
// 'x' for verbose mode.
// 'l' to make w and W locale dependent.
// 's' for dotall mode (‘.’ matches everything)
// 'u' to make w and W match unicode.
// http://mongoc.org/libbson/current/bson_append_regex.html
func (b *Bson) AppendRegexWLen(key, regex, options string) bool {
	return bool(C.bson_append_regex_w_len(
		(*C.bson_t)(b),
		C.CString(key),
		C.int(len(key)),
		C.CString(regex),
		C.int(len(regex)),
		C.CString(options),
	))
}
