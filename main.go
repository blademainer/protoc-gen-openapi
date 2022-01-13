// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"flag"

	"github.com/kollalabs/protoc-gen-openapi/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var flags flag.FlagSet

func main() {
	conf := generator.Configuration{
		Version:     flags.String("version", "0.0.1", "version number text, e.g. 1.2.3"),
		Title:       flags.String("title", "", "name of the API"),
		Description: flags.String("description", "", "description of the API"),
		Naming:      flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`),
		Validate:    flags.Bool("validate", false, "parse protoc-gen-validate options that are supported into openapi field options"),
	}

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		return generator.NewOpenAPIv3Generator(plugin, conf).Run()
	})
}
