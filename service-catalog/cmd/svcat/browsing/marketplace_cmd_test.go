/*
Copyright 2018 The Kubernetes Authors.

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

package browsing_test

import (
	"bytes"

	. "github.com/kubernetes-sigs/service-catalog/cmd/svcat/browsing"
	"github.com/kubernetes-sigs/service-catalog/cmd/svcat/command"
	svcattest "github.com/kubernetes-sigs/service-catalog/cmd/svcat/test"
	_ "github.com/kubernetes-sigs/service-catalog/internal/test"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/pkg/svcat"
	servicecatalog "github.com/kubernetes-sigs/service-catalog/pkg/svcat/service-catalog"
	servicecatalogfakes "github.com/kubernetes-sigs/service-catalog/pkg/svcat/service-catalog/service-catalogfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Register Command", func() {
	Describe("NewMarketplaceCmd", func() {
		It("Builds and returns a cobra command with the correct flags", func() {
			cxt := &command.Context{}
			cmd := NewMarketplaceCmd(cxt)
			Expect(*cmd).NotTo(BeNil())

			Expect(cmd.Use).To(Equal("marketplace"))
			Expect(cmd.Short).To(ContainSubstring("List available service offerings"))
			Expect(cmd.Example).To(ContainSubstring("svcat marketplace --namespace dev"))
			Expect(cmd.Aliases).To(ConsistOf("marketplace", "mp"))

			urlFlag := cmd.Flags().Lookup("namespace")
			Expect(urlFlag).NotTo(BeNil())
			Expect(urlFlag.Usage).To(ContainSubstring("If present, the namespace scope for this request"))
		})
	})
	Describe("Validate", func() {
	})
	Describe("Run", func() {
		It("Calls the pkg/svcat libs methods to retrieve all classes and plans and prints output to the user", func() {
			namespace := "banana"

			className := "foobarclass"
			classGUID := "abc123"
			classDescription := "This class foobars"
			className2 := "barbazclass"
			classGUID2 := "qwerty456"
			classDescription2 := "This class barbazs"
			planName := "foobarplan1"
			planGUID := "banana52"
			planName2 := "foobarplan2"
			planGUID2 := "banana53"
			planName3 := "barbazplan"
			planGUID3 := "banana54"
			classToReturn := &v1.ClusterServiceClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      classGUID,
				},
				Spec: v1.ClusterServiceClassSpec{
					CommonServiceClassSpec: v1.CommonServiceClassSpec{
						Description:  classDescription,
						ExternalName: className,
					},
				},
			}
			classToReturn2 := &v1.ClusterServiceClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      classGUID2,
				},
				Spec: v1.ClusterServiceClassSpec{
					CommonServiceClassSpec: v1.CommonServiceClassSpec{
						Description:  classDescription2,
						ExternalName: className2,
					},
				},
			}
			planToReturn := &v1.ClusterServicePlan{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      planGUID,
				},
				Spec: v1.ClusterServicePlanSpec{
					CommonServicePlanSpec: v1.CommonServicePlanSpec{
						ExternalName: planName,
					},
					ClusterServiceClassRef: v1.ClusterObjectReference{
						Name: classGUID,
					},
				},
			}
			planToReturn2 := &v1.ClusterServicePlan{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      planGUID2,
				},
				Spec: v1.ClusterServicePlanSpec{
					CommonServicePlanSpec: v1.CommonServicePlanSpec{
						ExternalName: planName2,
					},
					ClusterServiceClassRef: v1.ClusterObjectReference{
						Name: classGUID,
					},
				},
			}
			planToReturn3 := &v1.ClusterServicePlan{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      planGUID3,
				},
				Spec: v1.ClusterServicePlanSpec{
					CommonServicePlanSpec: v1.CommonServicePlanSpec{
						ExternalName: planName3,
					},
					ClusterServiceClassRef: v1.ClusterObjectReference{
						Name: classGUID2,
					},
				},
			}

			outputBuffer := &bytes.Buffer{}
			fakeApp, _ := svcat.NewApp(nil, nil, "default")
			fakeSDK := new(servicecatalogfakes.FakeSvcatClient)
			fakeSDK.RetrieveClassesReturns([]servicecatalog.Class{classToReturn, classToReturn2}, nil)
			fakeSDK.RetrievePlansReturns([]servicecatalog.Plan{planToReturn, planToReturn2, planToReturn3}, nil)
			fakeApp.SvcatClient = fakeSDK
			cmd := MarketplaceCmd{
				Namespaced: &command.Namespaced{Context: svcattest.NewContext(outputBuffer, fakeApp)},
				Formatted:  command.NewFormatted(),
			}
			cmd.Namespace = namespace

			err := cmd.Run()
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeSDK.RetrieveClassesCallCount()).To(Equal(1))
			scopeOpts, brokerFilter := fakeSDK.RetrieveClassesArgsForCall(0)
			Expect(scopeOpts).To(Equal(servicecatalog.ScopeOptions{
				Scope:     servicecatalog.AllScope,
				Namespace: namespace,
			}))
			Expect(brokerFilter).To((Equal("")))

			Expect(fakeSDK.RetrievePlansCallCount()).To(Equal(1))
			class, scopeOpts := fakeSDK.RetrievePlansArgsForCall(0)
			Expect(class).To(Equal(""))
			Expect(scopeOpts).To(Equal(servicecatalog.ScopeOptions{
				Scope:     servicecatalog.AllScope,
				Namespace: namespace,
			}))

			output := outputBuffer.String()
			Expect(output).To(ContainSubstring(className))
			Expect(output).To(ContainSubstring(planName))
			Expect(output).To(ContainSubstring(planName2))
			Expect(output).To(ContainSubstring(classDescription))
			Expect(output).To(ContainSubstring(className2))
			Expect(output).To(ContainSubstring(planName3))
			Expect(output).To(ContainSubstring(classDescription2))
		})
	})
})
