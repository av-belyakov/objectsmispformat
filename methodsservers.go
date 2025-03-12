package objectsmispformat

func (smisp ServersMispFormat) GetServersMisp() ServersMispFormat {
	return smisp
}

func (smisp *ServersMispFormat) SetValueNameServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Name = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueUrlServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Url = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueAuthkeyServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Authkey = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueOrgIdServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.OrgId = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueLastpulledidServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Lastpulledid = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueLastpushedidServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Lastpushedid = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueOrganizationServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Organization = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueRemoteOrgIdServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.RemoteOrgId = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePullRulesServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.PullRules = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePushRulesServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.PushRules = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueCertFileServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.CertFile = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueClientCertFileServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.ClientCertFile = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePriorityServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		smisp.Priority = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePushServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.Push = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePullServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.Pull = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePushSightingsServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.PushSightings = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePushGalaxyClustersServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.PushGalaxyClusters = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePullGalaxyClustersServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.PullGalaxyClusters = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValuePublishWithoutEmailServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.PublishWithoutEmail = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueUnpublishEventServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.UnpublishEvent = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueSelfSignedServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.SelfSigned = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueInternalServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.Internal = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueSkipProxyServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.SkipProxy = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueCachingEnabledServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.CachingEnabled = data

		isSuccess = true
	}

	return isSuccess
}

func (smisp *ServersMispFormat) SetValueCacheTimestampServersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		smisp.CacheTimestamp = data

		isSuccess = true
	}

	return isSuccess
}
