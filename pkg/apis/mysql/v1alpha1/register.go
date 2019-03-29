package v1alpha1

// This package will auto register types with the Kubernetes API

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder collects the scheme builder functions for the MySQL
	// Operator API.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme applies the SchemeBuilder functions to a specified scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name for the MySQL Operator API.
const GroupName = "paas.fit2cloud.com"

// SchemeGroupVersion  is the GroupVersion for the MySQL Operator API.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// Resource gets a MySQL Operator GroupResource for a specified resource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied
// scheme.
func addKnownTypes(s *runtime.Scheme) error {
	s.AddKnownTypes(SchemeGroupVersion,
		&MysqlBackup{},
		&MysqlBackupList{},
	)
	metav1.AddToGroupVersion(s, SchemeGroupVersion)
	return nil
}
