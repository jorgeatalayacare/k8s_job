# Kubernetes job sample

## Steps to deploy

### Environment Config  
If you’re using MAC silicon then this you have installed both [brew](https://brew.sh/) and rosetta. To install rosetta:

```bash
softwareupdate --install-rosetta
```

1. Install kind.
```bash
brew install kind
```

2. Deploy a multi-node cluster.
```bash
kind create cluster --config ./k8s/cluster.yaml
```

3. Register the image and deploy your job
```bash
docker build -t file-reader:0.1.0 .
kind load docker-image file-reader:0.1.0
kubectl apply -f ./k8s/job.yaml
```