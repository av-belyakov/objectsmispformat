package objectsmispformat

import (
	"slices"
)

func NewListFormatsMISP() *ListFormatsMISP {
	return &ListFormatsMISP{
		Event:      NewEventMisp(),
		Reports:    NewEventReports(),
		Attributes: []*AttributesMispFormat(nil),
		Objects:    map[int]*ObjectsMispFormat(nil),
		ObjectTags: &ListEventObjectTags{},
	}
}

// GetID уникальный идентификатор
func (o *ListFormatsMISP) GetID() string {
	return o.ID
}

// SetID уникальный идентификатор
func (o *ListFormatsMISP) SetID(v string) {
	o.ID = v
}

// ComparisonID выполняет сравнение уникальных идентификаторов
func (o *ListFormatsMISP) ComparisonID(v string) bool {
	return o.ID == v
}

// GetEvent объект Event
func (o *ListFormatsMISP) GetEvent() *EventsMispFormat {
	return o.Event
}

// SetEvent объект Event
func (o *ListFormatsMISP) SetEvent(v EventsMispFormat) {
	o.Event = &v
}

// ComparisonEvent выполняет сравнение свойств объекта Event
func (o *ListFormatsMISP) ComparisonEvent(v *EventsMispFormat) bool {
	return o.Event.Comparison(v)
}

// MatchingAndReplacementEvents выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementEvents(v EventsMispFormat) EventsMispFormat {
	return v.MatchingAndReplacement(*o.Event)
}

// GetReports объект Reports
func (o *ListFormatsMISP) GetReports() *EventReports {
	return o.Reports
}

// SetReports объект Reports
func (o *ListFormatsMISP) SetReports(v EventReports) {
	o.Reports = &v
}

// ComparisonReports выполняет сравнение свойств объекта Reports
func (o *ListFormatsMISP) ComparisonReports(v *EventReports) bool {
	return o.Reports.Comparison(v)
}

// MatchingAndReplacementReport выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementReport(v EventReports) EventReports {
	return v.MatchingAndReplacement(*o.Reports)
}

// GetAttributes объект Attributes
func (o *ListFormatsMISP) GetAttributes() []*AttributesMispFormat {
	return o.Attributes
}

// GetAttributes объект Attributes
func (o *ListFormatsMISP) SetAttributes(v []*AttributesMispFormat) {
	o.Attributes = v
}

// ComparisonAttributes выполняет сравнение свойств объекта Attributes
func (o *ListFormatsMISP) ComparisonAttributes(v []*AttributesMispFormat) bool {
	if len(o.Attributes) != len(v) {
		return false
	}

	for _, currentAttribute := range o.GetAttributes() {
		var isExist bool
		for _, addedAttribute := range v {
			if currentAttribute.ObjectId == addedAttribute.ObjectId {
				isExist = true

				if !currentAttribute.Comparison(addedAttribute) {
					return false
				}
			}
		}

		if !isExist {
			return false
		}
	}

	return true
}

// MatchingAndReplacementAttributes выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementAttributes(v []*AttributesMispFormat) []*AttributesMispFormat {
	currentAttributes := o.GetAttributes()
	for i := range v {
		for j := range currentAttributes {
			if v[i].ObjectId == currentAttributes[j].ObjectId {
				newAttr := currentAttributes[i].MatchingAndReplacement(*v[j])
				v[i] = &newAttr
			}
		}
	}

	return v
}

// GetObjects объект Objects
func (o *ListFormatsMISP) GetObjects() map[int]*ObjectsMispFormat {
	return o.Objects
}

// SetObjects объект Objects
func (o *ListFormatsMISP) SetObjects(v map[int]*ObjectsMispFormat) {
	o.Objects = v
}

// ComparisonObjects выполняет сравнение свойств объекта Objects
func (o *ListFormatsMISP) ComparisonObjects(v map[int]*ObjectsMispFormat) bool {
	if len(o.Objects) != len(v) {
		return false
	}

	for _, currentObject := range o.GetObjects() {
		var isExist bool
		for _, addedObject := range v {
			if currentObject.ID == addedObject.ID {
				isExist = true

				if !currentObject.Comparison(addedObject) {
					return false
				}
			}
		}

		if !isExist {
			return false
		}
	}

	return true
}

// MatchingAndReplacementObjects выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementObjects(v map[int]*ObjectsMispFormat) map[int]*ObjectsMispFormat {
	for key, obj := range v {
		for _, currentObject := range o.GetObjects() {
			if currentObject.ID == obj.ID {
				newAddObject := obj.MatchingAndReplacement(*currentObject)
				v[key] = &newAddObject
			}
		}
	}

	return v
}

// GetObjectTags объект ObjectTags
func (o *ListFormatsMISP) GetObjectTags() *ListEventObjectTags {
	return o.ObjectTags
}

// SetObjectTags объект ObjectTags
func (o *ListFormatsMISP) SetObjectTags(v ListEventObjectTags) {
	o.ObjectTags = &v
}

// ComparisonObjectTags выполняет сравнение свойств объекта ObjectTags
func (o *ListFormatsMISP) ComparisonObjectTags(v *ListEventObjectTags) bool {
	if len(*o.ObjectTags) != len(v.GetListTags()) {
		return false
	}

	for _, currentObject := range *o.ObjectTags {
		var isEqual bool
		for _, objectAdded := range v.GetListTags() {
			if currentObject == objectAdded {
				isEqual = true

				break
			}
		}

		if !isEqual {
			return false
		}
	}

	return true
}

// MatchingAndReplacementListEventObjectTags выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementListEventObjectTags(v ListEventObjectTags) ListEventObjectTags {
	for _, currentObject := range o.ObjectTags.GetListTags() {
		var isEqual bool
		if slices.Contains(v.GetListTags(), currentObject) {
			isEqual = true
		}

		if !isEqual {
			v.SetTag(currentObject)
		}
	}

	return v
}
