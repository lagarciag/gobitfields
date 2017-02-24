package gobitfields_test



import (
	"testing"
	log "github.com/Sirupsen/logrus"
	"os"
	"github.com/lagarciag/gobitfields"
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
	inputSlice := []byte{204, 160, 85, 78, 54, 69, 143, 176, 244, 1}

	field1 := gobitfields.GetField(inputSlice,38,15)
	t.Log("field 1: ",field1)


}


//[204 160 85 78 54 69 143 176 244 1]

func TestGetFieldSlice(t *testing.T) {

	bitfield := []byte{1,1,1,2,3,3,3,3,3,3,3,3,3,3,3,3}

	member1 := gobitfields.MemberMetaData{24,0}
	member2 := gobitfields.MemberMetaData{8,24}
	member3 := gobitfields.MemberMetaData{96,32}

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