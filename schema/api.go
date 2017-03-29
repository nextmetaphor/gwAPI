package schema

type SingleAPI struct {
	APIDefinition SingleAPIDefinition `json:"api_definition"`
}

type SingleAPIDefinition struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	APIID string `json:"api_id"`
	OrgID string `json:"org_id"`
	Proxy struct {
		ListenPath string `json:"listen_path"`
		TargetURL  string `json:"target_url"`
	} `json:"proxy"`
}

type MultipleAPIDefinition struct {
	APIs []struct {
		APIDefinition SingleAPIDefinition `json:"api_definition"`
	} `json:"apis"`
}
