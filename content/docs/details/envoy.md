---
weight: 9
title: The Envoy Proxy
---
## Envoy

[Envoy](https://envoyproxy.io) is an open source proxy that was designed for modern cloud-based applications. Envoy was created at Lyft and is widely used as both an **Edge Proxy** and a **Service Proxy**.

* An **edge proxy** is a network proxy that manages traffic into and out of an internal network. Typically that internal network contains multiple running services, and the edge proxy protects all of them by checking and verifying requests coming into the internal network.

* A **service proxy** is a network proxy that manages traffic into and out of a single service. Typically that service both accepts and makes requests, and all of this traffic goes through the service proxy. The service proxy protects the service by checking and verifying incoming requests, and it assists the service by routing outgoing requests, retrying them if necessary, and sometimes adding credentials.

Envoy has some key characteristics that make it an interesting part of an API proxy solution:
- Envoy is designed and tuned for high performance.
- Envoy includes many useful built-in features, including filters that can be enabled to support gRPC transcoding and JWT verification.
- Envoy supports extension mechanisms including [ext_authz](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter) and [ext_proc](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_proc_filter) that allow it to be almost arbitarily extended to solve many API management problems.

---
#### Go back to [Going Deeper](/docs/details).