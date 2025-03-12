package objectsmispformat

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/av-belyakov/objectsmispformat/internal/supportingfunctions"
)

func NewListAttributesMispFormat() *ListAttributesMispFormat {
	return &ListAttributesMispFormat{
		attributes: make(map[int]AttributesMispFormat)}
}

func NewAttributes() *AttributesMispFormat {
	return &AttributesMispFormat{}
}

func createNewAttributesMisp() AttributesMispFormat {
	return AttributesMispFormat{
		Category:       "Other",
		Type:           "other",
		Timestamp:      "0",
		Distribution:   getDistributionAttributes(),
		Uuid:           uuid.New().String(),
		FirstSeen:      fmt.Sprint(time.Now().Format(time.RFC3339)),
		LastSeen:       fmt.Sprint(time.Now().Format(time.RFC3339)),
		ToIds:          true,
		SharingGroupId: "1",
	}
}

func (lambda *ListAttributesMispFormat) GetCountList() int {
	return len(lambda.attributes)
}

func (lambda *ListAttributesMispFormat) CleanList() {
	lambda.mutex.Lock()
	defer lambda.mutex.Unlock()

	lambda.attributes = map[int]AttributesMispFormat{}
}

func (lambda *ListAttributesMispFormat) DelElementList(num int) (AttributesMispFormat, bool) {
	lambda.mutex.Lock()
	defer lambda.mutex.Unlock()

	var (
		ok   bool
		attr AttributesMispFormat
	)

	if attr, ok = lambda.attributes[num]; ok {
		delete(lambda.attributes, num)
	}

	return attr, ok
}

func (lamisp *ListAttributesMispFormat) GetList() map[int]AttributesMispFormat {
	return lamisp.attributes
}

// Comparison выполняет сравнение двух объектов типа AttributesMispFormat
func (a *AttributesMispFormat) Comparison(newAttributes *AttributesMispFormat) bool {
	//Uuid `json:"uuid"` не стал так как для каждого объекта с помощью
	//конструктора автоматически формируется свой идентификатор
	//
	//Timestamp `json:"timestamp"`
	//FirstSeen `json:"first_seen"`
	//LastSeen `json:"last_seen"`
	//а с временем вообще не ясно, что и когда может поменять TheHive

	if a.ToIds != newAttributes.ToIds {
		return false
	}

	if a.Deleted != newAttributes.Deleted {
		return false
	}

	if a.DisableCorrelation != newAttributes.DisableCorrelation {
		return false
	}

	if a.EventId != newAttributes.EventId {
		return false
	}

	if a.ObjectId != newAttributes.ObjectId {
		return false
	}

	if a.ObjectRelation != newAttributes.ObjectRelation {
		return false
	}

	if a.Category != newAttributes.Category {
		return false
	}

	if a.Type != newAttributes.Type {
		return false
	}

	if a.Value != newAttributes.Value {
		return false
	}

	if a.Distribution != newAttributes.Distribution {
		return false
	}
	if a.SharingGroupId != newAttributes.SharingGroupId {
		return false
	}
	if a.Comment != newAttributes.Comment {
		return false
	}

	return true
}

