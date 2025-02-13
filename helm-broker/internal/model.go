package internal

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/Masterminds/semver"
	"github.com/alecthomas/jsonschema"
	"github.com/fatih/structs"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"
	rafter "github.com/kyma-project/rafter/pkg/apis/rafter/v1"
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	chartv2 "k8s.io/helm/pkg/proto/hapi/chart"
)

// Index contains collection of all addons from the given repository
type Index struct {
	Entries map[AddonName][]IndexEntry `json:"entries"`
}

// IndexEntry contains information about single addon entry
type IndexEntry struct {
	// Name is set to index entry key name
	Name AddonName `json:"-"`
	// DisplayName is the entry name, currently treated by us as DisplayName
	DisplayName string       `json:"name"`
	Description string       `json:"description"`
	Version     AddonVersion `json:"version"`
}

// AddonID is a AddonWithCharts identifier as defined by Open Service Broker API.
type AddonID string

// AddonName is a AddonWithCharts name as defined by Open Service Broker API.
type AddonName string

// AddonVersion is a AddonWithCharts version which is defined in the index file
type AddonVersion string

// AddonPlanID is an identifier of AddonWithCharts plan as defined by Open Service Broker API.
type AddonPlanID string

// AddonPlanName is the name of the AddonWithCharts plan as defined by Open Service Broker API
type AddonPlanName string

// PlanSchemaType describes type of the schema file.
type PlanSchemaType string

// PlanSchema is schema definition used for creating parameters
type PlanSchema jsonschema.Schema

const (
	// SchemaTypeBind represents 'bind' schema plan
	SchemaTypeBind PlanSchemaType = "bind"
	// SchemaTypeProvision represents 'provision' schema plan
	SchemaTypeProvision PlanSchemaType = "provision"
	// SchemaTypeUpdate represents 'update' schema plan
	SchemaTypeUpdate PlanSchemaType = "update"
)

// ChartName is a type expressing name of the chart
type ChartName string

// ChartRef provide reference to addon's chart
type ChartRef struct {
	Name    ChartName
	Version semver.Version
}

// GobDecode is decoding chart info
func (cr *ChartRef) GobDecode(in []byte) error {
	var dto struct {
		Name    ChartName
		Version string
	}

	buf := bytes.NewReader(in)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&dto); err != nil {
		return errors.Wrap(err, "while decoding")
	}

	cr.Name = dto.Name

	ver, _ := semver.NewVersion(dto.Version)
	cr.Version = *ver

	return nil
}

// GobEncode implements GobEncoder for custom encoding
func (cr ChartRef) GobEncode() ([]byte, error) {
	dto := struct {
		Name    ChartName
		Version string
	}{
		Name:    cr.Name,
		Version: cr.Version.String(),
	}

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&dto); err != nil {
		return []byte{}, errors.Wrap(err, "while encoding")
	}

	return buf.Bytes(), nil
}

// ChartValues are used as container for chart's values.
// It's currently populated from yaml file or request parameters.
// TODO: switch to more concrete type
type ChartValues map[string]interface{}

func EmptyChartValues() ChartValues {
	return map[string]interface{}{}
}

// AddonPlanBindTemplate represents template used for helm chart installation
type AddonPlanBindTemplate []byte

// AddonPlan is a container for whole data of addon plan.
// Each addon needs to have at least one plan.
type AddonPlan struct {
	ID           AddonPlanID
	Name         AddonPlanName
	Description  string
	Schemas      map[PlanSchemaType]PlanSchema
	ChartRef     ChartRef
	ChartValues  ChartValues
	Metadata     AddonPlanMetadata
	Bindable     *bool
	Free         *bool
	BindTemplate AddonPlanBindTemplate
}

// AddonPlanMetadata provides metadata of the addon.
type AddonPlanMetadata struct {
	DisplayName string
}

// ToMap function is converting Metadata to format compatible with YAML encoder.
func (b AddonPlanMetadata) ToMap() map[string]interface{} {
	type mapped struct {
		DisplayName string `structs:"displayName"`
	}

	return structs.Map(mapped(b))
}

