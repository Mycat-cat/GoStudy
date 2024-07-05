package package1

import (
	"GoStudy/GoPrincipleAnalysis/ProgramInitialization/package2"
	"GoStudy/GoPrincipleAnalysis/ProgramInitialization/utils"
	"fmt"
)

var V1 = utils.TraceLog("init package1 value1", package2.Value1+10)
var V2 = utils.TraceLog("init package1 value2", package2.Value2+10)

func init() {
	fmt.Println("init func1 in package1")
}
