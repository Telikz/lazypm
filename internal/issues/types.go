package issues

import (
	"time"
)

type Issue struct {
	ID          int       `yaml:"id"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Status      string    `yaml:"status"`
	Labels      []string  `yaml:"labels,omitempty"`
	Assignees   []string  `yaml:"assignees,omitempty"`
	CreatedAt   time.Time `yaml:"created_at"`
	UpdatedAt   time.Time `yaml:"updated_at"`
}
