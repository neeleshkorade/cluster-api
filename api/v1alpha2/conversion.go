/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	"errors"
	"fmt"

	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *Cluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.Cluster)

	if err := Convert_v1alpha2_Cluster_To_v1alpha3_Cluster(src, dst, nil); err != nil {
		return err
	}

	// Manually convert Status.APIEndpoints to Spec.ControlPlaneEndpoint.
	if len(src.Status.APIEndpoints) > 0 {
		endpoint := src.Status.APIEndpoints[0]
		dst.Spec.ControlPlaneEndpoint.Host = endpoint.Host
		dst.Spec.ControlPlaneEndpoint.Port = int32(endpoint.Port)
	}

	return nil
}

// nolint
func (dst *Cluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.Cluster)

	if err := Convert_v1alpha3_Cluster_To_v1alpha2_Cluster(src, dst, nil); err != nil {
		return err
	}

	// Manually convert Spec.ControlPlaneEndpoint to Status.APIEndpoints.
	if !src.Spec.ControlPlaneEndpoint.IsZero() {
		dst.Status.APIEndpoints = []APIEndpoint{
			{
				Host: src.Spec.ControlPlaneEndpoint.Host,
				Port: int(src.Spec.ControlPlaneEndpoint.Port),
			},
		}
	}

	return nil
}

func (src *ClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.ClusterList)

	return Convert_v1alpha2_ClusterList_To_v1alpha3_ClusterList(src, dst, nil)
}

// nolint
func (dst *ClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.ClusterList)

	return Convert_v1alpha3_ClusterList_To_v1alpha2_ClusterList(src, dst, nil)
}

func (src *Machine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.Machine)

	return Convert_v1alpha2_Machine_To_v1alpha3_Machine(src, dst, nil)
}

// nolint
func (dst *Machine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.Machine)

	return Convert_v1alpha3_Machine_To_v1alpha2_Machine(src, dst, nil)
}

func (src *MachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.MachineList)

	return Convert_v1alpha2_MachineList_To_v1alpha3_MachineList(src, dst, nil)
}

// nolint
func (dst *MachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.MachineList)

	return Convert_v1alpha3_MachineList_To_v1alpha2_MachineList(src, dst, nil)
}

func (src *MachineSet) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.MachineSet)

	return Convert_v1alpha2_MachineSet_To_v1alpha3_MachineSet(src, dst, nil)
}

// nolint
func (dst *MachineSet) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.MachineSet)

	return Convert_v1alpha3_MachineSet_To_v1alpha2_MachineSet(src, dst, nil)
}

func (src *MachineSetList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.MachineSetList)

	return Convert_v1alpha2_MachineSetList_To_v1alpha3_MachineSetList(src, dst, nil)
}

// nolint
func (dst *MachineSetList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.MachineSetList)

	return Convert_v1alpha3_MachineSetList_To_v1alpha2_MachineSetList(src, dst, nil)
}

func (src *MachineDeployment) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.MachineDeployment)

	return Convert_v1alpha2_MachineDeployment_To_v1alpha3_MachineDeployment(src, dst, nil)
}

// nolint
func (dst *MachineDeployment) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.MachineDeployment)

	return Convert_v1alpha3_MachineDeployment_To_v1alpha2_MachineDeployment(src, dst, nil)
}

func (src *MachineDeploymentList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha3.MachineDeploymentList)

	return Convert_v1alpha2_MachineDeploymentList_To_v1alpha3_MachineDeploymentList(src, dst, nil)
}

// nolint
func (dst *MachineDeploymentList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha3.MachineDeploymentList)

	return Convert_v1alpha3_MachineDeploymentList_To_v1alpha2_MachineDeploymentList(src, dst, nil)
}

func Convert_v1alpha2_MachineSpec_To_v1alpha3_MachineSpec(in *MachineSpec, out *v1alpha3.MachineSpec, s apiconversion.Scope) error {
	if out.Bootstrap.ConfigRef.Kind == "KubeadmConfig" || out.Bootstrap.ConfigRef.Kind == "KubeadmConfigTemplate" {
		out.Bootstrap.ConfigRef.APIVersion = "v1alpha3"
	}

	if err := autoConvert_v1alpha2_MachineSpec_To_v1alpha3_MachineSpec(in, out, s); err != nil {
		return err
	}

	// Discards unused ObjectMeta

	return nil
}

func Convert_v1alpha2_MachineDeploymentSpec_To_v1alpha3_MachineDeploymentSpec(in *MachineDeploymentSpec, out *v1alpha3.MachineDeploymentSpec, s apiconversion.Scope) error {
	if out.Template.Spec.Bootstrap.ConfigRef.Kind == "KubeadmConfig" || out.Template.Spec.Bootstrap.ConfigRef.Kind == "KubeadmConfigTemplate" {
		out.Template.Spec.Bootstrap.ConfigRef.APIVersion = "v1alpha3"
	}

	return autoConvert_v1alpha2_MachineDeploymentSpec_To_v1alpha3_MachineDeploymentSpec(in, out, s)
}

