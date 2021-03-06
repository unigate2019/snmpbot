package snmp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type varBindStringTest struct {
	oid   OID
	value interface{}
	str   string
}

func testVarBindString(t *testing.T, test varBindStringTest) {
	var varBind = MakeVarBind(test.oid, test.value)

	assert.Equal(t, test.str, varBind.String())
}

func TestVarBindStringNull(t *testing.T) {
	testVarBindString(t, varBindStringTest{
		oid:   OID{1, 3, 6, 1, 2, 1, 1, 5, 0},
		value: nil,
		str:   "1.3.6.1.2.1.1.5.0",
	})
}

func TestVarBindString(t *testing.T) {
	testVarBindString(t, varBindStringTest{
		oid:   OID{1, 3, 6, 1, 2, 1, 1, 5, 0},
		value: []byte("qmsk-snmp test"),
		str:   "1.3.6.1.2.1.1.5.0=[113 109 115 107 45 115 110 109 112 32 116 101 115 116]",
	})
}
