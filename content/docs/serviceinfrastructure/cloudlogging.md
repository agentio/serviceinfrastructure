---
weight: 5
title: The Cloud Logging API
---
## The Cloud Logging API

[Cloud Logging](https://cloud.google.com/logging) is a Google service that provides centralized storage and analytics for application logs. Service Infrastructure users use it indirectly when Service Control API handlers write logs using the Cloud Logging API. This allows Service Infrastructure users to use the [Logs Explorer](https://cloud.google.com/logging/docs/view/logs-explorer-interface) to view their APIs and the Cloud Logging API to programmatically access their logs.

The Cloud Logging API is defined in the [googleapis](/docs/details/googleapis) repo in [logging_v2.yaml](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging_v2.yaml). It includes four services:

| Service | Purpose |
| ------- | ------- |
| `ConfigServiceV2` | A service for configuring sinks used to route log entries |
| [LoggingServiceV2](#the-loggingservicev2-service) | A service for ingesting and querying logs |
| `MetricsServiceV2` | A service for configuring logs-based metrics |
| `Operations` | A mix-in that handles long-running operations |

This is a large API with many uses beyond Service Infrastructure. Here we focus on the service most relevant to Service Infrastructure uses, the `LoggingServiceV2` service.

## The LoggingServiceV2 service

The [LoggingServiceV2](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging.proto#L39) service is defined in [logging.proto](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging.proto).

Documented as a "service for ingesting and querying logs," this will be a close focus for us as we look at how we can programmatically read the logs written by Service Infrastructure. Its methods are used by `ServiceControl` to write logs that users can view in the Cloud Console or query using other methods of the service.

The full names of these methods begin with `google.logging.v2.LoggingServiceV2.`

| Method | Description |
| ------ | ----------- |
| [ListLogs](#listlogs) | Lists the logs in projects, organizations, folders, or billing accounts |
| [WriteLogEntries](#writelogentries) | Writes log entries to Logging |
| [ListLogEntries](#listlogentries) | Lists log entries |
| [TailLogEntries](#taillogentries) | Streaming read of log entries as they are ingested |
| [DeleteLog](#deletelog) | Deletes all the log entries in a log |
| [ListMonitoredResourceDescriptors](#listmonitoredresourcedescriptors) | Lists the descriptors for monitored resource types used by Logging |

### **LogEntries**

[LogEntry](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/log_entry.proto#L38) is the most significant resource that we'll deal with in this API. A `LogEntry` usually appears as a single line in the Logs Explorer but can also be programmaticlly retrieved with the API. It has many optional properties that include an associated HTTP request and an arbitrary payload that can be either a binary-encoded protobuf message, a JSON value, or text.

### ListLogs

We discuss `ListLogs` first, because it provides us with the list of logs in our project. You'll see from the list below that a variety of logs exist, and we see several that are created by Cloud Run (with `run.googleapis.com` in their name). The log for our service is named with the service name, followed by %2F (a url-encoded forward slash), and `endpoints_log`.

```prompt
q logging list-logs projects/bobadojo
```

```
projects/bobadojo/logs/cloudaudit.googleapis.com%2Factivity
projects/bobadojo/logs/cloudaudit.googleapis.com%2Fsystem_event
projects/bobadojo/logs/run.googleapis.com%2F%2Fvar%2Flog%2Fnginx%2Ferror.log
projects/bobadojo/logs/run.googleapis.com%2Frequests
projects/bobadojo/logs/run.googleapis.com%2Fstderr
projects/bobadojo/logs/run.googleapis.com%2Fstdout
projects/bobadojo/logs/run.googleapis.com%2Fvarlog%2Fsystem
projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log
```

### WriteLogEntries

`WriteLogEntries` can be used to directly add entries to a log. When we use Service Infrastructure, the `ServiceControl` `Report` method calls this for us. In general, this requires a log name, an optional monitored resource, and a list of `LogEntry` values. We won't demonstrate this here.

### ListLogEntries

We can read the entries in a log with `ListLogEntries`. This can return many values, so here we display only the two most recent. You'll see below that the logs include API keys. We've replaced them with `REDACTED` in the text below, but in the original response these were raw (unencrypted) API key strings.

```prompt
q logging list-log-entries bobadojo stores.endpoints.bobadojo.cloud.goog/endpoints_log --limit 2
```

```
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
      "timestamp": 1731798005.4496007
    }
  },
  "timestamp": {
    "seconds": 1731798005,
    "nanos": 449600742
  },
  "receive_timestamp": {
    "seconds": 1731798007,
    "nanos": 60366065
  },
  "severity": 200,
  "insert_id": "a42b02cf-4901-45ba-8197-5148bed715ef4850908905639612439@b16",
  "http_request": {
    "request_method": "POST",
    "request_url": "/bobadojo.stores.v1.Stores/GetStore",
    "request_size": 422,
    "status": 200,
    "response_size": 166,
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
      "timestamp": 1731798005.4459765
    }
  },
  "timestamp": {
    "seconds": 1731798005,
    "nanos": 445976439
  },
  "receive_timestamp": {
    "seconds": 1731798007,
    "nanos": 60366065
  },
  "severity": 200,
  "insert_id": "a42b02cf-4901-45ba-8197-5148bed715ef4850908905639612439@b15",
  "http_request": {
    "request_method": "POST",
    "request_url": "/bobadojo.stores.v1.Stores/GetStore",
    "request_size": 422,
    "status": 200,
    "response_size": 184,
    "remote_ip": "10.1.9.100",
    "latency": {},
    "protocol": "grpc"
  }
}
```

`ListLogEntries` supports a rich language for filtering queries. It's described in detail at [Logging query language](https://cloud.google.com/logging/docs/view/logging-query-language) and we're actually already using it in the `q` command to filter our query to only return entries from the `endpoints_log`.

To illustrate a tighter filter, let's ask for the most recent call to `ListStores`.

```prompt
q logging list-log-entries bobadojo stores.endpoints.bobadojo.cloud.goog/endpoints_log --limit 1 --filter ' AND httpRequest.request_url = "/bobadojo.stores.v1.Stores/ListStores"'
```

```
{
  "log_name": "projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log",
  "resource": {
    "type": "api",
    "labels": {
      "location": "global",
      "method": "bobadojo.stores.v1.Stores.ListStores",
      "project_id": "bobadojo",
      "service": "stores.endpoints.bobadojo.cloud.goog",
      "version": "v1"
    }
  },
  "Payload": {
    "JsonPayload": {
      "api_key": "REDACTED",
      "api_key_state": "VERIFIED",
      "api_method": "bobadojo.stores.v1.Stores.ListStores",
      "api_name": "bobadojo.stores.v1.Stores",
      "api_version": "v1",
      "grpc_status_code": "OK",
      "http_status_code": 200,
      "log_message": "bobadojo.stores.v1.Stores.ListStores is called",
      "producer_project_id": "bobadojo",
      "response_code_detail": "via_upstream",
      "service_agent": "q/0.0.0",
      "service_config_id": "2024-10-18r0",
      "timestamp": 1730763427
    }
  },
  "timestamp": {
    "seconds": 1730763427,
    "nanos": 347017499
  },
  "receive_timestamp": {
    "seconds": 1730763428,
    "nanos": 821073787
  },
  "severity": 200,
  "insert_id": "5b8e82e2-b78e-4b5f-9b87-fe0a269330b34850908905639612439@a1",
  "http_request": {
    "request_method": "GET",
    "request_url": "/bobadojo.stores.v1.Stores/ListStores",
    "request_size": 10,
    "status": 200,
    "response_size": 10,
    "remote_ip": "10.1.1.1",
    "latency": {
      "seconds": 5
    },
    "protocol": "grpc"
  }
}
```

As one more example, let's ask for the most recent request that resulted in a non-OK (200) response code.

```prompt
q logging list-log-entries bobadojo stores.endpoints.bobadojo.cloud.goog/endpoints_log --limit 1 --filter ' AND httpRequest.status != 200'
```

```
{
  "log_name": "projects/bobadojo/logs/stores.endpoints.bobadojo.cloud.goog%2Fendpoints_log",
  "resource": {
    "type": "api",
    "labels": {
      "location": "us-west1",
      "method": "bobadojo.stores.v1.Stores.GetStore",
      "project_id": "bobadojo",
      "service": "stores.endpoints.bobadojo.cloud.goog",
      "version": "v1"
    }
  },
  "Payload": {
    "JsonPayload": {
      "api_key_state": "NOT CHECKED",
      "api_method": "bobadojo.stores.v1.Stores.GetStore",
      "api_name": "bobadojo.stores.v1.Stores",
      "api_version": "v1",
      "error_cause": "Method doesn't allow unregistered callers (callers without established identity). Please use API Key or other form of API consumer identity to call this API.",
      "http_status_code": 401,
      "location": "us-west1",
      "log_message": "bobadojo.stores.v1.Stores.GetStore is called",
      "producer_project_id": "bobadojo",
      "response_code_detail": "service_control_bad_request{MISSING_API_KEY}",
      "service_agent": "ESPv2/2.48.0",
      "service_config_id": "2024-10-18r0",
      "timestamp": 1731778552.495956
    }
  },
  "timestamp": {
    "seconds": 1731778552,
    "nanos": 495955841
  },
  "receive_timestamp": {
    "seconds": 1731778554,
    "nanos": 622525552
  },
  "severity": 500,
  "insert_id": "9b772b44-9b91-403b-be73-7e2d932ed94e4850908905639612439@a1",
  "http_request": {
    "request_method": "GET",
    "request_url": "/v1/stores/0",
    "request_size": 444,
    "status": 401,
    "response_size": 342,
    "remote_ip": "169.254.169.126",
    "latency": {},
    "protocol": "http"
  },
  "trace": "projects/bobadojo/traces/a6283db966f75ae660d3e7b740a1779e"
}
```

### TailLogEntries

`TailLogEntries` is a streaming gRPC method that allows us to watch a log for new entries. We generally give it the same arguments as `ListLogEntries` (omitting [order_by](https://github.com/googleapis/googleapis/blob/master/google/logging/v2/logging.proto#L291), since ordering is implicit).

As an example, here's a `q` command that tails the log entries for our API, filtering to only show requests to `FindStores`:

```
q logging tail-log-entries bobadojo stores.endpoints.bobadojo.cloud.goog/endpoints_log --limit 1000 --filter ' AND httpRequest.request_url = "/bobadojo.stores.v1.Stores/FindStores"'
```

### DeleteLog

The `DeleteLog` method can be used to delete all of the entries in a log. The log should be identified by its full resource name. We won't demonstrate this here.

### ListMonitoredResourceDescriptors

Finally, `ListMonitoredResourceDescriptors` lists the types of resources that can be used as monitored resources in Logging. You can view a list of their names with the following `q` command, but results are omitted because they are large and not relevant to our discussion.

```prompt
q logging list-monitored-resource-descriptors | jq .[].type -r
```

## Summarizing

Cloud Logging is a powerful API that is built for large-scale logging. Service Control uses it to store logs for our managed services, and if we take some time to understand the log structure (see the examples above) and the [filtering language](https://cloud.google.com/logging/docs/view/logging-query-language), we can build our own powerful queries that help us understand how our services are running and being used.

---
#### Continue with [the Cloud Monitoring API](/docs/serviceinfrastructure/cloudmonitoring).