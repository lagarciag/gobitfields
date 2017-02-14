package gobitfields

import (
	"github.com/lagarciag/bitwisebytes"
)

func GetField(inputSlice []byte,offset int, width int) (outputSlice []byte) {
	bytesWidth := width / 8
	bytesWidthModulus := width % 8
	if bytesWidthModulus != 0 {
		bytesWidth++
	}
	tmpField , _ := bitwisebytes.ShiftRight(inputSlice,uint(offset))
	outputSlice = tmpField[0:bytesWidth]
	maskSlice := bitwisebytes.MakeMask(uint(len(outputSlice)),uint(width),0)

	err := bitwisebytes.And(outputSlice,maskSlice)
	if err != nil {
		panic(err.Error())
	}
	return outputSlice
}


