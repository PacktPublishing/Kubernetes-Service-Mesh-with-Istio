# Installing Istio using Helm

In this video are going to install Istio using the Helm package manager.

## Prerequesites

You will need a running Kubernetes with enough resources. If you install
Istio locally in Docker Desktop make sure you followed the instructions
under https://istio.io/docs/setup/platform-setup/docker/

## Step 1: Setup and Verify

First, you need to download and setup the latest Istio release.
At the time of writing this is `1.3.4`, upgrade to latest version if desired.
```
$ curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.4 sh -

$ cd istio-1.3.4
$ export PATH=$PWD/bin:$PATH

$ istioctl verify-install
```

## Step 2: Install Helm

In this step we are installing and initializing the Helm package manager
in the Kubernetes cluster

```
# Install Helm
$ curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get | bash

# "Initialize Helm"
# add a service account within a namespace to segregate tiller
kubectl apply -f install/kubernetes/helm/helm-service-account.yaml

helm init --service-account tiller
helm repo update

# verify that helm is installed in the cluster
kubectl get deploy,svc tiller-deploy -n kube-system
```

## Step 3: Install Istio

```
# deploy Istio via Helm
$ helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.3.4/charts/

# Install the istio-init chart to bootstrap all the Istio CRDs
$ helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
$ kubectl get crds | grep 'istio.io' | wc -l

# install the istio chart using demo profile
$ helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml

# Verifying the installation
$ kubectl get pods -n istio-system
$ kubectl get svc -n istio-system
```

## (Optional) Step 4: Uninstall Istio

In case you want to uninstall Istio using Helm, issue the following commands:
```
$ helm delete --purge istio
$ helm delete --purge istio-init
$ kubectl label namespace default istio-injection-
```
