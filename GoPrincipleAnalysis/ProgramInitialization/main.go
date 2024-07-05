package main

import (
	"GoStudy/GoPrincipleAnalysis/ProgramInitialization/package1"
	"GoStudy/GoPrincipleAnalysis/ProgramInitialization/utils"
	"fmt"
)

func init() {
	fmt.Println("init func1 in main")
}

func init() {
	fmt.Println("init func2 in main")
}

var MainValue1 = utils.TraceLog("init M_v1", package1.V1+10)
var MainValue2 = utils.TraceLog("init M_v2", package1.V2+10)

func main() {
	fmt.Println("main func in main")
}
