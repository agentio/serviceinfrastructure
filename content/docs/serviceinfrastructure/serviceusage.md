---
weight: 7
title: The Service Usage API
---
## The Service Usage API

The Service Usage API is used to control access to Google APIs within Google Cloud projects.

The Service Usage API is defined in the [googleapis](/docs/details/googleapis) repo in [serviceusage.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage_v1.yaml). It includes two services.

| Service | Purpose |
| ------- | ------- |
| [ServiceUsage](#the-serviceusage-service) | Manage the services that are enabled for use in a Google Cloud project |
| [Operations](#the-operations-service) | A mix-in that handles long-running operations |

## The ServiceUsage service

The [ServiceUsage](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage.proto#L37) service is defined in [serviceusage.proto](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage.proto).

The full names of these methods begin with  `google.api.serviceusage.v1.ServiceUsage.`

| Method | Description |
| ------ | ----------- |
| [ListServices](#listservices) | List all services available to the specified project and the current state of those services with respect to the project |
| [GetService](#getservice) | Returns the service configuration and enabled state for a given service |
| [EnableService](#enableservice) | Enable a service so that it can be used with a project |
| [DisableService](#disableservice) | Disable a service so that it can no longer be used with a project |
| [BatchGetServices](#batchgetservices) | Returns the service configurations and enabled states for a given list of services |
| [BatchEnableServices](#batchenableservices) | Enable multiple services on a project |

### ListServices

From the proto: "List all services available to the specified project, and the current state of those services with respect to the project. The list includes all public services, all services for which the calling user has the `servicemanagement.services.bind` permission, and all services that have already been enabled on the project. The list can be filtered to only include services in a specific state, for example to only include services enabled on the project."

```
$ q service-usage list-services projects/nodal-time-442104-f1 | jq | more
[
  {
    "name": "projects/327402113844/services/analyticshub.googleapis.com",
    "parent": "projects/327402113844",
    "config": {
      "name": "analyticshub.googleapis.com",
      "title": "Analytics Hub API",
      "documentation": {
        "summary": "Exchange data and analytics assets securely and efficiently."
      },
      "quota": {},
      "authentication": {},
      "usage": {
        "requirements": [
          "serviceusage.googleapis.com/tos/cloud"
        ]
      },
      "monitoring": {}
    },
    "state": "ENABLED"
  },
  ... 20 services elided
]

$ q service-usage list-services projects/nodal-time-442104-f1 | jq .[].name -r
projects/327402113844/services/analyticshub.googleapis.com
projects/327402113844/services/bigquery.googleapis.com
projects/327402113844/services/bigqueryconnection.googleapis.com
projects/327402113844/services/bigquerydatapolicy.googleapis.com
projects/327402113844/services/bigquerymigration.googleapis.com
projects/327402113844/services/bigqueryreservation.googleapis.com
projects/327402113844/services/bigquerystorage.googleapis.com
projects/327402113844/services/cloudapis.googleapis.com
projects/327402113844/services/cloudresourcemanager.googleapis.com
projects/327402113844/services/cloudtrace.googleapis.com
projects/327402113844/services/dataform.googleapis.com
projects/327402113844/services/dataplex.googleapis.com
projects/327402113844/services/datastore.googleapis.com
projects/327402113844/services/logging.googleapis.com
projects/327402113844/services/monitoring.googleapis.com
projects/327402113844/services/servicemanagement.googleapis.com
projects/327402113844/services/serviceusage.googleapis.com
projects/327402113844/services/sql-component.googleapis.com
projects/327402113844/services/storage-api.googleapis.com
projects/327402113844/services/storage-component.googleapis.com
projects/327402113844/services/storage.googleapis.com

```

```
$ q service-usage list-services projects/nodal-time-442104-f1 --filter state:DISABLED | jq .[].name -r
projects/327402113844/services/a10-thunder-adc-601b150-byol.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-100mbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-10gbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-1gbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-200mbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-20mbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-500mbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-5gbps.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/a10-vthunder-adc-byol.endpoints.a10networks-public-396315.cloud.goog
projects/327402113844/services/aapl-miriinfotech-public.cloudpartnerservices.goog
projects/327402113844/services/ab-initio-cooperating-system.endpoints.ab-initio-419002.cloud.goog
projects/327402113844/services/ab-tasty-experimentation.endpoints.abtasty-public.cloud.goog
projects/327402113844/services/abacus.ai.endpoints.abacus-public.cloud.goog
projects/327402113844/services/abacus360-on-rcloud.endpoints.regnology-cloud-marketplace.cloud.goog
...

$ q service-usage list-services projects/nodal-time-442104-f1 --filter state:DISABLED | jq .[].name -r > DISABLED

$ wc -l DISABLED
6216 DISABLED

$ grep bobadojo DISABLED
projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog
```

### GetService

"Returns the service configuration and enabled state for a given service."

```
$ q service-usage get-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog | jq
{
  "name": "projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog",
  "parent": "projects/327402113844",
  "config": {
    "name": "stores.endpoints.bobadojo.cloud.goog",
    "title": "Boba Dojo Stores API",
    "apis": [
      {
        "name": "bobadojo.stores.v1.Stores",
        "methods": [
          {
            "name": "ListStores"
          },
          {
            "name": "FindStores"
          },
          {
            "name": "GetStore"
          }
        ],
        "version": "v1"
      }
    ],
    "documentation": {},
    "quota": {},
    "authentication": {},
    "usage": {},
    "endpoints": [
      {
        "name": "stores.endpoints.bobadojo.cloud.goog"
      }
    ],
    "monitoredResources": [
      {
        "type": "api",
        "labels": [
          {
            "key": "cloud.googleapis.com/location"
          },
          {
            "key": "cloud.googleapis.com/uid"
          },
          {
            "key": "serviceruntime.googleapis.com/api_version"
          },
          {
            "key": "serviceruntime.googleapis.com/api_method"
          },
          {
            "key": "serviceruntime.googleapis.com/consumer_project"
          },
          {
            "key": "cloud.googleapis.com/project"
          },
          {
            "key": "cloud.googleapis.com/service"
          }
        ]
      }
    ],
    "monitoring": {
      "consumerDestinations": [
        {
          "monitoredResource": "api",
          "metrics": [
            "serviceruntime.googleapis.com/api/consumer/request_count",
            "serviceruntime.googleapis.com/api/consumer/quota_used_count",
            "serviceruntime.googleapis.com/api/consumer/total_latencies",
            "serviceruntime.googleapis.com/api/consumer/request_sizes",
            "serviceruntime.googleapis.com/api/consumer/response_sizes"
          ]
        }
      ]
    }
  },
  "state": "DISABLED"
}
```

### EnableService

"Enable a service so that it can be used with a project."

```
$ q service-usage enable-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog 
{}

$ q service-usage get-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog | jq 
{
  "name": "projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog",
  "parent": "projects/327402113844",
  "config": {
    "name": "stores.endpoints.bobadojo.cloud.goog",
    "title": "Boba Dojo Stores API",
    "apis": [
      {
        "name": "bobadojo.stores.v1.Stores",
        "methods": [
          {
            "name": "ListStores"
          },
          {
            "name": "FindStores"
          },
          {
            "name": "GetStore"
          }
        ],
        "version": "v1"
      }
    ],
    "documentation": {},
    "quota": {},
    "authentication": {},
    "usage": {},
    "endpoints": [
      {
        "name": "stores.endpoints.bobadojo.cloud.goog"
      }
    ],
    "monitoredResources": [
      {
        "type": "api",
        "labels": [
          {
            "key": "cloud.googleapis.com/location"
          },
          {
            "key": "cloud.googleapis.com/uid"
          },
          {
            "key": "serviceruntime.googleapis.com/api_version"
          },
          {
            "key": "serviceruntime.googleapis.com/api_method"
          },
          {
            "key": "serviceruntime.googleapis.com/consumer_project"
          },
          {
            "key": "cloud.googleapis.com/project"
          },
          {
            "key": "cloud.googleapis.com/service"
          }
        ]
      }
    ],
    "monitoring": {
      "consumerDestinations": [
        {
          "monitoredResource": "api",
          "metrics": [
            "serviceruntime.googleapis.com/api/consumer/request_count",
            "serviceruntime.googleapis.com/api/consumer/quota_used_count",
            "serviceruntime.googleapis.com/api/consumer/total_latencies",
            "serviceruntime.googleapis.com/api/consumer/request_sizes",
            "serviceruntime.googleapis.com/api/consumer/response_sizes"
          ]
        }
      ]
    }
  },
  "state": "ENABLED"
}

$ q service-usage get-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog | jq .state -r
ENABLED

$ q service-usage list-services projects/nodal-time-442104-f1 | jq .[].name -r
projects/327402113844/services/analyticshub.googleapis.com
projects/327402113844/services/bigquery.googleapis.com
projects/327402113844/services/bigqueryconnection.googleapis.com
projects/327402113844/services/bigquerydatapolicy.googleapis.com
projects/327402113844/services/bigquerymigration.googleapis.com
projects/327402113844/services/bigqueryreservation.googleapis.com
projects/327402113844/services/bigquerystorage.googleapis.com
projects/327402113844/services/cloudapis.googleapis.com
projects/327402113844/services/cloudresourcemanager.googleapis.com
projects/327402113844/services/cloudtrace.googleapis.com
projects/327402113844/services/dataform.googleapis.com
projects/327402113844/services/dataplex.googleapis.com
projects/327402113844/services/datastore.googleapis.com
projects/327402113844/services/logging.googleapis.com
projects/327402113844/services/monitoring.googleapis.com
projects/327402113844/services/servicemanagement.googleapis.com
projects/327402113844/services/serviceusage.googleapis.com
projects/327402113844/services/sql-component.googleapis.com
projects/327402113844/services/storage-api.googleapis.com
projects/327402113844/services/storage-component.googleapis.com
projects/327402113844/services/storage.googleapis.com
projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog

```

### DisableService

"Disable a service so that it can no longer be used with a project. This prevents unintended usage that may cause unexpected billing charges or security leaks."

```
$ q service-usage disable-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog 
{}

$ q service-usage get-service projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog | jq .state -r
DISABLED
```

### BatchGetServices

"Returns the service configurations and enabled states for a given list of services."

### BatchEnableServices

"Enable multiple services on a project."

## The Operations service

This is the same service that we discussed for the Service Management API, so we won't discuss it in detail here.

## Usage Notes

### Disabling all of the APIs for a project

- Look at the default set of APIs that are enabled for a new project
- Disable all of them
- What happens with the project?

create a new project

![alt text](/screenshots/serviceusage-createproject.png)

view the project in the cloud console

![alt text](/screenshots/serviceusage-newproject.png)

![alt text](/screenshots/serviceusage-enableapis.png)
![alt text](/screenshots/serviceusage-enabledapis.png)

look at the project with gcloud

```
gcloud config set project dauntless-glow-441118-p5
```

```
$ gcloud services list --enabled
NAME                                TITLE
analyticshub.googleapis.com         Analytics Hub API
bigquery.googleapis.com             BigQuery API
bigqueryconnection.googleapis.com   BigQuery Connection API
bigquerydatapolicy.googleapis.com   BigQuery Data Policy API
bigquerymigration.googleapis.com    BigQuery Migration API
bigqueryreservation.googleapis.com  BigQuery Reservation API
bigquerystorage.googleapis.com      BigQuery Storage API
cloudapis.googleapis.com            Google Cloud APIs
cloudtrace.googleapis.com           Cloud Trace API
dataform.googleapis.com             Dataform API
dataplex.googleapis.com             Cloud Dataplex API
datastore.googleapis.com            Cloud Datastore API
logging.googleapis.com              Cloud Logging API
monitoring.googleapis.com           Cloud Monitoring API
servicemanagement.googleapis.com    Service Management API
serviceusage.googleapis.com         Service Usage API
sql-component.googleapis.com        Cloud SQL
storage-api.googleapis.com          Google Cloud Storage JSON API
storage-component.googleapis.com    Cloud Storage
storage.googleapis.com              Cloud Storage API
```

get the enabled services with `q`

```
q service-usage list-services projects/dauntless-glow-441118-p5 > services.json
```

```
$ jq .[].state < services.json  | wc -l
20
```

```
$ jq .[].name < services.json -r
projects/51662343665/services/analyticshub.googleapis.com
projects/51662343665/services/bigquery.googleapis.com
projects/51662343665/services/bigqueryconnection.googleapis.com
projects/51662343665/services/bigquerydatapolicy.googleapis.com
projects/51662343665/services/bigquerymigration.googleapis.com
projects/51662343665/services/bigqueryreservation.googleapis.com
projects/51662343665/services/bigquerystorage.googleapis.com
projects/51662343665/services/cloudapis.googleapis.com
projects/51662343665/services/cloudtrace.googleapis.com
projects/51662343665/services/dataform.googleapis.com
projects/51662343665/services/dataplex.googleapis.com
projects/51662343665/services/datastore.googleapis.com
projects/51662343665/services/logging.googleapis.com
projects/51662343665/services/monitoring.googleapis.com
projects/51662343665/services/servicemanagement.googleapis.com
projects/51662343665/services/serviceusage.googleapis.com
projects/51662343665/services/sql-component.googleapis.com
projects/51662343665/services/storage-api.googleapis.com
projects/51662343665/services/storage-component.googleapis.com
projects/51662343665/services/storage.googleapis.com
```

disable all of the services with a script

![alt text](/screenshots/serviceusage-disableservices.png)

```
$ sh DISABLE.sh 
{}Error: rpc error: code = FailedPrecondition desc = The service bigquery.googleapis.com is depended on by the following active service(s): bigquerystorage.googleapis.com,cloudapis.googleapis.com; Please specify disable_dependent_services=true if you want to proceed with disabling all services.
Help Token: AYJSUtmWySNobo0XEk_XraMjFzAFMofXuZufOd2VlXHlqhBNia8TlS4cNV0gRZ18_13N5W2oR4OGNf6WjtkLHAd5gjM-bpgPz-B4ZczdPmYoyNib
error details: name = ErrorInfo reason = COMMON_SU_SERVICE_HAS_DEPENDENT_SERVICES domain = serviceusage.googleapis.com metadata = map[service_name:bigquery.googleapis.com services:bigquerystorage.googleapis.com,cloudapis.googleapis.com]
error details: name = PreconditionFailure type = googleapis.com subj = ?error_code=100001&service_name=bigquery.googleapis.com&services=bigquerystorage.googleapis.com&services=cloudapis.googleapis.com desc =
Usage:
  q service-usage disable-service [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for disable-service

{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}
```

```
$ q service-usage disable-service projects/51662343665/services/bigquery.googleapis.com
```

```
$ q service-usage list-services projects/51662343665 | jq
[
  {
    "name": "projects/51662343665/services/serviceusage.googleapis.com",
    "parent": "projects/51662343665",
    "config": {
      "name": "serviceusage.googleapis.com",
      "title": "Service Usage API",
      "documentation": {
        "summary": "Enables services that service consumers want to use on Google Cloud Platform, lists the available or enabled services, or disables services that service consumers no longer use."
      },
      "quota": {},
      "authentication": {},
      "usage": {
        "requirements": [
          "serviceusage.googleapis.com/tos/cloud"
        ]
      },
      "monitoring": {}
    },
    "state": "ENABLED"
  }
]

```
disable the usage service

```
$ gcloud services list --project dauntless-glow-441118-p5
NAME                         TITLE
serviceusage.googleapis.com  Service Usage API
```

```
$ q service-usage disable-service projects/51662343665/services/serviceusage.googleapis.com
{}

$ gcloud services list --project dauntless-glow-441118-p5
Listed 0 items.

q service-usage list-services projects/51662343665 
[]
```

enable the bobadojo service

```
$ q service-usage enable-service projects/51662343665/services/stores.endpoints.bobadojo.cloud.goog 

$ gcloud services list --project dauntless-glow-441118-p5
NAME                                  TITLE
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API

```

![alt text](/screenshots/serviceusage-singleservice.png)

now the project has just a single service

```
$ gcloud auth application-default set-quota-project dauntless-glow-441118-p5
API [cloudresourcemanager.googleapis.com] not enabled on project [dauntless-glow-441118-p5]. Would you like to enable and retry (this will take a 
few minutes)? (y/N)?  y

Enabling service [cloudresourcemanager.googleapis.com] on project [dauntless-glow-441118-p5]...
ERROR: (gcloud.auth.application-default.set-quota-project) PERMISSION_DENIED: Service Usage API has not been used in project dauntless-glow-441118-p5 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=dauntless-glow-441118-p5 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry. This command is authenticated as None using the credentials in /home/tim/.config/gcloud/application_default_credentials.json, specified by the [auth/credential_file_override] property.
Google developers console API activation
https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=dauntless-glow-441118-p5
- '@type': type.googleapis.com/google.rpc.ErrorInfo
  domain: googleapis.com
  metadata:
    consumer: projects/dauntless-glow-441118-p5
    service: serviceusage.googleapis.com
  reason: SERVICE_DISABLED
```

### Granting access to an API to other users

Begin by giving the user the IAM role of "Service Consumer"

![alt text](/screenshots/serviceusage-grant.png)

The user can then look up your API in the Cloud Console

![alt text](/screenshots/serviceusage-granted.png)

The user can then enable your API and create API keys to use it.

![alt text](/screenshots/serviceusage-granted-detail.png)

What just happened? We set an iam policy on the service. View it with this `gcloud` command:
```
$ gcloud endpoints services get-iam-policy stores.endpoints.bobadojo.cloud.goog
bindings:
- members:
  - user:tim@mitra.so
  role: roles/servicemanagement.serviceConsumer
etag: BwYmbJiRfHM=
version: 1

```

We could do this from the `gcloud` with
```
gcloud endpoints services add-iam-policy-binding stores.endpoints.bobadojo.cloud.goog --member=user:tim@mitra.so --role=roles/servicemanagement.serviceConsumer
```

Let's remove the binding:

```
$ gcloud endpoints services remove-iam-policy-binding stores.endpoints.bobadojo.cloud.goog --member=user:tim@mitra.so --role=roles/servicemanagement.serviceConsumer
Updated IAM policy for service [stores.endpoints.bobadojo.cloud.goog].
etag: BwYmbUpFFGg=
version: 1
```

and now let's add it back:

```
$ gcloud endpoints services add-iam-policy-binding stores.endpoints.bobadojo.cloud.goog --member=user:tim@mitra.so --role=roles/servicemanagement.serviceConsumer
Updated IAM policy for service [stores.endpoints.bobadojo.cloud.goog].
bindings:
- members:
  - user:tim@mitra.so
  role: roles/servicemanagement.serviceConsumer
etag: BwYmbUvc2zk=
version: 1
```



---

```
$ q service-usage list-services projects/nodal-time-442104-f1  | jq .[].name -r > DISABLE.sh

#!/bin/sh

q service-usage disable-service projects/327402113844/services/analyticshub.googleapis.com
q service-usage disable-service projects/327402113844/services/bigquery.googleapis.com
q service-usage disable-service projects/327402113844/services/bigqueryconnection.googleapis.com
q service-usage disable-service projects/327402113844/services/bigquerydatapolicy.googleapis.com
q service-usage disable-service projects/327402113844/services/bigquerymigration.googleapis.com
q service-usage disable-service projects/327402113844/services/bigqueryreservation.googleapis.com
q service-usage disable-service projects/327402113844/services/bigquerystorage.googleapis.com
q service-usage disable-service projects/327402113844/services/cloudapis.googleapis.com
q service-usage disable-service projects/327402113844/services/cloudresourcemanager.googleapis.com
q service-usage disable-service projects/327402113844/services/cloudtrace.googleapis.com
q service-usage disable-service projects/327402113844/services/dataform.googleapis.com
q service-usage disable-service projects/327402113844/services/dataplex.googleapis.com
q service-usage disable-service projects/327402113844/services/datastore.googleapis.com
q service-usage disable-service projects/327402113844/services/logging.googleapis.com
q service-usage disable-service projects/327402113844/services/monitoring.googleapis.com
q service-usage disable-service projects/327402113844/services/servicemanagement.googleapis.com
q service-usage disable-service projects/327402113844/services/sql-component.googleapis.com
q service-usage disable-service projects/327402113844/services/storage-api.googleapis.com
q service-usage disable-service projects/327402113844/services/storage-component.googleapis.com
q service-usage disable-service projects/327402113844/services/storage.googleapis.com

#q service-usage disable-service projects/327402113844/services/serviceusage.googleapis.com

```

```
$ sh DISABLE.sh 
{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}
```

```
$ q service-usage list-services projects/nodal-time-442104-f1 | jq .[].name -r
projects/327402113844/services/serviceusage.googleapis.com
projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog
```

```
q service-usage disable-service projects/327402113844/services/serviceusage.googleapis.com
{}
```

```
$ q service-usage list-services projects/nodal-time-442104-f1 
[Error: rpc error: code = PermissionDenied desc = Service Usage API has not been used in project nodal-time-442104-f1 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/nodal-time-442104-f1 service:serviceusage.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1
Usage:
  q service-usage list-services PARENT [flags]

Flags:
      --filter string   filter (default "state:ENABLED")
      --format string   output format (default "json")
  -h, --help            help for list-services
```
```
$ q service-usage enable-service projects/327402113844/services/serviceusage.googleapis.com
Error: rpc error: code = PermissionDenied desc = Service Usage API has not been used in project nodal-time-442104-f1 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/nodal-time-442104-f1 service:serviceusage.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1
Usage:
  q service-usage enable-service [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for enable-service
```
```
$ gcloud services enable serviceusage.googleapis.com
Operation "operations/acat.p2-327402113844-d45c9e7a-440a-4bd0-8df5-b13840f681e9" finished successfully.

$ gcloud config list
[core]
account = tim@mitra.so
disable_usage_reporting = False
log_http_redact_token = false
project = nodal-time-442104-f1
[run]
region = us-west1

Your active configuration is: [default]

$ gcloud services list
NAME                                  TITLE
serviceusage.googleapis.com           Service Usage API
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API
```

```
$ gcloud services list
NAME                                  TITLE
serviceusage.googleapis.com           Service Usage API
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API

$ gcloud services disable serviceusage.googleapis.com
Operation "operations/acat.p17-327402113844-6e586555-42d3-4432-8e29-b4975d17e5f4" finished successfully.

$ gcloud services list
NAME                                  TITLE
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API

$ gcloud services enable serviceusage.googleapis.com
Operation "operations/acat.p2-327402113844-cf2ba8b5-4865-497e-a5cd-ea1b2aaea6d9" finished successfully.

$ gcloud services disable serviceusage.googleapis.com
Operation "operations/acat.p17-327402113844-e6e66b99-79d0-4e27-a711-762f9e498d62" finished successfully.

$ q service-usage enable-service projects/nodal-time-442104-f1/services/serviceusage.googleapis.com
Error: rpc error: code = PermissionDenied desc = Service Usage API has not been used in project nodal-time-442104-f1 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/nodal-time-442104-f1 service:serviceusage.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1
Usage:
  q service-usage enable-service [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for enable-service

```

## Summarizing

Service Usage is the simplest API that we discussed, but in an important way, it's the bridge that connects services to Google Cloud users. With the Service Usage API, Google Cloud users can find services to use in their projects, enable them, and also importantly, disable services that they don't need. This can held reduce security risks and accidental charges. If you want to share a service that you make with Google Cloud users, they'll need to know to use the Service Usage API to enable your service so that they can then generate API keys to use it.