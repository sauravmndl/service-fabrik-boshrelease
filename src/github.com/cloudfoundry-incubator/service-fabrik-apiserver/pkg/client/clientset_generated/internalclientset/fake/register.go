//TODO copyright header
package fake

import (
	backupinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	bindinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	deploymentinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	lockinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	AddToScheme(scheme)
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kuberentes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
func AddToScheme(scheme *runtime.Scheme) {
	backupinternalversion.AddToScheme(scheme)
	bindinternalversion.AddToScheme(scheme)
	deploymentinternalversion.AddToScheme(scheme)
	lockinternalversion.AddToScheme(scheme)

}
