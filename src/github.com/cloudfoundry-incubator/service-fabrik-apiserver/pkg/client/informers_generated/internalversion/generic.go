//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	"fmt"
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	bind "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	deployment "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	lock "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=backup.servicefabrik.io, Version=internalVersion
	case backup.SchemeGroupVersion.WithResource("defaultbackups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Backup().InternalVersion().DefaultBackups().Informer()}, nil

		// Group=bind.servicefabrik.io, Version=internalVersion
	case bind.SchemeGroupVersion.WithResource("directorbinds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bind().InternalVersion().DirectorBinds().Informer()}, nil
	case bind.SchemeGroupVersion.WithResource("dockerbinds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bind().InternalVersion().DockerBinds().Informer()}, nil

		// Group=deployment.servicefabrik.io, Version=internalVersion
	case deployment.SchemeGroupVersion.WithResource("directors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Deployment().InternalVersion().Directors().Informer()}, nil
	case deployment.SchemeGroupVersion.WithResource("dockers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Deployment().InternalVersion().Dockers().Informer()}, nil

		// Group=lock.servicefabrik.io, Version=internalVersion
	case lock.SchemeGroupVersion.WithResource("deploymentlocks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Lock().InternalVersion().DeploymentLocks().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
