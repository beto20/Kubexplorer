package knowledge

type SourceObject string

const (
	PODS        SourceObject = "Pods"
	SERVICES    SourceObject = "Services"
	DEPLOYMENTS SourceObject = "Deployments"
	SECRETS     SourceObject = "Secrets"
)
