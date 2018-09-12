package agent

const (
	ConnectionModeKey       = "connection"
	ConnectionModeLive      = "live"
	ConnectionModePotential = "potential"
)

// ExposedServices is a struct that holds list of entries each holding the
// information about an exposed service. JSON format of this struct is being
// sent back from a remote cluster's agent in response to an exposition request.
type ExposedServices struct {
	Services []*ExposedService
}

// ExposedService holds description of an exposed service that is visible to
// remote clusters.
type ExposedService struct {
	Name      string
	Namespace string
	Port      uint16
}

// ClusterConfig holds all the configuration information about the local
// cluster as well as the peered remote clusters.
type ClusterConfig struct {
	ID string `json:"id"`

	GatewayIP   string `json:"gatewayIP"`
	GatewayPort uint16 `json:"gatewayPort"`

	AgentIP   string `json:"agentIP"`
	AgentPort uint16 `json:"agentPort"`

	Peers        []ClusterConfig `json:"peers,omitempty"`
	TrustedPeers []string        `json:"trustedPeers,omitempty"`
}

// Ip is implementing the model.ClusterInfo interface
func (cc ClusterConfig) Ip(name string) string {
	if name == cc.ID {
		return cc.GatewayIP
	}
	for _, peer := range cc.Peers {
		if name == peer.ID {
			return peer.GatewayIP
		}
	}
	return "255.255.255.255" // dummy value for unknown clusters
}

// Port is implementing the model.ClusterInfo interface
func (cc ClusterConfig) Port(name string) uint32 {
	if name == cc.ID {
		return uint32(cc.GatewayPort)
	}
	for _, peer := range cc.Peers {
		if name == peer.ID {
			return uint32(peer.GatewayPort)
		}
	}
	return 8080 // dummy value for unknown clusters
}
