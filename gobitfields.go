package gobitfields

import (
	"github.com/lagarciag/bitwisebytes"
)

type MemberMetaData struct {
	Width  uint32
	Offset uint32
}

type StructMetaData struct {
	Width      uint32
	Members    []MemberMetaData
	MembersMap map[string]*MemberMetaData
}

//GetFields returns a bitfield member as specified by the provided offset and width
func GetField(inputSlice []byte,offset int, width int) (outputSlice []byte) {
	bytesWidth := getBytesSize(width)
	tmpField , _ := bitwisebytes.ShiftRight(inputSlice,uint(offset))
	outputSlice = tmpField[0:bytesWidth]
	maskSlice := bitwisebytes.MakeMask(uint(len(outputSlice)),uint(width),0)
	err := bitwisebytes.And(outputSlice,maskSlice)
	if err != nil {
		panic(err.Error())
	}
	return outputSlice
}

//PutField
func PutField(destSlice []byte, inputSlice []byte,offset int, width int) {
	tmpSlice := make([]byte,len(destSlice))

	for i , aByte := range inputSlice {
		tmpSlice[i] = aByte
	}

	tmpSlice2, err := bitwisebytes.ShiftLeft(tmpSlice,uint(offset))
	if err != nil {
		panic(err.Error())
	}

	bitwisebytes.Or(destSlice,tmpSlice2)
}


func PutAllFields(inputMatrix [][]byte,metadataList []MemberMetaData) (outputSlice []byte) {
	size := 0
	for _ , member := range metadataList {
		size += int(member.Width)
	}
	bytesSize := getBytesSize(size)

	outputSlice = make([]byte,bytesSize)

	for i , field := range inputMatrix {
		PutField(outputSlice,field,int(metadataList[i].Offset),int(metadataList[i].Width))
	}

	return outputSlice
}


//GetAllFieldsList returns a slice of all the fields specified in the metadata list
func GetAllFieldsList(inputSlice []byte, metadataList []MemberMetaData) (outputMatrix [][]byte) {
	outputMatrix = make([][]byte, len(metadataList))
	for i , metadata := range metadataList {
		field := GetField(inputSlice,int(metadata.Offset),int(metadata.Width))
		outputMatrix[i] = field
	}
	return outputMatrix
}

func getBytesSize(bitSize int) (bytesSize int) {
	bytesSize = bitSize / 8
	mod := bitSize % 8
	if mod != 0 {
		bytesSize++
	}
	return bytesSize

}

func ReverseMembers(inputBytes []byte,  metadataList []MemberMetaData) (reversed []byte) {

	fieldsMatrix := GetAllFieldsList(inputBytes, metadataList)

	reversedMetadata := ReverseMemberMetadataList(metadataList)

	reversedMatrix := reverseMembersMatrix(fieldsMatrix)
	reversed = PutAllFields(reversedMatrix,reversedMetadata)
	return reversed
}

func reverseMembersMatrix(input [][]byte) (output [][]byte) {
	length := len(input)
	output = make([][]byte, length)
	for i , data := range input {
		output[(length -1) - i] = data
	}
	return output
}

func ReverseBytes(input []byte) (output []byte) {
	length := len(input)
	output = make([]byte, length)
	for i , data := range input {
		output[(length -1) - i] = data
	}
	return output
}

func ReverseMemberMetadataList(input []MemberMetaData) (output []MemberMetaData) {
	length := len(input)
	output = make([]MemberMetaData, length)
	offset := 0
	for i , _ := range output {
		newData := MemberMetaData{}
		oldData := input[(length -1) - i]
		newData.Offset = uint32(offset)
		newData.Width = oldData.Width
		offset = offset + int(oldData.Width)
		output[i] = newData

	}
	return output
}

