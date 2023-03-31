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

package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha2"
)

var convertlog = logf.Log.WithName("convert-v1alpha1")

func (src *Desert) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert Desert To Hub")
	dst := dstRaw.(*dwsv1alpha.Desert)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Traveler = src.Spec.Traveler

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	return nil
}

func (dst *Desert) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert Desert From Hub")
	src := srcRaw.(*dwsv1alpha.Desert)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Traveler = src.Spec.Traveler

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	return nil
}

func (src *Vehicle) ConvertTo(dstRaw conversion.Hub) error {
	convertlog.Info("Convert Vehicle To Hub")
	dst := dstRaw.(*dwsv1alpha.Vehicle)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Make = src.Spec.Make

	dst.Status.Make = src.Status.Make

	return nil
}

func (dst *Vehicle) ConvertFrom(srcRaw conversion.Hub) error {
	convertlog.Info("Convert Vehicle From Hub")
	src := srcRaw.(*dwsv1alpha.Vehicle)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Make = src.Spec.Make

	dst.Status.Make = src.Status.Make

	return nil
}
