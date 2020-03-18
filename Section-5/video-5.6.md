# Authorization with JWT

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

Optionally, if you want to experiment with your own JWT and JWKS then follow these instructions to create your own credentials.

```bash
cd data/

# download the JWT generator from the Istio repository
$ wget https://raw.githubusercontent.com/istio/istio/release-1.5/security/tools/jwt/samples/gen-jwt.py

# (optionally) create your own key
$ openssl genrsa -out key.pem 2048

# generate a new JWKS and JWT data set
$ pip3 install jwcrypto
$ python3 gen-jwt.py key.pem --iss packtpub --sub demo --aud students --jwks=./jwks.json --expire=3153600000 --claims=publisher:packtpub > packtpub.jwt
```

## Running

Here, we will apply the JWT authentication and authorization policies to the Istio service mesh. Without the policies, everything should work as expected. Once we have applied the policies, all requests without a JWT are denied with a `403 Forbidden`. Only requests with the correct JWT bearer token are accepted.

```bash
# every call works
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# all requests without JWT are denied with 403
$ kubectl apply -f kubernetes/hello-istio-jwt-authz.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# prepare the JWT Bearer token
$ export BEARER_TOKEN="Bearer $(cat data/packtpub.jwt)"
$ echo $BEARER_TOKEN

# a request with correct JWT bearer token is accepted
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud Authorization:$BEARER_TOKEN
```