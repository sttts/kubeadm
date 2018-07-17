package main

import (
	"fmt"

	kubescheduler "github.com/sttts/kube-scheduler/config"
)

func main() {
	x := &kubescheduler.Foo{}
	x.Kind = "Foo"
	fmt.Println(x)
}
