## Cert Manager

[Install cert-manager](https://cert-manager.io/docs/installation/helm/):
```sh
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.17.2 \
  --set crds.enabled=true
```

```sh
kubectl apply -f cluster-issuer.yaml
```

## Keycloak

```sh
helm upgrade --install keycloak bitnami/keycloak \
  --version 24.6.1 \
  --values kc-values.yaml \
  --namespace keycloak \
  --create-namespace
```

## MongoDB

```sh
helm install my-mongodb bitnami/mongodb --version 16.5.13 -n message --create-namespace
```