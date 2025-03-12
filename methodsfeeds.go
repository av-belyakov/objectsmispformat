package objectsmispformat

func (fmisp FeedsMispFormat) GetFeedsMisp() FeedsMispFormat {
	return fmisp
}

func (fmisp *FeedsMispFormat) SetValueNameFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Name = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueProviderFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Provider = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueUrlFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Url = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueRulesFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Rules = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueDistributionFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Distribution = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueSharingGroupIdFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.SharingGroupId = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueTagIdFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.TagId = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueSourceFormatFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.SourceFormat = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueEventIdFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.EventId = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueInputSourceFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.InputSource = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueHeadersFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.Headers = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueOrgcIdFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		fmisp.OrgcId = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueEnabledFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.Enabled = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueFixedEventFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.FixedEvent = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueDeltaMergeFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.DeltaMerge = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValuePublishFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.Publish = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueOverrideIdsFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.OverrideIds = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueDeleteLocalFileFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.DeleteLocalFile = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueLookupVisibleFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.LookupVisible = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueCachingEnabledFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.CachingEnabled = data

		isSuccess = true
	}

	return isSuccess
}

func (fmisp *FeedsMispFormat) SetValueForceToIdsFeedsMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		fmisp.ForceToIds = data

		isSuccess = true
	}

	return isSuccess
}
