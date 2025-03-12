package objectsmispformat

func (umisp UsersMispFormat) GetUsersMisp() UsersMispFormat {
	return umisp
}

func (umisp *UsersMispFormat) SetValueOrgIdUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.OrgId = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueServerIdUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.ServerId = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueEmailUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.Email = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueAuthkeyUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.Authkey = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueInvitedByUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.InvitedBy = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueGpgkeyUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.Gpgkey = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueCertifPublicUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.CertifPublic = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueNidsSidUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.NidsSid = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueNewsreadUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.Newsread = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueRoleIdUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.RoleId = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueChangePwUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.ChangePw = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueExpirationUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.Expiration = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueCurrentLoginUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.CurrentLogin = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueLastLoginUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.LastLogin = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueDateCreatedUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.DateCreated = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueDateModifiedUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(string); ok {
		umisp.DateModified = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueAutoalertUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		umisp.Autoalert = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueTermsacceptedUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		umisp.Termsaccepted = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueContactalertUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		umisp.Contactalert = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueDisabledUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		umisp.Disabled = data

		isSuccess = true
	}

	return isSuccess
}

func (umisp *UsersMispFormat) SetValueForceLogoutUsersMisp(v any) bool {
	var isSuccess bool

	if data, ok := v.(bool); ok {
		umisp.ForceLogout = data

		isSuccess = true
	}

	return isSuccess
}
