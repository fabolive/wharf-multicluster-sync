
# Setting your environment to run the demo

To run this demo you will need three Kubernetes clusters.  They should all be available
as contexts in the same Kubernetes configuration.  For example, you should be able to
run `kubectl --context <ctx1> get pods`, for three contexts. 

If you are using $KUBECONFIG you can export multiple config files for this behavior, e.g.

```
export KUBECONFIG=$KUBECONFIG1:$KUBECONFIG2:$KUBECONFIG3
```

To setup the test environments run the following, replacing _ctx-ca_, _ctx1_, and _ctx2_ with your cluster names:

```
source ./demo_context.sh ctx-ca ctx1 ctx2
```

# Configuring Istio to use a common Citadel 

The _install_citadel.sh_ script will configure $CLUSTER1 and $CLUSTER2 to use Citadel on $ROOTCA_NAME.

```
./install_citadel.sh
```

# Run the Multi-Cluster agents on demo clusters

In this demo we have Cluster 1 watching exposed services on Cluster 2.
For this purpose we need to deploy the MC agent on both clusters and configure Cluster 1's agent
to peer with Cluster 2's agent.

We first deploy the agent to `$CLUSTER2` which doesn't watch any other clusters (donor only):

```
./deploy_cluster.sh $CLUSTER2
```
> We need to configure cluster 2 first because the assigned LoadBalancer IP address to the agent service needs to be used for configuring the agent on cluster 1.

We then configure and deploy the agent on `$CLUSTER1` and ask it to peer with `$CLUSTER2` (the 2nd argument):

```
./deploy_cluster.sh $CLUSTER1 $CLUSTER2
```

The script will get the relevant information (Istio Gateway and MC Agent IP addresses) from Cluster 1 and use it in the peer configuration.

