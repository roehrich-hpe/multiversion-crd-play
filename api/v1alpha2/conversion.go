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
	dwsv1alpha1 "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha1" // for names of annotations
	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var desertlog = logf.Log.WithName("desert-v1alpha2")

func (src *Desert) ConvertTo(dstRaw conversion.Hub) error {
	desertlog.Info("Convert To Hub")
	dst := dstRaw.(*dwsv1alpha.Desert)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Traveler = src.Spec.Traveler
	dst.Spec.Days = src.Spec.Days

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	// If the down-rev resource has been holding Spec.Tool in an
	// annotation, then copy it into the correct field in the hub.
	annotations := src.GetAnnotations()
	toolData, toolOk := annotations[dwsv1alpha1.ToolAnnotation]
	if !toolOk {
		// no tool value to preserve
		return nil
	}
	dst.Spec.Tool = toolData
	// Delete the annotation, so it isn't carried to the hub.
	delete(annotations, dwsv1alpha1.ToolAnnotation)
	src.SetAnnotations(annotations)

	return nil
}

func (dst *Desert) ConvertFrom(srcRaw conversion.Hub) error {
	desertlog.Info("Convert From Hub")
	src := srcRaw.(*dwsv1alpha.Desert)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Traveler = src.Spec.Traveler
	dst.Spec.Days = src.Spec.Days

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	// Save the hub's Spec.Tool in an annotation on the down-rev resource.
	annotations := dst.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[dwsv1alpha1.ToolAnnotation] = src.Spec.Tool
	dst.SetAnnotations(annotations)

	return nil
}
