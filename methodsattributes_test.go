package objectsmispformat_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/objectsmispformat"
)

func TestAttributes(t *testing.T) {
	auuid := uuid.NewString()

	mainAttributes := objectsmispformat.NewAttributes()
	mainAttributes.EventId = "te7d2"
	mainAttributes.ObjectId = "d1329-1s"
	mainAttributes.ObjectRelation = "some object relation"
	mainAttributes.Category = "category one"
	mainAttributes.Type = "some type"
	mainAttributes.Value = "some value"
	mainAttributes.Uuid = auuid
	mainAttributes.Distribution = "new distr"
	mainAttributes.SharingGroupId = "7899312"
	mainAttributes.Comment = "fast comment"
	mainAttributes.ToIds = true
	mainAttributes.Deleted = false
	mainAttributes.DisableCorrelation = false

	incomingAttributes := objectsmispformat.NewAttributes()
	incomingAttributes.EventId = "gsq6344"
	incomingAttributes.ObjectId = "avb29-2s"
	incomingAttributes.ObjectRelation = "some object relation with code"
	incomingAttributes.Category = "category two"
	incomingAttributes.Type = "my some type"
	incomingAttributes.Value = "my some value"
	incomingAttributes.Uuid = auuid
	incomingAttributes.Distribution = "new distr"
	incomingAttributes.SharingGroupId = "632848"
	incomingAttributes.Comment = "fast alfa comment"
	incomingAttributes.ToIds = true
	incomingAttributes.Deleted = true
	incomingAttributes.DisableCorrelation = true

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
