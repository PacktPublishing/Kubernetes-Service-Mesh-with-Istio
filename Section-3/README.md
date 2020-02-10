# Section 3: Service Resilience

Everything fails, all the time! In this section, you will see how to apply resiliency to the services of a service mesh. You will learn how to add and configure request timeouts and circuit breakers for your services. You will also learn to apply connections pools, retries and rate limiting.

- **Video 3.1**: [Adding a Circuit Breaker](video-3.1.md)
- **Video 3.2**: [Setting request timeouts](video-3.2.md)
- **Video 3.3**: [Connection Pools and Bulk Heading](video-3.3.md)
- **Video 3.4**: [Retries](video-3.4.md)
- **Video 3.5**: [Rate Limiting](video-3.5.md)

## Prerequisites

This section assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```