func Convert_v1alpha2_ClusterSpec_To_v1alpha3_ClusterSpec(in *ClusterSpec, out *v1alpha3.ClusterSpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ClusterSpec_To_v1alpha3_ClusterSpec(in, out, s); err != nil {
		return err
	}

	// Discards ControlPlaneRef

	return nil
}

func Convert_v1alpha2_ClusterStatus_To_v1alpha3_ClusterStatus(in *ClusterStatus, out *v1alpha3.ClusterStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ClusterStatus_To_v1alpha3_ClusterStatus(in, out, s); err != nil {
		return err
	}

	// Manually convert the Error fields to the Failure fields
	out.FailureMessage = in.ErrorMessage
	out.FailureReason = in.ErrorReason

	return nil
}

func Convert_v1alpha3_ClusterStatus_To_v1alpha2_ClusterStatus(in *v1alpha3.ClusterStatus, out *ClusterStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha3_ClusterStatus_To_v1alpha2_ClusterStatus(in, out, s); err != nil {
		return err
	}

	// Discards ControlPlaneReady

	// Manually convert the Failure fields to the Error fields
	out.ErrorMessage = in.FailureMessage
	out.ErrorReason = in.FailureReason

	return nil
}

func Convert_v1alpha2_MachineSetStatus_To_v1alpha3_MachineSetStatus(in *MachineSetStatus, out *v1alpha3.MachineSetStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_MachineSetStatus_To_v1alpha3_MachineSetStatus(in, out, s); err != nil {
		return err
	}

	// Manually convert the Error fields to the Failure fields
	out.FailureMessage = in.ErrorMessage
	out.FailureReason = in.ErrorReason

	return nil
}

func Convert_v1alpha3_MachineSetStatus_To_v1alpha2_MachineSetStatus(in *v1alpha3.MachineSetStatus, out *MachineSetStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha3_MachineSetStatus_To_v1alpha2_MachineSetStatus(in, out, s); err != nil {
		return err
	}

	// Manually convert the Failure fields to the Error fields
	out.ErrorMessage = in.FailureMessage
	out.ErrorReason = in.FailureReason

	return nil
}

func Convert_v1alpha2_MachineStatus_To_v1alpha3_MachineStatus(in *MachineStatus, out *v1alpha3.MachineStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_MachineStatus_To_v1alpha3_MachineStatus(in, out, s); err != nil {
		return err
	}

	// Manually convert the Error fields to the Failure fields
	out.FailureMessage = in.ErrorMessage
	out.FailureReason = in.ErrorReason

	return nil
}

func Convert_v1alpha3_ClusterSpec_To_v1alpha2_ClusterSpec(in *v1alpha3.ClusterSpec, out *ClusterSpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha3_ClusterSpec_To_v1alpha2_ClusterSpec(in, out, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha3_MachineStatus_To_v1alpha2_MachineStatus(in *v1alpha3.MachineStatus, out *MachineStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha3_MachineStatus_To_v1alpha2_MachineStatus(in, out, s); err != nil {
		return err
	}

	// Manually convert the Failure fields to the Error fields
	out.ErrorMessage = in.FailureMessage
	out.ErrorReason = in.FailureReason

	return nil
}

func Convert_v1alpha3_MachineDeploymentSpec_To_v1alpha2_MachineDeploymentSpec(in *v1alpha3.MachineDeploymentSpec, out *MachineDeploymentSpec, s apiconversion.Scope) error {
	return errors.New("cannot recover removed MachineDeploymentSpec Cluster Name")
}

func Convert_v1alpha3_MachineDeploymentStatus_To_v1alpha2_MachineDeploymentStatus(in *v1alpha3.MachineDeploymentStatus, out *MachineDeploymentStatus, s apiconversion.Scope) error {
	return fmt.Errorf("cannot recover removed MachineDeploymentStatus field Phase")
}

func Convert_v1alpha3_MachineSetSpec_To_v1alpha2_MachineSetSpec(in *v1alpha3.MachineSetSpec, out *MachineSetSpec, s apiconversion.Scope) error {
	return errors.New("cannot recover removed MachineSetSpec Cluster Name")
}

func Convert_v1alpha3_MachineSpec_To_v1alpha2_MachineSpec(in *v1alpha3.MachineSpec, out *MachineSpec, s apiconversion.Scope) error {
	return errors.New("cannot recover removed MachineSpec Cluster Name")
}

func Convert_v1alpha3_Bootstrap_To_v1alpha2_Bootstrap(in *v1alpha3.Bootstrap, out *Bootstrap, s apiconversion.Scope) error {
	// We need to fail early here given that we don't want to leak information from secrets back to the inline / plaintext field in v1alpha2.
	if in.Data == nil && in.DataSecretName != nil {
		return errors.New("cannot convert Machine's bootstrap data from Secret reference to inline field")
	}

	return autoConvert_v1alpha3_Bootstrap_To_v1alpha2_Bootstrap(in, out, s)
}
