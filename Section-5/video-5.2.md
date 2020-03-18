# Mutual TLS between services

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster.
Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

We will be using two different namespaces to demonstrate the security mTLS features.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

# The default namespace will not have the sidecar injection
$ kubectl label namespace default istio-injection=disabled
```

Next, we will apply the demo set of microservices. Per default, the mTLS between services is in PERMISSIVE mode, meaning that encrypted and unencrypted traffic is allowed in the service mesh.

```bash
$ kubectl apply -f kubernetes/demo/

$ kubectl apply -f kubernetes/hello-istio-secure.yaml
$ kubectl apply -f kubernetes/hello-istio-insecure.yaml

$ kubectl get all -n hello-istio
$ kubectl get all -n default
```

First, we check that we can call a service from the secure namespace and console pod:
```bash
$ kubectl exec -n hello-istio -c console -it hello-istio-secure-..... /bin/sh

$ wget hello-istio:8080/api/hello -S -O - | more
$ wget hello-istio.hello-istio.svc.cluster.local:8080/api/hello -S -O - | more
```

Next, we check that we can also call the service from an insecure namespace and console pod:
```bash
$ kubectl exec -it hello-istio-insecure-6969cf44bf-..... /bin/sh

$ wget hello-istio.hello-istio.svc.cluster.local:8080/api/hello -S -O - | more
```





