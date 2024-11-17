---
weight: 6
title: The Cloud Monitoring API
---
## The Cloud Monitoring API

The Cloud Monitoring API allows applications to read metrics describing API traffic.

The Cloud Monitoring API is defined in the [googleapis](/docs/details/googleapis) repo in [monitoring.yaml](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/monitoring.yaml). It includes nine services:

| Service | Purpose |
| ------- | ------- |
| `AlertPolicyService` | Manages alerting policies |
| `GroupService` | Inspects and manages groups of resources being monitored |
| [MetricService](#the-metricservice-service) | Manages metric descriptors, monitored resource descriptors, and time series data |
| `NotificationChannelService` | Configures how messages related to incidents are sent |
| `QueryService` | Manages time series data |
| `ServiceMonitoringService` | Manages service-oriented metrics |
| `SnoozeService` | Temporarily prevents alert policies from generating alerts |
| `UptimeCheckService` | Manages uptime check configurations |
| `Operations` | A mix-in that handles long-running operations |

And we thought the logging API was a lot! Here we'll only look at the `MetricService`.

## The MetricService service

The [MetricService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/metric_service.proto#L67) service is defined in [metric_service.proto](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/metric_service.proto).

We'll use it to read metrics written by `ServiceControl`. Using other services in the API, we could add alerts and other mechanisms to help us respond to unusual situations that we can observe with metrics.

The full names of these methods begin with `google.monitoring.v3.MetricService.`

| Method | Description |
| ------ | ----------- |
| [ListMonitoredResourceDescriptors](#listmonitoredresourcedescriptors) | Lists monitored resource descriptors that match a filter |
| [GetMonitoredResourceDescriptor](#getmonitoredresourcedescriptor) | Gets a single monitored resource descriptor |
| [ListMetricDescriptors](#listmetricdescriptors) | Lists metric descriptors that match a filter |
| [GetMetricDescriptor](#getmetricdescriptor) | Gets a single metric descriptor |
| [CreateMetricDescriptor](#createmetricdescriptor) | Creates a new metric descriptor |
| [DeleteMetricDescriptor](#deletemetricdescriptor) | Deletes a metric descriptor |
| [ListTimeSeries](#listtimeseries) | Lists time series that match a filter |
| [CreateTimeSeries](#createtimeseries) | Creates or adds data to one or more time series |
| [CreateServiceTimeSeries](#createservicetimeseries) | Creates or adds data to one or more service time series |

### ListMonitoredResourceDescriptors

`ListMonitoredResourceDescriptors` lists the resource types that can be specified as monitored resources.

This is also documented at [Monitored Resource Types](https://cloud.google.com/monitoring/api/resources).

```
$ q monitoring list-monitored-resource-descriptors projects/bobadojo | jq .[].name -r | wc -l
324

$ q monitoring list-monitored-resource-descriptors projects/bobadojo | jq .[].name -r | egrep "api$"
projects/bobadojo/monitoredResourceDescriptors/api
projects/bobadojo/monitoredResourceDescriptors/consumed_api
projects/bobadojo/monitoredResourceDescriptors/produced_api
```

### GetMonitoredResourceDescriptor

`GetMonitoredResourceDescriptor` lets us get individual descriptors. These are the same values that we get from `ListMonitoredResourceDescriptors`, so we won't exercise this in detail.

### ListMetricDescriptors

`ListMetricDescriptors` lists metrics that are defined for Cloud Monitoring APIs. There are lots of them!

```
$ q monitoring list-metric-descriptors projects/bobadojo | jq .[].name -r | wc -l
6440
```

For Service Infrastructure, we care about the ones associated with `serviceruntime.googleapis.com`:

```
 q monitoring list-metric-descriptors projects/bobadojo | jq .[].name -r | grep serviceruntime
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_count
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_latencies
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_latencies_backend
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_latencies_overhead
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_sizes
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/response_sizes
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/allocation/usage
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/concurrent/exceeded
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/concurrent/limit
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/concurrent/usage
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/exceeded
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/limit
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/quota/rate/net_usage
projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/reserved/metric1
```

These are described at [Google Cloud Metrics: serviceruntime](https://cloud.google.com/monitoring/api/metrics_gcp#gcp-serviceruntime), which closely matches the descriptors that we get from `ListMetricDescriptors` (hopefully because this documentation is automatically-generated!)

Here's an example. First, this is the online documentation:

![alt text](/screenshots/api-request-count.png)

This is the metric descriptor that we get from `ListMetricDescriptors`: 
```
 {
    "name": "projects/bobadojo/metricDescriptors/serviceruntime.googleapis.com/api/request_count",
    "type": "serviceruntime.googleapis.com/api/request_count",
    "labels": [
      {
        "key": "protocol",
        "description": "The protocol of the request, e.g. \"http\", \"grpc\""
      },
      {
        "key": "response_code",
        "description": "The HTTP response code for HTTP requests, or HTTP equivalent code for gRPC requests. See code mapping in https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto."
      },
      {
        "key": "response_code_class",
        "description": "The response code class for HTTP requests, or HTTP equivalent class for gRPC requests, e.g. \"2xx\", \"4xx\""
      },
      {
        "key": "grpc_status_code",
        "description": "The numeric gRPC response code for gRPC requests, or gRPC equivalent code for HTTP requests. See code mapping in https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto."
      }
    ],
    "metric_kind": 2,
    "value_type": 2,
    "unit": "1",
    "description": "The count of completed requests.",
    "display_name": "Request count",
    "metadata": {
      "launch_stage": 4,
      "sample_period": {
        "seconds": 60
      },
      "ingest_delay": {
        "seconds": 1800
      }
    },
    "launch_stage": 4,
    "monitored_resource_types": [
      "api",
      "consumed_api",
      "produced_api"
    ]
  },
```

### GetMetricDescriptor

Like `GetMonitoredResourceDescriptor`, `GetMetricDescriptor` just gives us the same values that we get from `ListMetricDescriptors`, so we won't exercise this in detail.

### CreateMetricDescriptor

`CreateMetricDescriptor` lets us create new metric descriptors. Since our focus is on reading values that are written by Service Control, we'll skip over this.

### DeleteMetricDescriptor

Similarly, we'll leave `DeleteMetricDescriptor` for future explorations.

### ListTimeSeries

`ListTimeSeries` is what we're after -- this is the method that we can call to extract metrics from Cloud Monitoring. `ListTimeSeries` requires a filter expression to specify what we want (a metric type and other optional qualifiers), and we can also specify a time window, an "aggregation", and other parameters to control the response. Here we'll just try a simple query to get the request counts in the last hour (this time range is a default value for the `q` subcommand).

```
$ q monitoring list-time-series bobadojo serviceruntime.googleapis.com/api/request_count | jq
[
  {
    "metric": {
      "type": "serviceruntime.googleapis.com/api/request_count",
      "labels": {
        "grpc_status_code": "0",
        "protocol": "grpc",
        "response_code": "200",
        "response_code_class": "2xx"
      }
    },
    "resource": {
      "type": "api",
      "labels": {
        "location": "global",
        "method": "bobadojo.stores.v1.Stores.FindStores",
        "project_id": "bobadojo",
        "service": "stores.endpoints.bobadojo.cloud.goog",
        "version": "v1"
      }
    },
    "metricKind": "DELTA",
    "valueType": "INT64",
    "points": [
      {
        "interval": {
          "endTime": "2024-11-12T03:17:00Z",
          "startTime": "2024-11-12T03:16:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:16:00Z",
          "startTime": "2024-11-12T03:15:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:12:00Z",
          "startTime": "2024-11-12T03:11:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:11:00Z",
          "startTime": "2024-11-12T03:10:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:07:00Z",
          "startTime": "2024-11-12T03:06:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:06:00Z",
          "startTime": "2024-11-12T03:05:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:02:00Z",
          "startTime": "2024-11-12T03:01:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T03:01:00Z",
          "startTime": "2024-11-12T03:00:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:57:00Z",
          "startTime": "2024-11-12T02:56:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:56:00Z",
          "startTime": "2024-11-12T02:55:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:52:00Z",
          "startTime": "2024-11-12T02:51:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:51:00Z",
          "startTime": "2024-11-12T02:50:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:47:00Z",
          "startTime": "2024-11-12T02:46:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:46:00Z",
          "startTime": "2024-11-12T02:45:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:42:00Z",
          "startTime": "2024-11-12T02:41:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:41:00Z",
          "startTime": "2024-11-12T02:40:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:37:00Z",
          "startTime": "2024-11-12T02:36:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:36:00Z",
          "startTime": "2024-11-12T02:35:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:32:00Z",
          "startTime": "2024-11-12T02:31:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:31:00Z",
          "startTime": "2024-11-12T02:30:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:27:00Z",
          "startTime": "2024-11-12T02:26:00Z"
        },
        "value": {
          "int64Value": "0"
        }
      },
      {
        "interval": {
          "endTime": "2024-11-12T02:26:00Z",
          "startTime": "2024-11-12T02:25:00Z"
        },
        "value": {
          "int64Value": "1"
        }
      }
    ]
  },
  ...
```

This was truncated! The actual response includes request counts for many other APIs. We can narrow the results by adding to the filter; filtering expressions are documented at [Monitoring filters](https://cloud.google.com/monitoring/api/v3/filters). Here we'll just ask for bobadojo APIs, and we'll use `jq` to just see which API methods are in our response:

```
$ q monitoring list-time-series bobadojo serviceruntime.googleapis.com/api/request_count --filter ' AND resource.labels.method = starts_with("bobadojo.stores")' | jq .[].resource.labels.method
"bobadojo.stores.v1.Stores.FindStores"
"bobadojo.stores.v1.Stores.GetStore"
"bobadojo.stores.v1.Stores.ListStores"
"bobadojo.stores.v1.Stores.FindStores"
"bobadojo.stores.v1.Stores.GetStore"
"bobadojo.stores.v1.Stores.FindStores"
"bobadojo.stores.v1.Stores.GetStore"
"bobadojo.stores.v1.Stores.ListStores"
```

For comparison, without the filter, we got the following:
```
$ q monitoring list-time-series bobadojo serviceruntime.googleapis.com/api/request_count | jq .[].resource.labels.method -r
bobadojo.stores.v1.Stores.FindStores
bobadojo.stores.v1.Stores.GetStore
bobadojo.stores.v1.Stores.ListStores
google.devtools.cloudtrace.v2.TraceService.BatchWriteSpans
google.monitoring.v3.MetricService.ListMetricDescriptors
google.monitoring.v3.MetricService.ListTimeSeries
google.monitoring.v3.MetricService.ListTimeSeries
google.monitoring.v3.MetricService.ListTimeSeries
google.cloud.location.Locations.ListLocations
google.cloud.run.v1.DomainMappings.ListDomainMappings
google.cloud.run.v1.Revisions.ListRevisions
google.cloud.run.v1.Routes.ListRoutes
google.cloud.run.v1.Services.GetService
google.cloud.run.v1.Services.ListServices
google.cloud.run.v1.Services.TestIamPermissions
google.api.servicecontrol.v1.ServiceController.Check
google.api.servicecontrol.v1.ServiceController.Report
google.api.servicecontrol.v1.ServiceController.Report
google.api.servicemanagement.v1.ServiceManager.GetServiceConfig
google.api.servicemanagement.v1.ServiceManager.ListServiceRollouts
google.api.servicemanagement.v1.ServiceManager.ListServices
google.api.servicemanagement.v1.ServiceManager.ListServices
bobadojo.stores.v1.Stores.FindStores
bobadojo.stores.v1.Stores.GetStore
bobadojo.stores.v1.Stores.FindStores
bobadojo.stores.v1.Stores.GetStore
bobadojo.stores.v1.Stores.ListStores
```

Notice that we're seeing calls to `servicemanagement` and `servicecontrol` along with a variety of other calls that have been happening either behind-the-scenes or in other experimentation.

We can also look at other metrics. Here we can see that we've gone over quota during some recent experiments with the Service Management APIs:

```
$ q monitoring list-time-series bobadojo serviceruntime.googleapis.com/quota/exceeded | jq | more
[
  {
    "metric": {
      "type": "serviceruntime.googleapis.com/quota/exceeded",
      "labels": {
        "limit_name": "DefaultRequestsPerMinutePerProject",
        "quota_metric": "servicemanagement.googleapis.com/default_requests"
      }
    },
    "resource": {
      "type": "consumer_quota",
      "labels": {
        "location": "global",
        "project_id": "bobadojo",
        "service": "servicemanagement.googleapis.com"
      }
    },
    "metricKind": "GAUGE",
    "valueType": "BOOL",
    "points": [
      {
        "interval": {
          "endTime": "2024-11-14T04:25:47.717304Z",
          "startTime": "2024-11-14T04:25:47.717304Z"
        },
        "value": {
          "boolValue": true
        }
      }
    ]
  }
]
```

### CreateTimeSeries

It's a bit out of scope for our purposes, but it's also possible to write time series data using our own calls to `MetricService`. Here as an example, we write two values to a custom metric and then read the resulting time series.

```
$ q monitoring create-time-series bobadojo custom.googleapis.com/stores/orders 10
Done writing time series data.

$ q monitoring create-time-series bobadojo custom.googleapis.com/stores/orders 25
Done writing time series data.

$ q monitoring list-time-series bobadojo custom.googleapis.com/stores/orders | jq
[
  {
    "metric": {
      "type": "custom.googleapis.com/stores/orders"
    },
    "resource": {
      "type": "global",
      "labels": {
        "project_id": "bobadojo"
      }
    },
    "metricKind": "GAUGE",
    "valueType": "DOUBLE",
    "points": [
      {
        "interval": {
          "endTime": "2024-11-14T04:58:15Z",
          "startTime": "2024-11-14T04:58:15Z"
        },
        "value": {
          "doubleValue": 25
        }
      },
      {
        "interval": {
          "endTime": "2024-11-14T04:58:07Z",
          "startTime": "2024-11-14T04:58:07Z"
        },
        "value": {
          "doubleValue": 10
        }
      }
    ]
  }
]
```
Naming conventions and other help for using custom-defined metrics can be found at [User-defined metrics overview](https://cloud.google.com/monitoring/custom-metrics).

### CreateServiceTimeSeries

`CreateServiceTimeSeries` is a variant of `CreateTimeSeries` that is intended for Google internal use only.