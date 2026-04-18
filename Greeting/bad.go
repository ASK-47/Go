package greeting

import (
	"fmt"
	"time"
)

func Fuck() {
	for i := 0; i < Int(); i++ {
		fmt.Println("Fuck!!!")
		time.Sleep(100 * time.Millisecond)
		fmt.Println(Int())
	}
}
