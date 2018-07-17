package bson

/*
#cgo CFLAGS: -I${SRCDIR}/../driver/include/libbson-1.0 -I../driver/include/libmongoc-1.0
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../driver/Darwin-lib
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/../driver/Linux-lib
#cgo LDFLAGS: -lbson-static-1.0
*/
import "C"
