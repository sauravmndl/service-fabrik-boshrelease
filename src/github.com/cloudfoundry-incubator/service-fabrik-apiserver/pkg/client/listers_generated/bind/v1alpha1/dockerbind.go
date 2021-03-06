//TODO copyright header

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DockerBindLister helps list DockerBinds.
type DockerBindLister interface {
	// List lists all DockerBinds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DockerBind, err error)
	// DockerBinds returns an object that can list and get DockerBinds.
	DockerBinds(namespace string) DockerBindNamespaceLister
	DockerBindListerExpansion
}

// dockerBindLister implements the DockerBindLister interface.
type dockerBindLister struct {
	indexer cache.Indexer
}

// NewDockerBindLister returns a new DockerBindLister.
func NewDockerBindLister(indexer cache.Indexer) DockerBindLister {
	return &dockerBindLister{indexer: indexer}
}

// List lists all DockerBinds in the indexer.
func (s *dockerBindLister) List(selector labels.Selector) (ret []*v1alpha1.DockerBind, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DockerBind))
	})
	return ret, err
}

// DockerBinds returns an object that can list and get DockerBinds.
func (s *dockerBindLister) DockerBinds(namespace string) DockerBindNamespaceLister {
	return dockerBindNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DockerBindNamespaceLister helps list and get DockerBinds.
type DockerBindNamespaceLister interface {
	// List lists all DockerBinds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DockerBind, err error)
	// Get retrieves the DockerBind from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DockerBind, error)
	DockerBindNamespaceListerExpansion
}

// dockerBindNamespaceLister implements the DockerBindNamespaceLister
// interface.
type dockerBindNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DockerBinds in the indexer for a given namespace.
func (s dockerBindNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DockerBind, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DockerBind))
	})
	return ret, err
}

// Get retrieves the DockerBind from the indexer for a given namespace and name.
func (s dockerBindNamespaceLister) Get(name string) (*v1alpha1.DockerBind, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dockerbind"), name)
	}
	return obj.(*v1alpha1.DockerBind), nil
}
