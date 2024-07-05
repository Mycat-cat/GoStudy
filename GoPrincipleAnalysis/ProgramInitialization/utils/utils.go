package utils

import "fmt"

func TraceLog(t string, v int) int {
	fmt.Printf("TraceLog----%s-------%d\n", t, v)
	return v
}