// AddonTag is a Tag attached to AddonWithCharts.
type AddonTag string

// AddonDocs contains data to create ClusterAssetGroup for every ClusterServiceClass.
type AddonDocs struct {
	Template rafter.CommonAssetGroupSpec
}

// Addon represents addon as defined by OSB API.
type Addon struct {
	ID                  AddonID
	Name                AddonName
	Version             semver.Version
	Description         string
	Plans               map[AddonPlanID]AddonPlan
	Metadata            AddonMetadata
	RepositoryURL       string
	Tags                []AddonTag
	Requires            []string
	Bindable            bool
	BindingsRetrievable bool
	PlanUpdatable       *bool
	Docs                []AddonDocs
	Status              v1alpha1.AddonStatus
	Reason              v1alpha1.AddonStatusReason
	Message             string
	SecretRef           corev1.SecretReference
}

// CommonAddon holds common addon configuration structs
type CommonAddon struct {
	Meta   v1.ObjectMeta
	Spec   v1alpha1.CommonAddonsConfigurationSpec
	Status v1alpha1.CommonAddonsConfigurationStatus
}

// IsReadyForInitialProcessing checks if the object is in the initial state - has never been processed.
func (ca *CommonAddon) IsReadyForInitialProcessing() bool {
	return ca.Status.ObservedGeneration == 0 || ca.Status.Phase == v1alpha1.AddonsConfigurationPending
}

func (ca *CommonAddon) IsReadyForReprocessing() bool {
	if ca.Meta.Generation > ca.Status.ObservedGeneration {
		return true
	}

	for _, r := range ca.Status.Repositories {
		if r.Status == v1alpha1.RepositoryStatusFailed && r.Reason == v1alpha1.RepositoryURLFetchingError {
			return true
		}
	}
	return false
}

// AddonWithCharts aggregates an addon with its chart(s)
type AddonWithCharts struct {
	Addon  *Addon
	Charts []*chart.Chart
}

// IsProvisioningAllowed determines addon can be provision on indicated namespace
// if addon has provisionOnlyOnce flag on true then check if addon already exist in this namespace
func (b Addon) IsProvisioningAllowed(namespace Namespace, collection []*Instance) bool {
	if !b.Metadata.ProvisionOnlyOnce {
		return true
	}

	for _, instance := range collection {
		if namespace != instance.Namespace {
			continue
		}
		if string(b.ID) == string(instance.ServiceID) {
			return false
		}
	}

	return true
}

// Labels are key-value pairs which add metadata information for addon.
type Labels map[string]string

// AddonMetadata holds addon metadata as defined by OSB API.
type AddonMetadata struct {
	DisplayName         string
	ProviderDisplayName string
	LongDescription     string
	DocumentationURL    string
	SupportURL          string
	ProvisionOnlyOnce   bool
	// ImageURL is graphical representation of the addon.
	// Currently SVG is required.
	ImageURL string
	Labels   Labels
}

// ToMap collect data from AddonMetadata to format compatible with YAML encoder.
func (b AddonMetadata) ToMap() map[string]interface{} {
	type mapped struct {
		DisplayName         string `structs:"displayName"`
		ProviderDisplayName string `structs:"providerDisplayName"`
		LongDescription     string `structs:"longDescription"`
		DocumentationURL    string `structs:"documentationUrl"`
		SupportURL          string `structs:"supportUrl"`
		ProvisionOnlyOnce   bool   `structs:"provisionOnlyOnce"`
		ImageURL            string `structs:"imageUrl"`
		Labels              Labels `structs:"labels"`
	}
	return structs.Map(mapped(b))
}

// DeepCopy returns a new AddonMetadata object.
func (b AddonMetadata) DeepCopy() AddonMetadata {
	out := b
	if b.Labels != nil {
		out.Labels = make(Labels, len(b.Labels))
		for k, v := range b.Labels {
			out.Labels[k] = v
		}

	}

	return out
}

