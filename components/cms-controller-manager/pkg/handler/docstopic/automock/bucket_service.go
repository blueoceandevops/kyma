// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"

import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"

// BucketService is an autogenerated mock type for the BucketService type
type BucketService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, namespacedName, private, labels
func (_m *BucketService) Create(ctx context.Context, namespacedName types.NamespacedName, private bool, labels map[string]string) error {
	ret := _m.Called(ctx, namespacedName, private, labels)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.NamespacedName, bool, map[string]string) error); ok {
		r0 = rf(ctx, namespacedName, private, labels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, namespace, labels
func (_m *BucketService) List(ctx context.Context, namespace string, labels map[string]string) ([]string, error) {
	ret := _m.Called(ctx, namespace, labels)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) []string); ok {
		r0 = rf(ctx, namespace, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, namespace, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
