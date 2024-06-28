package gozbx

type CommonGetParams struct {
	CountOutput   bool           `json:"countOutput,omitempty"`   // Return the number of records in the result instead of the actual data
	Editable      bool           `json:"editable,omitempty"`      // If set to true return only objects that the user has write permissions to.
	ExcludeSearch bool           `json:"excludeSearch,omitempty"` // Return results that do not match the criteria given in the search parameter.
	Filter        map[string]any `json:"filter,omitempty"`        // Filter the result.
	Limit         int            `json:"limit,omitempty"`         // Limit the number of records returned.
	Output        []string       `json:"output,omitempty"`        // Return only the given fields in the result.
	PreserveKeys  bool           `json:"preservekeys,omitempty"`  // Use IDs as keys in the resulting array.
	Search        []string       `json:"search,omitempty"`        // Search for records that match the given criteria.
	SearchByAny   bool           `json:"searchByAny,omitempty"`   // Search for records that match any of the given criteria.
	Sort          []string       `json:"sort,omitempty"`          // Sort the result.
}

type GroupId struct {
	GroupId string `json:"groupid"`
}

type HostId struct {
	HostId string `json:"hostid"`
}
