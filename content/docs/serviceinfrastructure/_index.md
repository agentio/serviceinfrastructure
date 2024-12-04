---
weight: 3
title: "Service Infrastructure"
bookFlatSection: false
bookCollapseSection: true
---
# Service Infrastructure

[Service Infrastructure](https://cloud.google.com/service-infrastructure/docs/overview) is described as Google's "foundational platform for creating, managing, securing, and consuming APIs and services across organizations."

```goat

  .-------------.                                  .--------------------.
  |             |                                  |                    |  o  Service
  | API clients +--------------------------------->| Service Deployment | -+-  Owner
  |             |           API Requests           |                    | / \
  .-------------.                                  .--------------------.
     o   o   o                                                ^
    -+- -+- -+-                                               |
    / \ / \ / \                                               |
   Service Users                                              |
                                                              v
  +---------------------------------------------------------------------+
  |                                                                     |
  |                      Service Infrastructure                         | Google
  |                                                                     |
  +----------------------------------------------+------------+---------+
            ^                  ^                 |            |
            | Service          | User            |            |
            | Configuration    | Authorizations  | Logs       | Metrics
            |                  |                 v            v
           o
          -+- Service Owner
          / \

```

Capabilities provided by Service Infrastructure include:
* authentication and authorization
* rate limiting
* logging, monitoring, analytics, and auditing

Virtually all of Google's APIs are built on Service Infrastructure, and Service Infrastructure can also be used for non-Google "third party" APIs. These APIs might be implemented using Google Cloud, but that's not necessary. Service Infrastructure can be used anywhere within reach of a Google data center, and the Service Infrastructure APIs are simple enough that if Google ever stopped providing them, they could be reimplemented and run anywhere.

Service Infrastructure is built on a central data format, called [Service Configuration](/docs/serviceinfrastructure/serviceconfig), and a small group of public Google APIs.

1. The [Service Management API](/docs/serviceinfrastructure/servicemanagement) manages descriptions of APIs and configurations that control their access.
2. The [Service Control API](/docs/serviceinfrastructure/servicecontrol) is used to check API requests from consumers and to monitor requests and responses.
3. The [API Keys API](/docs/serviceinfrastructure/apikeys) generates and manages API keys that consumers use to call APIs.
4. The [Cloud Logging API](/docs/serviceinfrastructure/cloudlogging) allows applications to read logs describing API traffic.
5. The [Cloud Monitoring API](/docs/serviceinfrastructure/cloudmonitoring) allows applications to read metrics describing API traffic.
6. The [Service Usage API](/docs/serviceinfrastructure/serviceusage) controls access to APIs within Google Cloud projects and is used to enable the APIs listed above and to share services with other Google Cloud users.

All of these APIs are described and built with Protocol Buffers, and we frequently refer to their Protocol Buffer descriptions as we explore them. The APIs themselves are available using either gRPC or HTTP/JSON. We use both but rely most heavily on gRPC.

Service Infrastructure is in use today, both by Google and by users of [Cloud Endpoints](/docs/endpoints), an API management product that Google released in 2016. Because Service Infrastructure is used by virtually all of Google's public APIs, it is possibly the highest-traffic API management system in existence.

---
#### Continue with [Service Configuration](/docs/serviceinfrastructure/serviceconfig).