package objectsmispformat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/objectsmispformat"
)

func TestEventReport(t *testing.T) {
	mainEventReports := objectsmispformat.NewEventReports()
	mainEventReports.SetName("Green report")
	mainEventReports.SetContent("anybody content")
	mainEventReports.SetDistribution("some string")

	incomingEventReports := objectsmispformat.NewEventReports()
	incomingEventReports.SetName("Red report")
	incomingEventReports.SetContent("everybody content")
	incomingEventReports.SetDistribution("some distribution string")

	t.Run("Test comparison", func(t *testing.T) {
		isEqual := incomingEventReports.Comparison(mainEventReports)
		assert.False(t, isEqual)

		isEqual = incomingEventReports.Comparison(incomingEventReports)
		assert.True(t, isEqual)
	})

	t.Run("Test matching and replacement", func(t *testing.T) {
		newMainEvent := incomingEventReports.MatchingAndReplacement(*mainEventReports)
		assert.Equal(t, newMainEvent.GetName(), "Red report")
		assert.Equal(t, newMainEvent.GetContent(), "everybody content")
		assert.Equal(t, newMainEvent.GetDistribution(), "some distribution string")
	})
}
