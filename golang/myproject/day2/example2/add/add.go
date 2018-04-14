package add

import (
	"fmt"
)

var Name string = "xxoo"
var Age int = 100

func init() {
	Name = "Hello world"
	Age = 10
	fmt.Println("add init.", Name, Age)
}
