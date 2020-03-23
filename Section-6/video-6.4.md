# Call Tracing with Jaeger

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

Check that Jaeger is running correctly and use `istioctl` to open the dashboard. Alternatively, use port-forwarding.

```bash
$ kubectl -n istio-system get svc jaeger-query

$ kubectl port-forward -n istio-system services/jaeger-query 16686
$ open http://localhost:16686

$ istioctl dashboard jaeger

# generate some traffic
$ hey -z 5s http://$INGRESS_HOST/api/hello Host:hello-istio.cloud
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```

For this to work, make sure to propagate the trace context HTTP headers to downstream services as described here: https://istio.io/docs/tasks/observability/distributed-tracing/overview/