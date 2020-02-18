# Injecting HTTP Delay Faults

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

First, make sure everything is running correctly without delays.
```
$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Next, configure a HTTP delay fault for any traffic to the hello-message v1 virtual service.

```yaml
- fault:
    delay:
      percentage:
        value: 100.0
      fixedDelay: 5s
  # optionally add header match here
  route:
  - destination:
      host: hello-message
      subset: v1
```

Apply the modified virtual service and check that the HTTP delay is configured correctly.

```
$ kubectl apply -f kubernetes/hello-message-v1-delay.yaml

# you should see an error message after 3s delay -> timeout working
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Bonus: play around with the HTTP delay value and the percentage.