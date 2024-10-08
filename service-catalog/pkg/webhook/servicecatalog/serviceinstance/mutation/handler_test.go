/*
Copyright 2019 The Kubernetes Authors.

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

package mutation_test

import (
	"context"
	"fmt"
	"testing"

	sc "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	scfeatures "github.com/kubernetes-sigs/service-catalog/pkg/features"
	"github.com/kubernetes-sigs/service-catalog/pkg/webhook/servicecatalog/serviceinstance/mutation"
	"github.com/kubernetes-sigs/service-catalog/pkg/webhookutil/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gomodules.xyz/jsonpatch/v2"
	admissionv1 "k8s.io/api/admission/v1"
	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func TestCreateUpdateHandlerHandleCreateSuccess(t *testing.T) {
	const fixUUID = "mocked-uuid-123-abc"
	tests := map[string]struct {
		givenRawObj []byte
		expPatches  []jsonpatch.Operation
	}{
		"Should set all default fields": {
			givenRawObj: []byte(`{
  				"apiVersion": "servicecatalog.k8s.io/v1",
  				"kind": "ServiceInstance",
  				"metadata": {
  				  "creationTimestamp": null,
  				  "name": "test-instance"
  				},
  				"spec": {
				  "updateRequests": 1,
				  "clusterServiceClassExternalName": "some-class",
				  "clusterServicePlanExternalName": "some-plan" 
  				}
			}`),
			expPatches: []jsonpatch.Operation{
				{
					Operation: "add",
					Path:      "/metadata/finalizers",
					Value: []interface{}{
						"kubernetes-incubator/service-catalog",
					},
				},
				{
					Operation: "add",
					Path:      "/spec/externalID",
					Value:     fixUUID,
				},
			},
		},
		"Should omit externalID and secretName if they are already set": {
			givenRawObj: []byte(`{
				"apiVersion": "servicecatalog.k8s.io/v1",
  				"kind": "ServiceInstance",
  				"metadata": {
  				  "creationTimestamp": null,
  				  "name": "test-instance"
  				},
  				"spec": {
				  "updateRequests": 1,
				  "clusterServiceClassExternalName": "some-class",
				  "clusterServicePlanExternalName": "some-plan",
				  "externalID": "my-external-id-123"
  				}
			}`),
			expPatches: []jsonpatch.Operation{
				{
					Operation: "add",
					Path:      "/metadata/finalizers",
					Value: []interface{}{
						"kubernetes-incubator/service-catalog",
					},
				},
			},
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			// given
			sc.AddToScheme(scheme.Scheme)
			decoder, err := admission.NewDecoder(scheme.Scheme)
			require.NoError(t, err)

			err = utilfeature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.OriginatingIdentity))
			require.NoError(t, err, "cannot disable OriginatingIdentity feature")
			// restore default state
			defer utilfeature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.OriginatingIdentity))

			fixReq := admission.Request{
				AdmissionRequest: admissionv1.AdmissionRequest{
					Operation: admissionv1.Create,
					Name:      "test-instance",
					Namespace: "system",
					Kind: metav1.GroupVersionKind{
						Kind:    "ServiceInstance",
						Version: "v1",
						Group:   "servicecatalog.k8s.io",
					},
					Object: runtime.RawExtension{Raw: tc.givenRawObj},
				},
			}

			handler := mutation.CreateUpdateHandler{
				UUID: func() types.UID { return fixUUID },
			}
			handler.InjectDecoder(decoder)

			// when
			resp := handler.Handle(context.Background(), fixReq)

			// then
			assert.True(t, resp.Allowed)
			require.NotNil(t, resp.PatchType)
			assert.Equal(t, admissionv1.PatchTypeJSONPatch, *resp.PatchType)

			// filtering out status cause k8s api-server will discard this too
			patches := tester.FilterOutStatusPatch(resp.Patches)

			require.Len(t, patches, len(tc.expPatches))
			for _, expPatch := range tc.expPatches {
				assert.Contains(t, patches, expPatch)
			}
		})
	}
}

func TestCreateUpdateHandlerHandleUpdateSuccess(t *testing.T) {
	const fixUUID = "mocked-uuid-123-abc"
	tests := map[string]struct {
		givenOldRawObj []byte
		givenNewRawObj []byte
		expPatches     []jsonpatch.Operation
	}{
		"Should clear out the ClusterServicePlanRef": {
			givenOldRawObj: []byte(`{
  				"apiVersion": "servicecatalog.k8s.io/v1",
  				"kind": "ServiceInstance",
  				"metadata": {
  				  "creationTimestamp": null,
  				  "name": "test-instance",
                  "generation": 1
  				},
  				"spec": {
				  "updateRequests": 1,
				  "clusterServiceClassExternalName": "some-class",
				  "clusterServicePlanExternalName": "some-plan",
                  "clusterServicePlanRef": {"name": "plan-01"},
                  "externalID": "external-id-001"
  				}
			}`),
			givenNewRawObj: []byte(`{
  				"apiVersion": "servicecatalog.k8s.io/v1",
  				"kind": "ServiceInstance",
  				"metadata": {
  				  "creationTimestamp": null,
  				  "name": "test-instance",
				  "generation": 1
  				},
  				"spec": {
				  "updateRequests": 1,
				  "clusterServiceClassExternalName": "some-class",
				  "clusterServicePlanExternalName": "some-plan-changed",
				  "clusterServicePlanRef": {"name": "plan-01"},
				  "externalID": "external-id-001"
  				}
			}`),
			expPatches: []jsonpatch.Operation{
				{
					Operation: "remove",
					Path:      "/spec/clusterServicePlanRef",
				},
			},
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			// given
			sc.AddToScheme(scheme.Scheme)
			decoder, err := admission.NewDecoder(scheme.Scheme)
			require.NoError(t, err)

			err = utilfeature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.OriginatingIdentity))
			require.NoError(t, err, "cannot disable OriginatingIdentity feature")
			// restore default state
			defer utilfeature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.OriginatingIdentity))

			fixReq := admission.Request{
				AdmissionRequest: admissionv1.AdmissionRequest{
					Operation: admissionv1.Update,
					Name:      "test-instance",
					Namespace: "system",
					Kind: metav1.GroupVersionKind{
						Kind:    "ServiceInstance",
						Version: "v1",
						Group:   "servicecatalog.k8s.io",
					},
					OldObject: runtime.RawExtension{Raw: tc.givenOldRawObj},
					Object:    runtime.RawExtension{Raw: tc.givenNewRawObj},
				},
			}

			handler := mutation.CreateUpdateHandler{
				UUID: func() types.UID { return fixUUID },
			}
			handler.InjectDecoder(decoder)

			// when
			resp := handler.Handle(context.Background(), fixReq)

			// then
			assert.True(t, resp.Allowed)
			require.NotNil(t, resp.PatchType)
			assert.Equal(t, admissionv1.PatchTypeJSONPatch, *resp.PatchType)

			// filtering out status cause k8s api-server will discard this too
			patches := tester.FilterOutStatusPatch(resp.Patches)

			require.Len(t, patches, len(tc.expPatches))
			for _, expPatch := range tc.expPatches {
				assert.Contains(t, patches, expPatch)
			}
		})
	}
}

func TestCreateUpdateHandlerHandleSetUserInfoIfOriginatingIdentityIsEnabled(t *testing.T) {
	// given
	sc.AddToScheme(scheme.Scheme)
	decoder, err := admission.NewDecoder(scheme.Scheme)
	require.NoError(t, err)

	// assumption that OriginatingIdentity feature is enabled by default

	reqUserInfo := authenticationv1.UserInfo{
		Username: "minikube",
		UID:      "123",
		Groups:   []string{"unauthorized"},
		Extra: map[string]authenticationv1.ExtraValue{
			"extra": {"val1", "val2"},
		},
	}

	fixReq := admission.Request{
		AdmissionRequest: admissionv1.AdmissionRequest{
			Operation: admissionv1.Create,
			Name:      "test-instance",
			Namespace: "system",
			Kind: metav1.GroupVersionKind{
				Kind:    "ServiceInstance",
				Version: "v1",
				Group:   "servicecatalog.k8s.io",
			},
			UserInfo: reqUserInfo,
			Object: runtime.RawExtension{Raw: []byte(`{
  				"apiVersion": "servicecatalog.k8s.io/v1",
  				"kind": "ServiceInstance",
  				"metadata": {
				  "finalizers": [ "kubernetes-incubator/service-catalog" ],
  				  "creationTimestamp": null,
  				  "name": "test-instance"
  				},
  				"spec": {
				  "updateRequests": 1,
				  "clusterServiceClassExternalName": "some-class",
				  "clusterServicePlanExternalName": "some-plan",
				  "externalID": "my-external-id-123"
  				}
			}`)},
		},
	}

	expPatches := []jsonpatch.Operation{
		{
			Operation: "add",
			Path:      "/spec/userInfo",
			Value: map[string]interface{}{
				"username": "minikube",
				"uid":      "123",
				"groups": []interface{}{
					"unauthorized",
				},
				"extra": map[string]interface{}{
					"extra": []interface{}{
						"val1", "val2",
					},
				},
			},
		},
	}

	handler := mutation.CreateUpdateHandler{}
	handler.InjectDecoder(decoder)

	// when
	resp := handler.Handle(context.Background(), fixReq)

	// then
	assert.True(t, resp.Allowed)
	require.NotNil(t, resp.PatchType)
	assert.Equal(t, admissionv1.PatchTypeJSONPatch, *resp.PatchType)

	// filtering out status cause k8s api-server will discard this too
	patches := tester.FilterOutStatusPatch(resp.Patches)

	require.Len(t, patches, len(expPatches))
	for _, expPatch := range expPatches {
		assert.Contains(t, patches, expPatch)
	}
}

func TestCreateUpdateHandlerHandleDecoderErrors(t *testing.T) {
	tester.DiscardLoggedMsg()

	for _, fn := range []func(t *testing.T, handler tester.TestDecoderHandler, kind string){
		tester.AssertHandlerReturnErrorIfReqObjIsMalformed,
		tester.AssertHandlerReturnErrorIfGVKMismatch,
	} {
		handler := mutation.CreateUpdateHandler{}
		fn(t, &handler, "ServiceInstance")
	}
}
