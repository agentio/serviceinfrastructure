---
weight: 3
title: The Service Control API
---
# The Service Control API

The Service Control API is used by the proxies, so you probably won't call it directly, but we describe it here so that we know exactly what the proxies are doing in case we want to make changes to them or build something new.

The Service Control API is used to check API requests from consumers and to monitor requests and responses.

## The Service Control API methods

The Service Control API is defined in the [googleapis](/docs/details/googleapis) repo in [servicecontrol.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/servicecontrol.yaml).
The methods specific to the API are defined 
by the [ServiceController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/service_controller.proto#L39) and
[QuotaController](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/quota_controller.proto#L37) services
in [service_controller.proto](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/service_controller.proto) and [quota_controller.proto](https://github.com/googleapis/googleapis/blob/master/google/api/servicecontrol/v1/quota_controller.proto).

Method names below are prefixed with `google.api.servicecontrol.v1.`

| Method | Description |
| ------ | ----------- |
| ServiceController.Check | Checks whether an operation on a service should be allowed to proceed based on the configuration of the service and related policies |
| QuotaController.AllocateQuota | Attempts to allocate quota for the specified consumer |
| ServiceController.Report | Reports operation results to Google Service Control, such as logs and metrics |

### Check

### AllocateQuota

### Report

## Usage Notes

Instead of calling these APIs individually, we'll set up a subcommand of `q` that pretends to be an API proxy and makes the Service Infrastructure calls for a sample API call. Our subcommand is called `q service-control mock`. Its key arguments are a service name, an operation name, and an API key that is being used to make our mock request.

We call it like this:
```
$ q service-control mock --service stores.endpoints.bobadojo.cloud.goog --operation bobadojo.stores.v1.Stores.ListStores --apikey $KEY
2024/11/04 15:37:06 calling check
{"operationId":"5b8e82e2-b78e-4b5f-9b87-fe0a269330b3","serviceConfigId":"2024-10-18r0","serviceRolloutId":"2024-11-04r1","checkInfo":{"unusedArguments":["caller_ip","private_caller_ip"],"consumerInfo":{"projectNumber":"1046800315646","type":"PROJECT","consumerNumber":"1046800315646"}}}
2024/11/04 15:37:07 calling allocate quota
{"operationId":"5b8e82e2-b78e-4b5f-9b87-fe0a269330b3","serviceConfigId":"2024-10-18r0"}
2024/11/04 15:37:07 constructing operation to report
2024/11/04 15:37:07 calling report
{"serviceConfigId":"2024-10-18r0","serviceRolloutId":"2024-11-04r4"}
```

What just happened? 

We just made the backend calls that a proxy would make to handle a request. First we call check to verify an API key, then allocate_quota to verify that the request is within quota limits, and then finally report to save logs and metrics.

In a minute or so, the results will show up in the log viewer.

![alt text](/screenshots/mock-log.png)

Compare this with logs from calls through our proxy and you'll see this is a pretty good match! Only the specific details vary because some of them are fake. But this demonstrates that we know how to call the Service Control APIs.

See the source code for [q service-control mock](https://github.com/agentio/q/blob/main/cmd/servicecontrol/mock.go) for details.

References:
- https://cloud.google.com/service-infrastructure/docs/service-control/getting-started