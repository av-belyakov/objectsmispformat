package objectsmispformat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/objectsmispformat"
)

func TestAttributes(t *testing.T) {
	mainAttributes := objectsmispformat.NewAttributes()
	mainAttributes.SetEventId("te7d2")
	mainAttributes.SetObjectId("d1329-1s")
	mainAttributes.SetObjectRelation("some object relation")
	mainAttributes.SetCategory("category one")
	mainAttributes.SetType("some type")
	mainAttributes.SetValue("some value")
	mainAttributes.SetDistribution("new distr")
	mainAttributes.SetSharingGroupId("7899312")
	mainAttributes.SetComment("fast comment")
	mainAttributes.SetToIds(true)
	mainAttributes.SetDeleted(false)
	mainAttributes.SetDisableCorrelation(false)

	incomingAttributes := objectsmispformat.NewAttributes()
	incomingAttributes.SetEventId("gsq6344")
	incomingAttributes.SetObjectId("avb29-2s")
	incomingAttributes.SetObjectRelation("some object relation with code")
	incomingAttributes.SetCategory("category two")
	incomingAttributes.SetType("my some type")
	incomingAttributes.SetValue("my some value")
	incomingAttributes.SetDistribution("new distr")
	incomingAttributes.SetSharingGroupId("632848")
	incomingAttributes.SetComment("fast alfa comment")
	incomingAttributes.SetToIds(true)
	incomingAttributes.SetDeleted(true)
	incomingAttributes.SetDisableCorrelation(true)

	t.Run("Test comparison", func(t *testing.T) {
		isEqual := incomingAttributes.Comparison(mainAttributes)
		assert.False(t, isEqual)

		isEqual = incomingAttributes.Comparison(incomingAttributes)
		assert.True(t, isEqual)
	})

	t.Run("Test matching and replacement", func(t *testing.T) {
		newMainAttributes := incomingAttributes.MatchingAndReplacement(*mainAttributes)
		assert.Equal(t, newMainAttributes.EventId, "gsq6344")
		assert.Equal(t, newMainAttributes.ObjectId, "avb29-2s")
		assert.Equal(t, newMainAttributes.ObjectRelation, "some object relation with code")
		assert.Equal(t, newMainAttributes.Category, "category two")
		assert.Equal(t, newMainAttributes.Type, "my some type")
		assert.Equal(t, newMainAttributes.Value, "my some value")
		assert.Equal(t, newMainAttributes.Distribution, "new distr")
		assert.Equal(t, newMainAttributes.SharingGroupId, "632848")
		assert.Equal(t, newMainAttributes.Comment, "fast alfa comment")
		assert.True(t, newMainAttributes.ToIds)
		assert.True(t, newMainAttributes.Deleted)
		assert.True(t, newMainAttributes.DisableCorrelation)
	})
}
