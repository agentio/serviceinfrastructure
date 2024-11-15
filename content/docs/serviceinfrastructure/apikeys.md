---
weight: 4
title: The API Keys API
---
## The API Keys API

The API Keys API generates and manages API keys that consumers use to call APIs.

The API Keys API is defined in the [googleapis](/docs/details/googleapis) repo in [apikeys_v2.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys_v2.yaml).
The methods specific to the API are defined
by the [ApiKeys](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys.proto#L37) service
in [apikeys.proto](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys.proto).

## The ApiKeys service

Method names below are prefixed with `google.api.apikeys.v2.ApiKeys.`

| Method | Description |
| ------ | ----------- |
| [CreateKey](#createkey) | Creates a new API key |
| [ListKeys](#listkeys) | Lists the API keys owned by a project |
| [GetKey](#getkey) | Gets the metadata for an API key |
| [GetKeyString](#getkeystring) | Get the key string for an API key |
| [UpdateKey](#updatekey) | Patches the modifiable fields of an API key |
| [DeleteKey](#deletekey) | Deletes an API key |
| [UndeleteKey](#undeletekey) | Undeletes an API key which was deleted within 30 days |
| [LookupKey](#lookupkey) | Find the parent project and resource name of the API key that matches the key string in the request |

### CreateKey

### ListKeys

### GetKey

### GetKeyString

### UpdateKey

### DeleteKey

### UndeleteKey

### LookupKey


## Usage Notes

### What happens if someone uses an invalid key?

We can test this by creating an API key with another GCP account and using that key to call our API.

```
$ curl 'https://stores-1046800315646.us-west1.run.app/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI'  -i
HTTP/2 403 
content-type: application/json
x-envoy-decorator-operation: ingress GetStore
x-cloud-trace-context: a8d45728e34f1d205dff18bf9a922ff4
date: Fri, 08 Nov 2024 15:49:12 GMT
server: Google Frontend
content-length: 117
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000

{"message":"PERMISSION_DENIED: API stores.endpoints.bobadojo.cloud.goog is not enabled for the project.","code":403}
```

When they search for the service, they won't find it. It is only visible in the organization that created it.
![alt text](/screenshots/service-not-found.png)


Delete the key and you'll get this response:
```
$ !curl
curl 'https://stores-1046800315646.us-west1.run.app/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI'  -i
HTTP/2 400 
content-type: application/json
x-envoy-decorator-operation: ingress GetStore
x-cloud-trace-context: 0efb7df218e906bf48255cfd072af922;o=1
date: Fri, 08 Nov 2024 15:58:13 GMT
server: Google Frontend
content-length: 91
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000

{"code":400,"message":"INVALID_ARGUMENT: API key not found. Please pass a valid API key."}
```