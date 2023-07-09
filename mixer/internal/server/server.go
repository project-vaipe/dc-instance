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

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	pubsub "cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/datacommonsorg/mixer/internal/parser/mcf"
	dcpubsub "github.com/datacommonsorg/mixer/internal/pubsub"
	"github.com/datacommonsorg/mixer/internal/server/resource"
	"github.com/datacommonsorg/mixer/internal/server/statvar"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/store/bigtable"
	"github.com/datacommonsorg/mixer/internal/translator/solver"
	"github.com/datacommonsorg/mixer/internal/translator/types"
	"googlemaps.github.io/maps"
)

const (
	mixerApiSecret       = "mixer-api-key"
	blockListSvgJsonPath = "/datacommons/svg/blocklist_svg.json"
)

func readLatestSecret(projectID, secretID string) (string, error) {
	// Create the context and the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secret manager client: %v", err)
	}

	// Build the request to access the latest secret version.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	// Access the secret version.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access latest secret version: %v", err)
	}

	// Return the secret payload as a string.
	return string(result.Payload.Data), nil
}

// Server holds resources for a mixer server
type Server struct {
	store      *store.Store
	metadata   *resource.Metadata
	cache      *resource.Cache
	mapsClient *maps.Client
	httpClient *http.Client
}

func (s *Server) updateBranchTable(ctx context.Context, branchTableName string) {
	branchTable, err := bigtable.NewBtTable(
		ctx,
		bigtable.BranchBigtableProject,
		bigtable.BranchBigtableInstance,
		branchTableName,
	)
	if err != nil {
		log.Printf("Failed to udpate branch cache Bigtable client: %v", err)
		return
	}
	s.store.BtGroup.UpdateBranchTable(
		bigtable.NewTable(branchTableName, branchTable, false /*isCustom=*/))
	log.Printf("Updated branch table to use %s", branchTableName)
}

// NewMetadata initialize the metadata for translator.
func NewMetadata(
	hostProject,
	bigQueryDataset,
	schemaPath,
	remoteMixerDomain string,
	foldRemoteRootSvg bool,
) (*resource.Metadata, error) {
	_, filename, _, _ := runtime.Caller(0)
	subTypeMap, err := solver.GetSubTypeMap(
		path.Join(path.Dir(filename), "../translator/table_types.json"))
	if err != nil {
		return nil, err
	}
	files, err := os.ReadDir(schemaPath)
	if err != nil {
		return nil, err
	}
	mappings := []*types.Mapping{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".mcf") {
			mappingStr, err := os.ReadFile(filepath.Join(schemaPath, f.Name()))
			if err != nil {
				return nil, err
			}
			mapping, err := mcf.ParseMapping(string(mappingStr), bigQueryDataset)
			if err != nil {
				return nil, err
			}
			mappings = append(mappings, mapping...)
		}
	}
	outArcInfo := map[string]map[string][]*types.OutArcInfo{}
	inArcInfo := map[string][]*types.InArcInfo{}

	var apiKey string
	if remoteMixerDomain != "" {
		apiKey, err = readLatestSecret(hostProject, mixerApiSecret)
		if err != nil {
			return nil, err
		}
	}

	return &resource.Metadata{
			Mappings:          mappings,
			OutArcInfo:        outArcInfo,
			InArcInfo:         inArcInfo,
			SubTypeMap:        subTypeMap,
			HostProject:       hostProject,
			BigQueryDataset:   bigQueryDataset,
			RemoteMixerDomain: remoteMixerDomain,
			RemoteMixerAPIKey: apiKey,
			FoldRemoteRootSvg: foldRemoteRootSvg,
		},
		nil
}

// SubscribeBranchCacheUpdate subscribe for branch cache update.
func (s *Server) SubscribeBranchCacheUpdate(ctx context.Context,
) error {
	return dcpubsub.Subscribe(
		ctx,
		bigtable.BranchBigtableProject,
		bigtable.BranchCacheSubscriberPrefix,
		bigtable.BranchCachePubsubTopic,
		func(ctx context.Context, msg *pubsub.Message) error {
			branchTableName := string(msg.Data)
			log.Printf("branch cache subscriber message received with table name: %s\n", branchTableName)
			s.updateBranchTable(ctx, branchTableName)
			return nil
		},
	)
}

type SearchOptions struct {
	UseSearch           bool
	BuildSvgSearchIndex bool
}

// NewCache initializes the cache for stat var hierarchy.
func NewCache(
	ctx context.Context,
	store *store.Store,
	searchOptions SearchOptions,
) (*resource.Cache, error) {
	var blocklistSvg []string
	// Read blocklisted svg from file.
	file, err := os.ReadFile(blockListSvgJsonPath)
	if err != nil {
		log.Printf("Could not read blocklist svg file. Using empty blocklist svg list.")
		blocklistSvg = []string{}
	} else {
		if err := json.Unmarshal(file, &blocklistSvg); err != nil {
			log.Printf("Could not unmarshal blocklist svg file. Using empty blocklist svg list.")
			blocklistSvg = []string{}
		}
	}
	rawSvg, err := statvar.GetRawSvg(ctx, store)
	if err != nil {
		return nil, err
	}
	parentSvgMap := statvar.BuildParentSvgMap(rawSvg)
	result := &resource.Cache{
		RawSvg:       rawSvg,
		ParentSvg:    parentSvgMap,
		BlockListSvg: map[string]struct{}{},
	}
	for _, svg := range blocklistSvg {
		statvar.RemoveSvg(rawSvg, parentSvgMap, svg)
		result.BlockListSvg[svg] = struct{}{}
	}
	if searchOptions.UseSearch {
		if searchOptions.BuildSvgSearchIndex {
			result.SvgSearchIndex = statvar.BuildStatVarSearchIndex(rawSvg, parentSvgMap, blocklistSvg)
		}
	}
	return result, nil
}

// NewMixerServer creates a new mixer server instance.
func NewMixerServer(
	store *store.Store,
	metadata *resource.Metadata,
	cache *resource.Cache,
	mapsClient *maps.Client,
) *Server {
	return &Server{
		store:      store,
		metadata:   metadata,
		cache:      cache,
		mapsClient: mapsClient,
		httpClient: &http.Client{},
	}
}
