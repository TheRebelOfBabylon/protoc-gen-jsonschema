package generator

import (
	"fmt"
	"regexp"
	"strings"

	protoc_gen_jsonschema "github.com/TheRebelOfBabylon/protoc-gen-jsonschema"
	"github.com/TheRebelOfBabylon/protoc-gen-jsonschema/config"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type JSONSchemaGenerator struct {
	cfg               *config.Config
	plugin            *protogen.Plugin
	linterRulePattern *regexp.Regexp
}

// NewJSONSchemaGenerator creates a new instance of the JSONSchemaGenerator struct
func NewJSONSchemaGenerator(plugin *protogen.Plugin, cfg *config.Config) *JSONSchemaGenerator {
	return &JSONSchemaGenerator{
		cfg:               cfg,
		plugin:            plugin,
		linterRulePattern: regexp.MustCompile(`\(-- .* --\)`),
	}
}

// Run runs the generator
func (g *JSONSchemaGenerator) Run() error {
	for _, file := range g.plugin.Files {
		if file.Generate {
			err := g.buildSchemasFromMessages(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// reformatComment reformats the protobuf comment string into a readable format
func (g *JSONSchemaGenerator) reformatComment(c protogen.Comments) string {
	comment := string(c)
	comment = strings.Replace(comment, "\n", "", -1)
	comment = g.linterRulePattern.ReplaceAllString(comment, "")
	return strings.TrimSpace(comment)
}

// createSchemaFromField creates a SchemaProperty struct
func (g *JSONSchemaGenerator) createSchemaFromField(fieldOpts *protoc_gen_jsonschema.FieldOptions, field *protogen.Field, arrayCheck bool) *SchemaProperty {
	propertySchema := &SchemaProperty{
		Description: g.reformatComment(field.Comments.Leading),
		Properties:  make(map[string]*SchemaProperty),
	}
	if fieldOpts != nil {
		// check if user specified field was required
		if fieldOpts.GetRequired() {
			propertySchema.IsRequired = true
		}
		// check if user specified field has a reference
		if ref := fieldOpts.GetRef(); ref != "" {
			propertySchema.Ref = ref
			return propertySchema
		}
		// check if user specified field has min length
		if minLength := fieldOpts.GetMinLength(); minLength != 0 {
			propertySchema.MinLength = minLength
		}
		// check if user specified field has max length
		if maxLength := fieldOpts.GetMaxLength(); maxLength != 0 {
			propertySchema.MaxLength = maxLength
		}
		// check if user specified field has min items
		if minItems := fieldOpts.GetMinItems(); minItems != 0 {
			propertySchema.MinItems = minItems
		}
		// check if user specified field has pattern
		if pattern := fieldOpts.GetPattern(); pattern != "" {
			propertySchema.Pattern = pattern
		}
		// check if user specified field has format
		if format := fieldOpts.GetFormat(); format != "" {
			propertySchema.Format = format
		}
	}
	// check for repeated key word
	if arrayCheck && field.Desc.IsList() {
		propertySchema.Type = "array"
		propertySchema.Items = g.createSchemaFromField(nil, field, false)
		if propertySchema.Items.IsRef {
			propertySchema.IsRef = true
		}
		return propertySchema
	} else if field.Desc.IsMap() {
		// maps will just be blank objects
		propertySchema.Type = "object"
		return propertySchema
	}
	// check type of field
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		propertySchema.Type = "boolean"
	case protoreflect.StringKind, protoreflect.BytesKind:
		propertySchema.Type = "string"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		propertySchema.Type = "integer"
		propertySchema.Format = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		propertySchema.Type = "integer"
		propertySchema.Format = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		propertySchema.Type = "integer"
		propertySchema.Format = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		propertySchema.Type = "integer"
		propertySchema.Format = "uint64"
	case protoreflect.FloatKind:
		propertySchema.Type = "number"
		propertySchema.Format = "float32"
	case protoreflect.DoubleKind:
		propertySchema.Type = "number"
		propertySchema.Format = "float64"
	case protoreflect.MessageKind:
		// check for google.protobuf.Struct or Any or stuff like that
		typeName := fmt.Sprintf("%s.%s", field.Message.Desc.ParentFile().Package(), field.Message.Desc.Name())
		switch typeName {
		case "google.protobuf.Struct":
			propertySchema.Type = "object"
		case "google.protobuf.Any":
			propertySchema.Type = "object"
			propertySchema.Properties["@type"] = &SchemaProperty{Type: "string"}
			propertySchema.Properties["value"] = &SchemaProperty{Type: "string"}
			propertySchema.Required = append(propertySchema.Required, []string{"@type", "value"}...)
		case "google.protobuf.Empty":
			return nil
		case "google.protobuf.Timestamp":
			propertySchema.Type = "string"
			propertySchema.Format = "date-time"
		default:
			if !arrayCheck {
				// this is the definition of the item so we don't want a redundant description
				propertySchema.Description = ""
			}
			propertySchema.Ref = fmt.Sprintf("#/definitions/%v", field.Message.Desc.Name())
			if !*g.cfg.RepeatedDefs {
				propertySchema.Ref = fmt.Sprintf("%v.json", field.Message.Desc.Name())
			}
			propertySchema.IsRef = true
		}
	case protoreflect.EnumKind:
		propertySchema.Type = "string"
		for _, value := range field.Enum.Values {
			propertySchema.Enum = append(propertySchema.Enum, string(value.Desc.Name()))
		}
	default:
		return nil
	}
	return propertySchema
}

// parseField parses a given protobuf field and creates a SchemaProperty struct
func (g *JSONSchemaGenerator) parseField(field *protogen.Field) *SchemaProperty {
	// check custom annotations
	if opt := proto.GetExtension(field.Desc.Options(), protoc_gen_jsonschema.E_FieldOptions); opt != nil {
		if fieldOpts, ok := opt.(*protoc_gen_jsonschema.FieldOptions); ok {
			// If we're ignoring it, return nil
			if fieldOpts.GetIgnore() {
				return nil
			}
			return g.createSchemaFromField(fieldOpts, field, true)
		}
	}
	return g.createSchemaFromField(nil, field, true)
}

// createSchemaFromMessage creates a Schema struct
func (g *JSONSchemaGenerator) createSchemaFromMessage(msgOpts *protoc_gen_jsonschema.MessageOptions, message *protogen.Message, schema *Schema) *Schema {
	if schema == nil {
		schema = NewSchema(
			fmt.Sprintf("%v.json", message.Desc.Name()),
			string(message.Desc.Name()),
			g.reformatComment(message.Comments.Leading),
			"object",
		)
	}
	if msgOpts != nil {
		// get Id annotion
		if newId := msgOpts.GetId(); newId != "" {
			schema.Id = newId
		}
	}
	for _, field := range message.Fields {
		// parse the field as a property
		if parsedField := g.parseField(field); parsedField != nil {
			schema.Properties[field.Desc.JSONName()] = parsedField
			if parsedField.IsRequired {
				schema.Required = append(schema.Required, field.Desc.JSONName())
			}
			// check if new field is a reference. Add to map of definitions IF cfg allows
			if parsedField.IsRef && *g.cfg.RepeatedDefs {
				newDefs := g.parseMessage(
					field.Message,
					&Schema{
						Type:        "object",
						Description: g.reformatComment(field.Message.Comments.Leading),
						Properties:  make(map[string]*SchemaProperty),
						Definitions: make(map[string]*Schema),
					},
				)
				// append all nested defs to this schema
				for name, newDef := range newDefs.Definitions {
					schema.Definitions[name] = newDef
				}
				// erase the nested defs from the new def and append
				newDefs.Definitions = make(map[string]*Schema)
				schema.Definitions[string(field.Message.Desc.Name())] = newDefs
			}
		}

	}
	// Override required parameter if stipulated in protobuf message definition
	if msgOpts != nil && msgOpts.GetAllFieldsRequired() {
		allFieldsRequired := []string{}
		for _, field := range message.Fields {
			allFieldsRequired = append(allFieldsRequired, field.Desc.JSONName())
		}
		schema.Required = allFieldsRequired[:]
	}
	return schema
}

// parseMessage will parse the protobuf Message definition and populate the Schema struct
func (g *JSONSchemaGenerator) parseMessage(message *protogen.Message, schema *Schema) *Schema {
	// check custom annotations
	if opt := proto.GetExtension(message.Desc.Options(), protoc_gen_jsonschema.E_MessageOptions); opt != nil {
		if msgOpts, ok := opt.(*protoc_gen_jsonschema.MessageOptions); ok {
			// If we're ignoring it, return nil
			if msgOpts.GetIgnore() {
				return nil
			}
			return g.createSchemaFromMessage(msgOpts, message, schema)
		}
	}
	return g.createSchemaFromMessage(nil, message, schema)
}

// buildSchemasFromMessages builds the JSON schema files from the messages inside the protobuf definition file
func (g *JSONSchemaGenerator) buildSchemasFromMessages(file *protogen.File) error {
	for _, message := range file.Messages {
		schema := g.parseMessage(message, nil)
		if schema != nil {
			outputFile := g.plugin.NewGeneratedFile(fmt.Sprintf("%s.json", message.Desc.Name()), "")
			outputFile.Write(schema.Json())
		}
	}
	return nil
}
