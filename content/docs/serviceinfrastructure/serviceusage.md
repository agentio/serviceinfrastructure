---
weight: 7
title: The Service Usage API
---
# The Service Usage API

The Service Usage API is used to control access to Google APIs within Google Cloud projects.

The Service Usage API is defined in the [googleapis](/docs/details/googleapis) repo in [serviceusage.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage_v1.yaml).
The methods specific to the API are defined by the [ServiceUsage](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage.proto#L37) service in [serviceusage.proto](https://github.com/googleapis/googleapis/blob/master/google/api/serviceusage/v1/serviceusage.proto).

## The ServiceUsage service

Method names below are prefixed with `google.api.serviceusage.v1.ServiceUsage.`

| Method | Description |
| ------ | ----------- |
| EnableService | Enable a service so that it can be used with a project |
| DisableService | Disable a service so that it can no longer be used with a project |
| GetService | Returns the service configuration and enabled state for a given service |
| ListServices | List all services available to the specified project and the current state of those services with respect to the project |
| BatchEnableServices | Enable multiple services on a project |
| BatchGetServices | Returns the service configurations and enabled states for a given list of services |

### EnableService

### DisableService

### GetService

### ListServices

### BatchEnableServices

### BatchGetServices

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
### log-http

There's a nice trick that we can do with `gcloud`... add the `--log-http` flag to see what API calls it makes.

```
$ gcloud endpoints services add-iam-policy-binding stores.endpoints.bobadojo.cloud.goog --member=user:tim@mitra.so --role=roles/servicemanagement.serviceConsumer --log-http
=======================
==== request start ====
uri: https://servicemanagement.googleapis.com/v1/services/stores.endpoints.bobadojo.cloud.goog:getIamPolicy?alt=json
method: POST
== headers start ==
b'accept': b'application/json'
b'accept-encoding': b'gzip, deflate'
b'authorization': --- Token Redacted ---
b'content-length': b'2'
b'content-type': b'application/json'
b'user-agent': b'google-cloud-sdk gcloud/498.0.0 command/gcloud.endpoints.services.add-iam-policy-binding invocation-id/906e0c85bd864d859fba8653af56baad environment/None environment-version/None client-os/LINUX client-os-ver/6.11.0 client-pltf-arch/x86_64 interactive/True from-script/False python/3.11.9 term/xterm-256color (Linux 6.11.0-8-generic)'
b'x-goog-api-client': b'cred-type/u'
== headers end ==
== body start ==
{}
== body end ==
==== request end ====
---- response start ----
status: 200
-- headers start --
Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
Cache-Control: private
Content-Encoding: gzip
Content-Type: application/json; charset=UTF-8
Date: Thu, 14 Nov 2024 06:33:26 GMT
Server: ESF
Transfer-Encoding: chunked
Vary: Origin, X-Origin, Referer
X-Content-Type-Options: nosniff
X-Debug-Tracking-ID: 10732980157617474798;o=0
X-Frame-Options: SAMEORIGIN
X-XSS-Protection: 0
-- headers end --
-- body start --
{
  "version": 1,
  "etag": "BwYm2aAT4GA="
}

-- body end --
total round trip time (request+response): 0.467 secs
---- response end ----
----------------------
=======================
==== request start ====
uri: https://servicemanagement.googleapis.com/v1/services/stores.endpoints.bobadojo.cloud.goog:setIamPolicy?alt=json
method: POST
== headers start ==
b'accept': b'application/json'
b'accept-encoding': b'gzip, deflate'
b'authorization': --- Token Redacted ---
b'content-length': b'151'
b'content-type': b'application/json'
b'user-agent': b'google-cloud-sdk gcloud/498.0.0 command/gcloud.endpoints.services.add-iam-policy-binding invocation-id/906e0c85bd864d859fba8653af56baad environment/None environment-version/None client-os/LINUX client-os-ver/6.11.0 client-pltf-arch/x86_64 interactive/True from-script/False python/3.11.9 term/xterm-256color (Linux 6.11.0-8-generic)'
b'x-goog-api-client': b'cred-type/u'
== headers end ==
== body start ==
{"policy": {"bindings": [{"members": ["user:tim@mitra.so"], "role": "roles/servicemanagement.serviceConsumer"}], "etag": "BwYm2aAT4GA=", "version": 1}}
== body end ==
==== request end ====
---- response start ----
status: 200
-- headers start --
Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
Cache-Control: private
Content-Encoding: gzip
Content-Type: application/json; charset=UTF-8
Date: Thu, 14 Nov 2024 06:33:27 GMT
Server: ESF
Transfer-Encoding: chunked
Vary: Origin, X-Origin, Referer
X-Content-Type-Options: nosniff
X-Debug-Tracking-ID: 8173729591266013879;o=0
X-Frame-Options: SAMEORIGIN
X-XSS-Protection: 0
-- headers end --
-- body start --
{
  "version": 1,
  "etag": "BwYm2aDCKOY=",
  "bindings": [
    {
      "role": "roles/servicemanagement.serviceConsumer",
      "members": [
        "user:tim@mitra.so"
      ]
    }
  ]
}

-- body end --
total round trip time (request+response): 0.633 secs
---- response end ----
----------------------
Updated IAM policy for service [stores.endpoints.bobadojo.cloud.goog].
bindings:
- members:
  - user:tim@mitra.so
  role: roles/servicemanagement.serviceConsumer
etag: BwYm2aDCKOY=
version: 1
```