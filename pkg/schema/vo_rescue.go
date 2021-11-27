package schema

import (
	"time"

	"gorm.io/gorm"
)

type RescueItem struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name      string    `json:"name"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Province  string    `json:"provice"`
	City      string    `json:"city"`
	District  string    `json:"district"`
	Is24Hour  *bool     `json:"is_24_hour"`
}

type RescueListRequest struct {
	Pagination
	Keyword string `form:"keyword"`
}

func (m *RescueListRequest) Valadate() error {
	if err := m.Pagination.Valadate(); err != nil {
		return err
	}

	return nil
}

type RescueListResponse struct {
	Data    []*RescueItem `json:"data"`
	Success bool          `json:"success"`
}
