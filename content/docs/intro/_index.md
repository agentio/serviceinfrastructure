---
weight: 1
title: Managing APIs
bookFlatSection: true
type: docs
---
# Managing APIs

[API management](/docs/glossary#api-management) is a thriving industry with many vendors, marketers, and definitions. Here we step back to the perspective of a software developer and say that API management _addresses the common concerns that arise when developers make and use networked APIs_.

## Networked APIs

Networked APIs are APIs that operate across computer networks. Networked APIs communicate using HTTP and similar protocols, and they are frequently produced by organizations that are different from the ones that consume them. This brings many technical and social concerns.

## Common API Management Concerns

For people making APIs, these concerns usually include:
* controlling who uses an API
* controlling how much an API is used
* monitoring the usage and performance of an API

People using APIs have similar concerns from the other side:
* getting access to an API, usually using credentials
* using an API without exceeding limits set by the producer
* monitoring the availability and performance of APIs that they are using

These are _common concerns_ because nearly all API producers and consumers have them. Developers often start by addressing these concerns in their service and client implementations, but soon they begin moving their solutions into libraries and frameworks, and eventually they turn to solutions outside of their code. Typically these external solutions run "out-of-process" and are part of the operating system or environment where the service or client runs.

## API Proxies and Gateways

Most frequently, the out-of-process part of the solution is called an _API proxy_. It is a _proxy_ because it represents someone else. When it represents someone who is calling an API, it is called a _forward proxy_. When it represents someone who is serving an API, it is called a _reverse proxy_. In other words, forward proxies help people make requests and reverse proxies help people serve requests.

```goat
.------------.     .---------------.                  .---------------.     .------------.
|            |     |               |                  |               |     |            |
| API client +---->| forward proxy +----------------->| reverse proxy +---->| API server |
|            |     |               |   API Request    |               |     |            |
.------------.     .---------------.                  .---------------.     .------------.
```

Most of the API management industry has focused on providing reverse proxies to organizations that make and offer APIs. Reverse proxies have evolved into _API gateways_ that provide powerful and sometimes sophisticated features including authentication, rate limiting, and monitoring. This is the focus of Service Infrastructure and our discussion here.

---
#### Continue with [Managing gRPC APIs](/docs/intro/grpc).