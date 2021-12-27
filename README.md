# ~~topokube~~ koozie
[![Build](https://github.com/zpratt/topokube/actions/workflows/docker-image.yml/badge.svg?branch=main)](https://github.com/zpratt/topokube/actions/workflows/docker-image.yml)

Tracking deployment metrics for applications deployed to kubernetes.

## What is this?

This is my idea for a potentially useful way to track the 4 Key Metrics (lead time, deployment frequency, mean time to recovery, change failure rate) automatically for applications deployed to k8s. 

## Project State

I'm still hacking on what this should look like. This is currently POC quality, though it can recognize when things are deployed to the cluster. 

## Prerequisite Steps:

1. install docker (*yes this is currently only working with docker*)
2. `brew update && brew install helm kind helmfile golangci-lint`
3. `helm plugin install https://github.com/databus23/helm-diff`

## Running In Kind:

1. `sudo sh -c "echo '127.0.0.1 topokube.local' >> /etc/hosts"`
2. `make inkind`
3. `open https://topokube.local:30443/ui/index.html`

## Cause a deployment and watch the output

1. `make cause-deploy`
2. `kubectl logs -n topokube -l app.kubernetes.io/name=node-app -c node-app`
