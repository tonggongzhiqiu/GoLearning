package package1

import (
	"GoLearning/package2"
	"GoLearning/util"
	"fmt"
)

var V1 = util.TraceLog("init package1 value1", package2.Value1+10)
var V2 = util.TraceLog("init package2 value2", package2.Value2+10)

func init() {
	fmt.Println("init() in package1")
}
