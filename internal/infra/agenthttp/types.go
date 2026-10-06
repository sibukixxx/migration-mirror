package agenthttp

import (
	"github.com/sibukixxx/migration-mirror/internal/domain/audit"
	"github.com/sibukixxx/migration-mirror/internal/domain/source"
)

type auditResponse struct {
	Audit    audit.Audit      `json:"audit"`
	Warnings []source.Warning `json:"warnings,omitempty"`
}
