package hghstring

import (
	"fmt"
	"time"
)

func ExampleParse() {
	timeDuration := (364 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := Parse(timeDuration).String()
	fmt.Println(duration)
}
