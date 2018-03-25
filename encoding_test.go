package main

import "testing"

func TestEncoding(t *testing.T) {
	path := "suse/sles/12-SP1/x86_64/yum/base/libstdc++48-devel-4.8.5-24.1.x86_64.rpm"
	encodedPath := encodePath(path)
	desiredPath := "suse/sles/12-SP1/x86_64/yum/base/libstdc%2B%2B48-devel-4.8.5-24.1.x86_64.rpm"
	if encodedPath != desiredPath {
		t.Errorf("Path was incorrect, got: \n%s, wanted \n%s\n", encodedPath, desiredPath)
	}
}
