# noddy
A docker host service running in the kubernestes cluster, to build/push the docker images in the cluster itself.

## Design
  A highlevel overview of the design plan:
  
  * Cluster side:
     * A receiver pod which would receive the build context and save it to a temp dir on hostPath of particular node (any sharable volume, need not to be Persistent one).
     * Another Pod with Kaniko container running in it, which would use the build context and build a docker image out of it and push it to a temp-docker registry.
     * temp-docker registry (Artifactory, nexus etc)
   
   * Client side:
      * CLI tool most probably kubectl subcommand `docker` with further subcommands like `init`(to install cluster receiver), `build` will (port-forward) send the context, and run kaniko container with docker args ,`detach` to break the established connection and `delete` to delete pod and service.
      
## Usage

Whenever the dev mode is up, otherwise down.
Use commands:
1. `kubectl docker init`
>  * deploy the receiver pod (kubectl apply -f noddy/services/receiver/deployment.yaml)
   * port-forward the pod in background (kubectl port-forward 6000:5000 &)
   * Print build-context receiver is ready..

2. `kubectl docker build --path /some/to/build-context -t my-app:latest`
>  * Tar the contents of --path `/some/to/build-context` with name of my-app.tar.gz
   * POST request with my-app.tar.gz
   * Deploy the DinD container pod (kubectl apply -f docker-agent.yaml)
   * Print the node on which DinD, receiver pods are running
        (The node has to be added as node-selector in my-app pod spec)
   * Delete the DinD pod(the src/ saved on hostPath of node must get clean-up ed).

3. `kubectl docker cleanup`
>  * Delete receiver deployment
   * stop port-forwarding
