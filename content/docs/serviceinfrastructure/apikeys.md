---
weight: 3
title: The API Keys API
---
# The API Keys API

The API Keys API generates and manages API keys that consumers use to call APIs.

## The API Keys API methods

The methods of the API Keys API are defined in the [googleapis](/googleapis) repo in [google/api/apikeys/v2/apikeys.proto](https://github.com/googleapis/googleapis/blob/9b94dba2f7f4b601f8232bc3a3f6ef32665279b9/google/api/apikeys/v2/apikeys.proto#L37).

| Method | Description |
| ------ | ----------- |
| CreateKey | Creates a new API key |
| ListKeys | Lists the API keys owned by a project |
| GetKey | Gets the metadata for an API key |
| GetKeyString | Get the key string for an API key |
| UpdateKey | Patches the modifiable fields of an API key |
| DeleteKey | Deletes an API key |
| UndeleteKey | Undeletes an API key which was deleted within 30 days |
| LookupKey | Find the parent project and resource name of the API key that matches the key string in the request |