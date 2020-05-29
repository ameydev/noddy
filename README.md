# noddy
A docker host service running in the kubernestes cluster's nodes, sharing and exposing cluster nodes docker daemon to kubectl clients

## Design
  A highlevel overview of the design plan:
  
  * Cluster side:
     * Pod running with a container (Docker in Docker), acting as docker host
     * A service which exposes the DOCKER_HOST daemon across cluster and to kubectl client machine using port-forward.
   
   * Client side:
      * CLI tool most probably kubectl subcommand `docker` with further subcommands like `init`(to install the pod and service in cluster), `attach` to establish connection between `DOCKER_HOST` with client machine (port-forwarding),`detach` to break the established connection and `delete` to delete pod and service.
      
## Usage

Once the connection is established, any docker commands executed on client machine should perform docker operations on remote host which the cluster node.
