/*
Copyright 2023.

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
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
	utilconversion "github.com/roehrich-hpe/multiversion-crd-play/github/cluster-api/util/conversion"
)

var convertlog = logf.Log.WithName("convert-v1alpha2")

func (src *Desert) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert Desert To Hub")
	dst := dstRaw.(*dwsv1alpha.Desert)

	if err := Convert_v1alpha2_Desert_To_v1alpha3_Desert(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &dwsv1alpha.Desert{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dst.Spec.Days = restored.Spec.Days

	return nil
}

func (dst *Desert) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert Desert From Hub")
	src := srcRaw.(*dwsv1alpha.Desert)

	if err := Convert_v1alpha3_Desert_To_v1alpha2_Desert(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

func (src *Vehicle) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert Vehicle To Hub")
	dst := dstRaw.(*dwsv1alpha.Vehicle)

	if err := Convert_v1alpha2_Vehicle_To_v1alpha3_Vehicle(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &dwsv1alpha.Vehicle{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dst.Status.Tires = restored.Status.Tires

	return nil
}

func (dst *Vehicle) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert Vehicle From Hub")
	src := srcRaw.(*dwsv1alpha.Vehicle)

	if err := Convert_v1alpha3_Vehicle_To_v1alpha2_Vehicle(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// The List-based ConvertTo/ConvertFrom routines are never used by the
// conversion webhook, but the conversion-verifier tool wants to see them.

func (src *DesertList) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert DesertList To Hub")
	dst := dstRaw.(*dwsv1alpha.DesertList)

	return Convert_v1alpha2_DesertList_To_v1alpha3_DesertList(src, dst, nil)
}

func (dst *DesertList) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert DesertList From Hub")
	src := srcRaw.(*dwsv1alpha.DesertList)

	return Convert_v1alpha3_DesertList_To_v1alpha2_DesertList(src, dst, nil)
}

func (src *VehicleList) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert VehicleList To Hub")
	dst := dstRaw.(*dwsv1alpha.VehicleList)

	return Convert_v1alpha2_VehicleList_To_v1alpha3_VehicleList(src, dst, nil)
}

func (dst *VehicleList) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert VehicleList From Hub")
	src := srcRaw.(*dwsv1alpha.VehicleList)

	return Convert_v1alpha3_VehicleList_To_v1alpha2_VehicleList(src, dst, nil)
}

func Convert_v1alpha3_DesertSpec_To_v1alpha2_DesertSpec(in *dwsv1alpha.DesertSpec, out *DesertSpec, s apiconversion.Scope) error {
	// Spec.Days was introduced in v1alpha2, so it needs a custom
	// conversion function.  The value will be preserved in an annotation,
	// allowing roundtrip without losing information.

	// The conversion-gen tool printed a warning about this.  Also see the
	// warning it placed in
	// autoConvert_v1alpha3_DesertSpec_To_v1alpha2_DesertSpec()
	// in zz_generated.conversion.go.

	// The conversion-gen tool creates all the parts, but in this case it
	// omitted Convert_v1alpha3_DesertSpec_To_v1alpha2_DesertSpec(),
	// forcing us to acknowledge that we are handling the conversion for
	// Spec.Days.

	return autoConvert_v1alpha3_DesertSpec_To_v1alpha2_DesertSpec(in, out, s)
}

func Convert_v1alpha3_VehicleStatus_To_v1alpha2_VehicleStatus(in *dwsv1alpha.VehicleStatus, out *VehicleStatus, s apiconversion.Scope) error {
	// Status.Tires was introduced in v1alpha2, so it needs a custom
	// conversion function.  The value will be preserved in an annotation,
	// allowing roundtrip without losing information.

	// The conversion-gen tool printed a warning about this.  Also see the
	// warning it placed in
	// autoConvert_v1alpha3_VehicleStatus_To_v1alpha2_VehicleStatus()
	// in zz_generated.conversion.go.

	// The conversion-gen tool creates all the parts, but in this case it
	// omitted Convert_v1alpha3_VehicleStatus_To_v1alpha2_VehicleStatus(),
	// forcing us to acknowledge that we are handling the conversion for
	// Status.Tires.

	return autoConvert_v1alpha3_VehicleStatus_To_v1alpha2_VehicleStatus(in, out, s)
}
