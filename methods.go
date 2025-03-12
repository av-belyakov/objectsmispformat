package objectsmispformat

import "slices"

func NewListFormatsMISP() *ListFormatsMISP {
	return &ListFormatsMISP{
		Event:      NewEventMisp(),
		Reports:    NewEventReports(),
		Attributes: []*AttributesMispFormat(nil),
		Objects:    map[int]*ObjectsMispFormat(nil),
		ObjectTags: &ListEventObjectTags{},
	}
}

// GetID возвращает уникальный идентификатор
func (o *ListFormatsMISP) GetID() string {
	return o.ID
}

// GetEvent возвращает объект Event
func (o *ListFormatsMISP) GetEvent() *EventsMispFormat {
	return o.Event
}

// GetReports возвращает объект Reports
func (o *ListFormatsMISP) GetReports() *EventReports {
	return o.Reports
}

// GetAttributes возвращает объект Attributes
func (o *ListFormatsMISP) GetAttributes() []*AttributesMispFormat {
	return o.Attributes
}

// GetObjects возвращает объект Objects
func (o *ListFormatsMISP) GetObjects() map[int]*ObjectsMispFormat {
	return o.Objects
}

// GetObjectTags возвращает объект ObjectTags
func (o *ListFormatsMISP) GetObjectTags() *ListEventObjectTags {
	return o.ObjectTags
}

// ComparisonID выполняет сравнение уникальных идентификаторов
func (o *ListFormatsMISP) ComparisonID(v string) bool {
	return o.ID == v
}

// ComparisonEvent выполняет сравнение свойств объекта Event
func (o *ListFormatsMISP) ComparisonEvent(v *EventsMispFormat) bool {
	return o.Event.Comparison(v)
}

// MatchingAndReplacementEvents выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementEvents(v EventsMispFormat) {
	newEvent := o.Event.MatchingAndReplacement(v)
	o.Event = &newEvent
}

// ComparisonReports выполняет сравнение свойств объекта Reports
func (o *ListFormatsMISP) ComparisonReports(v *EventReports) bool {
	return o.Reports.Comparison(v)
}

// MatchingAndReplacementReport выполняет сопоставление элементов объекта и их замену
func (o *ListFormatsMISP) MatchingAndReplacementReport(v EventReports) {
	newEventReport := o.Reports.MatchingAndReplacement(v)
	o.Reports = &newEventReport
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
func (o *ListFormatsMISP) MatchingAndReplacementAttributes(v []*AttributesMispFormat) {
	if len(o.Attributes) != len(v) {
		return
	}

	currentAttributes := o.GetAttributes()
	for i := 0; i < len(currentAttributes); i++ {
		for j := 0; j < len(v); j++ {
			if currentAttributes[i].ObjectId == v[j].ObjectId {
				newAttr := currentAttributes[i].MatchingAndReplacement(*v[j])
				currentAttributes[i] = &newAttr
			}
		}
	}
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
	for _, currentObject := range o.GetObjects() {
		for key, addedObject := range v {
			if currentObject.ID == addedObject.ID {
				newAddObject := currentObject.MatchingAndReplacement(*addedObject)
				v[key] = &newAddObject
			}
		}
	}

	return v
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
