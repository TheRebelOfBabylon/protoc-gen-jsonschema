package main

import (
	"flag"

	"github.com/TheRebelOfBabylon/protoc-gen-jsonschema/config"
	"github.com/TheRebelOfBabylon/protoc-gen-jsonschema/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var flags flag.FlagSet

func main() {
	cfg := &config.Config{
		EnumType: flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`),
		RepeatedDefs: flags.Bool("repeated_defs", true, `repeat definitions. If "true", repeats definitions across all files`),
	}

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		return generator.NewJSONSchemaGenerator(plugin, cfg).Run()
	})
}