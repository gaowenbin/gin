package swagger

type ApiSchema struct {
	RefStr string `json:"$ref,omitempty"`
}

type ApiResponse struct {
	Description string    `json:"description"`
	Schema      ApiSchema `json:"schema,omitempty"`
}

type ApiMethod struct {
	Tags []string `json:"tags,omitempty"`
	// this field SHOULD be less than 120 characters.
	Summary     string                 `json:"summary,omitempty"`
	Description string                 `json:"description"`
	OperationId string                 `json:"operationId,omitempty"`
	Consumes    []string               `json:"consumes,omitempty"`
	Produces    []string               `json:"produces,omitempty"`
	Parameters  []InParam              `json:"parameters,omitempty"`
	Responses   map[string]ApiResponse `json:"responses,omitempty"`
	Deprecated  bool                   `json:"deprecated,omitempty"`
	Security    []ScrRequirement       `json:"security,omitempty"`
}
