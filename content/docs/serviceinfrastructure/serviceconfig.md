---
weight: 1
title: Service Configuration
---
# Service Configuration

Service Configuration is the central format of Service Management and arose from Google internal API practice. It consists of the [Service](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/service.proto#L47) proto and the messages that it contains. API developers often provide this "service config" in YAML files that conform to the schema defined by the protos, but typically they only provide fragments of the full configuration and tools fill in the rest.

## An Example Service Configuration

Here is an example service configuration for a service managed with Endpoints:

[stores.endpoints.bobadojo.cloud.goog.2024-10-18r0.json](/examples/stores.endpoints.bobadojo.cloud.goog.2024-10-18r0.json)

Clearly this is not something that a developer would want to manually create! In practice, the service config is an intermediate internal format used by Cloud Endpoints. Developers instead provide fragments, like the one in [api_config.yaml](/docs/quickstart/files#api_configyaml), which contain only the parts of service config that can't be automatically derived from other known things.

## What's in the Service Configuration

### name

_The service name, which is a DNS-like logical identifier for the service, such as `calendar.googleapis.com`. The service name typically goes through DNS verification to make sure the owner of the service also owns the DNS name._

### title

_The product title for this service, it is the name displayed in Google Cloud Console._

### producer project id

_The Google project that owns this service._

### id

_A unique ID for a specific instance of this message, typically assigned by the client for tracking purpose. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If empty, the server may choose to generate one instead._

### apis

_A list of API interfaces exported by this service. Only the `name` field of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration author, as the remaining fields will be derived from the IDL during the normalization process. It is an error to specify an API interface here which cannot be resolved against the associated IDL files._

### types

_A list of all proto message types included in this API service. Types referenced directly or indirectly by the `apis` are automatically included.  Messages which are not referenced but shall be included, such as types used by the `google.protobuf.Any` type, should be listed here by name by the configuration author Example:_

```
types:
  - name: google.protobuf.Int32
```      
### enums

_A list of all enum types included in this API service. Enums referenced directly or indirectly by the `apis` are automatically included. Enums which are not referenced but shall be included should be listed here by name by the configuration author. Example:_
```
enums:
  - name: google.someapi.v1.SomeEnum
```

### documentation

_Additional API documentation._

### backend

_API backend configuration._

### http

_HTTP configuration._

### quota

_Quota configuration._

### authentication

_Auth configuration._

### context

_Context configuration._

### usage

_Configuration controlling usage of this service._

### endpoints

_Configuration for network endpoints.  If this is empty, then an endpoint with the same name as the service is automatically generated to service all defined APIs._

### control

_Configuration for the service control plane._

### logs

_Defines the logs used by this service._

### metrics

_Defines the metrics used by this service._

### monitored_resources

_Defines the monitored resources used by this service. This is required by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations._

### billing

_Billing configuration._

### logging

_Logging configuration._

### monitoring

_Monitoring configuration._

### system_parameters

_System parameter configuration._

### source_info

_Output only. The source information for this configuration if available._

### publishing

_Settings for [Google Cloud Client libraries](https://cloud.google.com/apis/docs/cloud-client-libraries) generated from APIs defined as protocol buffers._

### config_version

_Obsolete. Do not use. This field has no semantic meaning. The service config compiler always sets this field to `3`._
