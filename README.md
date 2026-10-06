# migration-mirror

<!-- oss-readme-intro:v1 -->
A **migration assessment and verification toolkit** for WordPress. It inspects a site before migration and verifies the destination afterward so you can detect missing content, SEO data, redirects, media and structural differences before cutover.

### Who is it for?

- engineers migrating WordPress sites
- agencies that need a repeatable migration checklist and evidence bundle
- teams that want CI-style verification instead of relying only on manual spot checks

### Install

```sh
git clone https://github.com/sibukixxx/migration-mirror.git
cd migration-mirror
make build
./bin/migration-mirror --help
```

### Typical workflow

```sh
migration-mirror assess --wp-root /var/www/html --out ./assessment
migration-mirror content snapshot wordpress --wp-root /var/www/html
migration-mirror content snapshot emdash --url https://new.example.com
migration-mirror content verify
```

### What can it check?

Technical assessment, migration risk signals, database planning, media inventory, SEO metadata and redirects, URL maps, environment readiness, privacy-preserving source/target snapshots and deterministic content verification.

### Use it when

Use it **before migration** to understand what must be preserved and **after migration** to prove what changed or went missing. It produces technical evidence; pricing, customer recommendations and CRM workflow stay outside this OSS boundary.

See [Quick Start](#quick-start), [Commands](#v01--v02--v04-commands), and the documentation under [docs/](docs/) for detailed operation.
<!-- /oss-readme-intro -->


> [!NOTE]
> **Canonical repository:** `migration-mirror` is the active source of truth. WordPress → EmDash is the first supported adapter pair, not the product boundary.


<!-- ecosystem-expansion:v1 -->
## Expansion contract

This repository is one capability owner in the TechVit ecosystem. New ideas should extend the existing graph before creating another silo.

**Canonical responsibility:** migration assessment, source/target integrity verification, and reproducible cutover evidence

When new work appears:

1. **Extend here** when the new behavior belongs to this repository's canonical responsibility.
2. **Integrate, don't copy** when another repository owns the concept. Exchange versioned API/artifact/event references; do not share databases or import another product's internals.
3. **Promote reusable capability only after real reuse appears.** Keep the first vertical slice close to its product; extract a common core when there are at least two real consumers or a deliberate public/private boundary.
4. **Create a new repository only** for a distinct product/vertical with its own lifecycle, deployment/security boundary, or OSS/commercial boundary. Its README must declare upstream/downstream owners from day one.
5. **Experiments may end without losing the asset.** Move validated rules, contracts, fixtures, and learnings into the canonical owner, then freeze/archive the duplicate implementation.

**Route new work:** migration固有のinspection/evidence/verificationをここへ追加する。WordPressやEmDash固有処理はadapterとして閉じ込め、business recommendationやcompany workflowはdownstreamへ渡す。

Cross-cutting owners: TechVit Admin = company operations; TechVit Platform = shared cloud/bootstrap/identity; Insight = research semantics; techvit-insight = managed research; Relayboard = work execution; AI Allowance Guard = privileged-action authority; EC Hub = commerce; AffiliateMediaOS/VideoForge = content production; ja-company-* = company-data/outreach producers; public/product sites = acquisition surfaces.

If ownership is unclear, implement the smallest vertical slice in the nearest existing owner first. Extract only when actual reuse proves the boundary.
<!-- /ecosystem-expansion:v1 -->

[![CI](https://github.com/sibukixxx/migration-mirror/actions/workflows/ci.yml/badge.svg)](https://github.com/sibukixxx/migration-mirror/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sibukixxx/migration-mirror.svg)](https://pkg.go.dev/github.com/sibukixxx/migration-mirror)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/sibukixxx/migration-mirror)](go.mod)
[![Sponsor](https://img.shields.io/badge/sponsor-%E2%9D%A4-ea4aaa?logo=githubsponsors)](https://github.com/sponsors/sibukixxx)

English | [日本語](README.ja.md) | [简体中文](README.zh-CN.md)

`migration-mirror` is a **WordPress Migration Assessment & Verification Toolkit**. It measures, detects, verifies, and emits evidence that can be consumed by people, CI, or other tools. Its existing WordPress → EmDash commands remain supported.

The public OSS boundary ends at technical facts, evidence, signals, deterministic reference scores, snapshots, and verification. It does not contain TechVit customer-diagnosis rules, sales recommendations, estimates, pricing, proposals, CRM, or lead automation.

> **EmDash imports your WordPress content. migration-mirror proves that nothing important was lost.** The official importer owns WXR/plugin import, Gutenberg conversion, schema creation, and media rewriting. `content verify` independently adds durable ID evidence, semantic HTML ↔ Portable Text comparison, mapped-field loss detection, and a reproducible CI/cutover gate.


<!-- current-implementation:2026-10-06 -->
## Current implementation snapshot — 2026-10-06

- Development was explicitly reactivated on 2026-10-04. The current implementation is an evidence-first WordPress migration assessment and verification toolkit rather than an EmDash-only importer.
- Implemented commands cover assessment/audit, DB planning, media inventory, SEO metadata/redirect extraction, URL-map comparison, environment-secret checks, source/target content snapshots and deterministic content verification.
- Independent verification can compare WordPress with EmDash using hashes and structural evidence without storing raw post bodies, producing CI/cutover findings and a resolved ID ledger.
- The toolkit intentionally stops at technical facts, signals, snapshots and verification. Customer recommendations, estimates, pricing, CRM and managed multi-customer workflows remain private/downstream concerns.
- `migration-mirror` is now canonical. WordPress → EmDash remains the first concrete adapter pair while the verification model is generalized to other migration sources and targets.
<!-- /current-implementation -->

<!-- role-boundary:v1 -->
## Role and boundaries

**Role:** migration全般のtechnical facts・signals・source/target snapshots・verification evidenceを生成する **evidence producer / verification toolkit**。

### Owns

- read-only migration assessment, with WordPress as the first source adapter
- deterministic technical signals and reference scores
- migration/content verification artifacts
- source provenance and reproducible cutover evidence

### Does not own

- customer-specific business recommendation
- pricing, proposal, CRM, or sales workflow
- generic research/hypothesis semantics
- privileged deployment authority
- a managed multi-customer control plane

### Integration

```text
WordPress / migration target
        ↓
    migration-mirror
        ↓ evidence / assessment artifact
human / research engine / private decision layer
```

Public artifacts should remain consumer-neutral. A downstream system may interpret the evidence, but this OSS repository must not encode a particular customer's commercial decision.
## Table of Contents

- [Why Small Commands](#why-small-commands)
- [Install](#install)
- [Quick Start](#quick-start)
- [Commands](#v01--v02--v04-commands)
- [Presets](#presets)
- [Architecture](#architecture)
- [Design Principles](#design-principles)
- [HTTP Agent Schema](#http-agent-schema)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Support the Project](#support-the-project)
- [License](#license)

```
migration-mirror assess                 -> collect a read-only, evidence-first assessment bundle
migration-mirror audit                  -> measure migration complexity and score risk
migration-mirror db plan                -> generate a DB migration plan from summary.json
migration-mirror media scan             -> build a JSON manifest of wp-content/uploads
migration-mirror report                 -> regenerate risk-report.md from summary.json
migration-mirror run --preset           -> execute a migration phase preset
migration-mirror secrets check          -> verify required secrets for a given preset/profile
migration-mirror seo extract-meta       -> dump per-post SEO metadata as JSON
migration-mirror seo extract-redirects  -> dump .htaccess + plugin redirects as JSON
migration-mirror seo url-map            -> diff two URL maps to find missing/added URLs
migration-mirror doctor                 -> check required external tools
migration-mirror content snapshot       -> fingerprint content without storing body text
migration-mirror content verify         -> gate cutover on source-to-target integrity
```

## Why Small Commands

Real migrations are usually phased. In practice, projects tend to split into stages such as:

```
minimum validation -> small production -> SEO-sensitive production -> media-heavy -> custom rebuild
```

`migration-mirror` acts as an orchestrator for those phases. It automates the mechanical parts while leaving human decisions explicit.

## Install

### From Source

Go 1.22+ is required.

```bash
git clone https://github.com/sibukixxx/migration-mirror.git
cd migration-mirror
make build
./bin/migration-mirror --help

# or
go install github.com/sibukixxx/migration-mirror/cmd/migration-mirror@latest
```

## Quick Start

Run on a WordPress host, or anywhere that has access to a WordPress install:

```bash
# Read-only full assessment (local WP-CLI, SSH, or the HTTP agent)
migration-mirror assess --wp-root /var/www/html --out ./assessment
migration-mirror assess --ssh user@example.com --wp-root /var/www/html --out ./assessment

# 1. Check external dependencies
migration-mirror doctor

# 2. Audit a WordPress site from a local path
migration-mirror audit --wp-root /var/www/html

# 3. Audit via HTTP agent
migration-mirror audit \
  --agent-url https://example.com/wp-json/migration-mirror/v1/audit \
# 4. Override the legacy risk-band policy (compatibility path)
migration-mirror audit \
  --wp-root /var/www/html \
  --risk-bands ./config/custom-risk-bands.json

# 5. Scan uploads locally
migration-mirror media scan --dir /var/www/html/wp-content/uploads --hash

# 6. Scan uploads via HTTP agent
migration-mirror media scan \
  --agent-url https://example.com/wp-json/migration-mirror/v1/media-scan \  --dir wp-content/uploads

# 7. Dry-run a preset, then apply it
migration-mirror run --preset minimal --wp-root /var/www/html --dry-run
migration-mirror run --preset minimal --wp-root /var/www/html --apply

# 8. Generate a DB migration plan from summary.json
migration-mirror db plan \
  --from migration-mirror-output/summary.json \
  --preset small-production

# 9. Check whether the current environment already has the required secrets
migration-mirror secrets check --profile small-production
migration-mirror secrets check --profile media-heavy --json

# 10. Run preset minimal via split agent endpoints
migration-mirror run --preset minimal \
  --agent-audit-url https://example.com/wp-json/migration-mirror/v1/audit \
  --agent-media-url https://example.com/wp-json/migration-mirror/v1/media-scan \  --wp-root /var/www/html \
  --apply

# 11. Use the same public risk-band policy during preset execution
migration-mirror run --preset minimal \
  --wp-root /var/www/html \
  --risk-bands ./config/custom-risk-bands.json \
  --apply
```

Artifacts are written to `migration-mirror-output/` by default:

- `summary.json`, `inventory.json`, `signals.json`, `technical-report.md` (`assess`)
- `summary.json`
- `risk-report.md`
- `media-manifest.json`

## v0.1 / v0.2 / v0.4 Commands

| Command | Purpose | Main Flags |
| --- | --- | --- |
| `assess` | Aggregate read-only facts, evidence-backed signals, timing, and technical score | `--wp-root` `--out` `--json` `--ssh` `--agent-url` |
| `doctor` | Check required tools such as `wp`, `wrangler`, and `git` | `--json` |
| `audit` | Measure 14 migration signals via local WP-CLI, SSH, or HTTP agent | `--wp-root` `--write` `--json` `--ssh` `--agent-url` `--risk-bands` |
| `db plan` | Generate a JSON/Markdown DB migration plan from `summary.json` | `--from` `--preset` `--write` `--json` |
| `media scan` | Build a JSON manifest via local path, SSH, or HTTP agent | `--dir` `--hash` `--max-files` `--histogram-only` `--ssh` `--agent-url` |
| `report` | Regenerate `risk-report.md` from `summary.json` | `--from` `--stdout` |
| `run --preset` | Execute one of the predefined migration presets | `--preset` `--wp-root` `--dry-run` `--apply` `--ssh` `--agent-audit-url` `--agent-media-url` `--risk-bands` |
| `secrets check` | Verify whether required secret env vars are already present | `--profile` `--json` |
| `seo extract-meta` | Dump per-post SEO metadata (Yoast / Rank Math / AIOSEO merged) | `--wp-root` `--write` `--ssh` |
| `seo extract-redirects` | Extract redirects from `.htaccess` and Redirection / SRM plugins | `--wp-root` `--write` `--ssh` |
| `seo url-map` | Diff two URL maps (matched / missing / new) | `--old` `--new` `--write` |
| `content snapshot wordpress` | Capture privacy-preserving WordPress fingerprints | `--wp-root` `--map` `--ssh` |
| `content snapshot emdash` | Capture EmDash fingerprints via the official CLI | `--url` `--map` `--jobs` |
| `content verify` | Compare snapshots and emit a cutover gate plus resolved ID ledger | `--expected` `--actual` `--map` `--policy` |

### Added in v0.2

- `db plan`
  Reads `summary.json` and emits a plan that classifies WordPress tables and metadata as `export`, `review`, `transform`, or `skip`. It does not dump or modify the database. The plan also carries a `target_notes` checklist for the destination side (EmDash / Cloudflare D1+R2): never copy auth tables across environments, verify the seed importer preserves source timestamps, D1's per-statement size ceiling, media completeness gating before cutover, and more. The lessons behind these notes are documented in [`docs/case-studies/newsmedia-13k-posts-lessons.md`](docs/case-studies/newsmedia-13k-posts-lessons.md) (Japanese).
- `secrets check`
  Checks existing environment variables only. It never generates or overwrites `.env`. Supported profiles are `small-production`, `seo-production`, `media-heavy`, `custom-rebuild`, and `agent`.

### Added in v0.4

- `seo extract-meta`
  Lists every published post / page via `wp post list`, then merges Yoast, Rank Math, and AIOSEO post meta keys into a single `seo-meta.json`. Plugin precedence is **Yoast > Rank Math > AIOSEO**, with the contributing plugin recorded in the `source` field. The `{prefix}yoast_indexable` table used by Yoast 14+ is auto-detected as well, so editor overrides are not lost when the legacy postmeta rows are empty or stale (explicit postmeta > indexable > core; provenance recorded as `source: "yoast_indexable"`).
- `seo extract-redirects`
  Combines three redirect sources into a single `seo-redirects.json`: `.htaccess` (`Redirect`, `RedirectMatch`, `RewriteRule [R=...]`), Redirection plugin (`wp_redirection_items`), and Safe Redirect Manager (`post_type=redirect_rule`).
- `seo url-map`
  Diffs two URL maps and reports `matched`, `only_in_old` (likely needs an explicit redirect), and `only_in_new`. Inputs may be JSON or plain text (one URL per line). URLs are normalised before comparison (scheme, trailing slash, fragment); path case is preserved.

### Independent content verification

Run the official EmDash import first, then capture both sides and verify them:

```bash
migration-mirror content snapshot wordpress --wp-root /var/www/html
migration-mirror content snapshot emdash --url https://new.example.com
migration-mirror content verify
```

Snapshots contain hashes and structural evidence, not raw post bodies. Verification checks identity, status, title, timestamps, visible text, headings, links, images, and configured custom fields. Explicit WordPress ID ↔ EmDash ID mappings win; unique slug matches are reported as warnings and written to `content-resolved-map.json` for subsequent deterministic runs. Critical/errors fail the command for CI and cutover use; `--policy` can override stable issue-code severities or allowed counts.

The core scoring rubric is additive. Legacy level bands remain replaceable through `--risk-bands path/to/custom.json` for compatibility. The bundled public policy does not provide estimates or pricing. New consumers should use `assess` and its evidence-backed technical score breakdown.


## Presets

`migration-mirror run --preset <name>` currently exposes five presets:

| Preset | Scope |
| --- | --- |
| `minimal` | PoC-level audit and migration feasibility report |
| `small-production` | Small production blog / landing page migration |
| `seo-production` | Production migration with SEO-sensitive content |
| `media-heavy` | Media-heavy migration with large uploads footprint |
| `custom-rebuild` | Rebuild-heavy migration involving theme, plugins, mu-plugins, and integrations |

`minimal` is still the only fully implemented preset. As of v0.2, `small-production`, `seo-production`, `media-heavy`, and `custom-rebuild` now include `db plan` and `secrets check`, while later steps such as `env generate`, deploy, and media sync/verify remain placeholder `todo` steps.

## Architecture

The repository is structured in three layers:

```text
cmd/
  migration-mirror/main.go
internal/
  cli/          cobra command definitions
  usecase/      orchestration per command
  domain/       pure data structures and rules
  infra/        adapters for external systems
  shell/        thin os/exec wrapper
test/
  e2e/          end-to-end helpers and fixtures
```

## Assessment scope and non-goals

`assess` requires local WP-CLI access, SSH access, or a compatible read-only HTTP agent. A public URL alone cannot expose private plugin state, database-backed custom fields, drafts, users, or filesystem evidence. A future `public-audit` command, if added, will explicitly identify its scope as **Public Surface Only** and will not be described as a full assessment.

The assessment pipeline is `Facts → Signals → Technical Score`. Missing probes are represented as `unavailable` rather than silently interpreted as zero or false. Every scored signal carries evidence, and scoring is deterministic for the same facts. See [`docs/assessment.md`](docs/assessment.md) for schema and compatibility details.

Non-goals include deciding whether WordPress should be replaced, recommending EmDash/Astro/headless CMS, estimating or pricing work, generating proposals or customer-facing sales reports, guaranteeing SEO rankings, and vulnerability exploitation.

Dependency direction: `cli -> usecase -> {domain, infra} -> shell`.

## Design Principles

- One command = one responsibility
- JSON / Markdown output
- Dry-run by default for destructive flows
- Thin wrappers around external tools
- Never generate or overwrite `.env`

## HTTP Agent Schema

`migration-mirror` can consume metrics from a read-only WordPress HTTP agent instead of SSH.

### `GET /wp-json/migration-mirror/v1/audit`

Headers:

```http
Authorization: Bearer <token>
Accept: application/json
```

Response:

```json
{
  "audit": {
    "site": {
      "home_url": "https://example.com",
      "site_url": "https://example.com",
      "wp_version": "6.5.0",
      "php_version": "8.2.12",
      "db_prefix": "wp_",
      "is_multisite": "no"
    },
    "content": {
      "posts": 120,
      "pages": 12,
      "drafts": 3,
      "private_posts": 1,
      "categories": 8,
      "tags": 15,
      "users": 4,
      "approved_comments": 22
    },
    "uploads": {
      "exists": true,
      "size": "12KB",
      "file_count": 3,
      "posts_with_uploads_paths": 7,
      "posts_with_http_urls": 2
    },
    "theme": {
      "active_theme": "example-theme",
      "php_files": 12,
      "css_files": 2,
      "js_files": 4,
      "page_templates": 3,
      "hook_like_occurrences": 12,
      "jquery_like_occurrences": 3
    },
    "plugins": {
      "active_count": 2,
      "has_acf": true,
      "has_woocommerce": false,
      "has_seo": true,
      "has_form": false,
      "has_redirect": true,
      "has_member": false,
      "has_multilingual": false,
      "has_cache": true
    },
    "customization": {
      "custom_post_type_count": 1,
      "custom_taxonomy_count": 1,
      "mu_plugin_count": 0,
      "mu_plugin_hook_like_occurrences": 0,
      "shortcode_post_count": 5,
      "seo_meta_count": 11,
      "serialized_meta_count": 9,
      "oversized_content_count": 0,
      "htaccess_redirect_like_lines": 0,
      "code_redirect_like_occurrences": 0,
      "external_integration_like_occurrences": 0
    }
  },
  "warnings": [
    {
      "code": "content.posts",
      "message": "probe failed"
    }
  ]
}
```

### `GET /wp-json/migration-mirror/v1/media-scan`

Query parameters:

- `dir`
- `hash=1`
- `max_files=200`
- `histogram_only=1`

Response:

```json
{
  "base_dir": "wp-content/uploads",
  "total_files": 3,
  "total_bytes": 12,
  "extensions": {
    "txt": 1
  },
  "files": [
    {
      "path": "2024/01/hello.txt",
      "size": 12,
      "sha256": "optional",
      "mime": "text/plain",
      "ext": "txt"
    }
  ]
}
```

Selection priority in Go is:

- `agent-url`
- `ssh`
- local

For `run --preset`, the preferred form is:

- `agent-audit-url`
- `agent-media-url`

`--agent-url` is still accepted as a backward-compatible fallback.

## Roadmap

| Version | Planned Scope | Status |
| --- | --- | --- |
| v0.1 | doctor / audit / media scan / report / run --preset minimal | done |
| v0.2 | `env generate`, `secrets check`, `db plan` | `secrets check` / `db plan` done; `env generate` pending |
| v0.3 | `media sync`, `media verify`, legacy uploads route worker scaffolding | `media sync` / `media verify` done; worker scaffolding pending |
| **v0.4 (current)** | `seo extract-meta`, `seo extract-redirects`, URL map comparison | done |
| v0.5 | `theme analyze`, `plugins analyze`, `mu-plugins analyze`, rebuild planning report | not started |
| v1.0 | Full implementation of all five presets plus Cloudflare-native CI / verification pipeline | not started |


## Contributing

Issues are part of the active development workflow again. Use them for reproducible WordPress migration problems, deterministic evidence gaps, roadmap work, and Cloudflare CI migration. Customer-specific recommendations, pricing, CRM, and private decision logic remain outside this public repository. See [CONTRIBUTING.md](CONTRIBUTING.md) for the PR flow and quality gates (`make build` / `make test` / `make vet`, `go test -race -count=1 ./...`).

## Support the Project

If `migration-mirror` saves you migration time, consider [sponsoring on GitHub Sponsors](https://github.com/sponsors/sibukixxx).

## License

MIT
