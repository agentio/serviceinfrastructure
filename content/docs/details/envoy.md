---
weight: 5
title: Envoy
---
# Envoy

Envoy is an open source proxy that was designed for modern cloud-based applications. Envoy was created at Lyft and is widely used as both an **Edge Proxy** and a **Service Proxy**.

* An edge proxy is a network proxy that manages traffic into and out of an internal network. Typically that internal network contains multiple running services, and the edge proxy protects all of them by checking and verifying requests coming into the internal network.

* A service proxy is a network proxy that manages traffic into and out of a single service. Typically that service both accepts and makes requests, and all of this traffic goes through the service proxy. The service proxy protects the service by checking and verifying incoming requests, and it assists the service by routing outgoing requests, retrying them if necessary, and sometimes adding credentials.

Envoy is designed and tuned for high performance.

Envoy includes many useful configurable features.

Envoy supports extension mechanisms that allow it to be used to solve many API management problems.