// MatchingAndReplacement выполняет сопоставление элементов объекта и их замену
func (a *AttributesMispFormat) MatchingAndReplacement(newAttributes AttributesMispFormat) AttributesMispFormat {
	if a.ToIds && a.ToIds != newAttributes.ToIds {
		newAttributes.ToIds = a.ToIds
	}

	if a.Deleted && a.Deleted != newAttributes.Deleted {
		newAttributes.Deleted = a.Deleted
	}

	if a.DisableCorrelation && a.DisableCorrelation != newAttributes.DisableCorrelation {
		newAttributes.DisableCorrelation = a.DisableCorrelation
	}

	if a.EventId != "" && a.EventId != newAttributes.EventId {
		newAttributes.EventId = a.EventId
	}

	if a.ObjectId != "" && a.ObjectId != newAttributes.ObjectId {
		newAttributes.ObjectId = a.ObjectId
	}

	if a.ObjectRelation != "" && a.ObjectRelation != newAttributes.ObjectRelation {
		newAttributes.ObjectRelation = a.ObjectRelation
	}

	if a.Category != "" && a.Category != newAttributes.Category {
		newAttributes.Category = a.Category
	}

	if a.Type != "" && a.Type != newAttributes.Type {
		newAttributes.Type = a.Type
	}

	if a.Value != "" && a.Value != newAttributes.Value {
		newAttributes.Value = a.Value
	}

	if a.Distribution != "" && a.Distribution != newAttributes.Distribution {
		newAttributes.Distribution = a.Distribution
	}

	if a.SharingGroupId != "" && a.SharingGroupId != newAttributes.SharingGroupId {
		newAttributes.SharingGroupId = a.SharingGroupId
	}

	if a.Comment != "" && a.Comment != newAttributes.Comment {
		newAttributes.Comment = a.Comment
	}

	return newAttributes
}

func (lamisp *ListAttributesMispFormat) SetValueObjectId(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.ObjectId = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueObjectRelation(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.ObjectRelation = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueCategory(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.Category = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) AutoSetValueCategory(v string, num int) {
	networkActivityCategory := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueCategory("Network activity", num)
	}
	payloadDeliveryCategory := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueCategory("Payload delivery", num)
	}

	collection := map[string]func(lamisp *ListAttributesMispFormat, num int){
		"snort_sid": networkActivityCategory,
		"url":       payloadDeliveryCategory,
		"domain":    networkActivityCategory,
		"md5":       payloadDeliveryCategory,
		"sha1":      payloadDeliveryCategory,
		"sha224":    payloadDeliveryCategory,
		"sha256":    payloadDeliveryCategory,
		"sha512":    payloadDeliveryCategory,
		"filename":  payloadDeliveryCategory,
		"ja3":       payloadDeliveryCategory,
		"ip_home":   networkActivityCategory,
	}

	if f, ok := collection[v]; ok {
		f(lamisp, num)
	}
}

func (lamisp *ListAttributesMispFormat) SetValueType(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.Type = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) AutoSetValueType(v string, num int) {
	snortSidType := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("snort", num)
	}
	urlType := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("url", num)
	}
	domainType := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("domain", num)
	}
	md5Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("md5", num)
	}
	sha1Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("sha1", num)
	}
	sha224Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("sha224", num)
	}
	sha256Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("sha256", num)
	}
	sha512Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("sha512", num)
	}
	filenameType := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("filename", num)
	}
	ja3Type := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("ja3-fingerprint-md5", num)
	}
	ipHome := func(lamisp *ListAttributesMispFormat, num int) {
		lamisp.SetValueType("other", num)
	}

	collection := map[string]func(lamisp *ListAttributesMispFormat, num int){
		"snort_sid": snortSidType,
		"url":       urlType,
		"domain":    domainType,
		"md5":       md5Type,
		"sha1":      sha1Type,
		"sha224":    sha224Type,
		"sha256":    sha256Type,
		"sha512":    sha512Type,
		"filename":  filenameType,
		"ja3":       ja3Type,
		"ip_home":   ipHome,
	}

	if f, ok := collection[v]; ok {
		f(lamisp, num)
	}

	//это для определения типа хеша
	if v == "hash" {
		if hashName, _, err := supportingfunctions.CheckStringHash(v); err == nil {
			lamisp.SetValueType(hashName, num)
		}
	}
}

