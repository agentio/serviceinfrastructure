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

```
  "name": "stores.endpoints.bobadojo.cloud.goog",
```

### title

_The product title for this service, it is the name displayed in Google Cloud Console._

```
  "title": "Boba Dojo Stores API",
```

### producer project id

_The Google project that owns this service._

```
  "producerProjectId": "bobadojo",
```

### id

_A unique ID for a specific instance of this message, typically assigned by the client for tracking purpose. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If empty, the server may choose to generate one instead._

```
  "id": "2024-10-18r0",
```

### apis

_A list of API interfaces exported by this service. Only the `name` field of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration author, as the remaining fields will be derived from the IDL during the normalization process. It is an error to specify an API interface here which cannot be resolved against the associated IDL files._

```
  "apis": [
    {
      "name": "bobadojo.stores.v1.Stores",
      "methods": [
        {
          "name": "ListStores",
          "requestTypeUrl": "type.googleapis.com/bobadojo.stores.v1.ListStoresRequest",
          "responseTypeUrl": "type.googleapis.com/bobadojo.stores.v1.ListStoresResponse",
          "options": [
            {
              "name": "google.api.http",
              "value": {
                "@type": "type.googleapis.com/google.api.HttpRule",
                "get": "/v1/stores"
              }
            }
          ]
        },
        {
          "name": "FindStores",
          "requestTypeUrl": "type.googleapis.com/bobadojo.stores.v1.FindStoresRequest",
          "responseTypeUrl": "type.googleapis.com/bobadojo.stores.v1.FindStoresResponse",
          "options": [
            {
              "name": "google.api.http",
              "value": {
                "@type": "type.googleapis.com/google.api.HttpRule",
                "get": "/v1/stores:find"
              }
            }
          ]
        },
        {
          "name": "GetStore",
          "requestTypeUrl": "type.googleapis.com/bobadojo.stores.v1.GetStoreRequest",
          "responseTypeUrl": "type.googleapis.com/bobadojo.stores.v1.Store",
          "options": [
            {
              "name": "google.api.method_signature",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "name"
              }
            },
            {
              "name": "google.api.http",
              "value": {
                "@type": "type.googleapis.com/google.api.HttpRule",
                "get": "/v1/{name=stores/*}"
              }
            }
          ]
        }
      ],
      "version": "v1",
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    }
  ],
```

### types

_A list of all proto message types included in this API service. Types referenced directly or indirectly by the `apis` are automatically included.  Messages which are not referenced but shall be included, such as types used by the `google.protobuf.Any` type, should be listed here by name by the configuration author Example:_

```
types:
  - name: google.protobuf.Int32
```

```
  "types": [
    {
      "name": "bobadojo.stores.v1.Store",
      "fields": [
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "name",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "IDENTIFIER"
              }
            }
          ],
          "jsonName": "name"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 4,
          "name": "type",
          "jsonName": "type"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 5,
          "name": "title",
          "jsonName": "title"
        },
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 6,
          "name": "location",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Location",
          "jsonName": "location"
        },
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 7,
          "name": "address",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Address",
          "jsonName": "address"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 8,
          "name": "store_hours",
          "jsonName": "storeHours"
        }
      ],
      "options": [
        {
          "name": "google.api.resource",
          "value": {
            "@type": "type.googleapis.com/google.api.ResourceDescriptor",
            "type": "stores.bobadojo.io/Store",
            "pattern": [
              "stores/{store}"
            ],
            "plural": "stores",
            "singular": "store"
          }
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.Address",
      "fields": [
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "street",
          "jsonName": "street"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "city",
          "jsonName": "city"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 3,
          "name": "state",
          "jsonName": "state"
        },
        {
          "kind": "TYPE_INT32",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 4,
          "name": "zip_code",
          "jsonName": "zipCode"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 5,
          "name": "region_code",
          "jsonName": "regionCode"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 6,
          "name": "county",
          "jsonName": "county"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.Location",
      "fields": [
        {
          "kind": "TYPE_FLOAT",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "latitude",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            }
          ],
          "jsonName": "latitude"
        },
        {
          "kind": "TYPE_FLOAT",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "longitude",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            }
          ],
          "jsonName": "longitude"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.BoundingBox",
      "fields": [
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "max",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Location",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            }
          ],
          "jsonName": "max"
        },
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "min",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Location",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            }
          ],
          "jsonName": "min"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.FindStoresRequest",
      "fields": [
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "bounds",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.BoundingBox",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            }
          ],
          "jsonName": "bounds"
        },
        {
          "kind": "TYPE_INT32",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "limit",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "OPTIONAL"
              }
            }
          ],
          "jsonName": "limit"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.FindStoresResponse",
      "fields": [
        {
          "kind": "TYPE_INT32",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "count",
          "jsonName": "count"
        },
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_REPEATED",
          "number": 2,
          "name": "stores",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Store",
          "jsonName": "stores"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.GetStoreRequest",
      "fields": [
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 1,
          "name": "name",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "REQUIRED"
              }
            },
            {
              "name": "google.api.resource_reference",
              "value": {
                "@type": "type.googleapis.com/google.api.ResourceReference",
                "type": "stores.bobadojo.com/Store"
              }
            }
          ],
          "jsonName": "name"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.ListStoresRequest",
      "fields": [
        {
          "kind": "TYPE_INT32",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "page_size",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "OPTIONAL"
              }
            }
          ],
          "jsonName": "pageSize"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 3,
          "name": "page_token",
          "options": [
            {
              "name": "google.api.field_behavior",
              "value": {
                "@type": "type.googleapis.com/google.protobuf.StringValue",
                "value": "OPTIONAL"
              }
            }
          ],
          "jsonName": "pageToken"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    },
    {
      "name": "bobadojo.stores.v1.ListStoresResponse",
      "fields": [
        {
          "kind": "TYPE_MESSAGE",
          "cardinality": "CARDINALITY_REPEATED",
          "number": 1,
          "name": "stores",
          "typeUrl": "type.googleapis.com/bobadojo.stores.v1.Store",
          "jsonName": "stores"
        },
        {
          "kind": "TYPE_STRING",
          "cardinality": "CARDINALITY_OPTIONAL",
          "number": 2,
          "name": "next_page_token",
          "jsonName": "nextPageToken"
        }
      ],
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3"
    }
  ],
```

