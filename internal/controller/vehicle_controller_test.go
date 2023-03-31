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

	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha2"
)

var _ = Describe("Vehicle Controller Test", func() {

	var (
		vehicle *dwsv1alpha.Vehicle
	)

	BeforeEach(func() {
		id := uuid.NewString()[0:8]
		vehicle = &dwsv1alpha.Vehicle{
			ObjectMeta: metav1.ObjectMeta{
				Name:      id,
				Namespace: corev1.NamespaceDefault,
			},
			Spec: dwsv1alpha.VehicleSpec{
				Make: "Jeep",
			},
		}

		Expect(k8sClient.Create(context.TODO(), vehicle)).To(Succeed())
	})

	It("can create a vehicle", func() {
		Eventually(func(g Gomega) {
			g.Expect(k8sClient.Get(context.TODO(), client.ObjectKeyFromObject(vehicle), vehicle)).To(Succeed())
			g.Expect(vehicle.Status.Make).To(Equal("Jeep"))
			g.Expect(vehicle.Status.Tires).To(Equal("New"))
		}).Should(Succeed())
	})
})
