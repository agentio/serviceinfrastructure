---
weight: 2
title: The Service Management API
---
## The Service Management API

API management starts with knowing what your APIs are, and the Service Management API is used to build a list and digest it into a form that allows API proxies to check and report API traffic.

The Service Management API is defined in the [googleapis](/docs/details/googleapis) repo in [servicemanagement_v1.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/servicemanagement/v1/servicemanagement_v1.yaml). It includes three services:

| Service | Purpose |
| ------- | ------- |
| [ServiceManager](#the-servicemanager-service) | Methods for managing service descriptions |
| [IAMPolicyService](#the-iampolicyservice-service) | A mix-in that configures access and sharing |
| [Operations](#the-operations-service) | A mix-in that handles long-running operations |

## The ServiceManager service

The [ServiceManager](https://github.com/googleapis/googleapis/blob/master/google/api/servicemanagement/v1/servicemanager.proto#L39) service ([Google documentation](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.api/servicemanagement.v1)) is defined in [servicemanager.proto](https://github.com/googleapis/googleapis/blob/master/google/api/servicemanagement/v1/servicemanager.proto) and
manages descriptions of APIs, focusing on the service configurations that control their usage.

The full names of these methods begin with `google.api.servicemanagement.v1.ServiceManager.`

| Method | Description |
| ------ | ----------- |
| [ListServices](#listservices) | Lists managed services |
| [GetService](#getservice) | Gets a managed service |
| [CreateService](#createservice) | Creates a new managed service |
| [DeleteService](#deleteservice) | Deletes a managed service |
| [UndeleteService](#undeleteservice) | Revives a previously deleted managed service |
| [ListServiceConfigs](#listserviceconfigs) | Lists the history of the service configuration for a managed service, from the newest to the oldest |
| [GetServiceConfig](#getserviceconfig) | Gets a service configuration (version) for a managed service |
| [CreateServiceConfig](#createserviceconfig) | Creates a new service configuration (version) for a managed service |
| [SubmitConfigSource](#submitconfigsource)  | Creates a new service configuration for a managed service based on user-supplied configuration source files |
| [ListServiceRollouts](#listservicerollouts) | Lists the history of the service configuration rollouts for a managed service, from the newest to the oldest |
| [GetServiceRollout](#getservicerollout) | Gets a service configuration rollout |
| [CreateServiceRollout](#createservicerollout) | Creates a new service configuration rollout |
| [GenerateConfigReport](#generateconfigreport) | Generates and returns a report of errors, warnings and changes from existing configurations |

The ServiceManager methods can be divided into three groups:
1. Managing Services. These methods handle the creation and deletion of registered services.
2. Managing Service Configurations. These methods handle Service Configurations and allow multiple revisions of configurations to be tracked for each service.
3. Managing Service Rollouts. These methods control how service configurations are deployed. New configurations can be rolled out all at once or fractionally.

To get to know them better, we'll call them with [q](/the-q-tool) for the example that we set up in the [quickstart](/docs/quickstart/demo).

### **Managed Services**

This first group of APIs operates on [ManagedService](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/servicemanagement/v1/resources.proto#L32C1-L42C2) resources. A `ManagedService` is very simple: just a service name and the id of the "producer project". The producer project is the Google Cloud project that owns the service listing. For us, it's the project that we've configured using `gcloud config set project PROJECTID`. So essentially, these methods are just for managing a list of services that we'll be managing.

### ListServices

`ListServices` lists the services that have been created within a project.

Let's try it. Like most methods, `ListServices` returns JSON, so we'll use `jq` to format the results.

```
$ q service-management list-services bobadojo | jq
[
  {
    "serviceName": "stores.endpoints.bobadojo.cloud.goog",
    "producerProjectId": "bobadojo"
  }
]
```

`jq` is pretty handy and worth getting to know (start [here](https://jqlang.github.io/jq/)). We can use `jq` to pull specific values from a JSON response. Here we'll use `jq` to get the name of the first service and set it to a variable.

```
$ SERVICE=$(q service-management list-services bobadojo  | jq .[0].serviceName -r)
$ echo $SERVICE
stores.endpoints.bobadojo.cloud.goog
```

### GetService

`GetService` gets details of a service. It returns a `ManagedService`, which doesn't tell us anything that we don't already know from calling `ListServices`, but let's try it to see it in action.

```
$ q service-management get-service $SERVICE | jq
{
  "serviceName": "stores.endpoints.bobadojo.cloud.goog",
  "producerProjectId": "bobadojo"
}
```

### CreateService 

We can create a new service with `CreateService`. This just takes a `ManagedService` as an argument, so we don't need to specify much, just a project id and a service id: 

```
$ q service-management create-service --help
Create service

Usage:
  q service-management create-service PROJECT SERVICE [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for create-service
```

But we have to be careful with our choice of service name. Let's try using "sample" as the service name to see what we get:

```
$ q service-management create-service bobadojo sample
Error: rpc error: code = PermissionDenied desc = Ownership for domain name 'sample' on project 'bobadojo' cannot be verified.
Usage:
  q service-management create-service PROJECTID SERVICE [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for create-service
```

That's an error, and from the error message, it seems that Google is expecting the service name to be a domain name.

What names are we allowed to use? We've already seen that we can use names of the form `SERVICEID.endpoints.PROJECTID.cloud.goog`. This is because Google controls the `cloud.goog` domain and makes its subdomains available for Endpoints customers. Also, Google App Engine users can access their services at `PROJECTID.appspot.com` and Google allows us to use these domains and their subdomains as service names.

We demonstrate this in the examples below by creating services that are subdomains of `.cloud.goog` and `.appspot.com`.

```
$ q service-management create-service sample.endpoints.bobadojo.cloud.goog bobadojo
service=Response { metadata: MetadataMap { headers: {"content-disposition": "attachment", "content-type": "application/grpc", "x-debug-tracking-id": "7176991653359097728;o=0", "date": "Sat, 13 Jul 2024 03:40:06 GMT", "alt-svc": "h3=\":443\"; ma=2592000,h3-29=\":443\"; ma=2592000", "grpc-status": "0"} }, message: Operation { name: "operations/services.sample.endpoints.bobadojo.cloud.goog-0", metadata: Some(Any { type_url: "type.googleapis.com/google.api.servicemanagement.v1.OperationMetadata", value: [10, 42, 115, 101, 114, 118, 105, 99, 101, 115, 47, 102, 111, 111, 46, 101, 110, 100, 112, 111, 105, 110, 116, 115, 46, 98, 111, 98, 97, 100, 111, 106, 111, 46, 99, 108, 111, 117, 100, 46, 103, 111, 111, 103, 34, 11, 8, 150, 238, 199, 180, 6, 16, 248, 207, 151, 118] }), done: false, result: Some(Response(Any { type_url: "type.googleapis.com/google.api.servicemanagement.v1.ManagedService", value: [18, 33, 102, 111, 111, 46, 101, 110, 100, 112, 111, 105, 110, 116, 115, 46, 98, 111, 98, 97, 100, 111, 106, 111, 46, 99, 108, 111, 117, 100, 46, 103, 111, 111, 103, 26, 8, 98, 111, 98, 97, 100, 111, 106, 111] })) }, extensions: Extensions }

$ q service-management create-service sample.bobadojo.appspot.com bobadojo
service=Response { metadata: MetadataMap { headers: {"content-disposition": "attachment", "content-type": "application/grpc", "x-debug-tracking-id": "9253787921377858843;o=0", "date": "Sat, 13 Jul 2024 03:40:22 GMT", "alt-svc": "h3=\":443\"; ma=2592000,h3-29=\":443\"; ma=2592000", "grpc-status": "0"} }, message: Operation { name: "operations/services.sample.bobadojo.appspot.com-0", metadata: Some(Any { type_url: "type.googleapis.com/google.api.servicemanagement.v1.OperationMetadata", value: [10, 33, 115, 101, 114, 118, 105, 99, 101, 115, 47, 102, 111, 111, 46, 98, 111, 98, 97, 100, 111, 106, 111, 46, 97, 112, 112, 115, 112, 111, 116, 46, 99, 111, 109, 34, 11, 8, 166, 238, 199, 180, 6, 16, 136, 218, 150, 112] }), done: false, result: Some(Response(Any { type_url: "type.googleapis.com/google.api.servicemanagement.v1.ManagedService", value: [18, 24, 102, 111, 111, 46, 98, 111, 98, 97, 100, 111, 106, 111, 46, 97, 112, 112, 115, 112, 111, 116, 46, 99, 111, 109, 26, 8, 98, 111, 98, 97, 100, 111, 106, 111] })) }, extensions: Extensions }
```

We aren't required to actually serve our APIs on these domains, but because Google controls both `cloud.goog` and `appspot.com`, we can get Google-hosted DNS if we use them. (This was announced in the [April 2017 Endpoints release notes](https://cloud.google.com/endpoints/docs/pricing-and-quotas#april_2017).)

But how does Google know what address to use for these names? The answer is that we specify that in Service Config, specifically in the [endpoints](/docs/serviceinfrastructure/serviceconfig/#endpoints) section. Here's an example for our test API, where we've substituted `WWW.XXX.YYY.ZZZ` for the actual endpoint address.

```
  "endpoints": [
    {
      "name": "stores.endpoints.bobadojo.cloud.goog",
      "target": "WWW.XXX.YYY.ZZZ"
    }
  ],
```
Now with this, Google serves DNS records that point to our server.
```
$ host stores.endpoints.bobadojo.cloud.goog
stores.endpoints.bobadojo.cloud.goog has address WWW.XXX.YYY.ZZZ

$ dig stores.endpoints.bobadojo.cloud.goog

; <<>> DiG 9.20.0-2ubuntu3-Ubuntu <<>> stores.endpoints.bobadojo.cloud.goog
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 53665
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;stores.endpoints.bobadojo.cloud.goog. IN A

;; ANSWER SECTION:
stores.endpoints.bobadojo.cloud.goog. 60 IN A	WWW.XXX.YYY.ZZZ

;; Query time: 284 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Mon Nov 04 11:02:31 PST 2024
;; MSG SIZE  rcvd: 81
```

To read more about service names and DNS, see [Configuring DNS on the cloud.goog domain](https://cloud.google.com/endpoints/docs/openapi/cloud-goog-dns-configure). Also, Google has additional suggestions for choosing service names in [Planning Your Cloud Projects](https://cloud.google.com/endpoints/docs/openapi/planning-cloud-projects).

Before we continue, let's explore one more way to name our services. In the original error message, we were told that `Ownership for domain name 'sample' on project 'bobadojo' cannot be verified.` So if we can prove to Google that we own a domain, can we use it as a service name?

The answer to that is yes, and we'll verify that here with an example. By following [these instructions](https://cloud.google.com/endpoints/docs/grpc/verify-domain-name), the `bobadojo` project was verified to own the `bobadojo.io` domain. As a test, we'll attempt to create a service named `sample.bobadojo.io`.

```
$ q service-management create-service bobadojo sample.bobadojo.io
operations/services.sample.bobadojo.io-0
```

That looks good. It returned an operation ID, and when the operation completes, we should have our new service.

If we check too soon, we'll get an error:
```
$ q service-management get-service sample.bobadojo.io
Error: rpc error: code = PermissionDenied desc = Service 'sample.bobadojo.io' not found or permission denied.
Usage:
  q service-management get-service SERVICE [flags]

Flags:
      --format string    output format (default "json")
  -h, --help             help for get-service
  -p, --project string   producer project
```

But after a minute or so, our new service is registered:
```
$ q service-management get-service sample.bobadojo.io 
{"serviceName":"sample.bobadojo.io","producerProjectId":"bobadojo"}
```

But do we get DNS for this? No, because `bobadojo.io` is controlled by another registrar.

```
$ host sample.bobadojo.com
Host sample.bobadojo.com not found: 3(NXDOMAIN)

$ dig sample.bobadojo.com

; <<>> DiG 9.20.0-2ubuntu3-Ubuntu <<>> sample.bobadojo.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 58636
;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;sample.bobadojo.com.		IN	A

;; AUTHORITY SECTION:
bobadojo.com.		10800	IN	SOA	dns1.namecheaphosting.com. cpanel.tech.namecheap.com. 1715572172 86400 7200 3600000 86400

;; Query time: 191 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Mon Nov 04 11:14:08 PST 2024
;; MSG SIZE  rcvd: 128
```

Of course, if we want to use this domain, we would add appropriate records with the registrar that we use to manage our domain.

### DeleteService

The last thing that we might want to do with a service is delete it... or undelete it if we change our minds! The Service Management API keeps deleted services in a "soft deleted" state for 30 days so that they can be undeleted.

To see this, let's first call the API that we set up in the [quickstart](/docs/quickstart/demo). Here `HOST` is set to the URL of the server that we saw in the Cloud Run console.
```
$ curl $HOST/v1/stores/0 -s -H "X-Api-Key: $KEY" | jq
{
  "name": "stores/0",
  "type": "office",
  "title": "Columbus, NM 88029",
  "location": {
    "latitude": 31.8301201,
    "longitude": -107.638199
  },
  "address": {
    "street": "South Main Street",
    "regionCode": "us"
  }
}
```

Now let's delete our service.
```
$ q service-management delete-service $SERVICE
operations/services.stores.endpoints.bobadojo.cloud.goog-26
```
That returned an operation. We have to wait for it to complete, but when it does, our service will be deleted.
```
$ q service-management get-service $SERVICE
Error: rpc error: code = PermissionDenied desc = Service 'stores.endpoints.bobadojo.cloud.goog' not found or permission denied.
Usage:
  q service-management get-service SERVICE [flags]

Flags:
      --format string    output format (default "json")
  -h, --help             help for get-service
  -p, --project string   producer project
```

Now if we try to call it, the requests will fail.
```
$ curl $HOST/v1/stores/0 -s -H "X-Api-Key: $KEY" | jq
{
  "code": 500,
  "message": "INTERNAL: Calling Google Service Control API failed with: 403 and body: \b\u0007\u0012_Permission 'servicemanagement.services.check' denied on service '<redacted_3rd_party_service>'."
}
```

### UndeleteService

Now let's change our minds and restore our deleted service.
```
$ q service-management undelete-service $SERVICE
operations/services.stores.endpoints.bobadojo.cloud.goog-27
```

Our service quickly becomes available again.
```
$ curl $HOST/v1/stores/0 -s -H "X-Api-Key: $KEY" | jq
{
  "name": "stores/0",
  "type": "office",
  "title": "Columbus, NM 88029",
  "location": {
    "latitude": 31.8301201,
    "longitude": -107.638199
  },
  "address": {
    "street": "South Main Street",
    "regionCode": "us"
  }
}
```

How did this respond so quickly? Our proxy is configured to use the "managed" rollout strategy, which causes it to check the Service Management API once every minute to see if the service configuration has changed. For more on this, see [Cloud Endpoints: Introducing a new way to manage API configuration rollout](https://cloud.google.com/blog/products/gcp/cloud-endpoints-introducing-a-new-way-to-manage-api-configuration-rollout).

### **Service Configurations**

The Managed Service records that we saw in the previous section don't contain much: just the name of a service and the owning project id. We actually configure our services with a separate resource, the "Service Configuration", and this is the "Service Config" that we discussed in the [Service Config](/docs/serviceinfrastructure/serviceconfig) section.

Separating `Service` from `ManagedService` has benefits:
- We can keep several revisions of `Service`; in fact, the Service Management API keeps the entire revision history.
- We can deploy multiple configurations of a service at once. The next group of methods manages rollouts, and rollouts can divide services between multiple configurations.
- We can (and do) have tighter access controls on `Service` configuration. We can list `ManagedService` resources outside of our control, but we can only see the `Service` resources that we own or have been explicitly granted access.

### ListServiceConfigs

`ListServiceConfigs` lists the service configurations associated with an service.

Try it.
```
$ q service-management list-service-configs $SERVICE
```

Yikes! This generates a lot of output. Let's use `jq` to just get one field out of the response, the ids of each service config:
```
$ q service-management list-service-configs $SERVICE | jq .[].id -r
2024-10-18r0
2024-10-17r0
2024-10-15r1
2024-10-15r0
2024-10-14r1
2024-10-14r0
2024-07-13r0
2024-07-12r11
2024-07-12r10
2024-07-12r9
2024-07-12r8
2024-07-12r7
2024-07-12r6
2024-07-12r5
2024-07-12r4
2024-07-12r3
2024-07-12r2
2024-07-12r1
2024-07-12r0
```

This is still a lot of output, but when you run it on your own service, it will probably only be one or two lines. This is showing the ids of every revision of service configuration that has been uploaded for this service. Let's set `CONFIG_ID` to the first one, which is also the most recent.
```
CONFIG_ID=$(q service-management list-service-configs $SERVICE | jq .[0].id -r)
```

### GetServiceConfig

`GetServiceConfig` gets an individual service configuration. We got a subset of the service config from the list call. This method will return everything, and if we set the `view` field to `FULL` in the [GetServiceConfigRequest](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/servicemanagement/v1/servicemanager.proto#L336) message, responses also include any source files that were uploaded when the configuration was created. We can do that with a command-line flag to `q`.

```
$ q service-management get-service-config $SERVICE $CONFIG_ID --full | jq
```
This returns a lot of output, and if you run it, you'll see that the source files are included as base64 strings. In this case, one of them is binary (`descriptor.pb`) but the other (`api_config.yaml`) is text and we can get it with `jq` and the `base64` command-line tool.
```
$ q service-management get-service-config $SERVICE $CONFIG_ID --full | jq .sourceInfo.sourceFiles.[0].fileContents -r | base64 -d
type: google.api.Service
config_version: 3

#
# Name of the service configuration.
#
name: stores.endpoints.bobadojo.cloud.goog

#
# API title to appear in the user interface (Google Cloud Console).
#
title: Boba Dojo Stores API
apis:
- name: bobadojo.stores.v1.Stores

#
# API usage restrictions.
#
usage:
  rules:
  # ListStores methods can be called without an API Key.
  - selector: bobadojo.stores.v1.Stores.ListStores
    allow_unregistered_calls: true
  - selector: bobadojo.stores.v1.Stores.GetStore
    allow_unregistered_calls: false
  - selector: bobadojo.stores.v1.Stores.FindStores
    allow_unregistered_calls: false

endpoints:
- name: stores.endpoints.bobadojo.cloud.goog
  target: "WWW.XXX.YYY.ZZZ"
```

### CreateServiceConfig

`CreateServiceConfig` is the first of two ways to create service configs. This one posts the full service config for a service.

```
q service-management service-config create-service-config
```

TODO: finish this

### SubmitConfigSource

The second way to create a service configuration is using `SubmitConfigSource` is to upload source files that include partial service config, protocol buffer file descriptor sets, and OpenAPI descriptions. Google will then compile these into the full service config.

```
q service-management submit-config-source stores.endpoints.bobadojo.cloud.goog ./stores-demo/api_config.yaml  ./stores-demo/descriptor.pb 

q service-management get-operation operations/serviceConfigs.stores.endpoints.bobadojo.cloud.goog:537024da-fdca-4c1c-8c33-5dbf1eb13e1a | jq .response.serviceConfig | head
```

TODO: finish this

### GenerateConfigReport

The `ServiceManager` service includes a function that we can use to compare two service configurations. By default it compares a configuration with its predecessor.

```
q service-management generate-config-report services/$SERVICE/configs/$CONFIG
```

Here's the report for a change that temporarily removed the API key requirement from the `FindStores` and `GetStore` API methods.
```
q service-management generate-config-report services/$SERVICE/configs/2024-10-14r0 | jq
{
  "serviceName": "stores.endpoints.bobadojo.cloud.goog",
  "id": "2024-10-14r0",
  "changeReports": [
    {
      "configChanges": [
        {
          "element": "usage.rules[selector==\"bobadojo.stores.v1.Stores.FindStores\"].allow_unregistered_calls",
          "newValue": "true",
          "changeType": "ADDED"
        },
        {
          "element": "usage.rules[selector==\"bobadojo.stores.v1.Stores.GetStore\"].allow_unregistered_calls",
          "newValue": "true",
          "changeType": "ADDED"
        }
      ]
    }
  ]
}
```

### **Service Rollouts**

The last few methods in the `ServiceManager` service handle rollouts. [Rollout](https://github.com/googleapis/googleapis/blob/d54f4e947e77b86ea2e0e243c92a174032098a54/google/api/servicemanagement/v1/resources.proto#L187) resources put service configurations into action. Rollouts are created automatically by calls to `gcloud endpoints deploy`, but when we create new service configurations with the `ServiceManager` service, we need to create rollouts explicitly.

### ListServiceRollouts

First let's list rollouts with `ListServiceRollouts`. The results below are truncated, because when this command is run, it returns the full history of rollouts for a service.

```
$ q service-management list-service-rollouts $SERVICE | jq
[
  {
    "rolloutId": "2024-11-04r1",
    "createTime": "2024-11-04T19:27:46.453Z",
    "createdBy": "timburks@gmail.com",
    "status": "SUCCESS",
    "trafficPercentStrategy": {
      "percentages": {
        "2024-10-18r0": 100
      }
    },
    "serviceName": "stores.endpoints.bobadojo.cloud.goog"
  },
  {
    "rolloutId": "2024-11-04r0",
    "createTime": "2024-11-04T19:25:28.279Z",
    "createdBy": "timburks@gmail.com",
    "status": "SUCCESS",
    "deleteServiceStrategy": {},
    "serviceName": "stores.endpoints.bobadojo.cloud.goog"
  },
  {
    "rolloutId": "2024-10-18r0",
    "createTime": "2024-10-18T04:48:51.188Z",
    "createdBy": "timburks@gmail.com",
    "status": "SUCCESS",
    "trafficPercentStrategy": {
      "percentages": {
        "2024-10-18r0": 100
      }
    },
    "serviceName": "stores.endpoints.bobadojo.cloud.goog"
  },
  ...
```

### GetServiceRollout

We can use `GetServiceRollout` to get an individual rollout by its id. This doesn't return anything that we didn't get from listing them, but when we look inside, we see that the `trafficPercentStrategy` refers to service config ids. For Cloud Endpoints, the percentages are always 100.

```
$ q service-management get-service-rollout $SERVICE 2024-11-04r1 | jq
{
  "rolloutId": "2024-11-04r1",
  "createTime": "2024-11-04T19:27:46.453Z",
  "createdBy": "timburks@gmail.com",
  "status": "SUCCESS",
  "trafficPercentStrategy": {
    "percentages": {
      "2024-10-18r0": 100
    }
  },
  "serviceName": "stores.endpoints.bobadojo.cloud.goog"
}
```

### CreateServiceRollout

We can create a new rollout with `CreateServiceRollout`. We just need to specify the service id and the id of the service config to be rolled out.

```
$ q service-management create-service-rollout $SERVICE 2024-10-18r0
operations/rollouts.stores.endpoints.bobadojo.cloud.goog:20355192-1186-4f42-a267-d76abb4dc35a
```

This creates a long running operation that will complete quickly, and our proxies will then start serving using the new configuraiton.

## The IAMPolicyService service

The [IAMPolicyService](https://github.com/googleapis/googleapis/blob/master/google/iam/v1/iam_policy.proto#L59) ([Google documentation](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.iam.v1)) is defined in [iam_policy.proto](https://github.com/googleapis/googleapis/blob/master/google/iam/v1/iam_policy.proto). It allows us to set IAM policies that allow other Google Cloud users to see the managed services that we create and enable them for use in their projects. When a user has enabled your service, they will be able to create API keys that they can use to call your service. When you check those keys with `ServiceControl`, you'll see which project owns them.

The full names of these methods begin with `google.iam.v1.IamPolicyService.`

| Method | Description |
| ------ | ----------- |
| [GetIamPolicy](#getiampolicy) | Gets the access control policy for a resource |
| [SetIamPolicy](#setiampolicy) | Sets the access control policy on the specified resource |
| [TestIamPermissions](#testiampermissions) | Returns permissions that a caller has on the specified resource |

### **Policies**

[Policies](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.iam.v1#policy) specify access controls, and the Service Management API uses them to give users access to services. A policy contains a list of "bindings" that grant a role to a list of users. 

Roles are collections of permissions, and permissions are typically fine-grained descriptions of capabilities. For the Service Management API, the supported roles and permissions are listed in [Service Management API Access Control](https://cloud.google.com/service-infrastructure/docs/service-management/access-control).

The specific role that we grant to share a service is `servicemanagement.serviceConsumer` and we refer to it as `roles/servicemanagement.serviceConsumer` in the binding. It includes a single permission, `servicemanagement.services.bind`. We'll see examples in the method descriptions below.

### GetIamPolicy

`GetIamPolicy` returns the policy associated with a service, which can be empty if nothing has been specified yet.

```
$ q service-management get-iam-policy services/stores.endpoints.bobadojo.cloud.goog | jq 
{
  "etag":"ACAB"
}
```

After we've added a few users, our policy might look like this:

```
$ q service-management get-iam-policy services/stores.endpoints.bobadojo.cloud.goog | jq 
{
  "version": 1,
  "bindings": [
    {:
      "role": "roles/servicemanagement.serviceConsumer",
      "members": [
        "user:tim@mitra.so",
        "user:tim@radtastical.com"
      ]
    }
  ],
  "etag": "BwYm5Z23vKM="
}
```

Next we'll see how to add users.

### SetIamPolicy

`SetIamPolicy` allows us to update the IAM policy associated with a service.

To try it, place the following in a file named `policy.json`, ideally using a user that you would like to share your service with:
```
{
  "version": 1,
  "bindings": [
    {
      "role": "roles/servicemanagement.serviceConsumer",
      "members": [
        "user:tim@mitra.so"
      ]
    }
  ]
}
```

Now use `q` to call `SetIamPolicy`:

```
$ q service-management set-iam-policy services/stores.endpoints.bobadojo.cloud.goog policy.json | jq
{
  "version": 1,
  "bindings": [
    {
      "role": "roles/servicemanagement.serviceConsumer",
      "members": [
        "user:tim@mitra.so"
      ]
    }
  ],
  "etag": "BwYm+t1o3Y4="
}
```

We can also add Google groups as members:

```
q service-management set-iam-policy services/stores.endpoints.bobadojo.cloud.goog policy.json | jq
{
  "version": 1,
  "bindings": [
    {
      "role": "roles/servicemanagement.serviceConsumer",
      "members": [
        "group:agentio@googlegroups.com",
        "user:tim@mitra.so"
      ]
    }
  ],
  "etag": "BwYm+uvcW4c="
}
```

But one thing that we can't do is add `allUsers` to make our API public. Here we try using `gcloud`, which makes the same calls that we've been making with `q`.

```
$ gcloud endpoints services add-iam-policy-binding stores.endpoints.bobadojo.cloud.goog --member=allUsers --role=roles/servicemanagement.serviceConsumer 
ERROR: Policy modification failed. For a binding with condition, run "gcloud alpha iam policies lint-condition" to identify issues in condition.
ERROR: (gcloud.endpoints.services.add-iam-policy-binding) INVALID_ARGUMENT: Member 'allUsers' is not allowed in role 'roles/servicemanagement.serviceConsumer'.
```

This is mentioned in [Cloud Endpoints Portal public access](https://groups.google.com/g/google-cloud-endpoints/c/-GrnDqyauiI).

### TestIamPermissions

The `TestIamPermissions` method allows us to see if one or more permissions is granted to the calling account.

To make these calls, we temporarily change our account by logging in again with `gcloud auth login`, in this case to the account that we used to share our service (here it's `tim@mitra.so`).

First we check the `servicemanagement.services.check` permission. We expect this first request to confirm that this permission is not included in the `servicemanagement.serviceConsumer` role.

```
$ q service-management test-iam-permissions services/stores.endpoints.bobadojo.cloud.goog servicemanagement.services.check
{}
```

Next we check the `servicemanagement.services.bind` permission. We expect this first request to return `servicemanagement.services.bind` since this permission is included in the `servicemanagement.serviceConsumer` role.

```
$ q service-management test-iam-permissions services/stores.endpoints.bobadojo.cloud.goog servicemanagement.services.bind
{"permissions":["servicemanagement.services.bind"]}
```

## The Operations service

The [Operations](https://github.com/googleapis/googleapis/blob/master/google/longrunning/operations.proto#L55) service ([Google Documentation](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.longrunning)) is defined in [operations.proto](https://github.com/googleapis/googleapis/blob/master/google/longrunning/operations.proto). It allows us to check the status of operations that complete asynchronously after their initiating request completes.

The full names of these methods begin with `google.longrunning.Operations.`

| Method | Description |
| ------ | ----------- |
| [ListOperations](#listoperations) | Lists operations that match the specified filter in the request |
| [GetOperation](#getoperation) | Gets the latest state of a long-running operation |
| [DeleteOperation](#deleteoperation) | Deletes a long-running operation |
| [CancelOperation](#canceloperation) | Starts asynchronous cancellation on a long-running operation |
| [WaitOperation](#waitoperation) |  Waits until the specified long-running operation is done or reaches at most a specified timeout, returning the latest state |

### **Operations**

`Operation` resources can be used to manage operations that complete asynchronously after their initiating request completes.

Quoting [Google documentation](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.longrunning#operations): "When an API method normally takes long time to complete, it can be designed to return Operation to the client, and the client can use this interface to receive the real response asynchronously by polling the operation resource, or pass the operation resource to another API (such as Google Cloud Pub/Sub API) to receive the response. Any API service that returns long-running operations should implement the Operations interface so developers can have a consistent client experience."

To get a sample operation, we will call the `CreateService` method, which returns an operation.

```
$ q service-management create-service bobadojo example.endpoints.bobadojo.cloud.goog
operations/services.example.endpoints.bobadojo.cloud.goog-0
```

`operations/services.example.endpoints.bobadojo.cloud.goog-0` is the name of the operation.

### ListOperations

`ListOperations` lets us see the operations that are currently active.

```
$ q service-management list-operations "" | jq
Error: rpc error: code = InvalidArgument desc = Must specify a filter.
```

That failed because this method requires a filter. We can learn more about that in the [ListOperationsRequest](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.longrunning#google.longrunning.ListOperationsRequest) documentation. Here we'll use a filter to select the service that we just attempted to create. The filter expression is `serviceName=example.endpoints.bobadojo.cloud.goog`.

```
$ q service-management list-operations "serviceName=example.endpoints.bobadojo.cloud.goog" | jq
[
  {
    "name": "operations/services.example.endpoints.bobadojo.cloud.goog-0",
    "metadata": {
      "@type": "type.googleapis.com/google.api.servicemanagement.v1.OperationMetadata",
      "resourceNames": [
        "services/example.endpoints.bobadojo.cloud.goog"
      ],
      "progressPercentage": 99,
      "startTime": "2024-11-15T22:37:25.944856Z"
    },
    "response": {
      "@type": "type.googleapis.com/google.api.servicemanagement.v1.ManagedService",
      "serviceName": "example.endpoints.bobadojo.cloud.goog",
      "producerProjectId": "bobadojo"
    }
  }
]
```
That worked, and we see that our operation is nearly (99%) complete.

### GetOperation

`GetOperation` lets us check on the status of a running operation. We can use it to directly check the operation that was returned above. We just provide the operation name that was returned from the original `CreateService` call.

```
q service-management get-operation operations/services.example.endpoints.bobadojo.cloud.goog-0 | jq
{
  "name": "operations/services.example.endpoints.bobadojo.cloud.goog-0",
  "metadata": {
    "@type": "type.googleapis.com/google.api.servicemanagement.v1.OperationMetadata",
    "resourceNames": [
      "services/example.endpoints.bobadojo.cloud.goog"
    ],
    "progressPercentage": 99,
    "startTime": "2024-11-15T22:37:25.944856Z"
  },
  "response": {
    "@type": "type.googleapis.com/google.api.servicemanagement.v1.ManagedService",
    "serviceName": "example.endpoints.bobadojo.cloud.goog",
    "producerProjectId": "bobadojo"
  }
}
```

Here again we see that our operation is 99% complete.

If we check a few minutes later, we'll see that the operation is complete.

```
$ q service-management get-operation operations/services.example.endpoints.bobadojo.cloud.goog-0 | jq
{
  "name": "operations/services.example.endpoints.bobadojo.cloud.goog-0",
  "metadata": {
    "@type": "type.googleapis.com/google.api.servicemanagement.v1.OperationMetadata",
    "resourceNames": [
      "services/example.endpoints.bobadojo.cloud.goog"
    ],
    "progressPercentage": 100,
    "startTime": "2024-11-15T22:37:25.944856Z"
  },
  "done": true,
  "response": {
    "@type": "type.googleapis.com/google.api.servicemanagement.v1.ManagedService",
    "serviceName": "example.endpoints.bobadojo.cloud.goog",
    "producerProjectId": "bobadojo"
  }
}
```

### DeleteOperation

`DeleteOperation` is defined in the Operations interface but is unsupported by the ServiceManagement service.

```
$ q service-management delete-operation operations/services.example.endpoints.bobadojo.cloud.goog-0 
Error: rpc error: code = Unimplemented desc = Method not found.
Usage:
  q service-management delete-operation OPERATION [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for delete-operation

```

### CancelOperation

`CancelOperation` is defined in the Operations interface but is unsupported by the ServiceManagement service.

```
$ q service-management cancel-operation operations/services.example.endpoints.bobadojo.cloud.goog-0 
Error: rpc error: code = Unimplemented desc = Method not found.
Usage:
  q service-management cancel-operation OPERATION [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for cancel-operation
```

### WaitOperation

`WaitOperation` is defined in the Operations interface but is unsupported by the ServiceManagement service.

```
$ q service-management wait-operation operations/services.example.endpoints.bobadojo.cloud.goog-0 
Error: rpc error: code = Unimplemented desc = WaitOperation is not supported now.
Usage:
  q service-management wait-operation OPERATION [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for wait-operation
```

## Summarizing

There's a lot in this section, but that's because we went over everything in great detail. To use Service Management, here's what you need to know:
1. Add your API by creating a managed service.
2. Describe your API by uploading a service config.
3. Activate your configuration by creating a rollout.
4. If you need to make changes, upload a new service config and roll it out.
5. To make your managed service visible to other Google Cloud users, set an IAM Policy on the service.

Now on to Service Control!