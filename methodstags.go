package objectsmispformat

func (tmisp TagsMispFormat) GetTagsMisp() TagsMispFormat {
	return tmisp
}

func (tmisp *TagsMispFormat) SetValueNameTagsMisp(v any) bool {
	var isSuccess bool

	if str, ok := v.(string); ok {
		tmisp.Name = str

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueColourTagsMisp(v any) bool {
	var isSuccess bool

	if str, ok := v.(string); ok {
		tmisp.Colour = str

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueOrgIdTagsMisp(v any) bool {
	var isSuccess bool

	if str, ok := v.(string); ok {
		tmisp.OrgId = str

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueUserIdTagsMisp(v any) bool {
	var isSuccess bool

	if str, ok := v.(string); ok {
		tmisp.UserId = str

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueNumericalValueTagsMisp(v any) bool {
	var isSuccess bool

	if str, ok := v.(string); ok {
		tmisp.NumericalValue = str

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueInheritedTagsMisp(v any) bool {
	var isSuccess bool

	switch v := v.(type) {
	case int, int32, int64:
		if data, ok := v.(int); ok {
			tmisp.Inherited = data

			isSuccess = true
		}

	case uint, uint32, uint64:
		if data, ok := v.(uint); ok {
			tmisp.Inherited = int(data)

			isSuccess = true
		}

	case float32, float64:
		if data, ok := v.(float64); ok {
			tmisp.Inherited = int(data)

			isSuccess = true
		}
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueHideTagTagsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		tmisp.HideTag = data

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueIsGalaxyTagsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		tmisp.IsGalaxy = data

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueExportableTagsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		tmisp.Exportable = data

		isSuccess = true
	}

	return isSuccess
}

func (tmisp *TagsMispFormat) SetValueIsCustomGalaxyTagsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		tmisp.IsCustomGalaxy = data

		isSuccess = true
	}

	return isSuccess
}
