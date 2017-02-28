package gobitfields_test



import (
	"testing"
	log "github.com/Sirupsen/logrus"
	"os"
	"github.com/lagarciag/gobitfields"
	"encoding/binary"
	"github.com/lagarciag/bitwisebytes"
)

func TestMain(t *testing.M) {
	log.SetLevel(log.DebugLevel)
	formatter := &log.TextFormatter{}
	formatter.ForceColors = true
	formatter.DisableTimestamp = true
	log.SetFormatter(formatter)
	v := t.Run()

	os.Exit(v)

}

func TestGetField(t *testing.T) {
	inputSlice := []byte{95,0}

	field1 := gobitfields.GetField(inputSlice,0,4)
	t.Log("field 1: ",field1)
	field2 := gobitfields.GetField(inputSlice,4,4)
	t.Log("field 2: ",field2)

	if field1[0] != 15 {
		t.Error("field must be 15")
	}

	if field2[0] != 5 {
		t.Error("field must be 5")
	}

	recoveredSlice := make([]byte,len(inputSlice))

	gobitfields.PutField(recoveredSlice,field1,0,4)
	gobitfields.PutField(recoveredSlice,field2,4,4)

	t.Log("recovered slice: ", recoveredSlice)

	for i, aByte := range inputSlice {
		if aByte != recoveredSlice[i] {
			t.Error("mistmatch byte : ", i)
		}
	}


	inputSlice = []byte{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	field3 := gobitfields.GetField(inputSlice,33,60)
	size := len(field3)
	t.Log("field 3: ",field3, size)

	if (field3[0] != 128) && (field3[6] != 128) {
		t.Error("mistmatch in field values")
	}
}


func TestGetFieldLong(t *testing.T) {
	//inputSlice := []byte{204, 160, 85, 78, 54, 69, 143, 176, 244, 1}
	inputSlice := []byte{12, 168, 204, 14, 64, 213, 0, 200, 212, 0}
	field1 := gobitfields.GetField(inputSlice,57,16)
	t.Log("field 1: ",field1)

}


//[204 160 85 78 54 69 143 176 244 1]

func TestGetFieldSlice(t *testing.T) {

	bitfield := []byte{1,1,1,2,3,3,3,3,3,3,3,3,3,3,3,3}

	member1 := gobitfields.MemberMetaData{24,0,""}
	member2 := gobitfields.MemberMetaData{8,24,""}
	member3 := gobitfields.MemberMetaData{96,32,""}

	metadataList := []gobitfields.MemberMetaData{member1,member2,member3}

	bitfieldMatrix := gobitfields.GetAllFieldsList(bitfield,metadataList)

	t.Log("matrix ", bitfieldMatrix)

	bitfieldRecovered := gobitfields.PutAllFields(bitfieldMatrix,metadataList)

	t.Log("recovered Matrix: ", bitfieldRecovered)

	for i , field := range bitfield {
		if bitfieldRecovered[i] != field {
			t.Error("mismatch")
		}
	}

	reversed := gobitfields.ReverseMembers(bitfieldRecovered,metadataList)
	t.Log("reversed: ", reversed)

}

func TestReverseFields(t *testing.T) {

	inputslice := []byte{1,2,3,4,5,6}
	outpuslice := gobitfields.ReverseBytes(inputslice)

	t.Log(outpuslice)

}


func TestReverseMembers(t *testing.T) {

	inputslice := []byte{1,2,3,4,5,6}
	outpuslice := gobitfields.ReverseBytes(inputslice)

	t.Log(outpuslice)

}

func TestPutfields(t *testing.T) {
	inputSlice := []byte{12,168,204,14,64,213,0,0,0,0}

	toPutSlice := []byte{100,106}

	gobitfields.PutField(inputSlice,toPutSlice,57,16)

	//[12 168 204 14 64 213 0 200 212 0]



	t.Log("result: ", inputSlice)
}


func TestPutfields2(t *testing.T) {
	inputSlice := []byte{140, 1 ,149, 217, 1, 168, 26, 0, 153, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	toPutSlice := []byte{43, 237, 72, 178, 124, 188, 140, 18}

	gobitfields.PutField(inputSlice,toPutSlice,78,64)

	t.Log("result: ", inputSlice)
}


func TestGetfields3(t *testing.T) {
	inputSlice := []byte{140, 1, 149, 217, 1, 168, 26, 0, 153, 218, 74, 59, 146, 44, 31, 47, 163, 0, 92, 44, 123, 160, 210, 22, 64, 84, 3, 62}

	out := gobitfields.GetField(inputSlice,5,73)
	t.Log("result: ", out)

	out2 := gobitfields.GetField(out,57,16)

	t.Log("result: ", out2)


	nextSlice := []byte{12,168,204,14,64,213,0,0,0,0}
	gobitfields.PutField(nextSlice,out2,57,16)

	t.Log(nextSlice)

	//[12 168 204 14 64 213 0 200 212 0]
}


func TestGetfields4(t *testing.T) {
	inputSlice := []byte{50, 59, 0}

	out := gobitfields.GetField(inputSlice,4,2)
	t.Log("input slice: ", inputSlice)
	t.Log("result: ", out)


	//[12 168 204 14 64 213 0 200 212 0]
}

//140 1 149 217 1 168 26 0 153 218 74 59 146 44 31 47 163 0 92 44 123 160 210 22 64 84 3 62]



func TestGetfields5(t *testing.T) {
	inputSlice := []byte{140, 1, 149, 217, 1, 168, 26, 0, 153, 218, 74, 59, 146, 44, 31, 47, 163, 0, 92, 44, 123, 160, 210, 22, 64, 84, 3, 62}
	offset := 142
	width := 55
	bytesWidth := gobitfields.GetBytesSize(width)

	out := gobitfields.GetField(inputSlice,offset,width)

	if len(out) != bytesWidth + 1 {
		t.Error("Bad size")
	}

	tmpField := make([]byte, 8)
	for  i, _ := range tmpField {
		tmpField[i ] = 0
	}
	for i , val := range out  {
		tmpField[i] = val
	}

	aword := binary.LittleEndian.Uint64(tmpField)

	putTmpField := make([]byte,8)

	binary.LittleEndian.PutUint64(putTmpField,6855775006536048)

	t.Log("input slice: ", inputSlice)
	t.Log("result: ", out)
	t.Log("aWord: ", aword)
	t.Log("gold:",  putTmpField)

	//[12 168 204 14 64 213 0 200 212 0]
}


func TestGetfields6(t *testing.T) {
	inputSlice := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	t.Log("inputSlice: ", inputSlice)

	putWord := uint64(6855775006536048)
	putTmpField := make([]byte,8)
	offset := 142
	width := 55
	binary.LittleEndian.PutUint64(putTmpField,putWord)

	shiftLeft, _ := bitwisebytes.ShiftLeft(putTmpField,6)
	t.Log("put: ", putTmpField)
	t.Log("shiftLeft ", shiftLeft)



	testWord := binary.LittleEndian.Uint64(putTmpField)

	t.Log("testWord: ", testWord)

	gobitfields.PutField(inputSlice, putTmpField,offset,width)

	bytesWidth := gobitfields.GetBytesSize(width)

	tmpGet := inputSlice[18:25]
	t.Log("mod :", offset % 8)

	t.Log("tmpGet", tmpGet)

	out := gobitfields.GetField(inputSlice,offset,width)

	t.Log("out : ", out)

	if len(out) != bytesWidth {
		t.Error("Bad size")
	}

	tmpField := make([]byte, 8)
	for  i, _ := range tmpField {
		tmpField[i ] = 0
	}
	for i , val := range out  {
		tmpField[i] = val
	}

	aword := binary.LittleEndian.Uint64(tmpField)

	t.Log("input slice: ", inputSlice)

	t.Log("aWord: ", aword)
	//[12 168 204 14 64 213 0 200 212 0]
}



func TestGetfields7(t *testing.T) {
	inputSlice := []byte{140, 1, 149, 217, 1, 168, 26, 0, 153, 218, 74, 59, 146, 44, 31, 47, 163, 0, 92, 44, 123, 160, 210, 22, 64, 84, 3, 62}
	t.Log("inputSlice: ", inputSlice)

	offset := 142
	width := 55

	bytesWidth := gobitfields.GetBytesSize(width)

	tmpGet := inputSlice[18:25]
	t.Log("mod :", offset % 8)

	t.Log("tmpGet", tmpGet)

	out := gobitfields.GetField(inputSlice,offset,width)

	t.Log("out : ", out)

	if len(out) != bytesWidth {
		t.Error("Bad size")
	}

	tmpField := make([]byte, 8)
	for  i, _ := range tmpField {
		tmpField[i ] = 0
	}
	for i , val := range out  {
		tmpField[i] = val
	}

	aword := binary.LittleEndian.Uint64(tmpField)

	t.Log("input slice: ", inputSlice)

	t.Log("aWord: ", aword)
	//[12 168 204 14 64 213 0 200 212 0]
}
