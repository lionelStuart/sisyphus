package utils

import "github.com/bwmarrin/snowflake"
import "testing"

func TestSnowFlake(t *testing.T) {
	n, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		t.Logf("id %s node %d step %d time %d", id, id.Node(), id.Step(), id.Time())
		t.Logf("Int64    : %#v", id.Int64())
		t.Logf("String   : %#v", id.String())
		t.Logf("Base2    : %#v", id.Base2())
		t.Logf("Base32   : %#v", id.Base32())
		t.Logf("Base36   : %#v", id.Base36())
		t.Logf("Base58   : %#v", id.Base58())
		t.Logf("Base64   : %#v", id.Base64())
		t.Logf("Bytes    : %#v", id.Bytes())
		t.Logf("IntBytes : %#v", id.IntBytes())
	}
}
