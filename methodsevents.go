package objectsmispformat

import (
	"fmt"

	"github.com/google/uuid"
)

func NewEventMisp() *EventsMispFormat {
	return &EventsMispFormat{
		Timestamp:         "0",
		Published:         getPublished(),
		PublishTimestamp:  "0",
		SightingTimestamp: "0",
		Uuid:              uuid.New().String(),
		Analysis:          getAnalysis(),
		Distribution:      getDistributionEvent(),
		ThreatLevelId:     getThreatLevelId(),
		SharingGroupId:    getSharingGroupId(),
	}
}

func (e *EventsMispFormat) CleanEventsMispFormat() {
	e.OrgId = ""
	e.OrgcId = ""
	e.Info = ""
	e.Uuid = uuid.New().String()
	e.Date = ""
	e.AttributeCount = ""
	e.ExtendsUuid = ""
	e.EventCreatorEmail = ""
	e.Published = getPublished()
	e.ProposalEmailLock = false
	e.Locked = false
	e.DisableCorrelation = false
	e.Analysis = getAnalysis()
	e.Distribution = getDistributionEvent()
	e.Timestamp = "0"
	e.ThreatLevelId = getThreatLevelId()
	e.PublishTimestamp = "0"
	e.SightingTimestamp = "0"
	e.SharingGroupId = getSharingGroupId()
}

// Comparison выполняет сравнение двух объектов типа EventsMispFormat
func (e *EventsMispFormat) Comparison(newEvents *EventsMispFormat) bool {
	if e.Analysis != newEvents.Analysis {
		return false
	}

	if e.Analysis != newEvents.Analysis {
		return false
	}

	if e.AttributeCount != newEvents.AttributeCount {
		return false
	}

	if e.OrgId != newEvents.OrgId {
		return false
	}

	if e.OrgcId != newEvents.OrgcId {
		return false
	}

	if e.Distribution != newEvents.Distribution {
		return false
	}

	if e.Info != newEvents.Info {
		return false
	}

	if e.Date != newEvents.Date {
		return false
	}

	if e.SharingGroupId != newEvents.SharingGroupId {
		return false
	}

	if e.ThreatLevelId != newEvents.ThreatLevelId {
		return false
	}

	if e.ExtendsUuid != newEvents.ExtendsUuid {
		return false
	}

	if e.EventCreatorEmail != newEvents.EventCreatorEmail {
		return false
	}

	if e.Published != newEvents.Published {
		return false
	}

	if e.ProposalEmailLock != newEvents.ProposalEmailLock {
		return false
	}

	if e.Locked != newEvents.Locked {
		return false
	}

	if e.DisableCorrelation != newEvents.DisableCorrelation {
		return false
	}

	// думаю время сравнивать не стоит, потому что большая вероятность получить идентичный
	//во всех параметрах объект у которого будет отличатся только время, что в данном случае не очень важно
	//Timestamp
	//PublishTimestamp
	//SightingTimestamp

	return true
}

// MatchingAndReplacement выполняет сопоставление элементов объекта и их замену
func (e *EventsMispFormat) MatchingAndReplacement(newEvents EventsMispFormat) EventsMispFormat {
	if e.Analysis != "" && e.Analysis != newEvents.Analysis {
		newEvents.Analysis = e.Analysis
	}

	if e.AttributeCount != "" && e.AttributeCount != newEvents.AttributeCount {
		newEvents.AttributeCount = e.AttributeCount
	}

	if e.OrgId != "" && e.OrgId != newEvents.OrgId {
		newEvents.OrgId = e.OrgId
	}

	if e.OrgcId != "" && e.OrgcId != newEvents.OrgcId {
		newEvents.OrgcId = e.OrgcId
	}

	if e.Distribution != "" && e.Distribution != newEvents.Distribution {
		newEvents.Distribution = e.Distribution
	}

	if e.Info != "" && e.Info != newEvents.Info {
		newEvents.Info = e.Info
	}

	if e.Uuid != "" && e.Uuid != newEvents.Uuid {
		newEvents.Uuid = e.Uuid
	}

	if e.Date != "" && e.Date != newEvents.Date {
		newEvents.Date = e.Date
	}

	if e.SharingGroupId != "" && e.SharingGroupId != newEvents.SharingGroupId {
		newEvents.SharingGroupId = e.SharingGroupId
	}

	if e.ThreatLevelId != "" && e.ThreatLevelId != newEvents.ThreatLevelId {
		newEvents.ThreatLevelId = e.ThreatLevelId
	}

	if e.ExtendsUuid != "" && e.ExtendsUuid != newEvents.ExtendsUuid {
		newEvents.ExtendsUuid = e.ExtendsUuid
	}

	if e.EventCreatorEmail != "" && e.EventCreatorEmail != newEvents.EventCreatorEmail {
		newEvents.EventCreatorEmail = e.EventCreatorEmail
	}

	if e.Published && e.Published != newEvents.Published {
		newEvents.Published = e.Published
	}

	if e.ProposalEmailLock && e.ProposalEmailLock != newEvents.ProposalEmailLock {
		newEvents.ProposalEmailLock = e.ProposalEmailLock
	}

	if e.Locked && e.Locked != newEvents.Locked {
		newEvents.Locked = e.Locked
	}

	if e.DisableCorrelation && e.DisableCorrelation != newEvents.DisableCorrelation {
		newEvents.DisableCorrelation = e.DisableCorrelation
	}

	return newEvents
}

