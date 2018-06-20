//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	lock "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
	internalclientset "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/lock/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DeploymentLockInformer provides access to a shared informer and lister for
// DeploymentLocks.
type DeploymentLockInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.DeploymentLockLister
}

type deploymentLockInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeploymentLockInformer constructs a new informer for DeploymentLock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentLockInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeploymentLockInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeploymentLockInformer constructs a new informer for DeploymentLock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeploymentLockInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Lock().DeploymentLocks(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Lock().DeploymentLocks(namespace).Watch(options)
			},
		},
		&lock.DeploymentLock{},
		resyncPeriod,
		indexers,
	)
}

func (f *deploymentLockInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeploymentLockInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deploymentLockInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lock.DeploymentLock{}, f.defaultInformer)
}

func (f *deploymentLockInformer) Lister() internalversion.DeploymentLockLister {
	return internalversion.NewDeploymentLockLister(f.Informer().GetIndexer())
}