// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golden

import (
	"context"
	"path"
	"runtime"
	"testing"

	pbs "github.com/datacommonsorg/mixer/internal/proto/service"
	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDerivedSeries(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "derived_series")

	testSuite := func(mixer pbs.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			entity     string
			variable   string
			goldenFile string
		}{
			{
				"geoId/06",
				"Count_Person_Female / Count_Person_Male",
				"case1.json",
			},
			{
				"geoId/06",
				"Count_Person_Female[mm=CensusACS5yrSurvey] / Count_Person_Male[mm=CensusACS5yrSurvey]",
				"case2.json",
			},
			{
				"geoId/06",
				"Count_Person - Count_Person_Female - Count_Person_Male",
				"case3.json",
			},
			{
				"geoId/06",
				"(Count_Person_Female - Count_Person_Male) / Count_Person",
				"case4.json",
			},
		} {
			resp, err := mixer.V2Observation(ctx,
				&pbv2.ObservationRequest{
					Select:   []string{"variable", "entity", "date", "value"},
					Variable: &pbv2.DcidOrExpression{Formula: c.variable},
					Entity:   &pbv2.DcidOrExpression{Dcids: []string{c.entity}},
				})
			if err != nil {
				t.Errorf("could not run mixer.V2Observation: %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			if test.GenerateGolden {
				test.UpdateGolden(resp, goldenPath, c.goldenFile)
				continue
			}
			var expected pbv2.ObservationResponse
			if err = test.ReadJSON(goldenPath, c.goldenFile, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file: %s", err)
				continue
			}

			if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("payload got diff: %v", diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"Observation(DerivedSeries)", &test.TestOption{}, testSuite); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
