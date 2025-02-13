---
title: Design of the Service Catalog
layout: docwithnav
---

This document covers the architecture and design of Service Catalog.

#### Table of Contents

- [Overview](#overview)
- [Terminology](#terminology)
- [Open Service Broker API](#open-service-broker-api)
- [Service Catalog Design](#service-catalog-design)

## Overview

The Service Catalog integrates with the 
[Open Service Broker API](https://github.com/openservicebrokerapi) (OSB API) to
translate the Kubernetes resource model into OSB API calls to a service broker.

It has the following high level features:

- Service Catalog users can register service brokers with Kubernetes
- Service Catalog can fetch the services and plans (called the catalog) from 
each service broker, and make it available to Kubernetes users
- Kubernetes users can request a new service from a broker by submitting a
`ServiceInstance` resource to Service Catalog
- Kubernetes users can link an application (i.e. one or more `Pod`s) to a 
service by creating a new `ServiceBinding` resource

These features provide a loose-coupling between Applications running in 
Kubernetes and services that they use.

Generally, the services that applications use are external (i.e. "black box")
to the Kubernetes cluster. For example, applications may decide to use a
cloud database service. 

Using Service Catalog and the appropriate service broker, application
developers can focus on their own business logic, leave development and
management of the service to someone else, and leave provisioning to the 
Service Catalog and broker.

## Terminology

- **Application**: Kubernetes uses the term "service" in a different way
  than Service Catalog does, so to avoid confusion the term *Application*
  will refer to the Kubernetes deployment artifact that will use a 
  `ServiceInstance`.
- **ClusterServiceBroker**: a Kubernetes resource that Service Catalog recognizes.
  This resource tells Service Catalog to fetch the catalog of a new broker,
  and merge the new catalog into the existing `ClusterServiceClass`es and 
  `ClusterServicePlan`s (see below). This resource is global - it doesn't have 
  a namespace.
- **ClusterServiceClass**: a Kubernetes resource that Service Catalog generates. After
  a user submits a `ClusterServiceBroker`, Service Catalog fetches the broker's 
  catalog and merges the new catalog entries into `ClusterServiceClass`es in 
  Kubernetes. Each `ClusterServiceClass` has one or more `ClusterServicePlan`s (see below).
  This resource is global - it doesn't have a namespace.
- **ClusterServicePlan**: a Kubernetes resource that Service Catalog generates. After
  a user submits a `ClusterServiceBroker`, Service Catalog fetches the broker's
  catalog and merges the new catalog entries into `ClusterServiceClass`es and 
  `ClusterServicePlan`s. Plans indicate variations on each `ClusterServiceClass` like
  cost, capacity, or quality of services (QoS).
  This resource is global - it doesn't have a namespace.
- **ServiceInstance**: a Kubernetes resource that Service Catalog recognizes.
  A Kubernetes user may submit a `ServiceInstance` to provision a new instance of a 
  service. The details of the service and the plan are listed in the resource, and 
  ServiceCatalog passes that information to the broker that can provision
  the service. This resource is not global - it has a namespace.
- **ServiceBinding**: a Kubernetes resource that Service Catalog recognizes.
  A Kubernetes user may submit a `ServiceBinding` to indicate their intent
  for their application to reference and use a `ServiceInstance`. Generally,
  a `ServiceBinding` will get credentials and a hostname that the application
  can use to talk to the service. After Service Catalog gets these credentials,
  it writes them into a Kubernetes `Secret`.
- **Credentials**: after Service Catalog receives the binding response, it 
  writes the response data into a Kubernetes `Secret`. This is the information
  needed for an application to talk to the service itself, and we call 
  it "Credentials."

## Open Service Broker API

The Service Catalog is an 
[Open Service Broker API](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md) 
(OSB API) client. The OSB API specification is the evolution of
the [Cloud Foundry Service Broker API](https://docs.cloudfoundry.org/services/api.html).

We're not going to detail the OSB API here; for more information, please see
the 
[Open Service Broker API Repository](https://github.com/openservicebrokerapi/servicebroker).

For the rest of this design document, we'll assume that you're familiar with 
the basic concepts of the OSB API.

## Service Catalog Design

![Service Catalog Design](images/sc-design.svg)

The above is the high level architecture of Service Catalog.
Service Catalog has two basic building blocks: a Webhook Server and a controller.


### Webhook Server

The Webhook Server uses [Admission Webhooks](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#what-are-admission-webhooks) 
to manage custom resources. Admission Webhook is a feature available in the Kubernetes API Server, 
that allows you to implement arbitrary control decisions, such as validation or mutation.

For every resource managed by Service Catalog (the ones listed in the previous section), there is a separate handler defined. You can see the structure at 
[`pkg/webhook/servicecatalog`](https://github.com/kubernetes-sigs/service-catalog/blob/master/pkg/webhook/servicecatalog).
The current version of all Service Catalog API resources is `v1`. The resources are defined here:
[`pkg/apis/servicecatalog/v1/types.go`](https://github.com/kubernetes-sigs/service-catalog/blob/master/pkg/apis/servicecatalog/v1/types.go).

If you want to learn more about Admission Webhooks, read [this](https://kubernetes.io/blog/2019/03/21/a-guide-to-kubernetes-admission-controllers/) document.

### Controller

The Service Catalog controller implements the behaviors of the service-catalog 
API. It monitors the API resources (by watching the stream of events from the
API server) and takes the appropriate actions to reconcile the current
state with the user's desired end state.

For example, if a user creates a `ClusterServiceBroker`, the Service Catalog 
controller will pick up the event and request the catalog from the 
broker listed in the resource.

For detailed information on a typical workflow, please see 
[the walkthrough](./walkthrough.md).
