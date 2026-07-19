package company

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name             string    `gorm:"not null" json:"name"`
	SiteURL          *string   `json:"site_url,omitempty"`
	Country          *string   `json:"country,omitempty"`
	City             *string   `json:"city,omitempty"`
	JobAggregatorURL *string   `json:"job_aggregator_url,omitempty"`
	CreatedAt        time.Time `gorm:"not null;default:now()" json:"created_at"`
}

func (Company) TableName() string {
	return "companies"
}
