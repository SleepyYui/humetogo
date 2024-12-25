package api

type ToolsList struct {
	PageNumber int64 `json:"page_number"`
	PageSize   int64 `json:"page_size"`
	TotalPages int64 `json:"total_pages"`
	ToolsPage  []Tool
}

type Tool struct {
	ToolType           string         `json:"tool_type"`
	ID                 string         `json:"id"`
	Version            int64          `json:"version"`
	VersionType        string         `json:"version_type"`
	Name               string         `json:"name"`
	CreatedOn          int64          `json:"created_on"`
	ModifiedOn         int64          `json:"modified_on"`
	Parameters         ToolParameters `json:"parameters"`
	VersionDescription string         `json:"version_description"`
	FallbackContent    string         `json:"fallback_content"`
	Description        string         `json:"description"`
}

type ToolParameters struct {
	Type       string                   `json:"type"`
	Properties ToolParametersProperties `json:"properties"`
	Required   []string                 `json:"required"`
}

type ToolParametersProperties struct {
	Location ToolParametersPropertiesLocation `json:"location"`
	Format   ToolParametersPropertiesFormat   `json:"format"`
}

type ToolParametersPropertiesLocation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ToolParametersPropertiesFormat struct {
	Type        string   `json:"type"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
}
