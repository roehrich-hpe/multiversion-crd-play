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

package controller

import (
	"context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	dwsv1alpha1 "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha1"
	//dwsv1alpha2 "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha2"
	dwsv1alphaHub "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
)

var _ = Describe("Desert Conversion Test", func() {

	var (
		desertHub *dwsv1alphaHub.Desert
	)

	BeforeEach(func() {
		id := uuid.NewString()[0:8]
		desertHub = &dwsv1alphaHub.Desert{
			ObjectMeta: metav1.ObjectMeta{
				Name:      id,
				Namespace: corev1.NamespaceDefault,
			},
			Spec: dwsv1alphaHub.DesertSpec{
				Type: "Semiarid",
				Days: 42,
				Tool: "Knife",
			},
		}

		Expect(k8sClient.Create(context.TODO(), desertHub)).To(Succeed())
	})

	PIt("reads a desert hub resource via the v1alpha1 spoke", func() {
		wantedKeysV1 := map[string]string{
			dwsv1alpha1.DaysAnnotation: "42",
			dwsv1alpha1.ToolAnnotation: "Knife",
		}
		desertV1 := &dwsv1alpha1.Desert{}

		// XXX
		// This fails because, while the conversion webhook is registered,
		// the CRDs do not have the conversion webhook configuration.
		// See the notes in envtest.Environment{} in suite_test.go.

		Eventually(func(g Gomega) map[string]string {
			g.Expect(k8sClient.Get(context.TODO(), client.ObjectKeyFromObject(desertHub), desertV1)).To(Succeed())
			return desertV1.GetAnnotations()
		}).Should(ContainElements(wantedKeysV1))
	})

})
