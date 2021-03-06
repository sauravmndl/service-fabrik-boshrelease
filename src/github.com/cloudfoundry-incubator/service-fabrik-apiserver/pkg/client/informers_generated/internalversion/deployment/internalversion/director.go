//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	deployment "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	internalclientset "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/deployment/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DirectorInformer provides access to a shared informer and lister for
// Directors.
type DirectorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.DirectorLister
}

type directorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDirectorInformer constructs a new informer for Director type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDirectorInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDirectorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDirectorInformer constructs a new informer for Director type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDirectorInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Deployment().Directors(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Deployment().Directors(namespace).Watch(options)
			},
		},
		&deployment.Director{},
		resyncPeriod,
		indexers,
	)
}

func (f *directorInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDirectorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *directorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deployment.Director{}, f.defaultInformer)
}

func (f *directorInformer) Lister() internalversion.DirectorLister {
	return internalversion.NewDirectorLister(f.Informer().GetIndexer())
}
