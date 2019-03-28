// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkins_io_v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx/pkg/client/listers/jenkins.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FactInformer provides access to a shared informer and lister for
// Facts.
type FactInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.FactLister
}

type factInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFactInformer constructs a new informer for Fact type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFactInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFactInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFactInformer constructs a new informer for Fact type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFactInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Facts(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Facts(namespace).Watch(options)
			},
		},
		&jenkins_io_v1.Fact{},
		resyncPeriod,
		indexers,
	)
}

func (f *factInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFactInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *factInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkins_io_v1.Fact{}, f.defaultInformer)
}

func (f *factInformer) Lister() v1.FactLister {
	return v1.NewFactLister(f.Informer().GetIndexer())
}
