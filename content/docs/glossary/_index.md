---
weight: 99
title: Service Infrapedia
bookFlatSection: true
type: docs
---
## Service Infrapedia

Here we run through a list of terms that are used in our discussion and provide definitions and extra detail that seems out of scope for the flow but worth mentioning.

### API

Application Programming Interface. ["An API divides and organizes the world of computing tasks in a particular way."](https://www.supremecourt.gov/opinions/20pdf/18-956_d18f.pdf) Recognizing that this term is sometimes defined loosely, we prefer to say that an API is an interface that is distinct from possibly many implementations ([API Servers](#api-server)) which may be running in instances that we call [Services](#service).

### API Client

Software that uses an API. Typically this is a computer program that could be anything from a low level tool like `curl` that directly calls APIs to a high level agent that organizes and makes complex sequences of API calls.

### API Gateway

A [reverse proxy](#reverse-proxy) that provides sophisticated features that might include routing requests to multiple backends, caching, load balancing, security, and [SSL termination](https://en.wikipedia.org/wiki/TLS_termination_proxy). 

### API Management

Discussed in the [Introduction](/docs/intro), API management is the set of activities and interests that arise when APIs are used at scale.

### API Server

Software that implements an API by accepting and responding to API requests. We might find an API server available as source code on GitHub.

### Cloud Endpoints

A Google product that provides [API Gateways](#api-gateway) using [Service Infrastructure](#service-infrapedia) and [Extensible Service Proxies](#extensible-service-proxy).

### Extensible Service Proxy

Either of two reverse proxy implementations available from Google that use [Service Infrastructure](#service-infrastructure) to control access to and monitor the operation of [Services](#service). The "Extensible" adjective suggests that these proxies can be extended to provide additional capabilities, and because they are open source and deployed by [Service Owners](#service-owner), this is theoretically possible, though not yet seen in practice.

### Forward Proxy

Software that provides proxying that represents an [API client](#api-client) by performing activities usually associated with clients in order to simplify or secure the operation of the clients that it represents.

### Google

The third-largest provider of cloud computing services after AWS and Azure.

### Google Cloud Run

A serverless platform that runs software in containers on Google's cloud infrastructure. See [Cloud Run | Google Cloud](https://cloud.google.com/run) for details.

### gRPC

["A high performance, open source universal RPC framework"](https://grpc.io) that was developed at Google and is now owned by the [Cloud Native Computing Foundation](https://cncf.io).

### Kubernetes

An open-source system for running software in containers. Kubernetes was created at Google and now has contributors and users from organizations all over the world. Hosted versions of Kubernetes are offered by nearly every cloud provider. Kubernetes is now owned by the [Cloud Native Computing Foundation](https://cncf.io).

### Protocol Buffers

["A free and open-source cross-platform data format used to serialize structured data."](https://en.wikipedia.org/wiki/Protocol_Buffers) Widely used outside of Google, but unlike gRPC and Kubernetes, Protocol Buffers was not donated to the CNCF and remains the property of Google, which shares it under a [permissive open-source license](https://github.com/protocolbuffers/protobuf/blob/main/LICENSE).

### Reverse Proxy

Software that provides proxying that represents an [API server](#api-server) by performing activities usually associated with servers in order to simplify or secure the operation of the servers that it represents.

### Service

An instance of an [API Server](#api-server) that listens on a network address and accepts and responds to API requests.

### Service Deployment

A deployment of an [API Server](#api-server) that runs on computing infrastructure that might be a dedicated physical server, a virtual machine, a system running managed containers like [Kubernetes](#kubernetes), or a serverless platform like [Google Cloud Run](#google-cloud-run).

### Service Infrastructure

Google's ["foundational platform for creating, managing, securing, and consuming APIs and services across organizations."](https://cloud.google.com/service-infrastructure/docs/overview) Service Infrastructure is not the only API management system that Google offers, but it is essentially the only one that Google uses: virtually all of Google's APIs are managed with Service Infrastructure.

### Service Owner

A person or organization that owns a [service](#service). The service owner is responsible for creating the service, keeping it operational, and may be held legally responsible for the operations of the server.

### Service User

A person or organization that uses a [service](#service). A service user typically uses one or more [API Clients](#api-client) to make requests to a service. A service user often has a business relationship with a [service owner](#service-owner) and is responsible for conforming to terms of use and for paying for usage, when required.