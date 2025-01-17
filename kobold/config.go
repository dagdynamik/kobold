package kobold

// in memory representation of the config used to setup kobold
// this is used so that the user facing config can have different
// formats. I.e. v1 and v2 or a flat kubernetes configmap
type NormalizedConfig struct {
	// List of endpoints to listen on
	Endpoints []EndpointSpec `json:"endpoints,omitempty"`
	// list of repositories to use
	Repositories []RepositorySpec `json:"repositories,omitempty"`
	// subscriptions link repositories to one or more endpoints
	Subscriptions []SubscriptionSpec `json:"subscriptions,omitempty"`
	// the registry auth is only used when using the k8s chain
	// it allows to configure where/how imagePullSecrets are obtained
	// this is only needed for registries that require an extra api call
	// such as dockerhub, which does not send the digest
	RegistryAuth RegistryAuthSpec `json:"registryAuth,omitempty"`
	// the commit message is used when kobold makes commits
	// both title and description are parsed as template string and executed
	// with an array of changes as context
	CommitMessage CommitMessageSpec `json:"commitMessage,omitempty"`
}

type CommitMessageSpec struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type EndpointType string

const (
	EndpointTypeGeneric   EndpointType = "generic"
	EndpointTypeACR       EndpointType = "acr"
	EndpointTypeDockerhub EndpointType = "dockerhub"
)

type Header struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type EndpointSpec struct {
	Name            string       `json:"name,omitempty"`
	Type            EndpointType `json:"type,omitempty"`
	Path            string       `json:"path,omitempty"`
	RequiredHeaders []Header     `json:"requiredHeaders,omitempty"`
}

type RegistryAuthSpec struct {
	ServiceAccount   string               `json:"serviceAccount,omitempty"`
	Namespace        string               `json:"namespace,omitempty"`
	ImagePullSecrets []ImagePullSecretRef `json:"imagePullSecrets,omitempty"`
}

type ImagePullSecretRef struct {
	Name string `json:"name,omitempty"`
}

type GitProvider string

const (
	ProviderGithub GitProvider = "github"
	ProviderAzure  GitProvider = "azure"
)

type RepositorySpec struct {
	Name     string      `json:"name,omitempty"`
	URL      string      `json:"url,omitempty"`
	Provider GitProvider `json:"provider,omitempty"`
	Username string      `json:"username,omitempty"`
	Password string      `json:"password,omitempty"`
}

type Strategy string

const (
	StrategyCommit      Strategy = "commit"
	StrategyPullRequest Strategy = "pull-request"
)

type FileTypeKind string

const (
	FileTypeKubernetes FileTypeKind = "kubernetes"
	FileTypeCompose    FileTypeKind = "docker-compose"
	FileTypeKo         FileTypeKind = "ko-build"
)

type FileTypeSpec struct {
	Kind    FileTypeKind
	Pattern string
}

var DefaultAssociations = []FileTypeSpec{
	{Kind: FileTypeKo, Pattern: ".ko.yaml"},
	{Kind: FileTypeCompose, Pattern: "*compose*.y?ml"},
	{Kind: FileTypeKubernetes, Pattern: "*"},
}

type EndpointRef struct {
	Name string `json:"name,omitempty"`
}

type RepositoryRef struct {
	Name string `json:"name,omitempty"`
}

type SubscriptionSpec struct {
	Name             string         `json:"name,omitempty"`
	EndpointRefs     []EndpointRef  `json:"endpointRefs,omitempty"`
	RepositoryRef    RepositoryRef  `json:"repositoryRef,omitempty"`
	Branch           string         `json:"branch,omitempty"`
	Strategy         Strategy       `json:"strategy,omitempty"`
	Scopes           []string       `json:"scopes,omitempty"`
	FileAssociations []FileTypeSpec `json:"fileAssociations,omitempty"`
}
