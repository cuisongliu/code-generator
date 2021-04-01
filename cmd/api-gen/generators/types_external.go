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
	"fmt"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

type typesExternalGenerator struct {
	generator.DefaultGen
	imports namer.ImportTracker
}

var _ generator.Generator = &typesExternalGenerator{}

func (g *typesExternalGenerator) Filter(_ *generator.Context, _ *types.Type) bool {
	return false
}
func (g *typesExternalGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, fmt.Sprintf("%s%s \"%s\"", "meta", "v1", "k8s.io/apimachinery/pkg/apis/meta/v1"))
	return imports
}
func (g *typesExternalGenerator) Finalize(context *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, context, "$", "$")
	m := map[string]interface{}{}
	sw.Do(typesExternalTypesTemplate, m)
	return sw.Error()
}

var typesExternalTypesTemplate = `
// TestSpec defines the desired state of Test
type TestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Project. Edit Project_types.go to remove/update
	Foo string ` + "`" + `json:"foo,omitempty" protobuf:"bytes,1,opt,name=foo" ` + "`" + `
}

// TestStatus defines the observed state of Test
type TestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,description="The creation date",JSONPath=` + "`" + `.metadata.creationTimestamp` + "`" + `,priority=0

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Test is the Schema for the tests API
type Test struct {
	metav1.TypeMeta   ` + "`" + `json:",inline"` + "`" + `
	metav1.ObjectMeta ` + "`" + `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"` + "`" + `

	Spec   TestSpec   ` + "`" + `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"` + "`" + `
	Status TestStatus ` + "`" + `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"` + "`" + `
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestList contains a list of Test
type TestList struct {
	metav1.TypeMeta ` + "`" + `json:",inline"` + "`" + `
	metav1.ListMeta ` + "`" + `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"` + "`" + `
	Items           []Test ` + "`" + `json:"items" protobuf:"bytes,2,opt,name=items"` + "`" + `
}
`