### enums

_A list of all enum types included in this API service. Enums referenced directly or indirectly by the `apis` are automatically included. Enums which are not referenced but shall be included should be listed here by name by the configuration author. Example:_
```
enums:
  - name: google.someapi.v1.SomeEnum
```

### documentation

_Additional API documentation._

```
  "documentation": {
    "rules": [
      {
        "selector": "bobadojo.stores.v1.Store",
        "description": "A store."
      },
      {
        "selector": "bobadojo.stores.v1.Store.name",
        "description": "Identifier. A unique id (e.g. store number)"
      },
      {
        "selector": "bobadojo.stores.v1.Store.type",
        "description": "An identifier indicating the type of store"
      },
      {
        "selector": "bobadojo.stores.v1.Store.title",
        "description": "Store name (human-readable)"
      },
      {
        "selector": "bobadojo.stores.v1.Store.location",
        "description": "Store location"
      },
      {
        "selector": "bobadojo.stores.v1.Store.address",
        "description": "Store address"
      },
      {
        "selector": "bobadojo.stores.v1.Store.store_hours",
        "description": "Store hours"
      },
      {
        "selector": "bobadojo.stores.v1.Address",
        "description": "An address."
      },
      {
        "selector": "bobadojo.stores.v1.Address.street",
        "description": "Street address"
      },
      {
        "selector": "bobadojo.stores.v1.Address.city",
        "description": "City"
      },
      {
        "selector": "bobadojo.stores.v1.Address.state",
        "description": "State"
      },
      {
        "selector": "bobadojo.stores.v1.Address.zip_code",
        "description": "Zip code"
      },
      {
        "selector": "bobadojo.stores.v1.Address.region_code",
        "description": "Country"
      },
      {
        "selector": "bobadojo.stores.v1.Address.county",
        "description": "County"
      },
      {
        "selector": "bobadojo.stores.v1.Location",
        "description": "A location in terrestrial coordinates."
      },
      {
        "selector": "bobadojo.stores.v1.Location.latitude",
        "description": "Required. Latitude of the location."
      },
      {
        "selector": "bobadojo.stores.v1.Location.longitude",
        "description": "Required. Longitude of the location."
      },
      {
        "selector": "bobadojo.stores.v1.BoundingBox",
        "description": "A bounding box in terrestrial coordinates."
      },
      {
        "selector": "bobadojo.stores.v1.BoundingBox.max",
        "description": "Required. Maximum coordinate values."
      },
      {
        "selector": "bobadojo.stores.v1.BoundingBox.min",
        "description": "Required. Minimum coordinate values."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresRequest",
        "description": "Request to FindStores."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresRequest.bounds",
        "description": "Required. Bounding box of the request."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresRequest.limit",
        "description": "Optional. Maximum number of results to return."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresResponse",
        "description": "Response from FindStores."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresResponse.count",
        "description": "Number of matching stores."
      },
      {
        "selector": "bobadojo.stores.v1.FindStoresResponse.stores",
        "description": "Matching stores."
      },
      {
        "selector": "bobadojo.stores.v1.GetStoreRequest",
        "description": "Request to GetStore."
      },
      {
        "selector": "bobadojo.stores.v1.GetStoreRequest.name",
        "description": "Required. The ID of the store resource to retrieve."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresRequest",
        "description": "Request to ListStores."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresRequest.page_size",
        "description": "Optional. Page size."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresRequest.page_token",
        "description": "Optional. Pagination token."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresResponse",
        "description": "Response from ListStores."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresResponse.stores",
        "description": "List of stores."
      },
      {
        "selector": "bobadojo.stores.v1.ListStoresResponse.next_page_token",
        "description": "Token of next page."
      },
      {
        "selector": "bobadojo.stores.v1.Stores",
        "description": "Get stores and related information."
      },
      {
        "selector": "bobadojo.stores.v1.Stores.ListStores",
        "description": "List all stores."
      },
      {
        "selector": "bobadojo.stores.v1.Stores.FindStores",
        "description": "Returns a list of all stores in a specified region."
      },
      {
        "selector": "bobadojo.stores.v1.Stores.GetStore",
        "description": "Returns a specific store."
      }
    ]
  },
```

