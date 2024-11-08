---
weight: 1
title: Service Configuration
---
# Service Configuration

Service Configuration is the central format of Service Management and arose from Google internal API practice. It consists of the [Service](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/service.proto#L47) proto and the messages that it contains. API developers often provide this "service config" in YAML files that conform to the schema defined by the protos, but typically they only provide fragments of the full configuration and tools fill in the rest.

## What's in the Service Configuration

An example full service config is in [stores.endpoints.bobadojo.cloud.goog.2024-10-18r0.json](/examples/stores.endpoints.bobadojo.cloud.goog.2024-10-18r0.json). Clearly this is not something that a developer would want to manually create! In practice, the service config is an intermediate internal format used by Cloud Endpoints. Developers instead provide fragments, like the one in [api_config.yaml](/docs/quickstart/files#api_configyaml), which contain only the parts of service config that can't be automatically derived from other known things.