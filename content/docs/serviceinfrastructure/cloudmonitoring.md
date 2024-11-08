---
weight: 6
title: The Cloud Monitoring API
---
# The Cloud Monitoring API

The Cloud Monitoring API allows applications to read metrics describing API traffic.

## The Cloud Monitoring API Methods

The methods of the Cloud Monitoring API are defined in the [googleapis](/googleapis) repo in [google/monitoring/v3/metric_service.proto](https://github.com/googleapis/googleapis/blob/9b94dba2f7f4b601f8232bc3a3f6ef32665279b9/google/monitoring/v3/metric_service.proto#L67).

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