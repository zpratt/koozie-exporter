# koozie-exporter
[![Build](https://github.com/zpratt/topokube/actions/workflows/docker-image.yml/badge.svg?branch=main)](https://github.com/zpratt/topokube/actions/workflows/docker-image.yml)

<img width="557" alt="logo" src="https://user-images.githubusercontent.com/5916561/147760399-363e7478-23af-4853-a1f1-9599ced7ec21.png">

Tracking the flow of idea to production for applications deployed to kubernetes.

## What is this?

This is my idea for a potentially useful way to track the 4 Key Metrics (lead time, deployment frequency, mean time to recovery, change failure rate) automatically for applications deployed to k8s. 

## Project State

I'm still hacking on what this should look like. This is currently POC quality, though it can recognize when things are deployed to the cluster. 

## Prerequisite Steps:

install docker (*yes this is currently only working with docker*)
```bash
brew update && brew install helm kind helmfile golangci-lint
helm plugin install https://github.com/databus23/helm-diff
```

## Running In Kind:

```bash
sudo sh -c "echo '127.0.0.1 topokube.local' >> /etc/hosts"
make inkind
open https://topokube.local:30443/ui/index.html
```

## Demo Time - Cause a deployment and watch the output

```bash
make cause-deploy
```

What does this do? It runs a trivial container, which simulates the deployment of an application. It then makes a request to /metrics to show the current deployment count.

Kubernetes creates a pod for you, which triggers an event that koozie-exporter watches for. Koozie uses the prometheus client to capture metrics. A metrics endpoint (/metrics) is exposed, so that you can scrape these metrics using your own prometheus instance.
