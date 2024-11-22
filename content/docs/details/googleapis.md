---
weight: 4
title: The googleapis Repository
---
## The googleapis Repository

Google publishes Protocol Buffer descriptions of public Google APIs in the [googleapis/googleapis](https://github.com/googleapis/googleapis) repository on GitHub. These can be used with Protocol Buffer and gRPC tools to generate clients, documentation, and any other support materials needed to work with Google's public APIs.

## What's in googleapis?

Generally, each API description is in a directory in a path ending with a version identifier. For example, the description of the [Cloud Translation API](https://cloud.google.com/translate) is in [google/cloud/translate/v3](https://github.com/googleapis/googleapis/tree/master/google/cloud/translate/v3). 

## An index of APIs

The googleapis repo includes an automatically-generated index of Google APIs in [api-index-v1.json](https://github.com/googleapis/googleapis/blob/master/api-index-v1.json).

Here we use `curl` and `jq` to count the number of APIs in this index. Note that strictly speaking, this is a list of API versions. The index currently contains two versions of the Cloud Translation API: `v3` and `v3beta1`.
```
$ curl -s https://raw.githubusercontent.com/googleapis/googleapis/refs/heads/master/api-index-v1.json | jq .apis.[].id -r | wc -l
383
```

### Service Configuration

We often think of Protocol Buffers as the way that Google describes its APIs, but the definition of a Google API starts in Service Config, which includes a list of the services in an API (each defined in a `.proto` file) and other properties of APIs that aren't in the `.proto` files.

For our Cloud Translation example, the service configuration is in [translate_v3.yaml](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translate_v3.yaml).

To get a broader sense of these files, we've copied all of them into this generated list of [googleapis Service Configurations](/docs/details/services).

### Protocol Buffers

The methods and messages of Google APIs are described with Protocol Buffers. Generally, the Protocol Buffer descriptions for an API are alongside the service config file for the API, and these often are split among multiple files.

The main Protocol Buffer description is in [translation_service.proto](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translation_service.proto).

### gRPC Service Configuration

Another file that you'll find for APIs in the `googleapis` repo is what the gRPC project unfortunately also calls [Service Config](https://grpc.io/docs/guides/service-config/). This is used to specify timeouts and retry behavior in gRPC clients that call the associated API.

The gRPC service configuration for the Cloud Translation API is in [translate_grpc_service_config.json](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translate_grpc_service_config.json).

---
#### Go back to [Going Deeper](/docs/details).