// GetOrgId возвращает значение OrgId
func (e *EventsMispFormat) GetOrgId() string {
	return e.OrgId
}

// SetAnyOrgId устанавливает значение для OrgId
func (e *EventsMispFormat) SetAnyOrgId(v any, num int) {
	if data, ok := v.(string); ok {
		e.OrgId = data
	}
}

// SetOrgId возвращает значение OrgId
func (e *EventsMispFormat) SetOrgId(v string) {
	e.OrgId = v
}

// GetOrgcId возвращает значение OrgcId
func (e *EventsMispFormat) GetOrgcId() string {
	return e.OrgcId
}

// SetAnyOrgcId устанавливает значение для OrgcId
func (e *EventsMispFormat) SetAnyOrgcId(v any, num int) {
	if data, ok := v.(string); ok {
		e.OrgcId = data
	}
}

// SetOrgcId устанавливает значение для OrgcId
func (e *EventsMispFormat) SetOrgcId(v string) {
	e.OrgcId = v
}

// GetDistribution возвращает значение Distribution
func (e *EventsMispFormat) GetDistribution() string {
	return e.Distribution
}

// SetAnyDistribution устанавливает значение для Distribution
func (e *EventsMispFormat) SetAnyDistribution(v any, num int) {
	if data, ok := v.(string); ok {
		e.Distribution = data
	}
}

// SetDistribution устанавливает значение для Distribution
func (e *EventsMispFormat) SetDistribution(v string) {
	e.Distribution = v
}

// GetInfo возвращает значение Info
func (e *EventsMispFormat) GetInfo() string {
	return e.Info
}

// SetAnyInfo устанавливает значение для Info
func (e *EventsMispFormat) SetAnyInfo(v any, num int) {
	if data, ok := v.(string); ok {
		e.Info = data
	}
}

// SetInfo устанавливает значение для Info
func (e *EventsMispFormat) SetInfo(v string) {
	e.Info = v
}

// GetUUID возвращает значение UUID
func (e *EventsMispFormat) GetUUID() string {
	return e.Uuid
}

// SetAnyUUID устанавливает значение для UUID
func (e *EventsMispFormat) SetAnyUUID(v any, num int) {
	if data, ok := v.(string); ok {
		e.Uuid = data
	}
}

// SetUUID устанавливает значение для UUID
func (e *EventsMispFormat) SetUUID(v string) {
	e.Uuid = v
}

// GetDate возвращает значение Date
func (e *EventsMispFormat) GetDate() string {
	return e.Date
}

// SetAnyDate устанавливает значение для Date
func (e *EventsMispFormat) SetAnyDate(v any, num int) {
	if data, ok := v.(string); ok {
		e.Date = data
	}
}

// SetDate устанавливает значение для Date
func (e *EventsMispFormat) SetDate(v string) {
	e.Date = v
}

// GetAnalysis возвращает значение Analysis
func (e *EventsMispFormat) GetAnalysis() string {
	return e.Analysis
}

