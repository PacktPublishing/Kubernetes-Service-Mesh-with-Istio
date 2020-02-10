# Connection Pools and Bulk Heading

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

Next, edit the destination rule definitions for `hello-istio` and  `hello-message` to configure the connection pools and bulk heads.

```yaml
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100             # maximum number of TCP conns
        connectTimeout: 5s              # TCP connection timeout
        tcpKeepalive:                   # keep alive settings
          time: 3600s
          interval: 60s
      http:
        maxRequestsPerConnection: 25    # max request per keep-alive
        http2MaxRequests: 5             # max number of HTTP2 conns
        http1MaxPendingRequests: 5      # max number of pending reqs
        maxRetries: 3                   # max number of retries
        idleTimeout: 60s                # idle timeout for connection
```

Now apply the modified destination rule definitions for the two services.

```bash
$ kubectl apply -f kubernetes/hello-istio-connection-pool.yaml
$ kubectl apply -f kubernetes/hello-message-connection-pool.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```