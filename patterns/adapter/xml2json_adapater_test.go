package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXML2JSONAdapter_ConvertXMLToJSON(t *testing.T) {
	xmlDto := XMLDTO{
		id: "1",
		data: []byte(`<note>
		<to>Maya</to>
		<from>Beka</from>
		<heading>Reminder</heading>
		<body>Miau!</body>
		</note>`),
	}
	jsonExpected := `{"To":"Maya","From":"Beka","Heading":"Reminder","Body":"Miau!"}`

	result, err := NewJSONAdapter(xmlDto).ToJSON()
	assert.Nil(t, err)
	assert.Equal(t, xmlDto.id, result.id)
	assert.Equal(t, jsonExpected, string(result.data))
}
