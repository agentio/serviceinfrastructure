---
weight: 4
bookFlatSection: true
title: "The Proxies"
---
# The Extensible Service Proxies

The Extensible Service Proxies run alongside API servers and are configured and controlled by [Service Infrastructure](/docs/serviceinfrastructure).

There are two Google-supported versions of these proxies:
1. [ESP](/docs/proxies/esp), a custom proxy based on nginx.
2. [ESPv2](/docs/proxies/espv2), a custom proxy based on [Envoy](/docs/details/envoy).

Recalling our overview discussion, the Extensible Service Proxy sits between API clients and an API server and handles the general-purpose API management needs of the server, allowing the server implementation to focus on its own special features. The proxy isn't fully self-contained, instead it calls out to Service Infrastructure for configuration, access control, and logging. This makes it lightweight and convenient to deploy alongside an API server.

Often the ESP and API server are built into their own containers that are deployed together in a single Kubernetes or Cloud Run deployment. This isn't the only way to use an ESP; it could also be run elsewhere and handle requests for multiple backends. But because it is lightweight, and because we generally like to minimize the network delay between the ESP and the API server, we prefer this [sidecar deployment](/docs/glossary#sidecar-deployment).

```goat

                                                   .------------------------------------------.
                                                   |                                          |
  .-------------.                                  |    .--------------.     .------------.   |
  |             |                                  |    |  Extensible  |     |            |   |  o  Service
  | API clients +----------------------------------+--->|   Service    +---->| API server |   | -+-  Owner
  |             |           API Requests           |    |    Proxy     |     |            |   | / \
  .-------------.                                  |    .---------+----.     .------------.   |
     o   o   o                                     |         ^    |                           |
    -+- -+- -+-                                    |         |    |        Service Deployment |
    / \ / \ / \                                    .---------+----+---------------------------.
   Service Users                                             |    |
                                                             |    |
                                                             |    | Request Handling:
                              Proxy Configuration            |    | - Check
                    +----------------------------------------+    | - Allocate Quota
                    |                                             | - Report
                    |                                             |
                    |                                             v
  +-----------------+---------------------------------------------+-------------------------------+
  |                                                                        Service Infrastructure |
  |                                                                                               | Google
  |                                                                                               |
  +-----------------+--------------------------+------------------+-------------------+-----------+
                    ^                          ^                  |                   |
                    | Service                  | User             |                   |
                    | Configuration            | Authorizations   | Logs              | Metrics
                    |                          |                  v                   v 
                 o      
                -+- Service Owner
                / \   

```