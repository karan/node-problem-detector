/*
Copyright 2019 The Kubernetes Authors All rights reserved.

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

package config

import (
	"reflect"
	"testing"

	"k8s.io/node-problem-detector/pkg/exporters/gkestackdriver/gke"
)

func TestApplyConfiguration(t *testing.T) {
	testCases := []struct {
		name           string
		originalConfig GKEStackdriverExporterConfig
		wantedConfig   GKEStackdriverExporterConfig
	}{
		{
			name: "normal",
			originalConfig: GKEStackdriverExporterConfig{
				ExportPeriod:          "60s",
				MetadataFetchTimeout:  "600s",
				MetadataFetchInterval: "10s",
				APIEndpoint:           "monitoring.googleapis.com:443",
				GKEMetadata: gke.Metadata{
					ProjectID:     "some-gcp-project",
					Zone:          "us-central1-a",
					InstanceID:    "56781234",
					InstanceName:  "some-gce-instance",
					ClusterName:   "foo-cluster",
					ClusterHash:   "a1b2c3d4",
					NodePoolName:  "cl1m4te-ch4ng3",
					OSVersion:     "cos-85-1234",
					KernelVersion: "5.4.0",
				},
			},
			wantedConfig: GKEStackdriverExporterConfig{
				ExportPeriod:          "60s",
				MetadataFetchTimeout:  "600s",
				MetadataFetchInterval: "10s",
				APIEndpoint:           defaultEndpoint,
				GKEMetadata: gke.Metadata{
					ProjectID:     "some-gcp-project",
					Zone:          "us-central1-a",
					InstanceID:    "56781234",
					InstanceName:  "some-gce-instance",
					ClusterName:   "foo-cluster",
					ClusterHash:   "a1b2c3d4",
					NodePoolName:  "cl1m4te-ch4ng3",
					OSVersion:     "cos-85-1234",
					KernelVersion: "5.4.0",
				},
			},
		},
		{
			name: "staging API endpoint",
			originalConfig: GKEStackdriverExporterConfig{
				ExportPeriod:          "60s",
				MetadataFetchTimeout:  "600s",
				MetadataFetchInterval: "10s",
				APIEndpoint:           "staging-monitoring.sandbox.googleapis.com:443",
				GKEMetadata: gke.Metadata{
					ProjectID:     "some-gcp-project",
					Zone:          "us-central1-a",
					InstanceID:    "56781234",
					InstanceName:  "some-gce-instance",
					ClusterName:   "foo-cluster",
					ClusterHash:   "a1b2c3d4",
					NodePoolName:  "cl1m4te-ch4ng3",
					OSVersion:     "cos-85-1234",
					KernelVersion: "5.4.0",
				},
			},
			wantedConfig: GKEStackdriverExporterConfig{
				ExportPeriod:          "60s",
				MetadataFetchTimeout:  "600s",
				MetadataFetchInterval: "10s",
				APIEndpoint:           "staging-monitoring.sandbox.googleapis.com:443",
				GKEMetadata: gke.Metadata{
					ProjectID:     "some-gcp-project",
					Zone:          "us-central1-a",
					InstanceID:    "56781234",
					InstanceName:  "some-gce-instance",
					ClusterName:   "foo-cluster",
					ClusterHash:   "a1b2c3d4",
					NodePoolName:  "cl1m4te-ch4ng3",
					OSVersion:     "cos-85-1234",
					KernelVersion: "5.4.0",
				},
			},
		},
		{
			name:           "empty",
			originalConfig: GKEStackdriverExporterConfig{},
			wantedConfig: GKEStackdriverExporterConfig{
				ExportPeriod:          "1m0s",
				MetadataFetchTimeout:  "10m0s",
				MetadataFetchInterval: "10s",
				APIEndpoint:           "monitoring.googleapis.com:443",
				GKEMetadata:           gke.Metadata{},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.originalConfig.ApplyConfiguration()
			if !reflect.DeepEqual(test.originalConfig, test.wantedConfig) {
				t.Errorf("Wanted: %+v. \nGot: %+v", test.wantedConfig, test.originalConfig)
			}
		})
	}
}
