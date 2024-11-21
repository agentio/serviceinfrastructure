---
weight: 2
title: "Managing gRPC APIs"
---
## Managing gRPC APIs

[gRPC](https://grpc.io) is an open source API framework that was [created at Google](https://grpc.io/blog/principles/) and [released in 2016](https://grpc.io/blog/ga-announcement/). gRPC makes it easier to produce and use high-performing APIs at scale, largely by constraining API designs to follow a certain style that allows API clients and servers to use highly-optimized tooling.

## What's different about gRPC?

### API Description

gRPC APIs are usually described with the [Protocol Buffers](https://protobuf.dev/) language, and request and response messages use the Protocol Buffer binary encoding. API designs are often expected to conform to style guides like Google's [API Improvement Proposals](https://google.aip.dev/), which increases their consistency and ability to be supported by tools and automation.

### API Implementation

gRPC clients and servers are more complex than other API implementations, so they are almost always built on generated code that is produced by tools that are usually open source. gRPC clients and server implementations are also complicated by using advanced networking technologies, including HTTP/2, streaming, and configurable automatic retry. To simplify API consumption, gRPC supports [transcoding](https://cloud.google.com/endpoints/docs/grpc/transcoding), which provides a simpler HTTP/JSON interface to gRPC APIs when their API descriptions have appropriate annotations and conform to style guidelines like [AIP-127](https://google.aip.dev/127).

## What do we want from gRPC API Management?

gRPC adds additional expectations and opportunities for API management. Here we list a few possibilties.

### Validation

gRPC APIs use Protocol Buffer encoding, so they are generally observable, allowing API management systems to look inside messages to validate or monitor requests and responses.

### Metadata

Protocol Buffer encoding isn't self-describing, so API management systems can make it easier to use gRPC APIs by providing standard metadata services like [gRPC Reflection](https://grpc.io/docs/guides/reflection/), which makes it easier for clients to serialize and deserialize API messages.

### Documentation

gRPC APIs are thoroughly and accurately described by their Protocol Buffer descriptions, so API management systems can easily provide generated documentation and even API client code.

### Transcoding

HTTP/JSON transcoding is well-defined (see [AIP-127](https://google.aip.dev/127)), so API management systems can provide transcoded versions of gRPC APIs that they manage.

### Transformation

Projects like [rejoiner](https://google.github.io/rejoiner/) show that gRPC APIs can be combined to create aggregated GraphQL APIs. This could be part of an advanced gRPC API management system.

---

Is there something else that you would like to get from gRPC API management? Tell us on [our issues page](https://github.com/agentio/serviceinfrastructure/issues).

---
#### Continue with [Get the Picture!](/docs/intro/picture).