// SetAnyAnalysis устанавливает значение для Analysis
func (e *EventsMispFormat) SetAnyAnalysis(v any, num int) {
	if data, ok := v.(string); ok {
		e.Analysis = data
	}
}

// SetAnalysis устанавливает значение для Analysis
func (e *EventsMispFormat) SetAnalysis(v string) {
	e.Analysis = v
}

// GetAttributeCount возвращает значение AttributeCount
func (e *EventsMispFormat) GetAttributeCount() string {
	return e.AttributeCount
}

// SetAnyAttributeCount устанавливает значение для AttributeCount
func (e *EventsMispFormat) SetAnyAttributeCount(v any, num int) {
	if data, ok := v.(string); ok {
		e.AttributeCount = data
	}
}

// SetAttributeCount устанавливает значение для AttributeCount
func (e *EventsMispFormat) SetAttributeCount(v string) {
	e.AttributeCount = v
}

// GetTimestamp возвращает значение Timestamp
func (e *EventsMispFormat) GetTimestamp() string {
	return e.Timestamp
}

// SetAnyTimestamp устанавливает значение для Timestamp
func (e *EventsMispFormat) SetAnyTimestamp(v any, num int) {
	if data, ok := v.(float64); ok {
		//e.Timestamp = fmt.Sprintf("%13.f", data)
		e.Timestamp = fmt.Sprintf("%13.f", data)[:10]
	}
}

// SetTimestamp устанавливает значение для Timestamp
func (e *EventsMispFormat) SetTimestamp(v string) {
	e.Timestamp = v
}

// GetSharingGroupId возвращает значение SharingGroupId
func (e *EventsMispFormat) GetSharingGroupId() string {
	return e.SharingGroupId
}

// SetAnySharingGroupId устанавливает значение для SharingGroupId
func (e *EventsMispFormat) SetAnySharingGroupId(v any, num int) {
	if data, ok := v.(string); ok {
		e.SharingGroupId = data
	}
}

// SetSharingGroupId устанавливает значение для SharingGroupId
func (e *EventsMispFormat) SetSharingGroupId(v string) {
	e.SharingGroupId = v
}

// GetThreatLevelId возвращает значение ThreatLevelId
func (e *EventsMispFormat) GetThreatLevelId() string {
	return e.ThreatLevelId
}

// SetAnyThreatLevelId устанавливает значение для ThreatLevelId
func (e *EventsMispFormat) SetAnyThreatLevelId(v any, num int) {
	if data, ok := v.(string); ok {
		e.ThreatLevelId = data
	}

	if data, ok := v.(float64); ok {
		e.ThreatLevelId = fmt.Sprint(data)
	}
}

// SetThreatLevelId устанавливает значение для ThreatLevelId
func (e *EventsMispFormat) SetThreatLevelId(v string) {
	e.ThreatLevelId = v
}

// GetPublishTimestamp возвращает значение PublishTimestamp
func (e *EventsMispFormat) GetPublishTimestamp() string {
	return e.PublishTimestamp
}

// SetAnyPublishTimestamp устанавливает значение для PublishTimestamp
func (e *EventsMispFormat) SetAnyPublishTimestamp(v any, num int) {
	if data, ok := v.(float64); ok {
		//e.PublishTimestamp = fmt.Sprintf("%13.f", data)
		e.PublishTimestamp = fmt.Sprintf("%13.f", data)[:10]
	}
}

// SetPublishTimestamp устанавливает значение для PublishTimestamp
func (e *EventsMispFormat) SetPublishTimestamp(v string) {
	e.PublishTimestamp = v
}

// GetSightingTimestamp возвращает значение SightingTimestamp
func (e *EventsMispFormat) GetSightingTimestamp() string {
	return e.SightingTimestamp
}

// SetAnySightingTimestamp устанавливает значение для SightingTimestamp
func (e *EventsMispFormat) SetAnySightingTimestamp(v any, num int) {
	if data, ok := v.(float64); ok {
		//e.SightingTimestamp = fmt.Sprintf("%13.f", data)
		e.SightingTimestamp = fmt.Sprintf("%13.f", data)[:10]
	}
}

// SetSightingTimestamp устанавливает значение для SightingTimestamp
func (e *EventsMispFormat) SetSightingTimestamp(v string) {
	e.SightingTimestamp = v
}

