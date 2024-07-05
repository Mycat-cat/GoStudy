package package2

import (
	"GoStudy/GoPrincipleAnalysis/ProgramInitialization/utils"
	"fmt"
)

var Value1 = utils.TraceLog("init package1 value1", 20)
var Value2 = utils.TraceLog("init package2 value2", 30)

func init() {
	fmt.Println("init func2 in package2")
}
