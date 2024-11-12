---
weight: 5
title: The Cloud Logging API
---
# The Cloud Logging API

The Cloud Logging API allows applications to read logs describing API traffic.

## The Cloud Logging API methods

The methods of the Cloud Logging API are defined in the [googleapis](/googleapis) repo in [google/logging/v2/logging.proto](https://github.com/googleapis/googleapis/blob/9b94dba2f7f4b601f8232bc3a3f6ef32665279b9/google/logging/v2/logging.proto#L39).

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
