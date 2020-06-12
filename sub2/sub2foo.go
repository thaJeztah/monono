package sub2

import (
	"fmt"
	"github.com/thaJeztah/monono/sub1"
)


const SayMyName = "sub2"

func SayHello() {
	fmt.Printf("Hello %s, and welcome %s\n", SayMyName, sub1.SayMyName)
}


