/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"
	fmt "fmt"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
	apply "k8s.io/client-go/util/apply"
)

// StatefulSetsGetter has a method to return a StatefulSetInterface.
// A group's client should implement this interface.
type StatefulSetsGetter interface {
	StatefulSets(namespace string) StatefulSetInterface
}

// StatefulSetInterface has methods to work with StatefulSet resources.
type StatefulSetInterface interface {
	Create(ctx context.Context, statefulSet *appsv1beta2.StatefulSet, opts v1.CreateOptions) (*appsv1beta2.StatefulSet, error)
	Update(ctx context.Context, statefulSet *appsv1beta2.StatefulSet, opts v1.UpdateOptions) (*appsv1beta2.StatefulSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, statefulSet *appsv1beta2.StatefulSet, opts v1.UpdateOptions) (*appsv1beta2.StatefulSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1beta2.StatefulSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*appsv1beta2.StatefulSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1beta2.StatefulSet, err error)
	Apply(ctx context.Context, statefulSet *applyconfigurationsappsv1beta2.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *appsv1beta2.StatefulSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, statefulSet *applyconfigurationsappsv1beta2.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *appsv1beta2.StatefulSet, err error)
	GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (*appsv1beta2.Scale, error)
	UpdateScale(ctx context.Context, statefulSetName string, scale *appsv1beta2.Scale, opts v1.UpdateOptions) (*appsv1beta2.Scale, error)
	ApplyScale(ctx context.Context, statefulSetName string, scale *applyconfigurationsappsv1beta2.ScaleApplyConfiguration, opts v1.ApplyOptions) (*appsv1beta2.Scale, error)

	StatefulSetExpansion
}

// statefulSets implements StatefulSetInterface
type statefulSets struct {
	*gentype.ClientWithListAndApply[*appsv1beta2.StatefulSet, *appsv1beta2.StatefulSetList, *applyconfigurationsappsv1beta2.StatefulSetApplyConfiguration]
}

// newStatefulSets returns a StatefulSets
func newStatefulSets(c *AppsV1beta2Client, namespace string) *statefulSets {
	return &statefulSets{
		gentype.NewClientWithListAndApply[*appsv1beta2.StatefulSet, *appsv1beta2.StatefulSetList, *applyconfigurationsappsv1beta2.StatefulSetApplyConfiguration](
			"statefulsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *appsv1beta2.StatefulSet { return &appsv1beta2.StatefulSet{} },
			func() *appsv1beta2.StatefulSetList { return &appsv1beta2.StatefulSetList{} },
			gentype.PrefersProtobuf[*appsv1beta2.StatefulSet](),
		),
	}
}

// GetScale takes name of the statefulSet, and returns the corresponding appsv1beta2.Scale object, and an error if there is any.
func (c *statefulSets) GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (result *appsv1beta2.Scale, err error) {
	result = &appsv1beta2.Scale{}
	err = c.GetClient().Get().
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("statefulsets").
		Name(statefulSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *statefulSets) UpdateScale(ctx context.Context, statefulSetName string, scale *appsv1beta2.Scale, opts v1.UpdateOptions) (result *appsv1beta2.Scale, err error) {
	result = &appsv1beta2.Scale{}
	err = c.GetClient().Put().
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("statefulsets").
		Name(statefulSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *statefulSets) ApplyScale(ctx context.Context, statefulSetName string, scale *applyconfigurationsappsv1beta2.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *appsv1beta2.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	request, err := apply.NewRequest(c.GetClient(), scale)
	if err != nil {
		return nil, err
	}

	result = &appsv1beta2.Scale{}
	err = request.
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("statefulsets").
		Name(statefulSetName).
		SubResource("scale").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}