// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/client-go/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Binding_To_servicecatalog_Binding,
		Convert_servicecatalog_Binding_To_v1alpha1_Binding,
		Convert_v1alpha1_BindingCondition_To_servicecatalog_BindingCondition,
		Convert_servicecatalog_BindingCondition_To_v1alpha1_BindingCondition,
		Convert_v1alpha1_BindingList_To_servicecatalog_BindingList,
		Convert_servicecatalog_BindingList_To_v1alpha1_BindingList,
		Convert_v1alpha1_BindingSpec_To_servicecatalog_BindingSpec,
		Convert_servicecatalog_BindingSpec_To_v1alpha1_BindingSpec,
		Convert_v1alpha1_BindingStatus_To_servicecatalog_BindingStatus,
		Convert_servicecatalog_BindingStatus_To_v1alpha1_BindingStatus,
		Convert_v1alpha1_Broker_To_servicecatalog_Broker,
		Convert_servicecatalog_Broker_To_v1alpha1_Broker,
		Convert_v1alpha1_BrokerCondition_To_servicecatalog_BrokerCondition,
		Convert_servicecatalog_BrokerCondition_To_v1alpha1_BrokerCondition,
		Convert_v1alpha1_BrokerList_To_servicecatalog_BrokerList,
		Convert_servicecatalog_BrokerList_To_v1alpha1_BrokerList,
		Convert_v1alpha1_BrokerSpec_To_servicecatalog_BrokerSpec,
		Convert_servicecatalog_BrokerSpec_To_v1alpha1_BrokerSpec,
		Convert_v1alpha1_BrokerStatus_To_servicecatalog_BrokerStatus,
		Convert_servicecatalog_BrokerStatus_To_v1alpha1_BrokerStatus,
		Convert_v1alpha1_Instance_To_servicecatalog_Instance,
		Convert_servicecatalog_Instance_To_v1alpha1_Instance,
		Convert_v1alpha1_InstanceCondition_To_servicecatalog_InstanceCondition,
		Convert_servicecatalog_InstanceCondition_To_v1alpha1_InstanceCondition,
		Convert_v1alpha1_InstanceList_To_servicecatalog_InstanceList,
		Convert_servicecatalog_InstanceList_To_v1alpha1_InstanceList,
		Convert_v1alpha1_InstanceSpec_To_servicecatalog_InstanceSpec,
		Convert_servicecatalog_InstanceSpec_To_v1alpha1_InstanceSpec,
		Convert_v1alpha1_InstanceStatus_To_servicecatalog_InstanceStatus,
		Convert_servicecatalog_InstanceStatus_To_v1alpha1_InstanceStatus,
		Convert_v1alpha1_ServiceClass_To_servicecatalog_ServiceClass,
		Convert_servicecatalog_ServiceClass_To_v1alpha1_ServiceClass,
		Convert_v1alpha1_ServiceClassList_To_servicecatalog_ServiceClassList,
		Convert_servicecatalog_ServiceClassList_To_v1alpha1_ServiceClassList,
		Convert_v1alpha1_ServicePlan_To_servicecatalog_ServicePlan,
		Convert_servicecatalog_ServicePlan_To_v1alpha1_ServicePlan,
	)
}

