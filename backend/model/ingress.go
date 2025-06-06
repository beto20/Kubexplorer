package model

type IngressDto struct {
	Name      string
	Namespace string
	Rules     []RuleDto
	Creation  string
	Labels    map[string]string
}

type RuleDto struct {
	Host string
	Path string
}
