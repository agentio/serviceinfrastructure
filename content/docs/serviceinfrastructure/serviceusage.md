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