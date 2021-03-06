// +build !ignore_autogenerated

//TODO copyright header

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package lock

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentLock) DeepCopyInto(out *DeploymentLock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentLock.
func (in *DeploymentLock) DeepCopy() *DeploymentLock {
	if in == nil {
		return nil
	}
	out := new(DeploymentLock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentLock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentLockList) DeepCopyInto(out *DeploymentLockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentLock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentLockList.
func (in *DeploymentLockList) DeepCopy() *DeploymentLockList {
	if in == nil {
		return nil
	}
	out := new(DeploymentLockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentLockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentLockSpec) DeepCopyInto(out *DeploymentLockSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentLockSpec.
func (in *DeploymentLockSpec) DeepCopy() *DeploymentLockSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentLockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentLockStatus) DeepCopyInto(out *DeploymentLockStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentLockStatus.
func (in *DeploymentLockStatus) DeepCopy() *DeploymentLockStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentLockStatus)
	in.DeepCopyInto(out)
	return out
}
