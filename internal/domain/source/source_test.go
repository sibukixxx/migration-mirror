package source_test

import (
	"testing"

	"github.com/sibukixxx/migration-mirror/internal/domain/source"
	"github.com/sibukixxx/migration-mirror/internal/infra/wpcli"
)

// TestWPCLIAuditorImplementsAuditor is a compile-time interface check.
// If wpcli.Auditor stops satisfying source.Auditor, this test fails to build.
func TestWPCLIAuditorImplementsAuditor(t *testing.T) {
	var _ source.Auditor = (*wpcli.Auditor)(nil)
	var _ source.WarningReporter = (*wpcli.Auditor)(nil)
}
