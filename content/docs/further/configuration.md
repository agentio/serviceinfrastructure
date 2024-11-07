---
weight: 3
title: Better Configuration
---
# Better Configuration

[probably the most-requested improvement to Endpoints](https://github.com/GoogleCloudPlatform/esp-v2/issues/500)

All we need is Service Config

To do our own Service Compilation, we need to generate service config from OpenAPI or other API descriptions

This could be based on open sourced (archived) API compiler

This is built into Service Management APIs… but doesn’t have to be

We can generate our own Service Config

This allows us to keep up with future versions of OpenAPI

But beware, there is no way to upload descriptor sets, so for gRPC transcoding, you have to use the Service Management APIs

