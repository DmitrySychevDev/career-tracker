package vacancy

import (
	"time"

	"github.com/google/uuid"
)

type Vacancy struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CompanyID   uuid.UUID  `gorm:"type:uuid;not null" json:"company_id"`
	Title       string     `gorm:"not null" json:"title"`
	StartSalary *int       `json:"start_salary,omitempty"`
	EndSalary   *int       `json:"end_salary,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	WorkType    *string    `json:"work_type,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	Skills      *string    `json:"skills,omitempty"`
	PublicDate  *time.Time `json:"public_date,omitempty"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
}

func (*Vacancy) TableName() string {
	return "vacancies"
}
