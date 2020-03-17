# Section 6: Diagnosability: Monitoring, Tracing, Visualization

Good diagnosability is a cornerstone of any cloud-native application and service mesh to have a holistic view on metrics, traces and the service graph. In this section, you will learn about the diagnosability tools Istio is integrated with.

- **Video 6.1**: The Diagnosability Triangle
- **Video 6.2**: [Metrics with Prometheus](video-6.2.md)
- **Video 6.3**: [Ops Dashboards with Grafana](video-6.3.md)
- **Video 6.4**: [Call Tracing with Jaeger](video-6.4.md)
- **Video 6.5**: [Access Logs with Envoy](video-6.5.md)
- **Video 6.6**: [Mesh Visualization with Kiali](video-6.6.md)

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
