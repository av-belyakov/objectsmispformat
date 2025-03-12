package objectsmispformat

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewListObjectsMispFormat() *ListObjectsMispFormat {
	return &ListObjectsMispFormat{objects: map[int]ObjectsMispFormat{}}
}

func NewObjectsMisp() *ObjectsMispFormat {
	return &ObjectsMispFormat{
		TemplateUUID:    uuid.NewString(),
		TemplateVersion: "1",
		FirstSeen:       fmt.Sprint(time.Now().UnixMicro()),
		Timestamp:       fmt.Sprint(time.Now().Unix()),
		MetaCategory:    "file",
		Distribution:    "5",
		Attribute:       []AttributeMispFormat{},
	}
}

// GetCountList количество значений в списке
func (lomf *ListObjectsMispFormat) GetCountList() int {
	return len(lomf.objects)
}

// CleanList очистить список
func (lomf *ListObjectsMispFormat) CleanList() {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	lomf.objects = map[int]ObjectsMispFormat{}
}

// GetList возвращает список объектов
func (lomf *ListObjectsMispFormat) GetList() map[int]ObjectsMispFormat {
	return lomf.objects
}

// Comparison выполняет сравнение двух объектов типа ObjectsMispFormat
func (o *ObjectsMispFormat) Comparison(newObjects *ObjectsMispFormat) bool {
	//TemplateUUID `json:"template_uuid"` не стал так как для каждого объекта
	//с помощью конструктора автоматически формируется свой идентификатор
	//
	//TemplateVersion string        `json:"template_version"`
	//FirstSeen       string        `json:"first_seen"`
	//Timestamp       string        `json:"timestamp"`
	//для этих объектов пока не стоит выполнять сравнение

	if o.Name != newObjects.Name {
		return false
	}

	if o.Description != newObjects.Description {
		return false
	}

	if o.Distribution != newObjects.Distribution {
		return false
	}

	if o.EventId != newObjects.EventId {
		return false
	}

	if o.MetaCategory != newObjects.MetaCategory {
		return false
	}

	if len(o.Attribute) != len(newObjects.Attribute) {
		return false
	}

	var isEqualList []bool
	for _, currentAttribute := range o.Attribute {
		var isEqual bool
		for _, newAttribute := range newObjects.Attribute {
			if currentAttribute.Value == newAttribute.Value {
				isEqual = true
			}
		}

		isEqualList = append(isEqualList, isEqual)
	}

	for _, v := range isEqualList {
		if !v {
			return false
		}
	}

	return true
}

// MatchingAndReplacement выполняет сопоставление элементов объекта и их замену
func (o *ObjectsMispFormat) MatchingAndReplacement(newObjects ObjectsMispFormat) ObjectsMispFormat {
	//FirstSeen       string        `json:"first_seen"`
	//Timestamp       string        `json:"timestamp"`
	//для этих объектов пока не стоит выполнять замену

	if o.Name != "" && o.Name != newObjects.Name {
		newObjects.Name = o.Name
	}

	if o.TemplateUUID != "" && o.TemplateUUID != newObjects.TemplateUUID {
		newObjects.TemplateUUID = o.TemplateUUID
	}

	if o.Description != "" && o.Description != newObjects.Description {
		newObjects.Description = o.Description
	}

	if o.Distribution != "" && o.Distribution != newObjects.Distribution {
		newObjects.Distribution = o.Distribution
	}

	if o.EventId != "" && o.EventId != newObjects.EventId {
		newObjects.EventId = o.EventId
	}

	if o.TemplateVersion != "" && o.TemplateVersion != newObjects.TemplateVersion {
		newObjects.TemplateVersion = o.TemplateVersion
	}

	if o.MetaCategory != "" && o.MetaCategory != newObjects.MetaCategory {
		newObjects.MetaCategory = o.MetaCategory
	}

	for _, currentAttribute := range o.Attribute {
		var isEqual bool
		for _, newAttribute := range newObjects.Attribute {
			if newAttribute.Value == currentAttribute.Value {
				isEqual = true

				break
			}
		}

		if !isEqual {
			newObjects.Attribute = append(newObjects.Attribute, currentAttribute)
		}
	}

	return newObjects
}

// SetValueId устанавливает значение ID
func (lomf *ListObjectsMispFormat) SetValueId(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	tmp.ID = fmt.Sprint(v)
	lomf.objects[num] = tmp
}

// SetValueEventId устанавливает значение EventId
func (lomf *ListObjectsMispFormat) SetValueEventId(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	tmp.EventId = fmt.Sprint(v)
	lomf.objects[num] = tmp
}

// SetValueName устанавливает значение Name
func (lomf *ListObjectsMispFormat) SetValueName(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	tmp.Name = fmt.Sprint(v)
	lomf.objects[num] = tmp
}

// SetValueDescription устанавливает значение Description
func (lomf *ListObjectsMispFormat) SetValueDescription(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	tmp.Description = fmt.Sprint(v)
	lomf.objects[num] = tmp
}

// SetValueFirstSeen устанавливает значение FirstSeen
func (lomf *ListObjectsMispFormat) SetValueFirstSeen(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	if dt, ok := v.(float64); ok {
		tmp.FirstSeen = time.UnixMilli(int64(dt)).Format(time.RFC3339)
	}

	lomf.objects[num] = tmp
}

// SetValueTimestamp устанавливает значение Timestamp
func (lomf *ListObjectsMispFormat) SetValueTimestamp(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	if dt, ok := v.(float64); ok {
		tmp.Timestamp = fmt.Sprintf("%10.f", dt)[:10]
	}

	lomf.objects[num] = tmp
}

// SetValueSize устанавливает значение Size
func (lomf *ListObjectsMispFormat) SetValueSize(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	tmp.Description = fmt.Sprintf("размер %v байт", v)
	lomf.objects[num] = tmp
}

// SetValueAttribute устанавливает значение Attribute
func (lomf *ListObjectsMispFormat) SetValueAttribute(v any, num int) {
	lomf.mutex.Lock()
	defer lomf.mutex.Unlock()

	tmp := lomf.getObjectMisp(num)
	if newSlice, ok := v.([]AttributeMispFormat); ok {
		tmp.Attribute = newSlice
		lomf.objects[num] = tmp
	}
}

func (lomf *ListObjectsMispFormat) getObjectMisp(num int) ObjectsMispFormat {
	var tmp ObjectsMispFormat

	if obj, ok := lomf.objects[num]; ok {
		tmp = obj
	} else {
		tmp = *NewObjectsMisp()
	}

	return tmp
}
