package bson

import (
	"fmt"
	"testing"
)

func TestMD5_MD5Finish(t *testing.T) {
	var m = MD5New()
	m.Append([]byte("happy"))
	var result = m.Finish()
	if result != "56ab24c15b72a457069c5ea42fcfc640" {
		t.Errorf("'happy' md5 result calc error")
	}
	fmt.Printf("'happy' md5 result: %s\n", result)
}
