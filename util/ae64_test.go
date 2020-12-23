package util

import (
	"fmt"
	"testing"
)

func TestAe64Encode(t *testing.T) {
	src := []byte("Qaa1jLayYg86BgxLDJrpqM9a7yhNUBvNkR")
	src2 := []byte("12345678")
	key := []byte("12345678")
	dst, _ := Ae64Encode(src, key)
	dst2, _ := Ae64Encode(src2, key)
	fmt.Println(dst)
	fmt.Println(dst2)
}

func TestAe64Decode(t *testing.T) {
	src := "Phgn95Se61hGId5ad70dNt9HFUhpZnFrzjMZLpPefcJq2MOPymyWRkXGsoiqQodt"
	src2 := "m0lxCSrfYVhmOhZcOhICrw=="
	key := []byte("12345678")
	orgi, _ := Ae64Decode(src, key)
	orgi2, _ := Ae64Decode(src2, key)
	fmt.Println(string(orgi))
	fmt.Println(string(orgi2))
}
