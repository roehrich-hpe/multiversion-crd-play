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
	dwsv1alpha2 "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha2"
	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
	utilconversion "github.com/roehrich-hpe/multiversion-crd-play/github/cluster-api/util/conversion"
)

var _ = Describe("Desert Controller Test", func() {

	var (
		desert *dwsv1alpha.Desert
	)

	BeforeEach(func() {
		id := uuid.NewString()[0:8]
		desert = &dwsv1alpha.Desert{
			ObjectMeta: metav1.ObjectMeta{
				Name:      id,
				Namespace: corev1.NamespaceDefault,
			},
			Spec: dwsv1alpha.DesertSpec{
				Type: "Semiarid",
			},
		}

		Expect(k8sClient.Create(context.TODO(), desert)).To(Succeed())
	})

	It("can create a desert", func() {
		Eventually(func(g Gomega) {
			g.Expect(k8sClient.Get(context.TODO(), client.ObjectKeyFromObject(desert), desert)).To(Succeed())
			g.Expect(desert.Spec.Traveler).To(Equal("Arriving"))
			g.Expect(desert.Spec.Days).To(Equal(5))
			g.Expect(desert.Status.Traveler).To(Equal("Arriving"))
			g.Expect(desert.GetAnnotations()).To(HaveLen(0))
		}).Should(Succeed())
	})

	It("reads a desert hub resource via the v1alpha1 spoke", func() {
		desertV1 := &dwsv1alpha1.Desert{}

		Eventually(func(g Gomega) {
			g.Expect(k8sClient.Get(context.TODO(), client.ObjectKeyFromObject(desert), desertV1)).To(Succeed())
			// Don't get deep into verifying the conversion.
			// We have api/v1alpha1/conversion_test.go that is
			// digging deep.
			anno := desertV1.GetAnnotations()
			g.Expect(anno).To(HaveLen(1))
			g.Expect(anno).Should(HaveKey(utilconversion.DataAnnotation))
		}).Should(Succeed())
	})

	It("reads a desert hub resource via the v1alpha2 spoke", func() {
		desertV2 := &dwsv1alpha2.Desert{}

		Eventually(func(g Gomega) {
			g.Expect(k8sClient.Get(context.TODO(), client.ObjectKeyFromObject(desert), desertV2)).To(Succeed())
			// Don't get deep into verifying the conversion.
			// We have api/v1alpha2/conversion_test.go that is
			// digging deep.
			anno := desertV2.GetAnnotations()
			g.Expect(anno).To(HaveLen(1))
			g.Expect(anno).Should(HaveKey(utilconversion.DataAnnotation))
		}).Should(Succeed())
	})
})
