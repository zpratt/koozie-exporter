# topokube
Topological view of your kubernetes cluster.

## What is this?

This is my idea for a potentially useful way to view the details of objects deployed to a target kuberentes cluster to help keep an inventory of what has been deployed, how long it's been there, and the dependencies between deployed objects (hence the usage of the term "topo").

See https://github.com/zpratt/topokube/issues/3 for feature ideas to include in v1.

## Project State

I'm still hacking on what this should look like, so there are no tests for anything. This is currently POC quality.

## Prerequisite Steps:

1. `brew update && brew install helm kind helmfile`
1. `helm plugin install https://github.com/databus23/helm-diff --version v3.0.0-rc.7`
   * at the time of writing this, there wasn't a stable release of the diff plugin, but this required to avoid errors with helm3 and using helmfile

## Running In Kind:

1. `sudo sh -c "echo '127.0.0.1 topokube.local' >> /etc/hosts"`
1. `make inkind`
1. `open http://topokube.local:30080/ui/index.html`