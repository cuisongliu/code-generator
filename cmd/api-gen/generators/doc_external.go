/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"io"
	clientgentypes "k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

type docExternalGenerator struct {
	generator.DefaultGen
	outputPackage string
	gv            clientgentypes.GroupVersion
}

var _ generator.Generator = &docExternalGenerator{}

func (g *docExternalGenerator) Filter(_ *generator.Context, _ *types.Type) bool {
	return false
}

func (g *docExternalGenerator) Finalize(context *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, context, "$", "$")
	m := map[string]interface{}{
		"groupName": g.gv.Group,
		"version":   g.gv.Version,
	}
	sw.Do(registerExternalTypesTemplate, m)
	return sw.Error()
}

var registerExternalTypesTemplate = `
// +k8s:deepcopy-gen=package,register
// +k8s:protobuf-gen=package
// +k8s:openapi-gen=false
// +k8s:defaulter-gen=TypeMeta

// +groupName=$.groupName$
`
