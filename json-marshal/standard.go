package json_marshal

import (
	"encoding/json"
	"strings"
)

type RandomStruct struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type InternalRandomStruct struct {
	name    string
	surname string
	id      int
	country string
}

func NewInternalRandomStruct(name string, surname string, id int, country string) *InternalRandomStruct {
	return &InternalRandomStruct{name: name, surname: surname, id: id, country: country}
}

type ModelTranslator interface {
	Translate(v InternalRandomStruct) RandomStruct
}

type StandardModelTranslator struct{}

func (s *StandardModelTranslator) Translate(v InternalRandomStruct) RandomStruct {
	name := strings.Join([]string{v.name, v.surname}, " ")
	return RandomStruct{Name: name, Id: v.id}
}

type StandardMarshaller struct {
	modelTranslator ModelTranslator
}

func NewStandardMarshaller(modelTranslator ModelTranslator) *StandardMarshaller {
	return &StandardMarshaller{modelTranslator: modelTranslator}
}

func (s *StandardMarshaller) Marshal(v InternalRandomStruct) ([]byte, error) {
	model := s.modelTranslator.Translate(v)
	result, err := json.Marshal(model)
	return result, err
}

type StandardUnmarshaller struct {
}

func (s *StandardUnmarshaller) Unmarshal(v []byte) (RandomStruct, error) {
	var model RandomStruct
	err := json.Unmarshal(v, &model)
	if err != nil {
		return RandomStruct{}, err
	}
	return model, nil
}
