## Welcome

We're really happy that you're considering joining us!
This challenge will help us understand your skills and will also be a starting point for the next interview.
We're not expecting everything to be done perfectly as we value your time but the more you share with us, the more we get to know about you!

This challenge is split into 3 parts:

1. Debugging
2. Implementation
3. Questions

If you find possible improvements to be done to this challenge please let us know in this readme and/or during the interview.

## The challenge

Pleo runs most of its infrastructure in Kubernetes.
It's a bunch of microservices talking to each other and performing various tasks like verifying card transactions, moving money around, paying invoices, etc.
This challenge is similar but (a lot) smaller :D

In this repo, we provide you with:

- `invoice-app/`: An application that gets invoices from a DB, along with its minimal `deployment.yaml`
- `payment-provider/`: An application that pays invoices, along with its minimal `deployment.yaml`
- `Makefile`: A file to organize commands.
- `deploy.sh`: A file to script your solution
- `test.sh`: A file to perform tests against your solution.

### Set up the challenge env

1. Fork this repository
2. Create a new branch for you to work with.
3. Install any local K8s cluster (ex: Minikube) on your machine and document your setup so we can run your solution.

### Part 1 - Fix the issue

The setup we provide has a :bug:. Find it and fix it! You'll know you have fixed it when the state of the pods in the namespace looks similar to this:

```
NAME                                READY   STATUS                       RESTARTS   AGE
invoice-app-jklmno6789-44cd1        1/1     Ready                        0          10m
invoice-app-jklmno6789-67cd5        1/1     Ready                        0          10m
invoice-app-jklmno6789-12cd3        1/1     Ready                        0          10m
payment-provider-abcdef1234-23b21   1/1     Ready                        0          10m
payment-provider-abcdef1234-11b28   1/1     Ready                        0          10m
payment-provider-abcdef1234-1ab25   1/1     Ready                        0          10m
```

#### Requirements

After building images and applying _deployment.yaml_ files pods were not starting giving errors (`kubectl describe pod name`): _"Error: container has runAsNonRoot and image will run as root kubernetes"_

```
NAME                           READY   STATUS                       RESTARTS   AGE
invoice-app-5c59847c9d-2q7hj   0/1     CreateContainerConfigError   0          3m49s
invoice-app-85bf5d4fbf-7c22t   0/1     CreateContainerConfigError   0          6m33s
invoice-app-85bf5d4fbf-df8gr   0/1     CreateContainerConfigError   0          6m33s
invoice-app-85bf5d4fbf-mxhvq   0/1     CreateContainerConfigError   0          6m33s
```

I have added line to Dockerfiles to let them run as non-root user and rebuild docker images.

```
$ k get pods
NAME                               READY   STATUS    RESTARTS   AGE
invoice-app-85bf5d4fbf-2r4tp       1/1     Running   0          22m
invoice-app-85bf5d4fbf-df8gr       1/1     Running   0          30m
invoice-app-85bf5d4fbf-h2bvj       1/1     Running   0          23m
payment-provider-fd68f8c7c-24zw7   1/1     Running   0          17m
payment-provider-fd68f8c7c-4hghs   1/1     Running   0          14m
payment-provider-fd68f8c7c-69wvw   1/1     Running   0          14m
```

### Part 2 - Setup the apps

We would like these 2 apps, `invoice-app` and `payment-provider`, to run in a K8s cluster and this is where you come in!

#### Requirements

1. `invoice-app` must be reachable from outside the cluster.
2. `payment-provider` must be only reachable from inside the cluster.
3. Update existing `deployment.yaml` files to follow k8s best practices. Feel free to remove existing files, recreate them, and/or introduce different technologies. Follow best practices for any other resources you decide to create.
4. Provide a better way to pass the URL in `invoice-app/main.go` - it's hardcoded at the moment
5. Complete `deploy.sh` in order to automate all the steps needed to have both apps running in a K8s cluster.
6. Complete `test.sh` so we can validate your solution can successfully pay all the unpaid invoices and return a list of all the paid invoices.

### Part 3 - Questions

Feel free to express your thoughts and share your experiences with real-world examples you worked with in the past.

#### Requirements

1. What would you do to improve this setup and make it "production ready"?
2. There are 2 microservices that are maintained by 2 different teams. Each team should have access only to their service inside the cluster. How would you approach this?
3. How would you prevent other services running in the cluster to communicate to `payment-provider`?

## What matters to us?

We expect the solution to run but we also want to know how you work and what matters to you as an engineer.
Feel free to use any technology you want! You can create new files, refactor, rename, etc.

Ideally, we'd like to see your progression through commits, verbosity in your answers and all requirements met.
Don't forget to update the README.md to explain your thought process.
