# Release and compatibility policy

## Status

`migration-mirror` is the canonical repository and product name.

The former `wp2emdash` repository/name is not a compatibility surface. No legacy binary alias, Go module alias, or dual-maintained implementation is provided.

## Pre-1.0 policy

Until the first stable 1.0 release:

- CLI flags and artifact schemas may change when required to establish a cleaner generic migration boundary.
- Breaking changes must be documented in the release notes and README.
- Deterministic evidence artifacts should carry explicit schema/version metadata before they are treated as stable downstream contracts.
- WordPress and EmDash are adapters. Their provider-specific behavior must not leak into generic domain contracts unless the concept is migration-generic.

## Stable surfaces to converge on

The intended stable public surfaces are:

1. `migration-mirror` CLI command semantics
2. versioned machine-readable evidence artifacts
3. deterministic source/target comparison behavior
4. exit codes suitable for CI/cutover gating

Internal Go packages remain implementation details unless a package is explicitly documented as a public SDK.

## Versioning

Use semantic versioning for public releases.

- patch: bug fixes with no intentional contract change
- minor (pre-1.0): features and documented contract adjustments
- major / 1.0: stable CLI and artifact compatibility commitment

The first public release should not be tagged until clean-clone build/test/lint/E2E verification succeeds in the canonical repository.
