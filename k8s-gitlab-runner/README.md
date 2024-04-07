# K8s GitLab Runners

GitLab runners to be run in a Kubernetes cluster.

## Setup

- External secrets - used to pull GitLab runner registration token from AWS Param Store and populate it as a secret
- Reloader - watch changes in Config Map and secret - then perform rolling upgrade of pods

- S3 bucket caching setup

