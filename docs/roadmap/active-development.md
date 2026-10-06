# Active development roadmap

migration-mirror is the TechVit OSS capability owner for **migration assessment / source-target integrity verification / cutover evidence**. WordPress → EmDash is the first supported adapter pair.

This file tracks the immediate issue set to open after repository Issues are enabled.

## P0

### Move GitHub Actions CI to Cloudflare-native CI

Parent: `sibukixxx/techvit-platform#41`.

- reproduce the current Go lint/test/race/E2E/govulncheck/build checks in Cloudflare Workflows + Sandbox
- preserve darwin/linux amd64/arm64 cross-build checks where they are meaningful
- store CI artifacts/logs outside GitHub Actions
- keep GitHub only as a transitional source/PR/status adapter
- remove `.github/workflows/ci.yml` only after parity
- update the README CI badge/documentation after cutover

### Complete non-minimal presets

Finish `small-production`, `seo-production`, `media-heavy`, and `custom-rebuild` without adding a monolithic migrate-all command.

- replace placeholder steps with typed/deterministic commands
- dry-run remains default
- each destructive action requires explicit apply/confirmation
- fixture coverage for every preset phase
- preserve evidence artifacts for downstream consumers

## P1

### Add theme / plugin / mu-plugin analysis

Implement the v0.5 roadmap:

- `theme analyze`
- `plugins analyze`
- `mu-plugins analyze`
- rebuild planning evidence

The public OSS must emit technical facts/evidence, not customer-specific recommendations.

### Strengthen migration verification

Extend content/cutover verification where real migrations expose gaps:

- custom fields and plugin-owned metadata
- taxonomy/author/status/date fidelity
- media completeness and attachment references
- redirect and canonical/SEO integrity
- explicit unknown/unavailable evidence rather than false pass

## P2

### Version the downstream evidence handoff

Define a stable, consumer-neutral artifact contract for downstream TechVit systems.

Expected consumers may include Insight / techvit-business, but migration-mirror must not import private business logic.

The artifact should contain:
- source provenance
- technical facts
- deterministic signals/findings
- verification outcomes
- schema/version metadata
- no customer-specific recommendation/pricing/CRM fields

## Product rule

Real customer work is an input to learning, not source code reuse.

Customer-specific information, credentials, content, business rules, and confidential data must not be copied into public fixtures. Generalize only the reusable technical pattern and reproduce it with synthetic fixtures.
