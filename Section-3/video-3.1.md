# Adding a Circuit Breaker

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

```bash
$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml

$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Next, edit the `kubernetes/hello-message-destination.yaml` and add the traffic policy and circuit breaker definitions.

```yaml
  trafficPolicy:
    outlierDetection:
      consecutiveErrors: 5    # 5 upstream errors (502, 503, 504)
      interval: 30s           # sliding window of 30s
      baseEjectionTime: 1m    # eject upstream for 1 minute
      maxEjectionPercent: 50  # max 50% of upstream hosts ejected
```

Now apply the virtual service and destination definitions for the `hello-message` service.

```bash
$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-destination.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```