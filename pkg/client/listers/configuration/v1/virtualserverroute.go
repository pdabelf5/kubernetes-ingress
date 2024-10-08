// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// VirtualServerRouteLister helps list VirtualServerRoutes.
// All objects returned here must be treated as read-only.
type VirtualServerRouteLister interface {
	// List lists all VirtualServerRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error)
	// VirtualServerRoutes returns an object that can list and get VirtualServerRoutes.
	VirtualServerRoutes(namespace string) VirtualServerRouteNamespaceLister
	VirtualServerRouteListerExpansion
}

// virtualServerRouteLister implements the VirtualServerRouteLister interface.
type virtualServerRouteLister struct {
	listers.ResourceIndexer[*v1.VirtualServerRoute]
}

// NewVirtualServerRouteLister returns a new VirtualServerRouteLister.
func NewVirtualServerRouteLister(indexer cache.Indexer) VirtualServerRouteLister {
	return &virtualServerRouteLister{listers.New[*v1.VirtualServerRoute](indexer, v1.Resource("virtualserverroute"))}
}

// VirtualServerRoutes returns an object that can list and get VirtualServerRoutes.
func (s *virtualServerRouteLister) VirtualServerRoutes(namespace string) VirtualServerRouteNamespaceLister {
	return virtualServerRouteNamespaceLister{listers.NewNamespaced[*v1.VirtualServerRoute](s.ResourceIndexer, namespace)}
}

// VirtualServerRouteNamespaceLister helps list and get VirtualServerRoutes.
// All objects returned here must be treated as read-only.
type VirtualServerRouteNamespaceLister interface {
	// List lists all VirtualServerRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error)
	// Get retrieves the VirtualServerRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualServerRoute, error)
	VirtualServerRouteNamespaceListerExpansion
}

// virtualServerRouteNamespaceLister implements the VirtualServerRouteNamespaceLister
// interface.
type virtualServerRouteNamespaceLister struct {
	listers.ResourceIndexer[*v1.VirtualServerRoute]
}
