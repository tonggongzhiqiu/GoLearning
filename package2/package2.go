package package2

import (
	"GoLearning/util"
	"fmt"
)

var Value1 = util.TraceLog("init package2 value1", 20)
var Value2 = util.TraceLog("init package2 value2", 30)

func init() {
	fmt.Println("init() in package2")
}
