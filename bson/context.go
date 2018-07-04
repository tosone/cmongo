package bson

/*
#include <bson.h>
*/
import "C"

type Context C.bson_context_t
type ContextFlag C.bson_context_flags_t

const ContextNone = C.BSON_CONTEXT_NONE
const ContextThreadSafe = C.BSON_CONTEXT_THREAD_SAFE
const ContextDisableHostCache = C.BSON_CONTEXT_DISABLE_HOST_CACHE
const ContextDisablePidCache = C.BSON_CONTEXT_DISABLE_PID_CACHE

func NewContext(ctx Context) Context {
	return C.bson_context_new(ctx)
}

func NewDefaultContext() Context {
	return C.bson_context_get_default()
}