// InstanceID is a service instance identifier.
type InstanceID string

// IsZero checks if InstanceID equals zero.
func (id InstanceID) IsZero() bool { return id == InstanceID("") }

// OperationID is used as binding operation identifier.
type OperationID string

// IsZero checks if OperationID equals zero
func (id OperationID) IsZero() bool { return id == OperationID("") }

// InstanceOperation represents single provisioning operation.
type InstanceOperation struct {
	InstanceID             InstanceID
	OperationID            OperationID
	Type                   OperationType
	State                  OperationState
	StateDescription       *string
	ProvisioningParameters *RequestParameters

	// CreatedAt points to creation time of the operation.
	// Field should be treated as immutable and is responsibility of storage implementation.
	// It should be set by storage Insert method.
	CreatedAt time.Time
}

// ReleaseName is the name of the Helm release.
type ReleaseName string

// ServiceID is an ID of the Service exposed via Service Catalog.
type ServiceID string

// IsZero checks if ServiceID equals zero
func (id ServiceID) IsZero() bool { return id == ServiceID("") }

// ServicePlanID is an ID of the Plan of Service exposed via Service Catalog.
type ServicePlanID string

// IsZero checks if ServicePlanID equals zero
func (id ServicePlanID) IsZero() bool { return id == ServicePlanID("") }

// Namespace is the name of namespace in k8s
type Namespace string

// ReleaseInfo contains additional data about release installed on instance provisioning.
type ReleaseInfo struct {
	Time         *google_protobuf.Timestamp
	ReleaseTime  time.Time
	Revision     int
	Config       *chartv2.Config
	ConfigValues map[string]interface{}
}

// RequestParameters wraps a map containing provided YAML with parameters from request
type RequestParameters struct {
	Data map[string]interface{}
}

const (
	// ClusterWide is a value which refers to cluster wide resources.
	ClusterWide Namespace = ""
)

// Instance contains info about Service exposed via Service Catalog.
type Instance struct {
	ID                     InstanceID
	ServiceID              ServiceID
	ServicePlanID          ServicePlanID
	ReleaseName            ReleaseName
	Namespace              Namespace
	ReleaseInfo            ReleaseInfo
	ProvisioningParameters *RequestParameters
	ParamsHash             string
}

// InstanceCredentials are created when we bind a service instance.
type InstanceCredentials map[string]string

// BindingID is used as Service Binding identifier
type BindingID string

// IsZero checks if BindingID equals zero
func (id BindingID) IsZero() bool { return id == BindingID("") }

// BindOperation represents single service binding operation.
type BindOperation struct {
	InstanceID       InstanceID
	BindingID        BindingID
	OperationID      OperationID
	Type             OperationType
	State            OperationState
	StateDescription *string

	// CreatedAt points to creation time of the operation.
	// Field should be treated as immutable and is responsibility of storage implementation.
	// It should be set by storage InsertBindOperation method.
	CreatedAt time.Time
}

// InstanceBindData contains data about service instance and it's credentials.
type InstanceBindData struct {
	InstanceID  InstanceID
	Credentials InstanceCredentials
}

// OperationState defines the possible states of an asynchronous request to a broker.
type OperationState string

// String returns state of the operation.
func (os OperationState) String() string {
	return string(os)
}

const (
	// OperationStateInProgress means that operation is in progress
	OperationStateInProgress OperationState = "in progress"
	// OperationStateSucceeded means that request succeeded
	OperationStateSucceeded OperationState = "succeeded"
	// OperationStateFailed means that request failed
	OperationStateFailed OperationState = "failed"
)

// OperationType defines the possible types of an asynchronous operation to a broker.
type OperationType string

const (
	// OperationTypeCreate means creating OperationType
	OperationTypeCreate OperationType = "create"
	// OperationTypeRemove means removing OperationType
	OperationTypeRemove OperationType = "remove"
	// OperationTypeUndefined means undefined OperationType
	OperationTypeUndefined OperationType = ""
)