func (lamisp *ListAttributesMispFormat) SetValueValue(v any, num int) {
	lamisp.mutex.Lock()

	value := fmt.Sprint(v)
	tmp := lamisp.getAttributesMisp(num)
	tmp.Value = fmt.Sprint(v)
	lamisp.attributes[num] = tmp

	//надо разблокировать Mutex до того как использовать lamisp.AutoSetValueCategoryAttributesMisp и
	//AutoSetValueType так как эти два метода используют методы AutoSetValueCategory и
	//AutoSetValueType вызывающие повторную блокировку Mutex
	lamisp.mutex.Unlock()

	//дополнительно, если значение подподает под рег. выражение типа "8030073:193.29.19.55"
	//то устанавливаем дополнительное значение типа в поле "object_relation"
	patter := regexp.MustCompile(`^[\d]+:((25[0-5]|2[0-4]\d|[01]?\d\d?)[.]){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	if patter.MatchString(value) {
		//выполняем автоматическое изменение значения свойства Category
		lamisp.AutoSetValueCategory("ip_home", num)

		//выполняем автоматическое изменение значения свойства Type
		lamisp.AutoSetValueType("ip_home", num)

		np := regexp.MustCompile(`^([\d]+):([\d]+\.[\d]+\.[\d]+\.[\d]+)$`)
		result := np.FindAllStringSubmatch(value, -1)
		if len(result) > 0 && len(result[0]) == 3 {
			lamisp.SetValueValue(result[0][2], num)
		}
	}
}

func (lamisp *ListAttributesMispFormat) SetValueUuid(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.Uuid = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueTimestamp(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if dt, ok := v.(float64); ok {
		tmp.Timestamp = fmt.Sprintf("%10.f", dt)[:10]
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueDistribution(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.Distribution = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueSharingGroupId(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.SharingGroupId = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueComment(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.Comment = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueFirstSeen(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if dt, ok := v.(float64); ok {
		tmp.FirstSeen = time.UnixMilli(int64(dt)).Format(time.RFC3339)
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueLastSeen(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if dt, ok := v.(float64); ok {
		tmp.LastSeen = time.UnixMilli(int64(dt)).Format(time.RFC3339)
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueToIds(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if data, ok := v.(bool); ok {
		tmp.ToIds = data
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueDeleted(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if data, ok := v.(bool); ok {
		tmp.Deleted = data
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) SetValueDisableCorrelation(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	if data, ok := v.(bool); ok {
		tmp.DisableCorrelation = data
	}

	lamisp.attributes[num] = tmp
}

func (lamisp *ListAttributesMispFormat) HandlingValueEventId(v any, num int) {
	lamisp.mutex.Lock()
	defer lamisp.mutex.Unlock()

	tmp := lamisp.getAttributesMisp(num)
	tmp.EventId = fmt.Sprint(v)
	lamisp.attributes[num] = tmp
}

// HandlingValueDataType изменяет некоторые поля объекта типа Attributes
// при этом на эти поля возможно вляиние других функций корректировщиков, например,
// функции getNewListAttributes, которая применяется для совмещения списков Attributes
// и Tags
func (lamisp *ListAttributesMispFormat) HandlingValueDataType(i any, num int) {
	v := fmt.Sprint(i)

	//выполняем автоматическое изменение значения свойства Category
	lamisp.AutoSetValueCategory(v, num)

	//выполняем автоматическое изменение значения свойства Type
	lamisp.AutoSetValueType(v, num)

	if v == "ip_home" {
		lamisp.SetValueObjectRelation(v, num)
	}
}

func (lamisp *ListAttributesMispFormat) getAttributesMisp(num int) AttributesMispFormat {
	var tmp AttributesMispFormat

	if attr, ok := lamisp.attributes[num]; ok {
		tmp = attr
	} else {
		tmp = createNewAttributesMisp()
	}

	return tmp
}

func HandlingListTags(l []string) [][2]string {
	nl := make([][2]string, 0, len(l))
	patter := regexp.MustCompile(`^misp:([\w\-].*)=\"([\w\-].*)\"$`)

	for _, v := range l {
		if !patter.MatchString(v) {
			continue
		}

		result := patter.FindAllStringSubmatch(v, -1)
		if len(result) > 0 && len(result[0]) == 3 {
			nl = append(nl, [2]string{result[0][1], result[0][2]})
		}
	}

	return nl
}

func getDistributionAttributes() string {
	return "2"
}
