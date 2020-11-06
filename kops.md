## Deploy Kubernetes CLuster on aws using KOPS

#### Configure awscli, kops & kubectl
```sh
sudo apt install awscl
which awscli 

wget https://github.com/kubernetes/kops/releases/download/v1.18.2/kops-linux-amd64
sudo mv kops-linux-amd64 /usr/local/bin/kops
chmod +x /usr/local/bin/kops
which kops
kops version --short

curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl version --short
```

#### Create KOPS Cluster
```sh
aws s3api create-bucket --bucket anjon-kops-bucket --region us-east-1
aws s3 ls
aws s3 ls anjon-kops-bucket
aws s3api put-bucket-versioning --bucket anjon-kops-bucket --region us-east-1 --versioning-configuration Status=Enabled
export KOPS_STATE_STORE=s3://anjon-kops-bucket
kops help
kops create cluster --name myops.k8s.local --zones us-east-1a --master-size t2.micro --node-size t2.micro --kubernetes-version 1.18.0
kops get cluster
kops edit cluster --name myops.k8s.local
kops update cluster --name myops.k8s.local --yes
kops validate cluster --name myops.k8s.local
```

#### Validate the Kubernetes Cluster
```sh 
kubectl get nodes 
kubectl get ns 
kubectl get sc
kubectl -n kube-system get all 
kubectl create deployment nginx --image=nginx --dry-run=client
kubectl get all 
kubectl port-forward nginx-xxxx 8080:80 
https://localhost:8080
kubectl delete deployment 
aws s3 ls anjon-kops-bucket
kops get cluster
```

## Upgrading Kubernetes Cluster Using KOPS

#### Upgrading Cluster to Specific Version
```sh
kops get cluster
kubectl version --short
kubectl get nodes
kops edit cluster --name myops.k8s.local
kubernetesversion: 1.18.3
kops update cluster --name myops.k8s.local --yes 
kops rolling-update cluster --name myops.k8s.local --yes 
kops validate cluster --name myops.k8s.local 
kubectl get no 
kubectl version --short
```

#### Upgrading Clustr to Latest Version
```sh
kops upgrade cluster --name myops.k8s.local --yes 
kops update cluster --name myops.k8s.local --yes 
kops rolling-update cluster --name myops.k8s.local --yes 
kops validate cluster --name myops.k8s.local 
kubectl get nodes
kubectl version --short 
```

## Kops Scaling Cluster
```sh
kops validate cluster --name myops.k8s.local
kubectl get nodes
kops get cluster 
kops get instancegroups --name myops.k8s.local 
kops edit instancegroups nodes --name myops.k8s.local 
set maxSize=1, minSize=1
kops update cluster --name myops.k8s.local --yes 
kops rolling-update cluster --name myops.k8s.local --yes 
kops validate cluster --name myops.k8s.local
kubectl get nodes 
```
