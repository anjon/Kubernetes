# Initialize the Helm Chart

**This is a basic steps for installing & initializing a helm chart**

```sh
# Installing helm
wget https://get.helm.sh/helm-v2.16.10-linux-amd64.tar.gz
tar -xvzf helm-v2.16.10-linux-amd64.tar.gz
cd linux-amd64/
sudo mv helm /usr/local/bin/
which helm
helm version
helm version --short --client
kubectl -n kube-system create serviceaccount tiller
kubectl create clusterrolebinding tiller --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl -n kube-system get sa,clusterrolebindings |grep -i tiller

# Initializing helm 
helm init --service-account tiller
helm help | less
helm home
ls .helm/
helm list
helm search jenkins
helm help imspect
helm inspect stable/jenkins
helm fetch stable/jenkins
tar -xvzf jenkins-2.5.2.tgz
cd jenkins/
ls -l templates/
rm -rf jenkins jenkins-2.5.2.tgz
kubectl -n kube-system get deploy,rs,po,sa,clusterrolebinding | grep -i tiller
helm help reset
helm reset --remove-helm-home
kubectl -n kube-system get deploy,rs,po,sa,clusterrolebinding | grep -i tiller
kubectl -n kube-system delete clusterrolebindings tiller
kubectl -n kube-system delete sa tiller
```
### Creating your own chart 
```sh
helm create mychart
tree mychart/
```
### Helm Migrating from v2 to v3 
```sh
# Install a Release to test during migration
helm install stable/phpmyadmin --name phpmyadmin
helm list

helm install stable/phpmyadmin --name phpmyadmin
tar xzf helm-v3.3.0-linux-amd64.tar.gz
sudo mv helm /usr/local/bin/helm3
which helm3
helm3 version --short
helm repo list
helm3 repo list
helm list
helm3 list
helm3 plugin install https://github.com/helm/helm-2to3
helm3 plugin list

helm3 2to3 move --help
helm3 2to3 move config --dry-run
helm3 2to3 move config

helm3 2to3 convert --help
helm3 2to3 convert phpmyadmin --dry-run
helm3 2to3 convert phpmyadmin

helm3 2to3 cleanup
sudo rm -rf /usr/local/bin/helm
sudo mv /usr/local/bin/{helm3,helm}

which helm
helm verison --short
helm search repo phpmyadmin
helm install phpmyadmin stable/phpmyadmin
watch kubectl get all 
```

### Create Custom Helm Chart form Scratch 
```sh
helm version --short
mkdir charts ; cd charts
mkdir my-nginx ; cd my-nginx
vim Chart.yaml
apiVersion: v1
name: my-nginx
description: My custom nginx chart
version: 0.4.0
appVersion: 1.0

mkdir templates
kubectl create deployment my-nginx --image=nginx --dry-run=client -o yaml > charts/my-nginx/templates/deployment.yaml
helm install --name my-nginx .
kubectl expose deployment my-nginx --port=80 --dry-run=client -o yaml > templates/service.yaml
helm list
helm rollback my-nginx 1
helm rollback my-nginx 2
helm delete --purge my-nginx
vim values.yaml
replicaCount: 1
service:
  type: NodePort

helm install --name myn-nginx . --set replicaCount=2
helm delete --purge myn-nginx
helm inspect values . > /tmp/my-nginx.values
helm install --name my-nginx . --values /tmp/my-nginx.values
helm upgrade my-nginx . --set replicaCount=2
helm upgrade my-nginx .
helm upgrade my-nginx . --set service.type=LoadBalancer
helm delete --purge my-nginx

helm create myapp
tree charts
helm list
```

### Local Helm Chart Repository
```sh
tree charts/
cd charts/
helm package my-nginx
helm serve --repo-path .
helm repo list
helm repo update
helm repo remove local
helm repo add local http://127.0.0.1:8879/charts
helm repo list
helm search local/

helm serve --repo-path . --address "0.0.0.0:5000"  # Serve for other cluster
helm repo add myrepo http://localhost:5000/charts
helm repo update
helm repo remove local
helm repo list
helm repo update
helm search myrepo/
helm install --name my-nginx myrepo/my-nginx
helm list
vim charts/my-nginx/Chart.yaml
helm package my-nginx
helm repo index .
helm repo update
helm search myrepo/
helm repo list
helm delete --purge my-nginx
```