func autoConvert_v1alpha1_Binding_To_servicecatalog_Binding(in *Binding, out *servicecatalog.Binding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_BindingSpec_To_servicecatalog_BindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BindingStatus_To_servicecatalog_BindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Binding_To_servicecatalog_Binding(in *Binding, out *servicecatalog.Binding, s conversion.Scope) error {
	return autoConvert_v1alpha1_Binding_To_servicecatalog_Binding(in, out, s)
}

func autoConvert_servicecatalog_Binding_To_v1alpha1_Binding(in *servicecatalog.Binding, out *Binding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_BindingSpec_To_v1alpha1_BindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_BindingStatus_To_v1alpha1_BindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_servicecatalog_Binding_To_v1alpha1_Binding(in *servicecatalog.Binding, out *Binding, s conversion.Scope) error {
	return autoConvert_servicecatalog_Binding_To_v1alpha1_Binding(in, out, s)
}

func autoConvert_v1alpha1_BindingCondition_To_servicecatalog_BindingCondition(in *BindingCondition, out *servicecatalog.BindingCondition, s conversion.Scope) error {
	out.Type = servicecatalog.BindingConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1alpha1_BindingCondition_To_servicecatalog_BindingCondition(in *BindingCondition, out *servicecatalog.BindingCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_BindingCondition_To_servicecatalog_BindingCondition(in, out, s)
}

func autoConvert_servicecatalog_BindingCondition_To_v1alpha1_BindingCondition(in *servicecatalog.BindingCondition, out *BindingCondition, s conversion.Scope) error {
	out.Type = BindingConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_servicecatalog_BindingCondition_To_v1alpha1_BindingCondition(in *servicecatalog.BindingCondition, out *BindingCondition, s conversion.Scope) error {
	return autoConvert_servicecatalog_BindingCondition_To_v1alpha1_BindingCondition(in, out, s)
}

func autoConvert_v1alpha1_BindingList_To_servicecatalog_BindingList(in *BindingList, out *servicecatalog.BindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.Binding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_BindingList_To_servicecatalog_BindingList(in *BindingList, out *servicecatalog.BindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_BindingList_To_servicecatalog_BindingList(in, out, s)
}

func autoConvert_servicecatalog_BindingList_To_v1alpha1_BindingList(in *servicecatalog.BindingList, out *BindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Binding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_servicecatalog_BindingList_To_v1alpha1_BindingList(in *servicecatalog.BindingList, out *BindingList, s conversion.Scope) error {
	return autoConvert_servicecatalog_BindingList_To_v1alpha1_BindingList(in, out, s)
}

func autoConvert_v1alpha1_BindingSpec_To_servicecatalog_BindingSpec(in *BindingSpec, out *servicecatalog.BindingSpec, s conversion.Scope) error {
	out.InstanceRef = in.InstanceRef
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.SecretName = in.SecretName
	out.ExternalID = in.ExternalID
	out.Checksum = (*string)(unsafe.Pointer(in.Checksum))
	return nil
}

func Convert_v1alpha1_BindingSpec_To_servicecatalog_BindingSpec(in *BindingSpec, out *servicecatalog.BindingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_BindingSpec_To_servicecatalog_BindingSpec(in, out, s)
}

func autoConvert_servicecatalog_BindingSpec_To_v1alpha1_BindingSpec(in *servicecatalog.BindingSpec, out *BindingSpec, s conversion.Scope) error {
	out.InstanceRef = in.InstanceRef
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.SecretName = in.SecretName
	out.ExternalID = in.ExternalID
	out.Checksum = (*string)(unsafe.Pointer(in.Checksum))
	return nil
}

func Convert_servicecatalog_BindingSpec_To_v1alpha1_BindingSpec(in *servicecatalog.BindingSpec, out *BindingSpec, s conversion.Scope) error {
	return autoConvert_servicecatalog_BindingSpec_To_v1alpha1_BindingSpec(in, out, s)
}

func autoConvert_v1alpha1_BindingStatus_To_servicecatalog_BindingStatus(in *BindingStatus, out *servicecatalog.BindingStatus, s conversion.Scope) error {
	out.Conditions = *(*[]servicecatalog.BindingCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_v1alpha1_BindingStatus_To_servicecatalog_BindingStatus(in *BindingStatus, out *servicecatalog.BindingStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_BindingStatus_To_servicecatalog_BindingStatus(in, out, s)
}

func autoConvert_servicecatalog_BindingStatus_To_v1alpha1_BindingStatus(in *servicecatalog.BindingStatus, out *BindingStatus, s conversion.Scope) error {
	out.Conditions = *(*[]BindingCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_servicecatalog_BindingStatus_To_v1alpha1_BindingStatus(in *servicecatalog.BindingStatus, out *BindingStatus, s conversion.Scope) error {
	return autoConvert_servicecatalog_BindingStatus_To_v1alpha1_BindingStatus(in, out, s)
}

func autoConvert_v1alpha1_Broker_To_servicecatalog_Broker(in *Broker, out *servicecatalog.Broker, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_BrokerSpec_To_servicecatalog_BrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BrokerStatus_To_servicecatalog_BrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Broker_To_servicecatalog_Broker(in *Broker, out *servicecatalog.Broker, s conversion.Scope) error {
	return autoConvert_v1alpha1_Broker_To_servicecatalog_Broker(in, out, s)
}

func autoConvert_servicecatalog_Broker_To_v1alpha1_Broker(in *servicecatalog.Broker, out *Broker, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_BrokerSpec_To_v1alpha1_BrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_BrokerStatus_To_v1alpha1_BrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_servicecatalog_Broker_To_v1alpha1_Broker(in *servicecatalog.Broker, out *Broker, s conversion.Scope) error {
	return autoConvert_servicecatalog_Broker_To_v1alpha1_Broker(in, out, s)
}

func autoConvert_v1alpha1_BrokerCondition_To_servicecatalog_BrokerCondition(in *BrokerCondition, out *servicecatalog.BrokerCondition, s conversion.Scope) error {
	out.Type = servicecatalog.BrokerConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1alpha1_BrokerCondition_To_servicecatalog_BrokerCondition(in *BrokerCondition, out *servicecatalog.BrokerCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_BrokerCondition_To_servicecatalog_BrokerCondition(in, out, s)
}

func autoConvert_servicecatalog_BrokerCondition_To_v1alpha1_BrokerCondition(in *servicecatalog.BrokerCondition, out *BrokerCondition, s conversion.Scope) error {
	out.Type = BrokerConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_servicecatalog_BrokerCondition_To_v1alpha1_BrokerCondition(in *servicecatalog.BrokerCondition, out *BrokerCondition, s conversion.Scope) error {
	return autoConvert_servicecatalog_BrokerCondition_To_v1alpha1_BrokerCondition(in, out, s)
}

func autoConvert_v1alpha1_BrokerList_To_servicecatalog_BrokerList(in *BrokerList, out *servicecatalog.BrokerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.Broker)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_BrokerList_To_servicecatalog_BrokerList(in *BrokerList, out *servicecatalog.BrokerList, s conversion.Scope) error {
	return autoConvert_v1alpha1_BrokerList_To_servicecatalog_BrokerList(in, out, s)
}

func autoConvert_servicecatalog_BrokerList_To_v1alpha1_BrokerList(in *servicecatalog.BrokerList, out *BrokerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Broker)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_servicecatalog_BrokerList_To_v1alpha1_BrokerList(in *servicecatalog.BrokerList, out *BrokerList, s conversion.Scope) error {
	return autoConvert_servicecatalog_BrokerList_To_v1alpha1_BrokerList(in, out, s)
}

func autoConvert_v1alpha1_BrokerSpec_To_servicecatalog_BrokerSpec(in *BrokerSpec, out *servicecatalog.BrokerSpec, s conversion.Scope) error {
	out.URL = in.URL
	out.AuthSecret = (*v1.ObjectReference)(unsafe.Pointer(in.AuthSecret))
	return nil
}

func Convert_v1alpha1_BrokerSpec_To_servicecatalog_BrokerSpec(in *BrokerSpec, out *servicecatalog.BrokerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_BrokerSpec_To_servicecatalog_BrokerSpec(in, out, s)
}

func autoConvert_servicecatalog_BrokerSpec_To_v1alpha1_BrokerSpec(in *servicecatalog.BrokerSpec, out *BrokerSpec, s conversion.Scope) error {
	out.URL = in.URL
	out.AuthSecret = (*v1.ObjectReference)(unsafe.Pointer(in.AuthSecret))
	return nil
}

func Convert_servicecatalog_BrokerSpec_To_v1alpha1_BrokerSpec(in *servicecatalog.BrokerSpec, out *BrokerSpec, s conversion.Scope) error {
	return autoConvert_servicecatalog_BrokerSpec_To_v1alpha1_BrokerSpec(in, out, s)
}

func autoConvert_v1alpha1_BrokerStatus_To_servicecatalog_BrokerStatus(in *BrokerStatus, out *servicecatalog.BrokerStatus, s conversion.Scope) error {
	out.Conditions = *(*[]servicecatalog.BrokerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_v1alpha1_BrokerStatus_To_servicecatalog_BrokerStatus(in *BrokerStatus, out *servicecatalog.BrokerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_BrokerStatus_To_servicecatalog_BrokerStatus(in, out, s)
}

func autoConvert_servicecatalog_BrokerStatus_To_v1alpha1_BrokerStatus(in *servicecatalog.BrokerStatus, out *BrokerStatus, s conversion.Scope) error {
	out.Conditions = *(*[]BrokerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_servicecatalog_BrokerStatus_To_v1alpha1_BrokerStatus(in *servicecatalog.BrokerStatus, out *BrokerStatus, s conversion.Scope) error {
	return autoConvert_servicecatalog_BrokerStatus_To_v1alpha1_BrokerStatus(in, out, s)
}

func autoConvert_v1alpha1_Instance_To_servicecatalog_Instance(in *Instance, out *servicecatalog.Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_InstanceSpec_To_servicecatalog_InstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_InstanceStatus_To_servicecatalog_InstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_Instance_To_servicecatalog_Instance(in *Instance, out *servicecatalog.Instance, s conversion.Scope) error {
	return autoConvert_v1alpha1_Instance_To_servicecatalog_Instance(in, out, s)
}

func autoConvert_servicecatalog_Instance_To_v1alpha1_Instance(in *servicecatalog.Instance, out *Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_InstanceSpec_To_v1alpha1_InstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_InstanceStatus_To_v1alpha1_InstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_servicecatalog_Instance_To_v1alpha1_Instance(in *servicecatalog.Instance, out *Instance, s conversion.Scope) error {
	return autoConvert_servicecatalog_Instance_To_v1alpha1_Instance(in, out, s)
}

func autoConvert_v1alpha1_InstanceCondition_To_servicecatalog_InstanceCondition(in *InstanceCondition, out *servicecatalog.InstanceCondition, s conversion.Scope) error {
	out.Type = servicecatalog.InstanceConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1alpha1_InstanceCondition_To_servicecatalog_InstanceCondition(in *InstanceCondition, out *servicecatalog.InstanceCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceCondition_To_servicecatalog_InstanceCondition(in, out, s)
}

func autoConvert_servicecatalog_InstanceCondition_To_v1alpha1_InstanceCondition(in *servicecatalog.InstanceCondition, out *InstanceCondition, s conversion.Scope) error {
	out.Type = InstanceConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_servicecatalog_InstanceCondition_To_v1alpha1_InstanceCondition(in *servicecatalog.InstanceCondition, out *InstanceCondition, s conversion.Scope) error {
	return autoConvert_servicecatalog_InstanceCondition_To_v1alpha1_InstanceCondition(in, out, s)
}

func autoConvert_v1alpha1_InstanceList_To_servicecatalog_InstanceList(in *InstanceList, out *servicecatalog.InstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.Instance)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_InstanceList_To_servicecatalog_InstanceList(in *InstanceList, out *servicecatalog.InstanceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceList_To_servicecatalog_InstanceList(in, out, s)
}

func autoConvert_servicecatalog_InstanceList_To_v1alpha1_InstanceList(in *servicecatalog.InstanceList, out *InstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Instance)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_servicecatalog_InstanceList_To_v1alpha1_InstanceList(in *servicecatalog.InstanceList, out *InstanceList, s conversion.Scope) error {
	return autoConvert_servicecatalog_InstanceList_To_v1alpha1_InstanceList(in, out, s)
}

func autoConvert_v1alpha1_InstanceSpec_To_servicecatalog_InstanceSpec(in *InstanceSpec, out *servicecatalog.InstanceSpec, s conversion.Scope) error {
	out.ServiceClassName = in.ServiceClassName
	out.PlanName = in.PlanName
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ExternalID = in.ExternalID
	out.Checksum = (*string)(unsafe.Pointer(in.Checksum))
	return nil
}

func Convert_v1alpha1_InstanceSpec_To_servicecatalog_InstanceSpec(in *InstanceSpec, out *servicecatalog.InstanceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceSpec_To_servicecatalog_InstanceSpec(in, out, s)
}

func autoConvert_servicecatalog_InstanceSpec_To_v1alpha1_InstanceSpec(in *servicecatalog.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	out.ServiceClassName = in.ServiceClassName
	out.PlanName = in.PlanName
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ExternalID = in.ExternalID
	out.Checksum = (*string)(unsafe.Pointer(in.Checksum))
	return nil
}

func Convert_servicecatalog_InstanceSpec_To_v1alpha1_InstanceSpec(in *servicecatalog.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	return autoConvert_servicecatalog_InstanceSpec_To_v1alpha1_InstanceSpec(in, out, s)
}

func autoConvert_v1alpha1_InstanceStatus_To_servicecatalog_InstanceStatus(in *InstanceStatus, out *servicecatalog.InstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]servicecatalog.InstanceCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.DashboardURL = (*string)(unsafe.Pointer(in.DashboardURL))
	return nil
}

func Convert_v1alpha1_InstanceStatus_To_servicecatalog_InstanceStatus(in *InstanceStatus, out *servicecatalog.InstanceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceStatus_To_servicecatalog_InstanceStatus(in, out, s)
}

func autoConvert_servicecatalog_InstanceStatus_To_v1alpha1_InstanceStatus(in *servicecatalog.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]InstanceCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.DashboardURL = (*string)(unsafe.Pointer(in.DashboardURL))
	return nil
}

func Convert_servicecatalog_InstanceStatus_To_v1alpha1_InstanceStatus(in *servicecatalog.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	return autoConvert_servicecatalog_InstanceStatus_To_v1alpha1_InstanceStatus(in, out, s)
}

func autoConvert_v1alpha1_ServiceClass_To_servicecatalog_ServiceClass(in *ServiceClass, out *servicecatalog.ServiceClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.BrokerName = in.BrokerName
	out.Bindable = in.Bindable
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make([]servicecatalog.ServicePlan, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ServicePlan_To_servicecatalog_ServicePlan(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Plans = nil
	}
	out.PlanUpdatable = in.PlanUpdatable
	out.ExternalID = in.ExternalID
	out.OSBTags = *(*[]string)(unsafe.Pointer(&in.OSBTags))
	out.OSBRequires = *(*[]string)(unsafe.Pointer(&in.OSBRequires))
	out.OSBMaxDBPerNode = (*string)(unsafe.Pointer(in.OSBMaxDBPerNode))
	out.OSBDashboardOAuth2ClientID = (*string)(unsafe.Pointer(in.OSBDashboardOAuth2ClientID))
	out.OSBDashboardSecret = (*string)(unsafe.Pointer(in.OSBDashboardSecret))
	out.OSBDashboardRedirectURI = (*string)(unsafe.Pointer(in.OSBDashboardRedirectURI))
	out.Description = in.Description
	out.OSBMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.OSBMetadata))
	return nil
}

func Convert_v1alpha1_ServiceClass_To_servicecatalog_ServiceClass(in *ServiceClass, out *servicecatalog.ServiceClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceClass_To_servicecatalog_ServiceClass(in, out, s)
}

func autoConvert_servicecatalog_ServiceClass_To_v1alpha1_ServiceClass(in *servicecatalog.ServiceClass, out *ServiceClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.BrokerName = in.BrokerName
	out.Bindable = in.Bindable
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make([]ServicePlan, len(*in))
		for i := range *in {
			if err := Convert_servicecatalog_ServicePlan_To_v1alpha1_ServicePlan(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Plans = nil
	}
	out.PlanUpdatable = in.PlanUpdatable
	out.ExternalID = in.ExternalID
	out.OSBTags = *(*[]string)(unsafe.Pointer(&in.OSBTags))
	out.OSBRequires = *(*[]string)(unsafe.Pointer(&in.OSBRequires))
	out.OSBMaxDBPerNode = (*string)(unsafe.Pointer(in.OSBMaxDBPerNode))
	out.OSBDashboardOAuth2ClientID = (*string)(unsafe.Pointer(in.OSBDashboardOAuth2ClientID))
	out.OSBDashboardSecret = (*string)(unsafe.Pointer(in.OSBDashboardSecret))
	out.OSBDashboardRedirectURI = (*string)(unsafe.Pointer(in.OSBDashboardRedirectURI))
	out.Description = in.Description
	out.OSBMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.OSBMetadata))
	return nil
}

func Convert_servicecatalog_ServiceClass_To_v1alpha1_ServiceClass(in *servicecatalog.ServiceClass, out *ServiceClass, s conversion.Scope) error {
	return autoConvert_servicecatalog_ServiceClass_To_v1alpha1_ServiceClass(in, out, s)
}

func autoConvert_v1alpha1_ServiceClassList_To_servicecatalog_ServiceClassList(in *ServiceClassList, out *servicecatalog.ServiceClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]servicecatalog.ServiceClass, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ServiceClass_To_servicecatalog_ServiceClass(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1alpha1_ServiceClassList_To_servicecatalog_ServiceClassList(in *ServiceClassList, out *servicecatalog.ServiceClassList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceClassList_To_servicecatalog_ServiceClassList(in, out, s)
}

func autoConvert_servicecatalog_ServiceClassList_To_v1alpha1_ServiceClassList(in *servicecatalog.ServiceClassList, out *ServiceClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceClass, len(*in))
		for i := range *in {
			if err := Convert_servicecatalog_ServiceClass_To_v1alpha1_ServiceClass(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_servicecatalog_ServiceClassList_To_v1alpha1_ServiceClassList(in *servicecatalog.ServiceClassList, out *ServiceClassList, s conversion.Scope) error {
	return autoConvert_servicecatalog_ServiceClassList_To_v1alpha1_ServiceClassList(in, out, s)
}

func autoConvert_v1alpha1_ServicePlan_To_servicecatalog_ServicePlan(in *ServicePlan, out *servicecatalog.ServicePlan, s conversion.Scope) error {
	out.Name = in.Name
	out.Description = in.Description
	out.ExternalID = in.ExternalID
	out.Bindable = (*bool)(unsafe.Pointer(in.Bindable))
	out.OSBFree = in.OSBFree
	out.OSBMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.OSBMetadata))
	return nil
}

func Convert_v1alpha1_ServicePlan_To_servicecatalog_ServicePlan(in *ServicePlan, out *servicecatalog.ServicePlan, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServicePlan_To_servicecatalog_ServicePlan(in, out, s)
}

func autoConvert_servicecatalog_ServicePlan_To_v1alpha1_ServicePlan(in *servicecatalog.ServicePlan, out *ServicePlan, s conversion.Scope) error {
	out.Name = in.Name
	out.ExternalID = in.ExternalID
	out.Bindable = (*bool)(unsafe.Pointer(in.Bindable))
	out.OSBFree = in.OSBFree
	out.Description = in.Description
	out.OSBMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.OSBMetadata))
	return nil
}

func Convert_servicecatalog_ServicePlan_To_v1alpha1_ServicePlan(in *servicecatalog.ServicePlan, out *ServicePlan, s conversion.Scope) error {
	return autoConvert_servicecatalog_ServicePlan_To_v1alpha1_ServicePlan(in, out, s)
}
