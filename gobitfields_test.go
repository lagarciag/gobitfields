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
	

}

