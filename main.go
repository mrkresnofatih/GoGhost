package main

import (
	"fmt"
	cider "mrkresnofatihdev/goghost/cider"
)

func main() {
	cidr, err := cider.CreateNewCider("0.0.0.0/1")
	if err != nil {
	}
	fmt.Println(cidr.RawString)
	fmt.Println(cidr.Size)
	fmt.Println(cidr.SubnetSize)

}
