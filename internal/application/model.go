package application

import (
	"time"

	"github.com/google/uuid"
)

type JobApplication struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID          uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	VacancyID       uuid.UUID `gorm:"type:uuid;not null" json:"vacancy_id"`
	ApplicationDate time.Time `gorm:"not null;default:now()" json:"application_date"`
	Status          string    `gorm:"not null" json:"status"`
	Note            *string   `json:"note,omitempty"`
	RecruiterName   *string   `json:"recruiter_name,omitempty"`
	RecruiterEmail  *string   `json:"recruiter_email,omitempty"`
	RecruiterTG     *string   `json:"recruiter_tg,omitempty"`
	CreatedAt       time.Time `gorm:"not null;default:now()" json:"created_at"`
}

func (*JobApplication) TableName() string {
	return "job_applications"
}
