/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen with arguments: --go-header-file=vendor/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt --input-dirs=[github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog,github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1] --internal-clientset-package=github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset --listers-package=github.com/kubernetes-incubator/service-catalog/pkg/client/listers --output-package=github.com/kubernetes-incubator/service-catalog/pkg/client/informers --versioned-clientset-package=github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset

package v1alpha1

import (
	servicecatalog_v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	clientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers/internalinterfaces"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/client/listers/servicecatalog/v1alpha1"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	cache "k8s.io/kubernetes/pkg/client/cache"
	runtime "k8s.io/kubernetes/pkg/runtime"
	watch "k8s.io/kubernetes/pkg/watch"
	time "time"
)

// InstanceInformer provides access to a shared informer and lister for
// Instances.
type InstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InstanceLister
}

type instanceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newInstanceInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ServicecatalogV1alpha1().Instances(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ServicecatalogV1alpha1().Instances(v1.NamespaceAll).Watch(options)
			},
		},
		&servicecatalog_v1alpha1.Instance{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *instanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.VersionedInformerFor(&servicecatalog_v1alpha1.Instance{}, newInstanceInformer)
}

func (f *instanceInformer) Lister() v1alpha1.InstanceLister {
	return v1alpha1.NewInstanceLister(f.Informer().GetIndexer())
}
