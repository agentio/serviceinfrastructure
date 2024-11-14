---
weight: 5
title: The Cloud Logging API
---
# The Cloud Logging API

[Cloud Logging](https://cloud.google.com/logging) is a Google service that provides centralized storage and analytics for application logs. Service Infrastructure users use it indirectly when Service Control API handlers write logs using the Cloud Logging API. This allows Service Infrastructure users to use the [Logs Explorer](https://cloud.google.com/logging/docs/view/logs-explorer-interface) to view their APIs and the Cloud Logging API to programmatically access their logs.

The Cloud Logging API is defined in the [googleapis](/docs/details/googleapis) repo in [logging_v2.yaml](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_v2.yaml).
The methods specific to the API are defined 
by the 
[ConfigServiceV2](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_config.proto#L50),
[LoggingServiceV2](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging.proto#L39), and
[MetricsServiceV2](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_metrics.proto#L38) in 
[logging_config.proto](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_config.proto),
[logging.proto](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging.proto), and
[logging_metrics.proto](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_metrics.proto).

This is a large API with many uses beyond Service Infrastructure. Here we list all of the available methods but focus on the ones most relevant to Service Infrastructure users.

## The ConfigServiceV2 service

Documented as a "service for configuring sinks used to route log entries", this service provides a rich set of capabilities for configuring log storage. Service Infrastructure users will find that these resources are automatically configured, so our survey of these methods will focus on listing the various resource types.

Method names below are prefixed with `google.logging.v2.ConfigServiceV2.`

| Method | Description |
| ------ | ----------- |
| ListBuckets | Lists log buckets |
| GetBucket | Gets a log bucket |
| CreateBucketAsync | Creates a log bucket asynchronously that can be used to store log entries |
| UpdateBucketAsync | Updates a log bucket asynchronously |
| CreateBucket | Creates a log bucket that can be used to store log entries |
| UpdateBucket | Updates a log bucket |
| DeleteBucket | Deletes a log bucket |
| UndeleteBucket | Undeletes a log bucket |
| ListViews | Lists views on a log bucket |
| GetView | Gets a view on a log bucket |
| CreateView | Creates a view over log entries in a log bucket |
| UpdateView | Updates a view on a log bucket |
| DeleteView | Deletes a view on a log bucket |
| ListSinks | Lists sinks |
| GetSink | Gets a sink |
| CreateSink | Creates a sink that exports specified log entries to a destination |
| UpdateSink | Updates a sink |
| DeleteSink | Deletes a sink |
| CreateLink | Asynchronously creates a linked dataset in BigQuery which makes it possible to use BigQuery to read the logs stored in the log bucket |
| DeleteLink | Deletes a link |
| ListLinks | Lists links |
| GetLink | Gets a link |
| ListExclusions | Lists all the exclusions on the _Default sink in a parent resource |
| GetExclusion | Gets the description of an exclusion in the _Default sink |
| CreateExclusion | Creates a new exclusion in the _Default sink in a specified parent resource |
| UpdateExclusion | Changes one or more properties of an existing exclusion in the _Default sink |
| DeleteExclusion | Deletes an exclusion in the _Default sink |
| GetCmekSettings | Gets the Logging CMEK settings for the given resource |
| UpdateCmekSettings | Updates the Log Router CMEK settings for the given resource |
| GetSettings | Gets the Log Router settings for the given resource |
| UpdateSettings | Updates the Log Router settings for the given resource |
| CopyLogEntries | Copies a set of log entries from a log bucket to a Cloud Storage bucket |

## The LoggingServiceV2 service

Documented as a "service for ingesting and querying logs," this will be a close focus for us as we look at how we can programmatically read the logs written by Service Infrastructure.

Method names below are prefixed with `google.logging.v2.LoggingServiceV2.`

| Method | Description |
| ------ | ----------- |
| DeleteLog | Deletes all the log entries in a log for the _Default Log Bucket |
| WriteLogEntries | Writes log entries to Logging |
| ListLogEntries | Lists log entries |
| ListMonitoredResourceDescriptors | Lists the descriptors for monitored resource types used by Logging |
| ListLogs | Lists the logs in projects, organizations, folders, or billing accounts |
| TailLogEntries | Streaming read of log entries as they are ingested |

### DeleteLog

### WriteLogEntries

### ListLogEntries

```
$ q logging list-log-entries bobadojo stores.endpoints.bobadojo.cloud.goog/endpoints_log
{
  "log_name": "projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log",
  "resource": {
    "type": "api",
    "labels": {
      "location": "global",
      "method": "bobadojo.stores.v1.Stores.GetStore",
      "project_id": "bobadojo",
      "service": "stores.endpoints.bobadojo.cloud.goog",
      "version": "v1"
    }
  },
  "Payload": {
    "JsonPayload": {
      "api_key": "REDACTED",
      "api_key_state": "VERIFIED",
      "api_method": "bobadojo.stores.v1.Stores.GetStore",
      "api_name": "bobadojo.stores.v1.Stores",
      "api_version": "v1",
      "grpc_status_code": "OK",
      "http_status_code": 200,
      "log_message": "bobadojo.stores.v1.Stores.GetStore is called",
      "producer_project_id": "bobadojo",
      "response_code_detail": "via_upstream",
      "service_agent": "ESPv2/2.48.0",
      "service_config_id": "2024-10-18r0",
      "timestamp": 1731381304.226605
    }
  },
  "timestamp": {
    "seconds": 1731381304,
    "nanos": 226605011
  },
  "receive_timestamp": {
    "seconds": 1731381306,
    "nanos": 106807545
  },
  "severity": 200,
  "insert_id": "ecb2ce7a-103f-4ffe-8f11-717c5ccfe0554850908905639612439@aq",
  "http_request": {
    "request_method": "POST",
    "request_url": "/bobadojo.stores.v1.Stores/GetStore",
    "request_size": 422,
    "status": 200,
    "response_size": 187,
    "remote_ip": "10.1.9.100",
    "latency": {},
    "protocol": "grpc"
  }
}
{
  "log_name": "projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log",
  "resource": {
    "type": "api",
    "labels": {
      "location": "global",
      "method": "bobadojo.stores.v1.Stores.GetStore",
      "project_id": "bobadojo",
      "service": "stores.endpoints.bobadojo.cloud.goog",
      "version": "v1"
    }
  },
  "Payload": {
    "JsonPayload": {
      "api_key": "REDACTED",
      "api_key_state": "VERIFIED",
      "api_method": "bobadojo.stores.v1.Stores.GetStore",
      "api_name": "bobadojo.stores.v1.Stores",
      "api_version": "v1",
      "grpc_status_code": "OK",
      "http_status_code": 200,
      "log_message": "bobadojo.stores.v1.Stores.GetStore is called",
      "producer_project_id": "bobadojo",
      "response_code_detail": "via_upstream",
      "service_agent": "ESPv2/2.48.0",
      "service_config_id": "2024-10-18r0",
      "timestamp": 1731381304.2231698
    }
  },
  "timestamp": {
    "seconds": 1731381304,
    "nanos": 223169837
  },
  "receive_timestamp": {
    "seconds": 1731381306,
    "nanos": 106807545
  },
  "severity": 200,
  "insert_id": "ecb2ce7a-103f-4ffe-8f11-717c5ccfe0554850908905639612439@ap",
  "http_request": {
    "request_method": "POST",
    "request_url": "/bobadojo.stores.v1.Stores/GetStore",
    "request_size": 422,
    "status": 200,
    "response_size": 206,
    "remote_ip": "10.1.9.100",
    "latency": {},
    "protocol": "grpc"
  }
}
{
  "log_name": "projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log",
  "resource": {
    "type": "api",
    "labels": {
      "location": "global",
      "method": "bobadojo.stores.v1.Stores.GetStore",
      "project_id": "bobadojo",
      "service": "stores.endpoints.bobadojo.cloud.goog",
      "version": "v1"
    }
  },
  "Payload": {
    "JsonPayload": {
      "api_key": "REDACTED",
      "api_key_state": "VERIFIED",
      "api_method": "bobadojo.stores.v1.Stores.GetStore",
      "api_name": "bobadojo.stores.v1.Stores",
      "api_version": "v1",
      "grpc_status_code": "OK",
      "http_status_code": 200,
      "log_message": "bobadojo.stores.v1.Stores.GetStore is called",
      "producer_project_id": "bobadojo",
      "response_code_detail": "via_upstream",
      "service_agent": "ESPv2/2.48.0",
      "service_config_id": "2024-10-18r0",
      "timestamp": 1731381304.2194657
    }
  },
  "timestamp": {
    "seconds": 1731381304,
    "nanos": 219465707
  },
  "receive_timestamp": {
    "seconds": 1731381306,
    "nanos": 106807545
  },
  "severity": 200,
  "insert_id": "ecb2ce7a-103f-4ffe-8f11-717c5ccfe0554850908905639612439@ao",
  "http_request": {
    "request_method": "POST",
    "request_url": "/bobadojo.stores.v1.Stores/GetStore",
    "request_size": 421,
    "status": 200,
    "response_size": 175,
    "remote_ip": "10.1.9.100",
    "latency": {},
    "protocol": "grpc"
  }
}

```

### ListMonitoredResourceDescriptors

### ListLogs

### TailLogEntries

## The MetricsServiceV2 service

This service is documented as a "service for configuring logs-based metrics". We'll focus on the read-only methods for observing any automatically-defined metrics.

Method names below are prefixed with `google.logging.v2.MetricsServiceV2.`

| Method | Description |
| ------ | ----------- |
| ListLogMetrics | Lists logs-based metrics |
| GetLogMetric | Gets a logs-based metric |
| CreateLogMetric | Creates a logs-based metric |
| UpdateLogMetric | Creates or updates a logs-based metric |
| DeleteLogMetric | Deletes a logs-based metric |