package wget

import (
	"fmt"
	"testing"
)

func TestWget(t *testing.T) {
	xmldata, err := Wget("http://www.google.com")
	if err != nil {
		t.Fail()
	}
	fmt.Printf("%s\n", string(xmldata))
}
