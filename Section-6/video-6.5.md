# Access Logs with Envoy

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

Also, make sure to deploy the demo files for this section.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl apply -f kubernetes/
```

## Running

First, we need to make sure that the access log feature for Istio has been enabled.

```bash
$ istioctl profile dump demo | grep accessLog

$ istioctl manifest apply --set profile=demo --set values.global.proxy.accessLogEncoding="JSON" --set values.global.proxy.accessLogFile="/dev/stdout"
```

Next, we generate some traffic on the demo microservices and check that that the access logging is working as expected.

```bash
# generate some traffic
$ hey -z 5s http://$INGRESS_HOST/api/hello Host:hello-istio.cloud
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# in a new window, use the kubectl tail plugin to watch the logs
$ kubectl logs -n hello-istio -c istio-proxy hello-istio-v1-..... -f
$ kubectl tail -n hello-istio -c istio-proxy
```
