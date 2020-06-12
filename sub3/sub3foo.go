package sub3

import (
	"fmt"
	"github.com/thaJeztah/monono/sub1"
	"github.com/thaJeztah/monono/sub2"
)

const SayMyName = "sub3"

func SayHello() {
	fmt.Printf("Hello %s\n", SayMyName)
	sub1.SayHello()
	sub2.SayHello()
}

