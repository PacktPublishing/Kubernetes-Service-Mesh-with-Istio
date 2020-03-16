NAME = istio-in-action
VERSION = 1.5.0

.PHONY: info

info:
	@echo "Kubernetes Service Meshes with Istio"

prepare:
	@gcloud config set compute/zone europe-west1-b
	@gcloud config set container/use_client_certificate False

cluster:
	@gcloud container clusters create $(NAME) --num-nodes=7 --enable-autoscaling --min-nodes=5 --max-nodes=10 --machine-type=n1-standard-2 --enable-network-policy --no-enable-autoupgrade
	@kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=$$(gcloud config get-value core/account)
	@kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
	@kubectl cluster-info

access-token:
	@gcloud config config-helper --format=json | jq .credential.access_token

dashboard:
	@kubectl proxy & 2>&1
	@sleep 3
	@gcloud config config-helper --format=json | jq .credential.access_token
	@open http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/

helm-install:
	@echo "Installing Helm"
	@curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get | bash

helm-init:
	@echo "Initialize Helm"

	# add a service account within a namespace to segregate tiller
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/helm-service-account.yaml

	@helm init --service-account tiller
	@helm repo update

	# verify that helm is installed in the cluster
	@kubectl get deploy,svc tiller-deploy -n kube-system

get-istio:
	@curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$(VERSION) sh -
	@echo "Make sure to export the PATH variable."

istio-manual:
	# deploy Istio manually
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-10.yaml
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-11.yaml
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-12.yaml
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-10.yaml
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-11.yaml
	@sleep 10
	@kubectl apply -f istio-$(VERSION)/install/kubernetes/istio-demo.yaml
	@sleep 5
	@kubectl get pods -n istio-system
	@kubectl get svc istio-ingressgateway -n istio-system

istio-helm:
	# deploy Istio via Helm
	@helm repo add istio.io https://storage.googleapis.com/istio-release/releases/$(VERSION)/charts/

	@helm install istio-$(VERSION)/install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
	@kubectl get crds | grep 'istio.io' | wc -l
	@sleep 10

	# install istio with demo template
	@helm install istio-$(VERSION)/install/kubernetes/helm/istio --name istio --namespace istio-system --values istio-$(VERSION)/install/kubernetes/helm/istio/values-istio-demo.yaml

	# Verifying the installation
	@kubectl get svc -n istio-system
	@kubectl get pods -n istio-system

istio-delete:
	@kubectl delete -f istio-$(VERSION)install/kubernetes/istio-demo.yaml
	@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-10.yaml
	@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-11.yaml
	@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-12.yaml
	@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-10.yaml
	@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-11.yaml
	@kubectl label namespace default istio-injection-

destroy:
	@gcloud container clusters delete $(NAME) --async --quiet

clean:
	@rm -rf istio-$(VERSION)
