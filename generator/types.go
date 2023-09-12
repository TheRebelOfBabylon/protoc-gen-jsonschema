package generator

import (
	"encoding/json"
)

type SchemaProperty struct {
	Type 		string 					   `json:"type,omitempty"`
	Format      string					   `json:"format,omitempty"`
	Description string 					   `json:"description,omitempty"`
	Ref		    string 					   `json:"$ref,omitempty"`
	Enum		[]string 				   `json:"enum,omitempty"`
	Properties  map[string]*SchemaProperty `json:"properties,omitempty"`
	Required    []string				   `json:"required,omitempty"`
	Items		*SchemaProperty			   `json:"items,omitempty"`
	MinItems    int32					   `json:"minItems,omitempty"`
	MinLength   int32					   `json:"minLength,omitempty"`
	MaxLength   int32					   `json:"maxLength,omitempty"`
	Pattern     string					   `json:"pattern,omitempty"`
	IsRequired  bool					   `json:"-"`
	IsRef		bool                       `json:"-"`
}

type Schema struct {
	Id 		  	string 					   `json:"$id,omitempty"`
	SchemaRef 	string  				   `json:"$schema,omitempty"`
	Title	  	string 					   `json:"title,omitempty"`
	Description string 					   `json:"description,omitempty"`
	Type		string 					   `json:"type,omitempty"`
	Properties  map[string]*SchemaProperty `json:"properties,omitempty"`
	Required    []string				   `json:"required,omitempty"`
	Definitions map[string]*Schema		   `json:"definitions,omitempty"`
	IsRequired  bool					   `json:"-"`
}

// NewSchema creates a NewSchema struct
func NewSchema(id, title, description, schemaType string) *Schema {
	return &Schema{
		Id: id,
		SchemaRef: "http://json-schema.org/draft-07/schema#",
		Title: title,
		Description: description,
		Type: schemaType,
		Properties: make(map[string]*SchemaProperty),
		Definitions: make(map[string]*Schema),
	}
}

// Json serializes the schema into JSON format
func (s *Schema) Json() []byte {
	bytes, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		panic(err)
	}
	return bytes[:]
}