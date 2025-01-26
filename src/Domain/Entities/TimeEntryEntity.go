package Entities

import (
	"fmt"
	"time"
)

type TimeEntryEntity struct {
	ProjectId   int       `json:"project_id"`
	Start       time.Time `json:"start"`
	duration    time.Duration
	Description string `json:"description"`
}

func (this *TimeEntryEntity) IsFromProject(projectId int) bool {
	return this.ProjectId == projectId
}

func (this *TimeEntryEntity) getDuration() string {
	this.duration.Round(time.Second)
	hour := this.duration / time.Hour
	this.duration -= hour * time.Hour
	minute := this.duration / time.Minute
	this.duration -= minute * time.Minute
	second := this.duration / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)

}

func (this *TimeEntryEntity) IsRunning() bool {
	return false
}
