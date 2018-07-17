package bson

/*
#include <bson.h>
*/
import "C"

// Context ..
type Context C.bson_context_t

// ContextFlag ..
type ContextFlag C.bson_context_flags_t

// ContextNone ..
const ContextNone = C.BSON_CONTEXT_NONE

// ContextThreadSafe ..
const ContextThreadSafe = C.BSON_CONTEXT_THREAD_SAFE

// ContextDisableHostCache ..
const ContextDisableHostCache = C.BSON_CONTEXT_DISABLE_HOST_CACHE

// ContextDisablePidCache ..
const ContextDisablePidCache = C.BSON_CONTEXT_DISABLE_PID_CACHE

// NewContext ..
func NewContext(ctx C.bson_context_flags_t) *C.bson_context_t {
	return C.bson_context_new(ctx)
}

// NewDefaultContext ..
func NewDefaultContext() *C.bson_context_t {
	return C.bson_context_get_default()
}
