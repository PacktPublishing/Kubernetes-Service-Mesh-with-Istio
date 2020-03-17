# Metrics with Prometheus

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

Check that Prometheus is running correctly and use `istioctl` to open the dashboard. Alternatively, use port-forwarding.

```bash
$ kubectl -n istio-system get service prometheus

$ kubectl port-forward -n istio-system service prometheus 9090
$ istioctl dashboard prometheus

$ open http://localhost:9090
```

In the Prometheus dashboard have a look at the different available metrics, execute queries and display some graphs. Here are some example queries:

- http_requests_total
- istio_requests_total{destination_app="hello-istio"}
- istio_requests_total{destination_app="hello-istio"}
- istio_requests_total{destination_app="hello-istio", destination_version="v2"}