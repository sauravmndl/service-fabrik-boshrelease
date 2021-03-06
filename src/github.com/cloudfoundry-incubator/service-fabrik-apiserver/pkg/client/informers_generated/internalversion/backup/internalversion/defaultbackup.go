//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	internalclientset "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/backup/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DefaultBackupInformer provides access to a shared informer and lister for
// DefaultBackups.
type DefaultBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.DefaultBackupLister
}

type defaultBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDefaultBackupInformer constructs a new informer for DefaultBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDefaultBackupInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDefaultBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDefaultBackupInformer constructs a new informer for DefaultBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDefaultBackupInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Backup().DefaultBackups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Backup().DefaultBackups(namespace).Watch(options)
			},
		},
		&backup.DefaultBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *defaultBackupInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDefaultBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *defaultBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&backup.DefaultBackup{}, f.defaultInformer)
}

func (f *defaultBackupInformer) Lister() internalversion.DefaultBackupLister {
	return internalversion.NewDefaultBackupLister(f.Informer().GetIndexer())
}