This is an instance of the [Documentation](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/documentation.proto#L80) proto. Reviewing the proto, we can see this can hold much more elaborate documentation than what we see here, which is automatically generated by Google's internal service compiler.

### backend

_API backend configuration._

```
  "backend": {
    "rules": [
      {
        "selector": "bobadojo.stores.v1.Stores.ListStores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.FindStores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.GetStore"
      }
    ]
  },
```

This is an instance of the [Backend](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/backend.proto#L26) proto. While not used here, this allows API methods to be assigned to different backends. This is respected by the Endpoints proxies.

### http

_HTTP configuration._

```
  "http": {
    "rules": [
      {
        "selector": "bobadojo.stores.v1.Stores.ListStores",
        "get": "/v1/stores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.FindStores",
        "get": "/v1/stores:find"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.GetStore",
        "get": "/v1/{name=stores/*}"
      }
    ]
  },
```

This is an instance of the [HTTP](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/http.proto#L29) proto. It contains HTTP rules that control HTTP/JSON transcoding and in the example above, is extracted from the proto annotations in the original API description.

### quota

_Quota configuration._

This is an instance of the [Quota](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/quota.proto#L76) proto. Not included in our example configuration, this allows service owners to define rate limits on individual API methods or groups of methods.

### authentication

_Auth configuration._

This is an instance of the [Authorization](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/auth.proto#L43) proto. It allows service owners to specify authorization requirements that include configurations that allow users to authenticate with JWT tokens.

### context

_Context configuration._

This is an instance of the [Context](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/context.proto#L61) proto. It can be use to list types of information that can be sent in requests as gRPC metadata.

### usage

_Configuration controlling usage of this service._

```
  "usage": {
    "rules": [
      {
        "selector": "bobadojo.stores.v1.Stores.ListStores",
        "allowUnregisteredCalls": true
      },
      {
        "selector": "bobadojo.stores.v1.Stores.FindStores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.GetStore"
      }
    ]
  },
```

This is an instance of the [Usage](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/usage.proto#L26) proto. It specifies which requests are accepted by an API and can be configured to allow unknown requests and requests without authentication to be accepted and forwarded to backends.

### endpoints

_Configuration for network endpoints.  If this is empty, then an endpoint with the same name as the service is automatically generated to service all defined APIs._

```
  "endpoints": [
    {
      "name": "stores.endpoints.bobadojo.cloud.goog",
      "target": "WWW.XXX.YYY.ZZZ"
    }
  ],
```

This is a repeated instance of the [Endpoints](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/endpoint.proto#L46) proto.

### control

_Configuration for the service control plane._

```
  "control": {
    "environment": "servicecontrol.googleapis.com"
  },
```

This is an instance of the [Control](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/control.proto#L33) proto. All Endpoints configurations should include this setting with this exact value. 

The Control proto also includes a repeated [MethodPolicy](https://github.com/googleapis/googleapis/blob/c72f219fedbb57d3f83c10550e135c4824b670eb/google/api/policy.proto#L70) field, which has no known public documentation beyond the comments in the proto.

### logs

_Defines the logs used by this service._

```
  "logs": [
    {
      "name": "endpoints_log"
    }
  ]
```

### metrics

_Defines the metrics used by this service._

```
  "metrics": [
    {
      "name": "serviceruntime.googleapis.com/api/consumer/request_count",
      "type": "serviceruntime.googleapis.com/api/consumer/request_count",
      "labels": [
        {
          "key": "/credential_id"
        },
        {
          "key": "/protocol"
        },
        {
          "key": "/response_code"
        },
        {
          "key": "/response_code_class"
        },
        {
          "key": "/status_code"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "INT64"
    },
    {
      "name": "serviceruntime.googleapis.com/api/consumer/total_latencies",
      "type": "serviceruntime.googleapis.com/api/consumer/total_latencies",
      "labels": [
        {
          "key": "/credential_id"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/request_count",
      "type": "serviceruntime.googleapis.com/api/producer/request_count",
      "labels": [
        {
          "key": "/protocol"
        },
        {
          "key": "/response_code"
        },
        {
          "key": "/response_code_class"
        },
        {
          "key": "/status_code"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "INT64"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/total_latencies",
      "type": "serviceruntime.googleapis.com/api/producer/total_latencies",
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/consumer/quota_used_count",
      "type": "serviceruntime.googleapis.com/api/consumer/quota_used_count",
      "labels": [
        {
          "key": "/credential_id"
        },
        {
          "key": "/quota_group_name"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "INT64"
    },
    {
      "name": "serviceruntime.googleapis.com/api/consumer/request_sizes",
      "type": "serviceruntime.googleapis.com/api/consumer/request_sizes",
      "labels": [
        {
          "key": "/credential_id"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/consumer/response_sizes",
      "type": "serviceruntime.googleapis.com/api/consumer/response_sizes",
      "labels": [
        {
          "key": "/credential_id"
        }
      ],
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/request_overhead_latencies",
      "type": "serviceruntime.googleapis.com/api/producer/request_overhead_latencies",
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/backend_latencies",
      "type": "serviceruntime.googleapis.com/api/producer/backend_latencies",
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/request_sizes",
      "type": "serviceruntime.googleapis.com/api/producer/request_sizes",
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    },
    {
      "name": "serviceruntime.googleapis.com/api/producer/response_sizes",
      "type": "serviceruntime.googleapis.com/api/producer/response_sizes",
      "metricKind": "DELTA",
      "valueType": "DISTRIBUTION"
    }
  ],
```

### monitored_resources

_Defines the monitored resources used by this service. This is required by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations._

```
  "monitoredResources": [
    {
      "type": "api",
      "labels": [
        {
          "key": "cloud.googleapis.com/location"
        },
        {
          "key": "cloud.googleapis.com/uid"
        },
        {
          "key": "serviceruntime.googleapis.com/api_version"
        },
        {
          "key": "serviceruntime.googleapis.com/api_method"
        },
        {
          "key": "serviceruntime.googleapis.com/consumer_project"
        },
        {
          "key": "cloud.googleapis.com/project"
        },
        {
          "key": "cloud.googleapis.com/service"
        }
      ]
    }
  ],
```

### billing

_Billing configuration._

### logging

_Logging configuration._

```
  "logging": {
    "producerDestinations": [
      {
        "monitoredResource": "api",
        "logs": [
          "endpoints_log"
        ]
      }
    ]
  },
```

### monitoring

_Monitoring configuration._

```
  "monitoring": {
    "producerDestinations": [
      {
        "monitoredResource": "api",
        "metrics": [
          "serviceruntime.googleapis.com/api/producer/request_count",
          "serviceruntime.googleapis.com/api/producer/total_latencies",
          "serviceruntime.googleapis.com/api/producer/request_overhead_latencies",
          "serviceruntime.googleapis.com/api/producer/backend_latencies",
          "serviceruntime.googleapis.com/api/producer/request_sizes",
          "serviceruntime.googleapis.com/api/producer/response_sizes"
        ]
      }
    ],
    "consumerDestinations": [
      {
        "monitoredResource": "api",
        "metrics": [
          "serviceruntime.googleapis.com/api/consumer/request_count",
          "serviceruntime.googleapis.com/api/consumer/quota_used_count",
          "serviceruntime.googleapis.com/api/consumer/total_latencies",
          "serviceruntime.googleapis.com/api/consumer/request_sizes",
          "serviceruntime.googleapis.com/api/consumer/response_sizes"
        ]
      }
    ]
  },
```

### system_parameters

_System parameter configuration._

```
  "systemParameters": {},
```

### source_info

_Output only. The source information for this configuration if available._

### publishing

_Settings for [Google Cloud Client libraries](https://cloud.google.com/apis/docs/cloud-client-libraries) generated from APIs defined as protocol buffers._

### config_version

_Obsolete. Do not use. This field has no semantic meaning. The service config compiler always sets this field to `3`._

```
  "configVersion": 3
```
