---
weight: 3
title: Better Configuration
---
# Better Configuration

Since our main focus has been gRPC API management, we haven't spent much time discussing Endpoints OpenAPI support, but we've noted that it is limited to OpenAPI 2.0. Updating this to later versions of OpenAPI is [probably the most-requested improvement to Endpoints](https://github.com/GoogleCloudPlatform/esp-v2/issues/500).

This is essentially a service compilation task. The backend Service Management APIs only support OpenAPI 2.0, but we can observe the service config that is generated for OpenAPI 2.0 APIs and generate our own equivalent service config that we upload directly. Is OpenAPI 3.x support still interesting to you? If so, let us know in the [issues](https://github.com/agentio/serviceinfrastructure/issues).