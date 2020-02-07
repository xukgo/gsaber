package jsonUtil
import (
	"fmt"
	"testing"

	"log"

	"github.com/buger/jsonparser"
	"github.com/hhh0pE/decimal"
)

type ObjIn struct {
	A string
	B int
}

type TestObject struct {
	A    string          `json:"a"`
	A2   int             `json:"A"`
	O    float64         `json:"o"`
	O2   string          `json:"O"`
	D    string          `json:"d"`
	D2   decimal.Decimal `json:"D"`
	Obj2 struct {
		Test string
	}
	Arr    []int
	ObjArr []ObjIn `json:"obj_items"`
}

func TestDecodeObject(t *testing.T) {
	var msg = []byte(`{"a":"test", "A":123, "o":12.32,"O":"Test","D":"123.134", "Obj2":{"Test":"Atata"}, "Arr":[1,20,200], "obj_items":[{"A":"test","B":10}]}`)

	var resultObject TestObject
	if decoding_err := Decode(msg, &resultObject); decoding_err != nil {
		t.Error(decoding_err)
	}
	if fmt.Sprintf("%v", resultObject) != `{test 123 12.32 Test 123.134 0 {Atata} [1 20 200] [{test 10}]}` {
		t.Error("Error!")
	}
	fmt.Printf("%v\n", resultObject)
	fmt.Println(string(msg))
}

func TestDecode(t *testing.T) {
	var res int
	Decode([]byte(`1`), &res)
	log.Println(res)
}

func TestJsonParser(t *testing.T) {
	var msg = []byte(`{"a":"test", "A":123, "o":12.32,"O":"Test"}`)

	val, _, _, _ := jsonparser.Get(msg, "o")
	fmt.Println(string(val))
	//fmt.Println(jsonparser.Get(msg, "a"))
}