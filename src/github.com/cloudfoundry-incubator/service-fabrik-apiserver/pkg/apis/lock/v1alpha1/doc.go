
//TODO copyright header


// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock
// +k8s:defaulter-gen=TypeMeta
// +groupName=lock.servicefabrik.io
package v1alpha1 // import "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock/v1alpha1"

