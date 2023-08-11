package conf

type Server struct {
	Listen         string `json:"listen" yaml:"listen"`
	IsUpload       bool   `json:"uploadSwitch" yaml:"uploadSwitch"`
	IsWildcard     bool   `json:"wildcardSwitch" yaml:"wildcardSwitch"`
	WildcardSuffix string `json:"wildcardSuffix" yaml:"wildcardSuffix"`
	SessionName    string `json:"sessionName" yaml:"sessionName"`
	StoreKeyPairs  string `json:"storeKeyPairs" yaml:"storeKeyPairs"`
}
