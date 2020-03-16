# Mutual TLS between services

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster.
Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo
```

## Running

We will be using two different namespaces to demonstrate the security mTLS features.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl label namespace default istio-injection=disabled
```
