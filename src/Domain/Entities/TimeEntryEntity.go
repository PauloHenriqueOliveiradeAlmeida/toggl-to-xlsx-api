package Entities

import (
	"fmt"
	"strings"
	"time"
)

type TimeEntryEntity struct {
	ProjectId   int       `json:"project_id"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Description string    `json:"description"`
}

func (this *TimeEntryEntity) IsFromProject(projectId int) bool {
	return this.ProjectId == projectId
}

func (this *TimeEntryEntity) GetDuration() time.Duration {
	return this.Stop.Sub(this.Start)
}

func (this *TimeEntryEntity) GetISODate() string {
	return strings.Split(this.Start.Format(time.RFC3339), "T")[0]
}

func (this *TimeEntryEntity) GetFormattedDuration() string {
	duration := this.GetDuration().Round(time.Second)
	hour := duration / time.Hour
	duration -= hour * time.Hour
	minute := duration / time.Minute
	duration -= minute * time.Minute
	second := duration / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)

}

func (this *TimeEntryEntity) IsRunning() bool {
	return false
}

func (this *TimeEntryEntity) ToMap() map[string]string {
	return map[string]string{
		"ProjectId":   fmt.Sprintf("%d", this.ProjectId),
		"Date":        this.GetISODate(),
		"Duration":    this.GetFormattedDuration(),
		"Description": this.Description,
	}
}
