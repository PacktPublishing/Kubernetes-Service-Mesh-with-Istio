# Using Envoy Filters

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-v1.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml

$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-destination.yaml
```

## Running

First, make sure everything is running correctly.

```
$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# check the container logs of the hello-message v1 pod
$ kubectl logs hello-message-v1-6dcc4fff9-hnbxs -c hello-message
```

Apply the prepared envoy filter manifest and check that everything is working as expected. It does take a while for the sidecar to pick up the new filter configuration.

```
$ kubectl apply -f kubernetes/hello-message-v1-filter.yaml

$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ kubectl logs hello-message-v1-6dcc4fff9-hnbxs -c hello-message
```