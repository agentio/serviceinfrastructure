---
weight: 6
title: The Cloud Monitoring API
---
# The Cloud Monitoring API

The Cloud Monitoring API allows applications to read metrics describing API traffic.

## The Cloud Monitoring API Methods

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

### CreateServiceTimeSeries