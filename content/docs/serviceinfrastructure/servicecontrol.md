---
weight: 3
title: The Service Control API
---
## The Service Control API

The Service Control API is used by the proxies, so you probably won't call it directly, but we describe it here so that we know exactly what the proxies are doing in case we want to make changes to them or to build something new.

The Service Control API is defined in the [googleapis](/docs/details/googleapis) repo in [servicecontrol.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/servicecontrol.yaml). It is used to check API requests from consumers and to monitor requests and responses. It includes two services:

| Service | Purpose |
| ------- | ------- |
| [ServiceController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/service_controller.proto#L39) | Check and report on requests made to a managed service |
| [QuotaController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/quota_controller.proto#L37) | Provide quota controls for requests a managed service | 

## The ServiceController service

The [ServiceController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/service_controller.proto#L39) service is defined in [service_controller.proto](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/service_controller.proto) and provides two methods that are used by proxies to check and report on requests to a managed service.

The full names of these methods begin with `google.api.servicecontrol.v1.ServiceController.`

| Method | Description |
| ------ | ----------- |
| [Check](#check) | Checks whether an operation on a service should be allowed to proceed based on the configuration of the service and related policies |
| [Report](#report) | Reports operation results to Google Service Control, such as logs and metrics |

### Check

The `Check` method should be called by proxies after an operation is received and before it is processed. Typically `Check` request messages are small and contain an operation id, an API key, and just a few other details.

### Report

`Report` is called by proxies after operations are processed. `Report` request messages include a significant amount of detail about the processed API including information to be logged and tracked in system-defined metrics.

## The QuotaController service

The [QuotaController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/quota_controller.proto#L37) service is defined in [quota_controller.proto](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/quota_controller.proto) and provides a single method that is used by proxies to implement quotas on managed services.

The full names of these methods begin with `google.api.servicecontrol.v1.QuotaController.`

| Method | Description |
| ------ | ----------- |
| [AllocateQuota](#allocatequota) | Attempts to allocate quota for the specified consumer |

### AllocateQuota

`AllocateQuota` is called by proxies after `Check` and before `Report`. It sends "quota metrics" which are charged against a user's quota and its response indicates whether an operation should proceed or fail with an out-of-quota error.

## Usage Notes

### Mocking calls to Service Control

Instead of calling these APIs individually, we've set up a subcommand of `q` that pretends to be an API proxy and makes the Service Infrastructure calls for a sample API call. Our subcommand is called `q service-control mock`. Its key arguments are a service name, an operation name, and an API key that is being used to make our mock request.

We call it like this:
```prompt
q service-control mock --service stores.endpoints.bobadojo.cloud.goog --operation bobadojo.stores.v1.Stores.ListStores --apikey $KEY
```

```
2024/11/04 15:37:06 calling check
{"operationId":"5b8e82e2-b78e-4b5f-9b87-fe0a269330b3","serviceConfigId":"2024-10-18r0","serviceRolloutId":"2024-11-04r1","checkInfo":{"unusedArguments":["caller_ip","private_caller_ip"],"consumerInfo":{"projectNumber":"1046800315646","type":"PROJECT","consumerNumber":"1046800315646"}}}
2024/11/04 15:37:07 calling allocate quota
{"operationId":"5b8e82e2-b78e-4b5f-9b87-fe0a269330b3","serviceConfigId":"2024-10-18r0"}
2024/11/04 15:37:07 constructing operation to report
2024/11/04 15:37:07 calling report
{"serviceConfigId":"2024-10-18r0","serviceRolloutId":"2024-11-04r4"}
```

What just happened? 

We just made the backend calls that a proxy would make to handle a request. First we call `Check` to verify an API key, then `AllocateQuota` to verify that the request is within quota limits, and then finally we call `Report` to save logs and metrics.

In a minute or so, the results will show up in the log viewer.

![alt text](/screenshots/mock-log.png)

Compare this with logs from calls through our proxy and you'll see this is a pretty good match! Only the specific details vary because some of them are fake. But this demonstrates that we know how to call the Service Control APIs.

See the source code for [q service-control mock](https://github.com/agentio/q/blob/main/cmd/servicecontrol/mock.go) for details.

### Service Control v2

In this section, we've been describing the Service Control v1 API, the API used by the [Extensible Service Proxies](/docs/proxies). [Getting Started with the Service Control API](https://cloud.google.com/service-infrastructure/docs/service-control/getting-started) describes a new version of Service Control that is currently in Private Preview and [only available for approved services](https://cloud.google.com/service-infrastructure/docs/service-control/reference/rpc/google.api/servicecontrol.v2). It's unclear if or when this will be made available for general use, and since the proxies use the v1 API, we can expect that to continue to be available. The two versions are very similar, so if we ever do need to update, we expect that will be easy.

## Summarizing

Here we've seen how it works, but for a typical Service Infrastructure user, the Service Control API is a behind-the-scenes helper, taking requests from the [Extensible Service Proxy](/docs/proxies), checking them with the [API Keys API](/docs/serviceinfrastructure/apikeys), logging them with the [Cloud Logging API](/docs/serviceinfrastructure/cloudlogging), and collecting metrics on them with the [Cloud Monitoring API](/docs/serviceinfrastructure/cloudmonitoring). We'll look at these three APIs next to see how we can get the most out of our Service Infrastructure usage.

---
#### Continue with [the API Keys API](/docs/serviceinfrastructure/apikeys).