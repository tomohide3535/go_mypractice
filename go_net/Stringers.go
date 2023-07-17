package main

import "fmt"

type IPaddr [4]byte


func (p IPaddr) replace_stiring() string{
	return fmt.Sprintf("%v.%v.%v.%v", p[0], p[1], p[2], p[3])
}


func main(){
	hosts := map[string]IPaddr{
		"loopback":{127, 0, 0, 0},
		"googleDNS":{8, 8, 8, 8},
	}
	for name, ip := range hosts{
		fmt.Printf("%v: %v\n", name, ip.replace_stiring())
	}
}