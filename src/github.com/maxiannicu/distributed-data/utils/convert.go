package utils

import (
	"encoding/json"
	"encoding/xml"
)

func Serialize(format byte, source interface{}) ([]byte,error) {
	switch format {
	case XMLFormat:
		return xml.Marshal(source)
	case JsonFormat:
		return json.Marshal(source)
	default:
		return nil, NewInvalidArgumentError("Invalid format given")
	}
}

func Deserealize(format byte, bytes []byte, destination interface{}) error {
	switch format {
	case XMLFormat:
		return xml.Unmarshal(bytes, destination)
	case JsonFormat:
		return json.Unmarshal(bytes, destination)
	default:
		return NewInvalidArgumentError("Invalid format given")
	}
}

