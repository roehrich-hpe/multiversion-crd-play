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
	"fmt"

	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var desertlog = logf.Log.WithName("desert-v1alpha1")
var DaysAnnotation = "dws.cray.hpe.com/days"
var ToolAnnotation = "dws.cray.hpe.com/tool"

func (src *Desert) ConvertTo(dstRaw conversion.Hub) error {
	desertlog.Info("Convert To Hub")
	dst := dstRaw.(*dwsv1alpha.Desert)

	dst.ObjectMeta = src.ObjectMeta

	dst.Spec.Foo = src.Spec.Foo
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Traveler = src.Spec.Traveler

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	// If the down-rev resource has been holding Spec.Days in an
	// annotation, then copy it into the correct field in the hub.
	// Same for Spec.Tool.
	annotations := src.GetAnnotations()
	dayData, dayOk := annotations[DaysAnnotation]
	toolData, toolOk := annotations[ToolAnnotation]
	if !dayOk && !toolOk {
		// no days or tool values to preserve
		return nil
	}
	if dayOk {
		days := 0
		_, err := fmt.Sscanf(dayData, "%d", &days)
		if err != nil {
			desertlog.Info("unable to convert days", "%v", err)
			return err
		}
		dst.Spec.Days = days
		// Delete the annotation, so it isn't carried to the hub.
		delete(annotations, DaysAnnotation)
	}
	if toolOk {
		dst.Spec.Tool = toolData
		// Delete the annotation, so it isn't carried to the hub.
		delete(annotations, ToolAnnotation)
	}
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

	dst.Status.Traveler = src.Status.Traveler
	dst.Status.WaterLevel = src.Status.WaterLevel

	// Save the hub's Spec.Days in an annotation on the down-rev resource.
	// Same for Spec.Tool.
	annotations := dst.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[DaysAnnotation] = fmt.Sprintf("%d", src.Spec.Days)
	annotations[ToolAnnotation] = src.Spec.Tool
	dst.SetAnnotations(annotations)

	return nil
}
