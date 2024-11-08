---
weight: 4
title: "Service Infrastructure"
bookFlatSection: true
---
# Service Infrastructure

[Service Infrastructure](https://cloud.google.com/service-infrastructure/docs/overview) is described as Google's "foundational platform for creating, managing, securing, and consuming APIs and services across organizations."

Capabilities provided by Service Infrastructure include:
* authentication and authorization
* rate limiting
* logging, monitoring, analytics, and auditing

Virtually all of Google's APIs are built on Service Infrastructure, and Service Infrastructure can also be used for non-Google "third party" APIs. These APIs might be implemented using Google Cloud, but that's not necessary. Service Infrastructure can be used anywhere within reach of a Google Cloud data center, and the Service Infrastructure APIs are simple enough that they can be reimplemented and run anywhere.

Service Infrastructure is built on a central data format, called [Service Configuration](/docs/serviceinfrastructure/serviceconfig) and a small(ish) group of public Google APIs.

1. The [Service Management API](/docs/serviceinfrastructure/servicemanagement) manages descriptions of APIs and configurations that control their access.
2. The [Service Control API](/docs/serviceinfrastructure/servicecontrol) is used to check API requests from consumers and to monitor requests and responses.
3. The [API Keys API](/docs/serviceinfrastructure/apikeys) generates and manages API keys that consumers use to call APIs.
4. The [Cloud Logging API](/docs/serviceinfrastructure/cloudlogging) allows applications to read logs describing API traffic.
5. The [Cloud Monitoring API](/docs/serviceinfrastructure/cloudmonitoring) allows applications to read metrics describing API traffic.
6. The [Service Usage API](/docs/serviceinfrastructure/serviceusage), is used to control access to Google APIs within Google Cloud projects, and is used to enable the other APIs listed above.

In 2016, Google released the [Cloud Endpoints](/docs/endpoints) API management product, which uses Service Infrastructure to manage APIs on Google Cloud and external platforms.
