# Setup
```bash
brew install kubectl
brew install kustomize
```

# development クラスタの作成
ingressを使えるようにセットアップ

See: https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress

```bash
cat <<EOF | kind create cluster --name development --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 10080
    protocol: TCP
  - containerPort: 443
    hostPort: 10443
    protocol: TCP
EOF

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

# production クラスタの作成
ingressを使えるようにセットアップ

See: https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress

```bash
cat <<EOF | kind create cluster --name production --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 20080
    protocol: TCP
  - containerPort: 443
    hostPort: 20443
    protocol: TCP
EOF

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

# デプロイ

## development
imageのビルド
```bash
docker build -t my-app:v1 .
```
imageをkindにload
```bash
kind load docker-image -n development my-app:v1
```

デプロイ
```bash
kubectl config use-context kind-development
kustomize build manifests/overlayes/development | kubectl apply -f -
```

## production
imageのビルド
```bash
docker build -t my-app:v1 .
```
imageをkindにload
```bash
kind load docker-image -n production my-app:v1
```

デプロイ
```bash
kubectl config use-context kind-production
kustomize build manifests/overlayes/production | kubectl apply -f -
```

# 動作確認

## development
```bash
curl localhost:10080
```

## production
```bash
curl localhost:20080
```
