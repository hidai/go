package amazon

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/hidai/go/wget"
	"testing"
)

var secure_key *string = flag.String("secure_key", "", "Secure Key")

func TestBrowseNodeUnmarshal(t *testing.T) {
	flag.Parse()
	var key_and_tag = KeyAndTag{
		"AKIAJVUKFPDB27IBTYQQ",
		*secure_key,
		"hzb-20",
	}
	url := GetBrowseNodeLookupUrl(492168, key_and_tag)
	fmt.Println(url)
	xmldata, err := wget.Wget(url)
	if err != nil {
		t.Fail()
	}
	var v BrowseNodeLookupResponse
	xml.Unmarshal(xmldata, &v)
	fmt.Printf("%+v\n", v)
}
