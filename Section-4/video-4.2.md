# Injecting HTTP Abort Faults

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

First, make sure everything is running correctly without any faults.
```
$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Next, configure a HTTP abort fault for any traffic to the hello-message v1 virtual service.

```yaml
- fault:
    abort:
      httpStatus: 500
      percentage:
        value: 100.0
  # optionally add header match here
  route:
  - destination:
      host: hello-message
      subset: v1
```

Apply the modified virtual service and check that the HTTP abort fault is configured correctly.

```
$ kubectl apply -f kubernetes/hello-message-v1-abort.yaml

# you should see an error message about the abort filter
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Bonus: play around with the HTTP fault code and the percentage.
