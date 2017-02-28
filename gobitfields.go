package gobitfields

import (
	"github.com/lagarciag/bitwisebytes"
//	"math"
	//"fmt"
)

type MemberMetaData struct {
	Width  uint32
	Offset uint32
	Name string
}

type StructMetaData struct {
	Width      uint32
	Members    []MemberMetaData
	MembersMap map[string]*MemberMetaData
}

//GetFields returns a bitfield member as specified by the provided offset and width
func GetField(inputSlice []byte,offset int, width int) (outputSlice []byte) {
	bytesWidth := GetBytesSize(width)
	smallShift := offset % 8
	offSetByte := offset / 8


	byteStart := offSetByte
	byteEnd := offSetByte + bytesWidth

	if smallShift > 0 {
		byteEnd++
	}

	tmpSlice := inputSlice[byteStart:byteEnd]
	//fmt.Println("GET FIELD:", tmpSlice)

	outputSlice , err  := bitwisebytes.ShiftRight(tmpSlice,uint(smallShift))
	if err != nil {
		panic(err.Error())
	}
	return outputSlice[0:bytesWidth]
}

//PutField
func PutField(destSlice []byte, inputSlice []byte,offset int, width int) {

	smallShift := offset % 8

	smallShiftedSlice , err := bitwisebytes.ShiftLeft(inputSlice,uint(smallShift))

	if err != nil {
		panic(err.Error())
	}

	offSetByte := offset / 8

	for _ , aByte := range smallShiftedSlice {
		destSlice[offSetByte] = destSlice[offSetByte] | aByte
		offSetByte++
	}

	/*
	tmpSlice := make([]byte,len(destSlice))

	for i , aByte := range inputSlice {
		tmpSlice[i] = aByte
	}

	tmpSlice2, err := bitwisebytes.ShiftLeft(tmpSlice,uint(offset))
	if err != nil {
		panic(err.Error())
	}


	bitwisebytes.Or(destSlice,tmpSlice2)
	*/
}


func PutAllFields(inputMatrix [][]byte,metadataList []MemberMetaData) (outputSlice []byte) {
	size := 0
	for _ , member := range metadataList {
		size += int(member.Width)
	}
	bytesSize := GetBytesSize(size)

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

func GetBytesSize(bitSize int) (bytesSize int) {
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
		newData.Name = oldData.Name
		offset = offset + int(oldData.Width)
		output[i] = newData

	}
	return output
}

