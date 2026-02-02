# AGENTS.md - Development Guide for Automated Agents

This document provides essential information for automated agents working on the `go-booklore` project.

## Project Overview

**go-booklore** is a type-safe Go client SDK for the Booklore API, automatically generated from an OpenAPI 3.1.0 specification.

- **Language**: Go 1.25.5+
- **Type**: SDK/Library
- **Module**: `github.com/amalgamated-tools/go-booklore`
- **Purpose**: Provide developers with a strongly-typed interface to the Booklore book management API

## Key Characteristics

### 1. Auto-Generated Code
- The SDK is **automatically generated** from the OpenAPI specification
- **Do NOT manually edit** `pkg/client/client.gen.go` or model files
- Changes should be made to:
  - The OpenAPI specification itself (api-docs.json)
  - The `generate.sh` script (for spec preprocessing)
  - Configuration in `cfg.yaml`

### 2. Generation Pipeline
```
api-docs.json (from server)
    ↓
generate.sh (jq preprocessing)
    ↓
oapi-codegen (code generation)
    ↓
pkg/client/client.gen.go (generated SDK)
```

The `generate.sh` script:
1. Downloads the latest OpenAPI spec
2. Fixes common issues with the spec using jq:
   - Converts `exclusiveMinimum: 0` to `true`
   - Converts `*/*` content types to `application/json`
   - Removes `/**` path patterns
   - Adds missing path parameters
   - **Removes duplicate XML content types** (keeps only one)
3. Runs `go generate` to create the client code

### 3. Special XML Content Type Handling
The project includes special handling for endpoints with multiple XML content types (e.g., OPDS search):
- Multiple XML types (application/opensearchdescription+xml, application/atom+xml, text/xml) are consolidated
- Only `application/xml` is kept to avoid duplicate struct fields
- See the `generate.sh` script for implementation

## Project Structure

```
go-booklore/
├── main.go                    # Example usage / test file
├── generate.go                # Code generation directives (go:generate comment)
├── generate.sh                # Preprocessing script for OpenAPI spec
├── cfg.yaml                   # oapi-codegen configuration
├── api-docs.json              # OpenAPI specification (downloaded)
├── go.mod                      # Go module definition
├── go.sum                      # Dependency checksums
├── README.md                   # User documentation
├── RELEASING.md               # Release procedures
├── SDK_USAGE.md               # SDK usage examples
├── AGENTS.md                  # This file
├── .github/
│   └── workflows/
│       ├── release.yml        # Release automation (tag → binaries → release)
│       └── test.yml           # CI/CD pipeline (tests, lint, build)
├── scripts/
│   └── release.sh             # Helper script for creating releases
└── pkg/
    └── client/
        ├── client.gen.go      # Auto-generated client code (DO NOT EDIT)
        ├── types.go           # Auto-generated types (DO NOT EDIT)
        └── models.go          # Auto-generated models (DO NOT EDIT)
```

## Common Tasks for Agents

### 1. Regenerating the SDK

**When to regenerate:**
- After the upstream Booklore API changes its OpenAPI spec
- To test if spec preprocessing works correctly
- Before releasing a new version

**How to regenerate:**
```bash
./generate.sh
```

This will:
1. Download the latest OpenAPI spec from `https://demo.booklore.org/api/v1/api-docs`
2. Preprocess the spec with jq fixes
3. Run `go generate ./...` to create new client code

**After regeneration:**
- Review changes to `pkg/client/client.gen.go`
- Ensure no compilation errors: `go build ./...`
- Run tests: `go test -v ./...`
- Commit changes: `git add . && git commit -m "chore: regenerate SDK from updated OpenAPI spec"`

### 2. Updating Dependencies

**View current dependencies:**
```bash
go list -m all
```

**Update dependencies:**
```bash
go get -u ./...
go mod tidy
```

**Check for security vulnerabilities:**
```bash
go list -json -m all | nancy sleuth
```

**Commit dependency updates:**
```bash
git add go.mod go.sum
git commit -m "chore: update dependencies"
```

### 3. Testing the SDK

**Run all tests:**
```bash
go test -v ./...
```

**Run tests with race detector:**
```bash
go test -v -race ./...
```

**Run tests with coverage:**
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Test a specific example:**
```bash
go run main.go
```

### 4. Linting and Code Quality

**Run golangci-lint locally:**
```bash
golangci-lint run --timeout=5m
```

**Build the project:**
```bash
go build -v ./...
```

**Format code:**
```bash
go fmt ./...
goimports -w .
```

## Release Process

**For creating releases, see [RELEASING.md](./RELEASING.md)**

