## To Create a EKS (Amazon Elastic Kubernetes Service)

#### Step 01:
To launch a eks cluster in aws we need 3 component to prepare 
...awscli
b) eksctl
c) kubectl

```sh
sudo apt install awscli
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
which eksctl
eksctl version
curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.18.9/2020-11-02/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin
which kubectl
kubectl version --short --client
. <(eksctl completion bash)  ---> add to the .bashrc for bash auto completion
sudo kubectl completion bash >/etc/bash_completion.d/kubectl
```
#### Step 02:
Now as our environment is ready we can proceed to the cluster creation
```sh
aws configure
eksctl create cluster --help
eksctl create cluster --name first-eks --region us-east-1 --nodegroup-name standard-nodes --node-type t3.small --managed
kubectl get nodes
kubectl -n kube-system get pods
kubectl create deployment nginx --image=nginx --dry-run=client -o yaml > simple-deployment.yaml
kubectl create -f simple-deployment.yaml
```
