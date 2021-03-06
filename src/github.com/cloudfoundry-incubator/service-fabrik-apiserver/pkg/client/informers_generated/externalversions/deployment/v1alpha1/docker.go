//TODO copyright header

// This file was automatically generated by informer-gen

package v1alpha1

import (
	deployment_v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	clientset "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/deployment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DockerInformer provides access to a shared informer and lister for
// Dockers.
type DockerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DockerLister
}

type dockerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDockerInformer constructs a new informer for Docker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDockerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDockerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDockerInformer constructs a new informer for Docker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDockerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeploymentV1alpha1().Dockers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DeploymentV1alpha1().Dockers(namespace).Watch(options)
			},
		},
		&deployment_v1alpha1.Docker{},
		resyncPeriod,
		indexers,
	)
}

func (f *dockerInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDockerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dockerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deployment_v1alpha1.Docker{}, f.defaultInformer)
}

func (f *dockerInformer) Lister() v1alpha1.DockerLister {
	return v1alpha1.NewDockerLister(f.Informer().GetIndexer())
}
