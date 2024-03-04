# cdktf-tests

This repository allows you to deploy, infrastructures and apps, with CDK for Terraform (CDKTF).

## Prerequisites

* Install cdktf CLI

## Docker

This repository allows you to deploy, through CDKTF, two apps in containers and a network to communicate to each others.

frontend -> backend

```bash
cd docker
cdk deploy
```

## OVHcloud

```bash
cd ovhcloud
cdktf deploy
```
