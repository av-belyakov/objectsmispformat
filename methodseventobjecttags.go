package objectsmispformat

import "strings"

func NewListEventObjectTags() *ListEventObjectTags {
	return &ListEventObjectTags{}
}

func (leot *ListEventObjectTags) SetTag(v string) {
	listPattern := []string{
		"ats",
		"sensor",
		"misp-galaxy",
		"class-attack",
	}

	searchName := func(str string) bool {
		for _, value := range listPattern {
			if strings.Contains(str, value) {
				return true
			}
		}

		return false
	}

	if searchName(strings.ToLower(v)) {
		*leot = append(*leot, v)
	}
}

func (leot *ListEventObjectTags) GetListTags() ListEventObjectTags {
	return *leot
}

func (leot *ListEventObjectTags) CleanListTags() {
	leot = &ListEventObjectTags{}
}
