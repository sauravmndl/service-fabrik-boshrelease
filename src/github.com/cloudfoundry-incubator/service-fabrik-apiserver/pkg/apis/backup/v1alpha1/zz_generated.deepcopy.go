// +build !ignore_autogenerated

//TODO copyright header

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultBackup) DeepCopyInto(out *DefaultBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultBackup.
func (in *DefaultBackup) DeepCopy() *DefaultBackup {
	if in == nil {
		return nil
	}
	out := new(DefaultBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultBackupList) DeepCopyInto(out *DefaultBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultBackupList.
func (in *DefaultBackupList) DeepCopy() *DefaultBackupList {
	if in == nil {
		return nil
	}
	out := new(DefaultBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultBackupSpec) DeepCopyInto(out *DefaultBackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultBackupSpec.
func (in *DefaultBackupSpec) DeepCopy() *DefaultBackupSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultBackupStatus) DeepCopyInto(out *DefaultBackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultBackupStatus.
func (in *DefaultBackupStatus) DeepCopy() *DefaultBackupStatus {
	if in == nil {
		return nil
	}
	out := new(DefaultBackupStatus)
	in.DeepCopyInto(out)
	return out
}
