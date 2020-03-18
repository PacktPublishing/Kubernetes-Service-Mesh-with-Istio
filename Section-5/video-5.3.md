# Enabling Strict Mode

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster.
Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

Also, this video assumes that you have followed the instructions from the previous video.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled
$ kubectl label namespace default istio-injection=disabled

$ kubectl apply -f kubernetes/demo/

$ kubectl apply -f kubernetes/hello-istio-secure.yaml -n hello-istio
$ kubectl apply -f kubernetes/hello-istio-insecure.yaml -n default

$ kubectl get all -n hello-istio
$ kubectl get all -n default
```

## Running

In order to enable strict mTLS mode, you have to create and enable a `PeerAuthentication` object on a namespace basis.

```yaml
apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: "default"
spec:
  mtls:
    mode: STRICT
```

One you applied the strict mTLS configuration, you will notice that all pods are in an unready state.

```bash
$ kubectl apply -f kubernetes/hello-istio-strict-mtls.yaml -n hello-istio

# check all pods and deployments -> all are not ready anymore
$ kubectl get all -n hello-istio

# enable strict mTLS for the whole service mesh
$ kubectl apply -f kubernetes/hello-istio-strict-mtls.yaml -n kube-system
```

In order to fix the problem with the readyness and liveness probes you have two options: either you separate the probes from the HTTP traffic by specifying a dedicated port, or you use annotations to
rewrite the HTTP probe automatically during sidecar injection.

```bash
$ kubectl apply -f kubernetes/hello-istio-mtls-ports.yaml
$ kubectl apply -f kubernetes/hello-istio-mtls-annotations.yaml
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
