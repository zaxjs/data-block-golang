package data_block

type Model struct {
	ID                    string             `json:"$id,omitempty"`
	Type                  string             `json:"type,omitempty"`
	CanDrag               bool               `json:"canDrag,omitempty"`
	Omitted               bool               `json:"omitted,omitempty"`
	Children              []interface{}      `json:"children,omitempty"`
	Disabled              bool               `json:"disabled,omitempty"`
	Expanded              bool               `json:"expanded,omitempty"`
	FieldName             string             `json:"fieldName,omitempty"`
	FieldType             string             `json:"fieldType,omitempty"`
	FieldLabel            string             `json:"fieldLabel,omitempty"`
	FieldValues           FieldValues        `json:"fieldValues,omitempty"`
	FieldSettings         FieldSettings      `json:"fieldSettings,omitempty"`
	FieldDescription      string             `json:"fieldDescription,omitempty"`
	FieldValidations      []FieldValidations `json:"fieldValidations,omitempty"`
	FieldNamePlaceholder  string             `json:"fieldNamePlaceholder,omitempty"`
	FieldLabelPlaceholder string             `json:"fieldLabelPlaceholder,omitempty"`
}

type FieldSettings struct {
	Localizations []string `json:"localizations,omitempty"`
}

type FieldValidations struct {
	Max       int    `json:"max,omitempty"`
	Min       int    `json:"min,omitempty"`
	Type      string `json:"type,omitempty"`
	Required  bool   `json:"required,omitempty"`
	ScaleType string `json:"scaleType,omitempty"`
}

type FieldValues struct {
	DefaultValue   string `json:"defaultValue,omitempty"`
	DatasetValues  string `json:"datasetValues,omitempty"`
	PlaceholderTip string `json:"placeholderTip,omitempty"`
}
