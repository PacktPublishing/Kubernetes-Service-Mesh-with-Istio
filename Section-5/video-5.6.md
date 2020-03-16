# Authorization with JWT

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Also, make sure you have followed the setup instructions from the previous videos.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml
$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-destination.yaml

$ kubectl apply -f kubernetes/hello-istio-secure.yaml
$ kubectl apply -f kubernetes/hello-istio-insecure.yaml

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

Optionally, if you want to experiment with your own JWT and JWKS then follow these instructions to create your own credentials.

```bash
# download the JWT generator from the Istio repository
$ wget https://raw.githubusercontent.com/istio/istio/master/security/tools/jwt/samples/gen-jwt.py -O data/gen-jwt.py

$ cd data/
$ openssl genrsa -out key.pem 2048
$ pip3 install jwcrypto
$ python3 gen-jwt.py key.pem --iss packtpub --sub demo --aud students --jwks=./jwks.json --expire=3153600000 --claims=publisher:packtpub > packtpub.jwt
```

## Running

