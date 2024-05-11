# kcd-dhaka-gitops-demo

## ArgoCD Installations (this is needed only once initially)

- create `argocd` namespace: `kubectl create namespace argocd`
- install `argocd`: `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`:
  - make sure pods are running: `kubectl get pods -n argocd`
  - check the services: `kubectl get svc -n argocd`
- port forward of the `argocd` server (to access the `argocd` UI): `kubectl port-forward -n argocd svc/argocd-server 8080:443`
- login to argocd UI(`localhost:8080`):
  - username: `admin`
  - password: get by `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo`
- connect the git repository from argocd UI:
  - go to settings -> repositories -> add github (project: default, )
  - you may need to provide git personal access token for the respective git repo
- apply argocd application: `kubectl apply -f argocd/application.yaml`
