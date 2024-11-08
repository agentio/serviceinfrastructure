---
weight: 7
title: The Service Usage API
---
# The Service Usage API

The Service Usage API is used to control access to Google APIs within Google Cloud projects.

## The Service Usage API methods

The methods of the Service Usage API are defined in the [googleapis](/googleapis) repo in [google/api/serviceusage/v1/serviceusage.proto](https://github.com/googleapis/googleapis/blob/9b94dba2f7f4b601f8232bc3a3f6ef32665279b9/google/api/serviceusage/v1/serviceusage.proto#L37).

| Method | Description |
| ------ | ----------- |
| EnableService | Enable a service so that it can be used with a project |
| DisableService | Disable a service so that it can no longer be used with a project |
| GetService | Returns the service configuration and enabled state for a given service |
| ListServices | List all services available to the specified project and the current state of those services with respect to the project |
| BatchEnableServices | Enable multiple services on a project |
| BatchGetServices | Returns the service configurations and enabled states for a given list of services |

Exercise:
- Look at the default set of APIs that are enabled for a new project
- Disable all of them
- What happens with the project?



create a new project

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