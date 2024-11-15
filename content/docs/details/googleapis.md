---
weight: 4
title: The googleapis Repository
---
## The googleapis Repository

Google publishes Protocol Buffer descriptions of public Google APIs in the [googleapis/googleapis](https://github.com/googleapis/googleapis) repository on GitHub. These can be used with Protocol Buffer and gRPC tools to generate clients, documentation, and any other support materials needed to work with Google's public APIs.

The googleapis repo includes an automatically-generated index of Google APIs in [api-index-v1.json](https://github.com/googleapis/googleapis/blob/master/api-index-v1.json). Generally, each API is in a directory in a path ending with a version identifier. For example, the description of the [Cloud Translation API](https://cloud.google.com/translate) is in [google/cloud/translate/v3](https://github.com/googleapis/googleapis/tree/master/google/cloud/translate/v3). The main Protocol Buffer description is in [translation_service.proto](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translation_service.proto) and the service configuration is in [translate_v3.yaml](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translate_v3.yaml).

We can also find many examples of [Service Configuration](/docs/serviceinfrastructure/serviceconfig) in the googleapis repo. To make it easier to explore them, we've copied them into this generated list of [googleapis Service Configurations](/docs/details/services).