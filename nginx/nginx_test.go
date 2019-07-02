/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nginx

import (
	"github.com/buildpack/libbuildpack/buildplan"
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestUnitNGINX(t *testing.T) {
	spec.Run(t, "NGINX", testNGINX, spec.Report(report.Terminal{}))
}

func testNGINX(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	when("buildpack.yml", func() {
		var f *test.DetectFactory

		it.Before(func() {
			f = test.NewDetectFactory(t)
		})

		it("can load an empty buildpack.yaml", func() {
			test.WriteFile(t, filepath.Join(f.Detect.Application.Root, "buildpack.yml"), "")

			loaded, err := LoadBuildpackYAML(f.Detect.Application.Root)

			Expect(err).To(Succeed())
			Expect(loaded).To(Equal(BuildpackYAML{}))
		})

		it("can load a version & web server", func() {
			yaml := "{'nginx': {'version': 1.0.0}}"
			test.WriteFile(t, filepath.Join(f.Detect.Application.Root, "buildpack.yml"), yaml)

			loaded, err := LoadBuildpackYAML(f.Detect.Application.Root)
			actual := BuildpackYAML{
				Config: Config{
					Version: "1.0.0",
				},
			}

			Expect(err).To(Succeed())
			Expect(loaded).To(Equal(actual))
		})

		it("can load mainline version", func() {
			found := LoadMainlineVersion(buildplan.Metadata{"version-lines": map[string]string{"mainline": "1.0.0"}})
			Expect(found).To(Equal("1.0.0"))
		})

		it("can load stable version", func() {
			found := LoadStableVersion(buildplan.Metadata{"version-lines": map[string]string{"stable": "1.0.0"}})
			Expect(found).To(Equal("1.0.0"))
		})
	})
}