// GetExtendsUuid возвращает значение ExtendsUuid
func (e *EventsMispFormat) GetExtendsUuid() string {
	return e.ExtendsUuid
}

// SetAnyExtendsUUID устанавливает значение для ExtendsUUID
func (e *EventsMispFormat) SetAnyExtendsUUID(v any, num int) {
	if data, ok := v.(string); ok {
		e.ExtendsUuid = data
	}
}

// SetExtendsUUID устанавливает значение для ExtendsUUID
func (e *EventsMispFormat) SetExtendsUUID(v string) {
	e.ExtendsUuid = v
}

// GetEventCreatorEmail возвращает значение EventCreatorEmail
func (e *EventsMispFormat) GetEventCreatorEmail() string {
	return e.EventCreatorEmail
}

// SetAnyEventCreatorEmail устанавливает значение для EventCreatorEmail
func (e *EventsMispFormat) SetAnyEventCreatorEmail(v any, num int) {
	if data, ok := v.(string); ok {
		e.EventCreatorEmail = data
	}
}

// SetEventCreatorEmail устанавливает значение для EventCreatorEmail
func (e *EventsMispFormat) SetEventCreatorEmail(v string) {
	e.EventCreatorEmail = v
}

// GetPublished возвращает значение Published
func (e *EventsMispFormat) GetPublished() bool {
	return e.Published
}

// SetAnyPublished устанавливает значение для Published
func (e *EventsMispFormat) SetAnyPublished(v any, num int) {
	if data, ok := v.(bool); ok {
		e.Published = data
	}
}

// SetPublished устанавливает значение для Published
func (e *EventsMispFormat) SetPublished(v bool) {
	e.Published = v
}

// GetProposalEmailLock возвращает значение ProposalEmailLock
func (e *EventsMispFormat) GetProposalEmailLock() bool {
	return e.ProposalEmailLock
}

// SetAnyProposalEmailLock устанавливает значение для ProposalEmailLock
func (e *EventsMispFormat) SetAnyProposalEmailLock(v any, num int) {
	if data, ok := v.(bool); ok {
		e.ProposalEmailLock = data
	}
}

// SetProposalEmailLock устанавливает значение для ProposalEmailLock
func (e *EventsMispFormat) SetProposalEmailLock(v bool) {
	e.ProposalEmailLock = v
}

// GetLocked возвращает значение Locked
func (e *EventsMispFormat) GetLocked() bool {
	return e.Locked
}

// SetAnyLocked устанавливает значение для Locked
func (e *EventsMispFormat) SetAnyLocked(v any, num int) {
	if data, ok := v.(bool); ok {
		e.Locked = data
	}
}

// SetLocked устанавливает значение для Locked
func (e *EventsMispFormat) SetLocked(v bool) {
	e.Locked = v
}

// GetDisableCorrelation возвращает значение DisableCorrelation
func (e *EventsMispFormat) GetDisableCorrelation() bool {
	return e.DisableCorrelation
}

// SetAnyDisableCorrelation устанавливает значение для DisableCorrelation
func (e *EventsMispFormat) SetAnyDisableCorrelation(v any, num int) {
	if data, ok := v.(bool); ok {
		e.DisableCorrelation = data
	}
}

// SetDisableCorrelation устанавливает значение для DisableCorrelation
func (e *EventsMispFormat) SetDisableCorrelation(v bool) {
	e.DisableCorrelation = v
}

func getAnalysis() string {
	return "2"
}

func getDistributionEvent() string {
	return "1"
}

func getThreatLevelId() string {
	return "4"
}

func getSharingGroupId() string {
	return "1"
}

func getPublished() bool {
	return false
}

/*func getTagTLP(tlp int) datamodels.TagsMispFormat {
	tag := datamodels.TagsMispFormat{Name: "tlp:red", Colour: "#cc0033"}

	switch tlp {
	case 0:
		tag.Name = "tlp:white"
		tag.Colour = "#ffffff"
	case 1:
		tag.Name = "tlp:green"
		tag.Colour = "#339900"
	case 2:
		tag.Name = "tlp:amber"
		tag.Colour = "#ffc000"
	}

	return tag
}*/
