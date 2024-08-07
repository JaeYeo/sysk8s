/*
Copyright The Kubernetes Authors.

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

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_Job = map[string]string{
	"":         "Job represents the configuration of a single job.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
	"status":   "Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (Job) SwaggerDoc() map[string]string {
	return map_Job
}

var map_JobCondition = map[string]string{
	"":                   "JobCondition describes current state of a job.",
	"type":               "Type of job condition, Complete or Failed.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastProbeTime":      "Last time the condition was checked.",
	"lastTransitionTime": "Last time the condition transit from one status to another.",
	"reason":             "(brief) reason for the condition's last transition.",
	"message":            "Human readable message indicating details about last transition.",
}

func (JobCondition) SwaggerDoc() map[string]string {
	return map_JobCondition
}

var map_JobList = map[string]string{
	"":         "JobList is a collection of jobs.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of Jobs.",
}

func (JobList) SwaggerDoc() map[string]string {
	return map_JobList
}

var map_JobSpec = map[string]string{
	"":                        "JobSpec describes how the job execution will look like.",
	"parallelism":             "Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
	"completions":             "Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
	"activeDeadlineSeconds":   "Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer",
	"backoffLimit":            "Specifies the number of retries before marking this job failed. Defaults to 6",
	"selector":                "A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
	"manualSelector":          "manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector",
	"template":                "Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
	"ttlSecondsAfterFinished": "ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.",
}

func (JobSpec) SwaggerDoc() map[string]string {
	return map_JobSpec
}

var map_JobStatus = map[string]string{
	"":               "JobStatus represents the current state of a Job.",
	"conditions":     "The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
	"startTime":      "Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"completionTime": "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"active":         "The number of actively running pods.",
	"succeeded":      "The number of pods which reached phase Succeeded.",
	"failed":         "The number of pods which reached phase Failed.",
}

func (JobStatus) SwaggerDoc() map[string]string {
	return map_JobStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
