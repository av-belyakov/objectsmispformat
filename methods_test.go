package objectsmispformat_test

import (
	"testing"

	"github.com/av-belyakov/objectsmispformat"
	"github.com/stretchr/testify/assert"
)

func TestComparisonAttributes(t *testing.T) {
	attributes := make([]*objectsmispformat.AttributesMispFormat, 0, 3)
	attributes = append(attributes, &objectsmispformat.AttributesMispFormat{
		EventId:            "",
		ObjectId:           "~88679125064",
		ObjectRelation:     "",
		Category:           "Other",
		Type:               "other",
		Value:              "45.3.6.66",
		Uuid:               "a69d30d6-8775-4a46-be97-b3e6d81c0c4c",
		Timestamp:          "1729076104",
		Distribution:       "2",
		SharingGroupId:     "1",
		Comment:            "simple ip address",
		FirstSeen:          "2024-10-16T13:55:04+03:00",
		LastSeen:           "2025-03-19T15:31:02+03:00",
		ToIds:              true,
		Deleted:            false,
		DisableCorrelation: true,
	})
	attributes = append(attributes, &objectsmispformat.AttributesMispFormat{
		EventId:            "",
		ObjectId:           "~89858891776",
		ObjectRelation:     "",
		Category:           "Other",
		Type:               "other",
		Value:              "",
		Uuid:               "5a1805ef-0103-40b9-a210-6152cf9e5aa5",
		Timestamp:          "1734000134",
		Distribution:       "2",
		SharingGroupId:     "1",
		Comment:            "картинка с тюленем",
		FirstSeen:          "2024-12-12T13:42:14+03:00",
		LastSeen:           "2025-03-19T15:31:02+03:00",
		ToIds:              true,
		Deleted:            false,
		DisableCorrelation: true,
	})

	attributes = append(attributes, &objectsmispformat.AttributesMispFormat{
		EventId:            "",
		ObjectId:           "~90607099992",
		ObjectRelation:     "",
		Category:           "Other",
		Type:               "other",
		Value:              "THEHIVE-SESSION=eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImF1dGhDb250ZXh0Ijoie1widXNlcklkXCI6XCJhLmJlbHlha292QGNsb3VkLmdjbVwiLFwidXNlck5hbWVcIjpcIkJlbHlha292ICpJUkUgQXJ0ZW1peVwiLFwib3JnYW5pc2F0aW9uXCI6XCJ-NDE5MlwiLFwicGVybWlzc2lvbnNcIjpbXCJtYW5hZ2VTaGFyZVwiLFwibWFuYWdlQW5hbHlzZVwiLFwibWFuYWdlVGFza1wiLFwibWFuYWdlQ2FzZVRlbXBsYXRlXCIsXCJtYW5hZ2VDYXNlXCIsXCJtYW5hZ2VVc2VyXCIsXCJtYW5hZ2VQcm9jZWR1cmVcIixcIm1hbmFnZVBhZ2VcIixcIm1hbmFnZU9ic2VydmFibGVcIixcIm1hbmFnZVRhZ1wiLFwibWFuYWdlQ29uZmlnXCIsXCJtYW5hZ2VBbGVydFwiLFwiYWNjZXNzVGhlSGl2ZUZTXCIsXCJtYW5hZ2VBY3Rpb25cIl19IiwiZXhwaXJlIjoiMTczMjUzMzYwOTI4NiIsIndhcm5pbmciOiIxNzMyNTMzMzA5Mjg2In0sIm5iZiI6MTczMjUzMDAwOSwiaWF0IjoxNzMyNTMwMDA5fQ.N_ZvqUAw0wt3AeSQmovk6NlFJ_UZkilxGvlh_UGnCX4",
		Uuid:               "e926b00e-1b9d-4948-b72d-5fbae024d0cd",
		Timestamp:          "1732530100",
		Distribution:       "2",
		SharingGroupId:     "1",
		Comment:            "efefe",
		FirstSeen:          "2024-11-25T13:21:40+03:00",
		LastSeen:           "2025-03-19T15:31:02+03:00",
		ToIds:              true,
		Deleted:            false,
		DisableCorrelation: true,
	})

	listFormatsMISP := objectsmispformat.NewListFormatsMISP()
	listFormatsMISP.SetAttributes(attributes)

	isEqual := listFormatsMISP.ComparisonAttributes(attributes)
	assert.True(t, isEqual)
}
