package common

//InArgs :
type InArgs struct {
	DirPath    *string
	GitURL     *string
	GitBranch  *string
	CommitID   *string
	OutputFile *string
}

//Resource :
type Resource struct {
	GitURL    string `json:"git_url"`
	GitBranch string `json:"git_branch"`
	CommitID  string `json:"commitid"`

	Resource struct {
		Name         string            `json:"name, omitempty"`
		Namespace    string            `json:"namespace,omitempty"`
		Selectors    []string          `json:"selectors, omitempty"`
		Labels       map[string]string `json:"labels, omitempty"`
		FilePath     string            `json:"filepath, omitempty"`
		Kind         string            `json:"kind, omitempty"`
		ReplicaCount int               `json:"replica,omitempty"`
		Image        struct {
			ID string `json:"id, omitempty"`
		} `json:"image"`
		Network []NetworkAttr `json:"network"`
		Envs    []string
	} `json:"resource,  omitempty`
}

//NetworkAttr :
type NetworkAttr struct {
	HostPort      int    `json:"host_port,omitempty"`
	ContainerPort int    `json:"container_url,omitempty"`
	Protocol      string `json:"protocol,omitempty "`
}

//SvcNetworkAttr :
type SvcNetworkAttr struct {
	Port       int    `json:"port,omitempty"`
	TargetPort int    `json:"target_port,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
}

//Service :
type Service struct {
	GitURL    string `json:"git_url"`
	GitBranch string `json:"git_branch"`
	CommitID  string `json:"commitid"`
	Resource  struct {
		Name      string   `json:"name, omitempty"`
		Namespace string   `json:"namespace, omitempty"`
		Selectors []string `json:"selectors, omitempty"`
		// Labels    map[string]string `json:"labels, omitempty"`
		Type     string           `json:"type, omitempty"`
		FilePath string           `json:"filepath, omitempty"`
		Kind     string           `json:"kind, omitempty"`
		Network  []SvcNetworkAttr `json:"network, omitempty"`
	} `json:"resource, omitempty"`
}

//Connections :
type Connections struct {
	Source Resource `json:"source, omitempty"`
	Target Resource `json:"target"`
	Link   Service  `json:"link"`
}

const (
	//ServiceCtx :
	ServiceCtx = "service"

	//DeployCtx :
	DeployCtx = "deployment"
)
