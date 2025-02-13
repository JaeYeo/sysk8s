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

package output

import (
	"io"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/pkg/svcat/service-catalog"
)

func getBrokerScope(broker servicecatalog.Broker) string {
	if broker.GetNamespace() != "" {
		return servicecatalog.NamespaceScope
	}
	return servicecatalog.ClusterScope
}

func getBrokerStatusCondition(status v1.CommonServiceBrokerStatus) v1.ServiceBrokerCondition {
	if len(status.Conditions) > 0 {
		return status.Conditions[len(status.Conditions)-1]
	}
	return v1.ServiceBrokerCondition{}
}

func getBrokerStatusShort(status v1.CommonServiceBrokerStatus) string {
	lastCond := getBrokerStatusCondition(status)
	return formatStatusShort(string(lastCond.Type), lastCond.Status, lastCond.Reason)
}

func getBrokerStatusFull(status v1.CommonServiceBrokerStatus) string {
	lastCond := getBrokerStatusCondition(status)
	return formatStatusFull(string(lastCond.Type), lastCond.Status, lastCond.Reason, lastCond.Message, lastCond.LastTransitionTime)
}

func writeBrokerListTable(w io.Writer, brokers []servicecatalog.Broker) {
	t := NewListTable(w)
	t.SetHeader([]string{
		"Name",
		"Namespace",
		"URL",
		"Status",
	})
	for _, broker := range brokers {
		t.Append([]string{
			broker.GetName(),
			broker.GetNamespace(),
			broker.GetURL(),
			getBrokerStatusShort(broker.GetStatus()),
		})
	}
	t.Render()
}

// WriteBrokerList prints a list of brokers in the specified output format.
func WriteBrokerList(w io.Writer, outputFormat string, brokers ...servicecatalog.Broker) {
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, brokers)
	case FormatYAML:
		writeYAML(w, brokers, 0)
	case FormatTable:
		writeBrokerListTable(w, brokers)
	}
}

// WriteBroker prints a broker in the specified output format.
func WriteBroker(w io.Writer, outputFormat string, broker servicecatalog.Broker) {
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, broker)
	case FormatYAML:
		writeYAML(w, broker, 0)
	case FormatTable:
		writeBrokerListTable(w, []servicecatalog.Broker{broker})
	}
}

// WriteBrokerDetails prints details for a single broker.
func WriteBrokerDetails(w io.Writer, broker servicecatalog.Broker) {
	t := NewDetailsTable(w)
	table := [][]string{}
	table = append(table, []string{"Name:", broker.GetName()})
	table = append(table, []string{"Scope:", getBrokerScope(broker)})
	if broker.GetNamespace() != "" {
		table = append(table, []string{"Namespace:", broker.GetNamespace()})
	}
	table = append(table, []string{"URL:", broker.GetURL()})
	table = append(table, []string{"Status:", getBrokerStatusFull(broker.GetStatus())})
	t.AppendBulk(table)
	t.Render()
}
