package json_marshal

import "testing"

func TestStandardMarshal(t *testing.T) {
	internal := NewInternalRandomStruct("test", "test", 1, "test")
	translator := &StandardModelTranslator{}
	marshaller := NewStandardMarshaller(translator)
	res, err := marshaller.Marshal(*internal)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	expected := []byte(`{"name":"test test","id":1}`)
	if string(res) != string(expected) {

		t.Errorf("Expected %v, got %v", string(expected), string(res))
	}
}

func TestStandardUnmarshal(t *testing.T) {
	model := []byte(`{"name":"test test","id":1}`)
	unmarshaler := &StandardUnmarshaller{}
	res, err := unmarshaler.Unmarshal(model)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	expected := RandomStruct{Id: 1, Name: "test test"}
	if expected != res {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
