package bson

/*
#include <bson.h>
*/
import "C"

// Context ..
type Context C.bson_context_t

// ContextFlag ..
type ContextFlag C.bson_context_flags_t

var (
	// ContextFlagNone ..
	ContextFlagNone ContextFlag = C.BSON_CONTEXT_NONE

	// ContextFlagThreadSafe ..
	ContextFlagThreadSafe ContextFlag = C.BSON_CONTEXT_THREAD_SAFE

	// ContextFlagDisableHostCache ..
	ContextFlagDisableHostCache ContextFlag = C.BSON_CONTEXT_DISABLE_HOST_CACHE

	// ContextFlagDisablePidCache ..
	ContextFlagDisablePidCache ContextFlag = C.BSON_CONTEXT_DISABLE_PID_CACHE
)

// NewContext ..
func NewContext(ctx ContextFlag) *C.bson_context_t {
	return C.bson_context_new(C.bson_context_flags_t(ctx))
}

// NewDefaultContext ..
func NewDefaultContext() *C.bson_context_t {
	return C.bson_context_get_default()
}
