package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := []byte(`
		call hello_server alice
		hello_server bob
		call hello_server eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		for _, v := range s {
			fmt.Println(string(v))
		}
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
