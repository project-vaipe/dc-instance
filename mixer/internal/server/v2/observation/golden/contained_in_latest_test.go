// Copyright 2022 Google LLC
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

func TestFetchContainInLatest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "contained_in_latest")

	testSuite := func(mixer pbs.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			variables        []string
			entityExpression string
			goldenFile       string
		}{
			{
				[]string{
					"dummy",
					"Count_Person",
					"Median_Age_Person",
					"NumberOfMonths_WetBulbTemperature_35COrMore_RCP45_MinRelativeHumidity",
				},
				"geoId/06<-containedInPlace+{typeOf:County}",
				"CA_County.json",
			},
			{
				[]string{"Median_Age_Person"},
				"country/USA<-containedInPlace+{typeOf:County}",
				"USA_County.json",
			},
			{
				[]string{
					"Count_Person_Employed",
					"Annual_Generation_Electricity",
					"UnemploymentRate_Person",
					"Count_Person_FoodInsecure",
				},
				"country/USA<-containedInPlace+{typeOf:State}",
				"US_State.json",
			},
			{
				[]string{
					"Count_Person_Literate",
				},
				"country/IND<-containedInPlace+{typeOf:AdministrativeArea1}",
				"IND_AA1.json",
			},
			{
				[]string{"Count_Person"},
				"Earth<-containedInPlace+{typeOf:Country}",
				"Country.json",
			},
			{
				[]string{"Max_Temperature_RCP45"},
				"geoId/06085<-containedInPlace+{typeOf:City}",
				"max_temprature.json",
			},
			{
				[]string{"Annual_Emissions_GreenhouseGas_NonBiogenic"},
				"geoId/06<-containedInPlace+{typeOf:EpaReportingFacility}",
				"epa_facility.json",
			},
			{
				[]string{"Count_Person"},
				"nuts/FR412<-containedInPlace+{typeOf:AdministrativeArea4}",
				"read_series_bt.json",
			},
			{
				[]string{"Count_Person"},
				"country/FRA<-containedInPlace+{typeOf:DummyType}",
				"dummy_type.json",
			},
		} {
			goldenFile := c.goldenFile
			resp, err := mixer.V2Observation(ctx, &pbv2.ObservationRequest{
				Select:   []string{"variable", "entity", "date", "value"},
				Variable: &pbv2.DcidOrExpression{Dcids: c.variables},
				Entity:   &pbv2.DcidOrExpression{Expression: c.entityExpression},
				Date:     "LATEST",
			})
			if err != nil {
				t.Errorf("could not run V2Observation (contained_in): %s", err)
				continue
			}
			if latencyTest {
				continue
			}
			if test.GenerateGolden {
				test.UpdateGolden(resp, goldenPath, goldenFile)
				continue
			}
			var expected pbv2.ObservationResponse
			if err = test.ReadJSON(goldenPath, goldenFile, &expected); err != nil {
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
		"FetchContainInLatest",
		&test.TestOption{UseMemdb: true},
		testSuite,
	); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
