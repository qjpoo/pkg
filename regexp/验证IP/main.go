package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	//fmt.Println(os.Args[0], os.Args[1], os.Args[2])
	fmt.Println("IP: ", os.Args[1])
	if IsIP(os.Args[1]) {
		fmt.Println("ok ...")
	}else {
		fmt.Println("no ...")
	}
}
