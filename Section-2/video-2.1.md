# Controlling Ingress Traffic

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled
```

## Running

```
$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

$ echo $INGRESS_HOST

# deploy sample application
$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl get all

# create ingress gateway and route traffic to microservices
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```
