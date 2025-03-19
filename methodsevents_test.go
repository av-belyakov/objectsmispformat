package objectsmispformat_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/objectsmispformat"
)

func TestEventMisp(t *testing.T) {
	mainEvent := objectsmispformat.NewEventMisp()
	mainEvent.SetOrgId("633fba874")
	mainEvent.SetOrgcId("fa6f432")
	mainEvent.SetDistribution("some string")
	mainEvent.SetInfo("some info")
	mainEvent.SetAnalysis("1")
	mainEvent.SetAttributeCount("2233")
	mainEvent.SetSharingGroupId("tw7-6e2-4we")
	mainEvent.SetThreatLevelId("1y-13-44")
	mainEvent.SetExtendsUUID(uuid.NewString())
	mainEvent.SetEventCreatorEmail("some.a@email.net")
	mainEvent.SetPublished(true)
	mainEvent.SetProposalEmailLock(true)
	mainEvent.SetLocked(false)
	mainEvent.SetDisableCorrelation(true)

	incomingEvent := objectsmispformat.NewEventMisp()
	incomingEvent.SetOrgId("633fba874")
	incomingEvent.SetOrgcId("a629df")
	incomingEvent.SetDistribution("some distribution string")
	incomingEvent.SetInfo("new some info")
	incomingEvent.SetAnalysis("2")
	incomingEvent.SetAttributeCount("--2233")
	incomingEvent.SetSharingGroupId("55e-p32-4we")
	incomingEvent.SetThreatLevelId("1y-13-44")
	incomingEvent.SetExtendsUUID(uuid.NewString())
	incomingEvent.SetEventCreatorEmail("some.b@email.net")
	incomingEvent.SetPublished(true)
	incomingEvent.SetProposalEmailLock(false)
	incomingEvent.SetLocked(true)
	incomingEvent.SetDisableCorrelation(true)

	t.Run("Test comparison", func(t *testing.T) {
		isEqual := incomingEvent.Comparison(mainEvent)
		assert.False(t, isEqual)

		isEqual = incomingEvent.Comparison(incomingEvent)
		assert.True(t, isEqual)
	})

	t.Run("Test matching and replacement", func(t *testing.T) {
		newMainEvent := incomingEvent.MatchingAndReplacement(*mainEvent)
		assert.Equal(t, newMainEvent.GetOrgId(), "633fba874")
		assert.Equal(t, newMainEvent.GetOrgcId(), "a629df")
		assert.Equal(t, newMainEvent.GetDistribution(), "some distribution string")
		assert.Equal(t, newMainEvent.GetInfo(), "new some info")
		assert.Equal(t, newMainEvent.GetAnalysis(), "2")
		assert.Equal(t, newMainEvent.GetAttributeCount(), "--2233")
		assert.Equal(t, newMainEvent.GetSharingGroupId(), "55e-p32-4we")
		assert.Equal(t, newMainEvent.GetEventCreatorEmail(), "some.b@email.net")
		//По свойствам типа bool, если входящий объект содержит false то замена не производится
		assert.True(t, newMainEvent.GetPublished())
		assert.True(t, newMainEvent.GetProposalEmailLock())
		assert.True(t, newMainEvent.Locked)
		assert.True(t, newMainEvent.GetDisableCorrelation())
	})
}
