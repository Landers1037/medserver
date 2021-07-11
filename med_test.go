package main

import "testing"

func TestInitCMD(t *testing.T) {
	t.Log("start create empty cmd")
	e := InitCMD()
	if e != nil {
		t.Error(e)
	}
}
