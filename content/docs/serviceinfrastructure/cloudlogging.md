---
weight: 4
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
