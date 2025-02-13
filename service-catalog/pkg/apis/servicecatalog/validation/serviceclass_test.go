/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	servicecatalog "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
)

func validClusterServiceClass() *servicecatalog.ClusterServiceClass {
	return &servicecatalog.ClusterServiceClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-clusterserviceclass",
		},
		Spec: servicecatalog.ClusterServiceClassSpec{
			CommonServiceClassSpec: servicecatalog.CommonServiceClassSpec{
				Bindable:     true,
				ExternalName: "test-clusterserviceclass",
				ExternalID:   "1234-4354a-49b",
				Description:  "service description",
			},
			ClusterServiceBrokerName: "test-clusterservicebroker",
		},
	}
}

func TestValidateClusterServiceClass(t *testing.T) {
	cases := []struct {
		name         string
		serviceClass *servicecatalog.ClusterServiceClass
		valid        bool
	}{
		{
			name:         "valid serviceClass",
			serviceClass: validClusterServiceClass(),
			valid:        true,
		},
		{
			name: "valid serviceClass - uppercase in externalID",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalID = "40D-0983-1b89"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - period in externalID",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalID = "4315f5e1-0139-4ecf-9706-9df0aff33e5a.plan-name"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - underscore in ExternalID",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalID = "4315f5e1-0139-4ecf-9706-9df0aff33e5a_plan-name"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - period in externalName",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalName = "abc.com"
				return s
			}(),
			valid: true,
		},
		{
			name: "invalid serviceClass - has namespace",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Namespace = "test-ns"
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - empty externalID",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalID = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - missing description",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.Description = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - invalid externalName",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalName = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - missing externalName",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalName = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - valid but weird externalName1",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalName = "-"
				return s
			}(),
			valid: true,
		},
		{
			name: "invalid serviceClass - valid but weird externalName2",
			serviceClass: func() *servicecatalog.ClusterServiceClass {
				s := validClusterServiceClass()
				s.Spec.ExternalName = "0"
				return s
			}(),
			valid: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := ValidateClusterServiceClass(tc.serviceClass)
			if len(errs) != 0 && tc.valid {
				t.Errorf("unexpected error: %v", errs)
			} else if len(errs) == 0 && !tc.valid {
				t.Error("unexpected success")
			}
		})
	}
}

func validServiceClass() *servicecatalog.ServiceClass {
	return &servicecatalog.ServiceClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-clusterserviceclass",
			Namespace: "test-ns",
		},
		Spec: servicecatalog.ServiceClassSpec{
			CommonServiceClassSpec: servicecatalog.CommonServiceClassSpec{
				Bindable:     true,
				ExternalName: "test-clusterserviceclass",
				ExternalID:   "1234-4354a-49b",
				Description:  "service description",
			},
			ServiceBrokerName: "test-clusterservicebroker",
		},
	}
}

func TestValidateServiceClass(t *testing.T) {
	cases := []struct {
		name         string
		serviceClass *servicecatalog.ServiceClass
		valid        bool
	}{
		{
			name:         "valid serviceClass",
			serviceClass: validServiceClass(),
			valid:        true,
		},
		{
			name: "valid serviceClass - uppercase in externalID",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalID = "40D-0983-1b89"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - period in externalID",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalID = "4315f5e1-0139-4ecf-9706-9df0aff33e5a.plan-name"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - period in externalName",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalName = "abc.com"
				return s
			}(),
			valid: true,
		},
		{
			name: "valid serviceClass - underscore in ExternalID",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalID = "4315f5e1-0139-4ecf-9706-9df0aff33e5a_plan-name"
				return s
			}(),
			valid: true,
		},
		{
			name: "invalid serviceClass - has no namespace",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Namespace = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - empty externalID",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalID = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - missing description",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.Description = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - invalid externalName",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalName = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - missing externalName",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalName = ""
				return s
			}(),
			valid: false,
		},
		{
			name: "invalid serviceClass - valid but weird externalName1",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalName = "-"
				return s
			}(),
			valid: true,
		},
		{
			name: "invalid serviceClass - valid but weird externalName2",
			serviceClass: func() *servicecatalog.ServiceClass {
				s := validServiceClass()
				s.Spec.ExternalName = "0"
				return s
			}(),
			valid: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := ValidateServiceClass(tc.serviceClass)
			if len(errs) != 0 && tc.valid {
				t.Errorf("unexpected error: %v", errs)
			} else if len(errs) == 0 && !tc.valid {
				t.Error("unexpected success")
			}
		})
	}
}
