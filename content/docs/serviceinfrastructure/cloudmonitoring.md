---
weight: 6
title: The Cloud Monitoring API
---
# The Cloud Monitoring API

The Cloud Monitoring API allows applications to read metrics describing API traffic.

The Cloud Monitoring API is defined in the [googleapis](/docs/details/googleapis) repo in [monitoring.yaml](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/monitoring.yaml).
The methods specific to the API are defined 
by the 
[AlertPolicyService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/alert_service.proto#L44),
[GroupService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/group_service.proto#L48),
[MetricService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/metric_service.proto#L67),
[NotificationChannelService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/notification_service.proto#L38),
[QueryService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/query_service.proto#L34),
[ServiceMonitoringService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/service_service.proto#L39),
[SnoozeService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/snooze_service.proto#L37), and
[UptimeCheckService](https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/uptime_service.proto#L43).

This is a large API. Our discussion will be limited to the `MetricService`.

## The MetricService service

Method names below are prefixed with `google.monitoring.v3.MetricService.`

| Method | Description |
| ------ | ----------- |
| ListMonitoredResourceDescriptors | Lists monitored resource descriptors that match a filter |
| GetMonitoredResourceDescriptor | Gets a single monitored resource descriptor |
| ListMetricDescriptors | Lists metric descriptors that match a filter |
| GetMetricDescriptor | Gets a single metric descriptor |
| CreateMetricDescriptor | Creates a new metric descriptor |
| DeleteMetricDescriptor | Deletes a metric descriptor |
| ListTimeSeries | Lists time series that match a filter |
| CreateTimeSeries | Creates or adds data to one or more time series |
| CreateServiceTimeSeries | Creates or adds data to one or more service time series |

### ListMonitoredResourceDescriptors
 
### GetMonitoredResourceDescriptor

### ListMetricDescriptors

### GetMetricDescriptor

### CreateMetricDescriptor

### DeleteMetricDescriptor

### ListTimeSeries

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
### CreateTimeSeries

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

### CreateServiceTimeSeries

https://cloud.google.com/monitoring/api/metrics_gcp#gcp-serviceruntime

## Usage Notes

```
$ q monitoring list-metric-descriptors projects/bobadojo > descriptors.json
$ jq < descriptors.json .[].name -r | grep serviceruntime
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

https://cloud.google.com/monitoring/api/v3/filters

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