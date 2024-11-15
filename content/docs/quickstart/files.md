---
weight: 3
title: Quickstart Files
bookHidden: true
---
## Demo Files

These are the files that are automatically generated in the [quickstart demo](/docs/quickstart/demo) by `q demo`.

These files were created for the `bobadojo` project in the `us-west1` region. Your files will vary slightly to match the `gcloud` configuration in the shell where you run `q demo`.

---

### api_config.yaml

This file contains a fragment of Service Configuration that specifies important aspects of the demo API: the service name, the specific API methods that it includes, and usage restrictions on specific methods.

```
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
```

---

### CHECK.sh

This is a utility script that can be optionally run after `SETUP.sh` to verify that access to the Cloud Run instance is correctly configured.

```
#/bin/sh

gcloud projects get-iam-policy bobadojo \
--flatten='bindings[].members' \
--format='table(bindings.role)' \
--filter='bindings.members:stores@bobadojo.iam.gserviceaccount.com'
```

---

### descriptor.pb

This is the binary-encoded FileDescriptorSet that describes the demo API.

```
binary data
```

---

### iam.yaml

This is the IAM configuration that allows calls to be made directly to the Cloud Run instance.

```
- members:
  - allUsers
  role: roles/run.invoker
```

---

### service.yaml

This is the Cloud Run service description. It describes a deployment containing two containers: one with the stores server and one containing the endpoints proxy (ESPv2) with appropriate configuration options.

```
version: 1
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: stores
spec:
  template:
    spec:
      serviceAccountName: stores@bobadojo.iam.gserviceaccount.com
      containers:
      - image: gcr.io/endpoints-release/endpoints-runtime:2
        name: espv2
        args:
        - --listener_port=8081
        - --backend=grpc://localhost:8080
        - --service=stores.endpoints.bobadojo.cloud.goog
        - --rollout_strategy=managed
        ports:
        - name: http1
          containerPort: 8081
      - image: us-west1-docker.pkg.dev/bobadojo/stores/stores:latest
        name: stores
```

---

### SETUP.sh

This does everything needed to get the service running on Cloud Run with Cloud Endpoints management.

```
#!/bin/sh

echo 'Enabling the Service Control API.'
gcloud services enable servicecontrol.googleapis.com

echo 'Enabling the Service Management API.'
gcloud services enable servicemanagement.googleapis.com

echo 'Enabling the Cloud Run Admin API.'
gcloud services enable run.googleapis.com

echo 'Creating the service from the service descriptor and API config file.'
gcloud endpoints services deploy descriptor.pb api_config.yaml

echo 'Creating a service account to run the server container.'
gcloud iam service-accounts create stores

echo 'Giving the service account roles to call the Service Control APIs.'
gcloud projects add-iam-policy-binding bobadojo \
	--member serviceAccount:stores@bobadojo.iam.gserviceaccount.com \
	--role roles/servicemanagement.serviceController
gcloud projects add-iam-policy-binding bobadojo \
	--member serviceAccount:stores@bobadojo.iam.gserviceaccount.com \
  --role roles/cloudtrace.agent

echo 'Creating the Cloud Run container.'
gcloud run services replace service.yaml

echo 'Configuring IAM to allow outside access to the container.'
gcloud run services set-iam-policy --quiet stores iam.yaml
```
