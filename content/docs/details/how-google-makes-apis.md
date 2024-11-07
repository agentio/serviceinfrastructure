---
weight: 3
title: How Google Makes APIs
---
# How Google Makes APIs

Looking at things that can be observed externally, we can learn a lot about how Google makes APIs and about large-scale API design in general. 

As I wrote in [I got a golden ticket: what I learned about APIs in my first year at Google](https://medium.com/apis-and-digital-transformation/i-got-a-golden-ticket-what-i-learned-about-apis-in-my-first-year-at-google-556e1f02f9ab), Google makes a lot of APIs and Google APIs get called a lot. These two dimensions of scale, in development and operations, have had a big influence on API practice at Google and led to things that might not be obvious to someone using an online tutorial to code their first API (which is a great first step!)

To get an idea of how many APIs Google produces, we could start by just asking Google directly by calling the [API Discovery Service](https://developers.google.com/discovery/v1/reference/apis/list)

```
$ curl -s https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[].name' -r | wc -l
463

$ curl -s https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[].name' -r | sort | uniq | wc -l
286

$ curl https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[] | "\(.name) \(.version)"' -r

```

Google also publishes a directory of API descriptions at [github.com/googleapis/googleapis](https://github.com/googleapis/googleapis). These are all public Google APIs, and the descriptions are in the Protocol Buffers format. They can be used with Protocol Buffer and gRPC support tools to generate clients, documentation, and any other materials needed to work with Google's public APIs.

The googleapis repo includes an automatically-generated index of Google APIs in [api-index-v1.json](https://github.com/googleapis/googleapis/blob/master/api-index-v1.json).

```
$ curl -s https://raw.githubusercontent.com/googleapis/googleapis/refs/heads/master/api-index-v1.json | jq .apis.[].id -r | wc -l
383
```

Generally, each API in the `googleapis` repo is in a directory in a path ending with a version identifier. For example, the description of the [Cloud Translation API](https://cloud.google.com/translate) is in [google/cloud/translate/v3](https://github.com/googleapis/googleapis/tree/master/google/cloud/translate/v3). The main Protocol Buffer description is in [translation_service.proto](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translation_service.proto) and the service configuration is in [translate_v3.yaml](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translate_v3.yaml).

### Protocol Buffers

The message serialization mechanism.
The language for describing APIs.
The methodology for implementing and using APIs.

Our sample API is described with Protocol Buffers. Here are the files that describe it:
```
bobadojo/stores/v1/stores.proto
```

### Service Configuration

Protocol Buffers files describe the API surface, but don't contain much about the deployment and operation of APIs. This is in a structure called Service Configuration, which is described by a Protocol Buffer description of the [service message](https://github.com/googleapis/googleapis/blob/master/google/api/service.proto#L80). The `googleapis` repo contains service configuration for each API. Another way to count Google APIs would be to count the service config files in the googleapis repo. These are yaml files that start with the line `type: google.api.Service`.

https://github.com/search?q=repo%3Agoogleapis%2Fgoogleapis%20allow_without_credential&type=code

### gRPC Service Configuration

Another file that you'll find for each API in the `googleapis` repo is what the gRPC team unfortunately also calls "Service Config". This is used to specify timeouts and retry behavior in gRPC clients that call the associated API.

https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translate_grpc_service_config.json

### Methodology 

Notes on good practice:
- proto files are best defined in a directory hierarchy because they are rarely in isolation (proto files often include other files).
- version names are in the proto path

services and messages

support code is almost always generated

The googleapis repository contains public protos for Google APIs. It's effectively a monorepo of API descriptions.

### gRPC

An RPC framework for HTTP/2 (and beyond) that works great with protocol buffers and Google APIs.

support code is generated

why does Google prefer gRPC?
- performance
- client-side influence (retry)

### Transcoding

Transcoding is a standard mechanism for making gRPC APIs available using HTTP+JSON.

https://cloud.google.com/endpoints/docs/grpc/transcoding

https://google.aip.dev/127

### Client libraries

Google provides client libraries for most Google APIs.

(Aside: Google doesn't do this for all APIs, but people are trying!)

#### Discovery-based Client Libraries

Google API Discovery Service

#### gRPC Client Libraries

http://github.com/googleapis/googleapis

### The gcloud command-line interface

gcloud can call Google APIs and often provides higher-level capabilities.

### AIPs and api-linter

Much of the above requires consistency! For this Google teams spent years internally developing standards that are now shared publicly as AIPs.

https://cloud.google.com/apis/design

https://aip.dev

To see the value of AIPs, consider pagination, as described in [AIP-158](https://google.aip.dev/158).

For more, see [API Design Patterns](https://www.manning.com/books/api-design-patterns) by JJ Geewax.