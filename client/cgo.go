package bson

/*
#cgo CFLAGS: -I../driver/include/libbson-1.0 -I../driver/include/libmongoc-1.0
#cgo darwin,amd64 LDFLAGS: -L../driver/Darwin-lib
#cgo linux,amd64 LDFLAGS: -L../driver/Linux-lib
#cgo LDFLAGS: -lmongoc-static-1.0 -lbson-static-1.0
*/
import "C"
