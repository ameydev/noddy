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

`kubectl docker init`

`kubectl docker build -t test-data .`

`kubectl docker cleanup`
