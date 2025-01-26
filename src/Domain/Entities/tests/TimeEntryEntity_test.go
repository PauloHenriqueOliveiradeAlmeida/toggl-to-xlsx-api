package Tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"toggl-xlsx-back/src/Domain/Entities"
)

func TestIfTimeEntryIsAProvidedProject(test *testing.T) {
	projectId := 1

	timeEntry := Entities.TimeEntryEntity{
		ProjectId: projectId,
	}

	isFromProject := timeEntry.IsFromProject(projectId)
	assert.Equal(test, true, isFromProject)
}
