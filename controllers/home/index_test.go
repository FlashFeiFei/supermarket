/*
   Created by jinhan on 17-10-18.
   Tip:
   Update:
*/
package home

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flashfeifei/supermarket/conf"
	"github.com/flashfeifei/supermarket/models/util"
)

func init() {
	conf.ForTestInitConfig()
	util.Connect()
}

func TestGetNav(t *testing.T) {
	a := GetNav(0, 0)
	b, _ := json.Marshal(a)
	fmt.Printf("%v", string(b))
}
