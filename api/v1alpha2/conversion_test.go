/*
 * Copyright 2023 Hewlett Packard Enterprise Development LP
 * Other additional copyright holders may be indicated within.
 *
 * The entirety of this work is licensed under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 *
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha2

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"

	dwsv1alpha "github.com/roehrich-hpe/multiversion-crd-play/api/v1alpha3"
	utilconversion "github.com/roehrich-hpe/multiversion-crd-play/github/cluster-api/util/conversion"
)

func TestFuzzyConversion(t *testing.T) {
	t.Run("for Desert", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &dwsv1alpha.Desert{},
		Spoke: &Desert{},
	}))
	t.Run("for Vehicle", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:   &dwsv1alpha.Vehicle{},
		Spoke: &Vehicle{},
	}))
}

// Just touch ginkgo, so it's here to interpret any ginkgo args from
// "make test", so that doesn't fail on this test file.
var _ = BeforeSuite(func() {})
