package jsonUtil

import (
	"testing"

	"encoding/json"
	"github.com/hhh0pE/decimal"
	"github.com/json-iterator/go"
)

type TestStructInner struct {
	Name  string
	Value int
}

type TestStruct struct {
	Int1        int
	Int2        int64   `json:"int_2"`
	Float1      float64 `json:"float_1"`
	Float2      float32
	Str         string `json:"custom_name"`
	Str2        string
	Ignore      string `json:"-"`
	InnerStruct TestStructInner
	ArrValues   []int64 `json:"int_values"`
	ArrValues2  []decimal.Decimal
}

func BenchmarkDecode(b *testing.B) {
	var msg = []byte(`{"Int1":100,"int_2":-123,"float_1":123.1013213,"Float2":1.231e-7,"custom_name":"Test struct atatat","Str2":"Case sensetive json unmarshal based on jsonparser lib with \"quote\"","InnerStruct":{"Name":"Name","Value":90923},"int_values":[1,23,554,9354,0,-123,545,1114],"ArrValues2":[1.023,0.23023,0.000023,11313.22]}`)

	for i := 0; i < b.N; i++ {
		var resObj TestStruct
		if decoding_err := Decode(msg, &resObj); decoding_err != nil {
			b.Error("Decoding error", decoding_err)
		}
	}
}

func BenchmarkDecodeStd(b *testing.B) {
	var msg = []byte(`{"Int1":100,"int_2":-123,"float_1":123.1013213,"Float2":1.231e-7,"custom_name":"Test struct atatat","Str2":"Case sensetive json unmarshal based on jsonparser lib with \"quote\"","InnerStruct":{"Name":"Name","Value":90923},"int_values":[1,23,554,9354,0,-123,545,1114],"ArrValues2":[1.023,0.23023,0.000023,11313.22]}`)
	for i := 0; i < b.N; i++ {
		var resObj TestStruct
		if decoding_err := json.Unmarshal(msg, &resObj); decoding_err != nil {
			b.Error("Decoding error", decoding_err)
		}
	}
}

func BenchmarkJsoniter(b *testing.B) {
	var msg = []byte(`{"Int1":100,"int_2":-123,"float_1":123.1013213,"Float2":1.231e-7,"custom_name":"Test struct atatat","Str2":"Case sensetive json unmarshal based on jsonparser lib with \"quote\"","InnerStruct":{"Name":"Name","Value":90923},"int_values":[1,23,554,9354,0,-123,545,1114],"ArrValues2":[1.023,0.23023,0.000023,11313.22]}`)
	for i := 0; i < b.N; i++ {
		var resObj TestStruct
		if decoding_err := jsoniter.Unmarshal(msg, &resObj); decoding_err != nil {
			b.Error("Decoding error", decoding_err)
		}
	}
}