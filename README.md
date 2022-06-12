# rke

All steps is run on bastion host.

Please update file hosts.ini and bastion.ini to adjust ip and server access credential

### 1. Setup Bastion

$ ansible-playbook -i bastion.ini setup-bastion.yml

### 2. Create RKE User

$ ansible-playbook -i hosts.ini create-user-rke.yml

### 3. Install pre setup kubernetes

$ ansible-playbook -i hosts.ini pre-setup-kubernetes.yml

### 4. Spin UP RKE Kubernetes cluster

Please update file cluster.yml to adjust ip and server kubernetes that you have. and then run :

$ rke up

### 5. Running kubectl non-root
To make kubectl work for your non-root user, run these commands, which are also part of the kubeadm init output.

$ mkdir -p $HOME/.kube

$ sudo cp -i /rke/admin.conf $HOME/.kube/config

$ sudo chown $(id -u):$(id -g) $HOME/.kube/config


## Install Rook Ceph

Run all step below under directory rook/deploy/examples and make sure all pods running before running every the next step, "kubectl get po -n rook-ceph"

$ kubectl create -f crds.yaml

$ kubectl create -f common.yaml

$ kubectl create -f operator.yaml

$ kubectl config set-context --current --namespace rook-ceph

$ kubectl create -f cluster.yaml

$ kubectl  apply  -f toolbox.yaml

$ kubectl create -f filesystem.yaml

$ kubectl create -f csi/cephfs/storageclass.yaml
