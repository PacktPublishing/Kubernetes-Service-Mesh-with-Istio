# Retries

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

```
$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml
$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-destination.yaml

$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

Next, edit the virtual service definitions for `hello-istio` and  `hello-message` to configure the retry behaviour.

```yaml
    retries:
      attempts: 3
      perTryTimeout: 500ms
      retryOn: gateway-error,connect-failure,refused-stream
```

Now apply the modified destination rule definitions for the two services.

```bash
$ kubectl apply -f kubernetes/hello-istio-retries.yaml
$ kubectl apply -f kubernetes/hello-message-retries.yaml
```
