package objectsmispformat_test

import (
	"testing"

	"github.com/av-belyakov/objectsmispformat"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestObjects(t *testing.T) {
	ouuid := uuid.NewString()

	mainObjectsMisp := objectsmispformat.NewObjectsMisp()
	mainObjectsMisp.ID = ouuid
	mainObjectsMisp.TemplateUUID = "r621-dw"
	mainObjectsMisp.TemplateVersion = "template_version"
	mainObjectsMisp.Name = "any name"
	mainObjectsMisp.Description = "my all description"
	mainObjectsMisp.EventId = "y72e-r83"
	mainObjectsMisp.MetaCategory = "my-meta-category"
	mainObjectsMisp.Distribution = "any distribution"

	incomingObjectsMisp := objectsmispformat.NewObjectsMisp()
	incomingObjectsMisp.ID = ouuid
	incomingObjectsMisp.TemplateUUID = "72r4-dw"
	incomingObjectsMisp.TemplateVersion = "template version 11.2"
	incomingObjectsMisp.Name = "every names"
	incomingObjectsMisp.Description = "old description"
	incomingObjectsMisp.EventId = "dw93-r3"
	incomingObjectsMisp.MetaCategory = "my meta category"
	incomingObjectsMisp.Distribution = "prod distribution"

	t.Run("Test comparison", func(t *testing.T) {
		isEqual := incomingObjectsMisp.Comparison(mainObjectsMisp)
		assert.False(t, isEqual)

		isEqual = incomingObjectsMisp.Comparison(incomingObjectsMisp)
		assert.True(t, isEqual)
	})

	t.Run("Test matching and replacement", func(t *testing.T) {
		newObjectsMisp := incomingObjectsMisp.MatchingAndReplacement(*mainObjectsMisp)
		assert.Equal(t, newObjectsMisp.TemplateUUID, "72r4-dw")
		assert.Equal(t, newObjectsMisp.TemplateVersion, "template version 11.2")
		assert.Equal(t, newObjectsMisp.Name, "every names")
		assert.Equal(t, newObjectsMisp.Description, "old description")
		assert.Equal(t, newObjectsMisp.EventId, "dw93-r3")
		assert.Equal(t, newObjectsMisp.MetaCategory, "my meta category")
		assert.Equal(t, newObjectsMisp.Distribution, "prod distribution")
	})
}
