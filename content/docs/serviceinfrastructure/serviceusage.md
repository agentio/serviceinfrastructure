---
weight: 7
title: The Service Usage API
---
## The Service Usage API

The Service Usage API is used to control access to Google APIs within Google Cloud projects. Google Cloud users regularly use it to enable services for use, and astute ones also use it to disable services that they don't need. Service Infrastructure users need to know that if they want to share their services with other Google Cloud users, those users will use Service Usage to find and enable them.

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

We'll be exploring Service Usage from the perspective of a user of a managed service. If you're sharing an API with another Google Cloud user, this is your consumer. So for the demos that follow, we've created a new Google Cloud account with a new user (our consumer), we've used [SetIamPolicy](/docs/serviceinfrastructure/servicemanagement/#setiampolicy) to share our managed service with our consumer's email address, and then we logged into `gcloud` and the Cloud Console with our consumer's account, and created a new project.

Here's what it looks like to create a new project in the cloud console:

![alt text](/screenshots/serviceusage-createproject.png)

For the examples that follow, we set the `CLIENT` environment variable to the project ID of the new project and `PROJECT` for the id of the project that owns our service. For us, that's done below. To use the examples that follow, do the same with your own values.

```prompt
CLIENT=nodal-time-442104-f1
```
```prompt
PROJECT=bobadojo
```

### **Services**

The Service Usage API defines it's own [Service](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/resources.proto#L38) type that allows it to add some additional information about services:
- The service name is relative to the owning project (e.g. `projects/$CLIENT/services/serviceusage.googleapis.com`). This allows us to selectively enable and disable services for use within individual projects.
- The `state` field contains the state of the service (`ENABLED` or `DISABLED`)

`Service` also includes a pared-down `ServiceConfig` type that contains a subset of the full service config. This subset is mainly focused on the needs of service consumers.

### ListServices

The `ListServices` method lets us list services available to our current project. This includes all the public services, all the ones that we own, and all the services for which we've been given the `servicemanagement.services.bind` permission, typically by being granted the `servicemanagement.serviceConsumer` role.

We can also filter services based on their state (enabled or disabled), and we've implemented our `q` subcommand to default to listing only enabled services. Let's try it:

```prompt
q service-usage list-services projects/$CLIENT | jq | more
```
```
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
```

That's a lot of detail! We're getting a list of service configurations that have been filtered to only contain information appropriate for service consumers.
To get a better overview, let's filter the result with `jq` to look at just the service names:

```prompt
q service-usage list-services projects/$CLIENT | jq .[].name -r
```
```
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

Here we see the list of services that Google enables by default in new projects. We also see that the full name of the service includes the id of the containing project. So an enabled service is enabled in the context of that project.

Our list corresponds to this view in the Google Cloud Console:

![alt text](/screenshots/serviceusage-enabledapis.png)

To see the services that we might enable, let's rerun with a different filter.

```prompt
q service-usage list-services projects/$CLIENT \
    --filter state:DISABLED | jq .[].name -r
```
```
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
```

Wow this is a huge list! It's usually not displayed in the Cloud Console, which instead presents us with a search interface:

![alt text](/screenshots/serviceusage-enableapis.png)


Let's put our list in a file so we can easily count the number of services.

```prompt
q service-usage list-services projects/$CLIENT \
    --filter state:DISABLED | jq .[].name -r > DISABLED
```
```prompt
wc -l DISABLED
```
```
6216 DISABLED
```

Scanning the output, we can see that it includes our shared service:
```prompt
grep $PROJECT DISABLED
```
```
projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog
```

### GetService

`GetService` allows us to get the service configuration and state for a service. Let's use it to look at our shared service.

```prompt
q service-usage get-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog | jq
```
```
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

Here we see that the state is "DISABLED". Let's use the API to make it available for use in our project.

### EnableService

The `EnableService` method can be used to enable a service. Let's use it for our Stores service:
```prompt
q service-usage enable-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog 
```
```
{}
```

Now we can get our service again and see that it is enabled.
```prompt
q service-usage get-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog | jq 
```
```
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
```

```prompt
q service-usage get-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog \
    | jq .state -r
```
```
ENABLED
```

It also shows up now in our list of enabled services.
```prompt
q service-usage list-services projects/$CLIENT | jq .[].name -r
```
```
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

If we change our mind, we can use `DisableService` to disable a service within our project. Quoting the doc strings in the proto: "This prevents unintended usage that may cause unexpected billing charges or security leaks."

```prompt
q service-usage disable-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog 
```
```
{}
```

```prompt
q service-usage get-service \
    projects/$CLIENT/services/stores.endpoints.$PROJECT.cloud.goog \
    | jq .state -r
```
```
DISABLED
```

### BatchGetServices

`BatchGetServices` lets us check the enable state for a list of services. It is just like `GetService` but takes a list of service names and returns a list of service configs.

### BatchEnableServices

`BatchEnableServices` lets us enable a list of services. It is just like `EnableService` but takes a list of service names.

## The Operations service

This is the same service that we discussed for the Service Management API, so we won't discuss it in detail here.

Reviewing the service protos, we see that `Operations` are returned by the methods that enable and disable services, and although these typically complete quickly, in production we might want to check the returned operations to be sure they have completed.

## Usage Notes

### Disabling all of the APIs for a project

Since disabling unneeded APIs seems to be a recommended practice, let's use the Service Usage API to disable everything in our consumer project other than the service that we've shared.

We'll start by putting all of our enabled services in a file. We call this file `DISABLE.sh` because we're going to edit it to be a script that disables these services.

```prompt
q service-usage list-services projects/$CLIENT \
    | jq .[].name -r > DISABLE.sh
```

Next we edit the file, putting `q service-usage disable-service ` at the start of each line:

```
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

Note that we've deleted the line containing the stores service that we shared with this user. We've also commented out the call that disables the Service Usage API so that we can explore that case separately. 

```prompt
sh DISABLE.sh 
```
```
{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}
```

Now when we list our services, we see that we have just two services:

```prompt
q service-usage list-services projects/$CLIENT | jq .[].name -r
```
```
projects/327402113844/services/serviceusage.googleapis.com
projects/327402113844/services/stores.endpoints.bobadojo.cloud.goog
```

Now let's disable the Service Usage API.
```prompt
q service-usage disable-service \
    projects/$CLIENT/services/serviceusage.googleapis.com
```
```
{}
```

```prompt
q service-usage list-services projects/$CLIENT 
```
```
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

Oops! Now that we've disabled Service Usage, we can no longer list our services. That's not a big surprise.

```prompt
q service-usage enable-service \
    projects/$CLIENT/services/serviceusage.googleapis.com
```
```
Error: rpc error: code = PermissionDenied desc = Service Usage API has not been used in project nodal-time-442104-f1 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/nodal-time-442104-f1 service:serviceusage.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1
Usage:
  q service-usage enable-service [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for enable-service
```

We also can't reenable it. It seems that we've locked ourselves out of this project. But let's try again with `gcloud`:

```prompt
gcloud services enable serviceusage.googleapis.com
```
```
Operation "operations/acat.p2-327402113844-d45c9e7a-440a-4bd0-8df5-b13840f681e9" finished successfully.
```
```prompt
gcloud config list
```
```
[core]
account = tim@mitra.so
disable_usage_reporting = False
log_http_redact_token = false
project = nodal-time-442104-f1
[run]
region = us-west1

Your active configuration is: [default]
```
```prompt
gcloud services list
```
```
NAME                                  TITLE
serviceusage.googleapis.com           Service Usage API
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API
```

Wow, somehow `gcloud` was able to do this when `q` couldn't. We suspect that someone on the `gcloud` team thought of this possibility and is using a backup identity to perform this action. 

We can use gcloud to manage service usage for a project, even when we've disabled the Service Usage API in that project.

```prompt
gcloud services list
```
```
NAME                                  TITLE
serviceusage.googleapis.com           Service Usage API
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API
```
```prompt
gcloud services disable serviceusage.googleapis.com
```
```
Operation "operations/acat.p17-327402113844-6e586555-42d3-4432-8e29-b4975d17e5f4" finished successfully.
```
```prompt
gcloud services list
```
```
NAME                                  TITLE
stores.endpoints.bobadojo.cloud.goog  Boba Dojo Stores API
```
```prompt
gcloud services enable serviceusage.googleapis.com
```
```
Operation "operations/acat.p2-327402113844-cf2ba8b5-4865-497e-a5cd-ea1b2aaea6d9" finished successfully.
```
```prompt
gcloud services disable serviceusage.googleapis.com
```
```
Operation "operations/acat.p17-327402113844-e6e66b99-79d0-4e27-a711-762f9e498d62" finished successfully.
```

And this works even though `q`, which uses Application Default Credentials associated with the managed project, is unable to make Service Usage calls.

```prompt
q service-usage enable-service \
    projects/$CLIENT/services/serviceusage.googleapis.com
```
```
Error: rpc error: code = PermissionDenied desc = Service Usage API has not been used in project nodal-time-442104-f1 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/nodal-time-442104-f1 service:serviceusage.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=nodal-time-442104-f1
Usage:
  q service-usage enable-service [flags]

Flags:
      --format string   output format (default "json")
  -h, --help            help for enable-service
```

Coming out of this digression, we check the Cloud Console and see that it shows only one service enabled for this project:

![alt text](/screenshots/serviceusage-singleservice.png)

### Granting access to an API to other users

In the Service Management section, we showed how we can use the `IamPolicyService` to share our service with another user. Here's how it looks in the Cloud Console. Note that now we have gone back and logged in with our producer identity. 

Begin by giving the user the IAM role of "Service Consumer". For this screen, we are logged into the Cloud Console with the service owner's identity.

![alt text](/screenshots/serviceusage-grant.png)

The user can then look up your API in the Cloud Console. Here we are logged in as the service consumer:

![alt text](/screenshots/serviceusage-granted.png)

The user can then enable your API and create API keys to use it.

![alt text](/screenshots/serviceusage-granted-detail.png)

What just happened? Just as we did in our Service Management discussion, we set the IAM policy on the service. Another way to do this is with `gcloud endpoints services` subcommands. First we'll view the policy:

```prompt
gcloud endpoints services get-iam-policy stores.endpoints.$PROJECT.cloud.goog
```
```
bindings:
- members:
  - user:tim@mitra.so
  role: roles/servicemanagement.serviceConsumer
etag: BwYmbJiRfHM=
version: 1

```

Next we update the policy binding:
```prompt
gcloud endpoints services add-iam-policy-binding \
    stores.endpoints.$PROJECT.cloud.goog \
    --member=user:tim@mitra.so \
    --role=roles/servicemanagement.serviceConsumer
```

We can also remove the binding:
```prompt
gcloud endpoints services remove-iam-policy-binding \
    stores.endpoints.$PROJECT.cloud.goog \
    --member=user:tim@mitra.so \
    --role=roles/servicemanagement.serviceConsumer
```
```
Updated IAM policy for service [stores.endpoints.bobadojo.cloud.goog].
etag: BwYmbUpFFGg=
version: 1
```

And if we change our minds, we can easily add it back:
```prompt
gcloud endpoints services add-iam-policy-binding \
    stores.endpoints.$PROJECT.cloud.goog \
    --member=user:tim@mitra.so \
    --role=roles/servicemanagement.serviceConsumer
```
```
Updated IAM policy for service [stores.endpoints.bobadojo.cloud.goog].
bindings:
- members:
  - user:tim@mitra.so
  role: roles/servicemanagement.serviceConsumer
etag: BwYmbUvc2zk=
version: 1
```

## Summarizing

Service Usage is the simplest API that we discussed, but in an important way, it's the bridge that connects services to Google Cloud users. With the Service Usage API, Google Cloud users can find services to use in their projects, enable them, and also importantly, disable services that they don't need. This can held reduce security risks and accidental charges. If you want to share a service that you make with Google Cloud users, they'll need to know to use the Service Usage API to enable your service so that they can then generate API keys to use it.

---
#### Continue with [the Extensible Service Proxies](/docs/proxies).