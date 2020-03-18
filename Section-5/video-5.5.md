# Authorization for HTTP traffic

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Also, make sure you have followed the setup instructions from the previous videos.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl apply -f kubernetes/demo/

$ kubectl apply -f kubernetes/hello-istio-secure.yaml
$ kubectl apply -f kubernetes/hello-istio-insecure.yaml

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

Any HTTP traffic inside the service mesh can be enabled to disabled using an `AuthorizationPolicy`.

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: hello-message-http-policy
  namespace: hello-istio
spec:
  selector:
    matchLabels:
      app: hello-message
  # toggle between DENY and ALLOW
  action: DENY
  rules:
  - to:
      - operation:
          methods: ["GET"]
    from:
      - source:
          namespaces:
            - "hello-istio"
```

Apply the above policy and check that the HTTP communication within the `hello-istio` namespace is not possible anymore but it is possible from outside the namespace.

```bash
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ kubectl apply -f kubernetes/hello-message-http-policy.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ kubectl exec -it hello-istio-insecure-6969cf44bf-.... /bin/sh
$ wget hello-message.hello-istio.svc.cluster.local:8080/api/message/hello -S -O - | more

# do some cleanup
$ kubectl delete AuthorizationPolicy -n hello-istio --all
```


