# K8S Gitops manager

# Features 

## Gitops

TODO

## Based on tools you love 

K8S Gitops Manager is using Kubernetes and Helm to package and deploy your applications. It means that you can import your existing charts.

## Template composition

Your usual values.yaml handling is adapted to Gitops. It means that you can use some built-in values like `{{.Branch}}` or `{{.BranchUrlSafe}}`. They will be computed by K8S gitops manager before being packaged and installed by Tiller.

## Based and webhooks

Unlike Weave Flux, branches will be triggered by Webhooks. It means it will be instant to deploy changes to Kubernetes. Github will the first one to be implemented.

## Secrets (Step 2)

Just like [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets), this tool will encrypt strings before being saved to Github, and will decrypt them when creating a new Helm release.

## Configuration/secret management (Step 2)

The tool will contain an UI to handle configuration and secrets the Gitops way. More on this later.

## Next 
- Deployment management/visualisation
- Canary deployments


### Development

#### Prerequesites

Install [air](https://github.com/cosmtrek/air) (Golang live reload) and [kubefwd](https://github.com/txn2/kubefwd) (port forwarding to Helm Tiller). 

```
$ sudo kubefwd services -n kube-system => started port forwarding

$ GITHUB_PAT=*** air => runs golang server

$ cd frontend && npm install && npm start => runs frontend server
```