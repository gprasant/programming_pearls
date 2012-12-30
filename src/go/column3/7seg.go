package main

import (
	"fmt"
	"strings"
)

func InitSegmentMap() map[string]byte {
	segmentMap := make(map[string]byte)
	segmentMap["0"] = 0x7D
	segmentMap["1"] = 0x50
	segmentMap["2"] = 0x37
	segmentMap["3"] = 0x53
	segmentMap["4"] = 0x5A
	segmentMap["5"] = 0x4F
	segmentMap["6"] = 0x6F
	segmentMap["7"] = 0x54
	segmentMap["8"] = 0x7F
	segmentMap["9"] = 0x5F
	return segmentMap
}

func main() {
	var n string
	var output []byte
	seven_segmentMap := InitSegmentMap()
	fmt.Print("Enter the number : $ ")
	fmt.Scanf("%s\n", &n)
	digits := strings.Split(n, "")
	for _, v := range digits {
		output = append(output, seven_segmentMap[v])
	}
	fmt.Println(output)

}
