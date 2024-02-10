package adapter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type note struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

type XMLDTO struct {
	id   string
	data []byte
}

type JSONDTO struct {
	id   string
	data []byte
}

type JSONAdapter struct {
	xmlDTO XMLDTO
}

func NewJSONAdapter(xmldto XMLDTO) *JSONAdapter {
	return &JSONAdapter{xmlDTO: xmldto}
}

func (x *JSONAdapter) ToJSON() (*JSONDTO, error) {
	var noteToConvert note
	if err := xml.Unmarshal(x.xmlDTO.data, &noteToConvert); err != nil {
		return nil, err
	}

	b, err := json.Marshal(noteToConvert)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &JSONDTO{
		id:   x.xmlDTO.id,
		data: b,
	}, nil
}