Quick reference:
```bash
# Option 1: Use helper script
./scripts/release.sh v1.0.0

# Option 2: Manual git tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

The release workflow will:
1. Validate version format (vX.Y.Z)
2. Build binaries for Linux, macOS, and Windows
3. Create a GitHub release with binaries
4. Generate release notes from commits

## Important Files and Their Purposes

### `generate.sh`
**Purpose**: Preprocess OpenAPI spec and generate SDK code
**When to modify**: When adding new spec fixes or preprocessing logic
**DO NOT modify**: The actual client generation command (last line)

Key fixes in this file:
- Line 12: Fix `exclusiveMinimum` format issues
- Lines 13-20: Convert `*/*` to `application/json`
- Lines 21-23: Remove `/**` paths
- Lines 24-53: Add missing path parameters
- Lines 16-28: Remove duplicate XML content types

### `cfg.yaml`
**Purpose**: Configuration for oapi-codegen
**When to modify**: If you need to change code generation behavior
**Common changes**:
- Package name
- Output paths
- Generation options

### `go.mod`
**Purpose**: Define Go module and dependencies
**When to modify**: When updating Go version or dependencies
**Current Go version**: 1.25.5

### `main.go`
**Purpose**: Example usage and testing
**When to modify**: To demonstrate new SDK features or test changes
**DO NOT commit breaking changes** without updating examples

## Common Issues and Solutions

### Issue: golangci-lint version mismatch
**Error**: "Go language version (go1.24) used to build golangci-lint is lower than targeted Go version (1.25.5)"
**Solution**: Update `.github/workflows/test.yml` to use v4 with version v1.62.2+

### Issue: Duplicate struct fields in generated code
**Error**: Multiple `XML200` or `JSON200` fields in response structs
**Cause**: Multiple content types with the same media type
**Solution**: Add preprocessing in `generate.sh` to consolidate content types (already done for XML)

### Issue: Missing path parameters
**Error**: Generated client methods missing required path parameters
**Solution**: Add parameter definitions in `generate.sh` (see lines 24-53)

### Issue: Build fails with new OpenAPI spec
**Cause**: OpenAPI spec has errors or unsupported patterns
**Solution**: Add new jq filters to `generate.sh` to fix the issue

## Code Generation Details

### oapi-codegen
- **Tool**: `github.com/oapi-codegen/oapi-codegen/v2`
- **Invoked by**: `go generate generate.go`
- **Configuration**: `cfg.yaml`
- **Output**: `pkg/client/client.gen.go`

### Response Types
All `WithResponse` methods return structs with:
- `StatusCode()` method for HTTP status
- Status-specific fields: `JSON200`, `JSON201`, `JSON400`, etc.
- `HTTPResponse` field for raw HTTP response
- `Body` field for raw response body

Example generated response:
```go
type GetBooksResponse struct {
    Body         []byte
    HTTPResponse *http.Response
    JSON200      *[]Book
    JSON400      *ErrorResponse
}
```

### Request Types
Request methods accept:
- Context for cancellation/timeouts
- Request body (if applicable)
- Query/path parameters (if applicable)

Example generated method:
```go
func (c *ClientWithResponses) GetBooksWithResponse(
    ctx context.Context,
    params *GetBooksParams,
) (*GetBooksResponse, error)
```

## Dependencies

### Core Dependencies
- `github.com/getkin/kin-openapi` - OpenAPI spec parsing
- `github.com/oapi-codegen/runtime` - Generated code runtime support
- `github.com/oapi-codegen/oapi-codegen/v2` - Code generation tool

### Development Dependencies
- `golangci-lint` - Linting (v1.62.2+ for Go 1.25.5)
- Go standard library

## Git Workflow

### Branch Strategy
- `main` - Stable, released code
- `develop` - Development branch
- Feature branches - Named like `feature/description`

### Commit Messages
- Use conventional commits: `type: description`
- Types: `feat`, `fix`, `chore`, `docs`, `test`, `refactor`
- Examples:
  - `fix: handle multiple XML content types in OPDS endpoint`
  - `feat: add support for new API endpoint`
  - `chore: regenerate SDK from updated spec`
  - `docs: update README with new examples`

### Pull Requests
- Always create PRs against `develop` or `main`
- Include what changed and why
- Reference related issues
- Wait for CI/CD to pass before merging

## Useful Commands Reference

```bash
# Development
go mod tidy                    # Clean up dependencies
go build ./...                 # Build the project
go test -v ./...              # Run tests
go test -v -race ./...        # Run tests with race detection
go fmt ./...                   # Format code
golangci-lint run             # Run linter

# Generation
./generate.sh                  # Regenerate SDK from OpenAPI spec

# Releases
./scripts/release.sh v1.0.0    # Create a release

# Information
go list -m all                 # List all dependencies
go version                     # Check Go version
git log --oneline -10          # View recent commits
```

## When to Involve Humans

Contact the maintainers (DO NOT attempt alone) when:
1. **Breaking API changes** - Changes to public interfaces
2. **Major version bumps** - Significant functionality changes
3. **Security issues** - Vulnerability fixes or disclosures
4. **Upstream breaking changes** - API breaking changes requiring SDK redesign
5. **Release approval** - Before pushing to main/releasing
6. **Documentation disputes** - When unsure about project direction

## Testing Checklist for Agents

Before committing any changes:
- [ ] `go build ./...` succeeds
- [ ] `go test -v ./...` passes
- [ ] `golangci-lint run` passes (if not excluded)
- [ ] No uncommitted files in working directory
- [ ] Commit message follows conventional commits format
- [ ] Code is properly formatted (`go fmt ./...`)
- [ ] Generated code is up-to-date (if spec changed)

## Useful Resources

- [OpenAPI 3.1 Spec](https://spec.openapis.org/oas/v3.1.0)
- [oapi-codegen Documentation](https://github.com/oapi-codegen/oapi-codegen)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [Semantic Versioning](https://semver.org/)
- [Conventional Commits](https://www.conventionalcommits.org/)

## Questions and Support

For questions about the project:
1. Check README.md and SDK_USAGE.md
2. Review RELEASING.md for release procedures
3. Check this file (AGENTS.md) for agent-specific guidance
4. Look at existing issues and PRs for similar problems
5. Consult git history for context on design decisions

---

**Last Updated**: February 1, 2025
**Go Version**: 1.25.5
**OpenAPI Version**: 3.1.0
