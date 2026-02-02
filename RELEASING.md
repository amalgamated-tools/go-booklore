# Releasing go-booklore

This document describes how to create and publish a new release of go-booklore.

## Release Process

### Option 1: Automatic via Git Tag (Recommended)

1. **Create and push a git tag:**
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

2. **GitHub Actions will automatically:**
   - Validate the version format (must be `vX.Y.Z`)
   - Build binaries for all supported platforms
   - Create a GitHub release with the built binaries
   - Generate release notes automatically from commits since the last release

### Option 2: Manual Trigger via GitHub UI

1. Go to **Actions** tab in the GitHub repository
2. Select **"Release"** workflow on the left
3. Click **"Run workflow"** button
4. Enter the version number (e.g., `v1.0.0`)
5. Click **"Run workflow"**

The workflow will:
- Validate the version format
- Build binaries for all platforms
- Create a GitHub release
- Tag the repository with the new version

## Version Format

Versions must follow semantic versioning and use the format: **`vX.Y.Z`**

Examples:
- `v1.0.0` - First release
- `v1.1.0` - Minor version bump (new features)
- `v1.0.1` - Patch version bump (bug fixes)
- `v2.0.0` - Major version bump (breaking changes)

## Supported Platforms

The build workflow creates binaries for:
- **Linux**: amd64, arm64
- **macOS**: amd64, arm64 (Apple Silicon)
- **Windows**: amd64

Each binary is named with the format: `go-booklore-vX.Y.Z-{OS}-{ARCH}[.exe]`

Example:
- `go-booklore-v1.0.0-linux-amd64`
- `go-booklore-v1.0.0-darwin-arm64`
- `go-booklore-v1.0.0-windows-amd64.exe`

## Pre-Release Checklist

Before releasing, ensure:

- [ ] All tests pass locally: `go test -v ./...`
- [ ] Code is linted: `golangci-lint run`
- [ ] Build succeeds: `go build ./...`
- [ ] Update `CHANGELOG.md` with release notes (if applicable)
- [ ] All commits are pushed to main branch
- [ ] The commit you want to release is the HEAD of the main branch

## After Release

Once a release is created:

1. **GitHub Release** - A release page is created with:
   - Automated release notes (commit summary)
   - Pre-built binaries for all platforms
   - Direct download links

2. **pkg.go.dev** - The release is automatically indexed on [pkg.go.dev](https://pkg.go.dev/github.com/amalgamated-tools/go-booklore)
   - Usually available within a few minutes
   - Go users can `go get` the specific version

3. **GitHub Releases API** - The release is accessible via:
   ```bash
   gh release view v1.0.0
   ```

## Downloading Released Binaries

Users can download pre-built binaries from:
- [GitHub Releases Page](https://github.com/amalgamated-tools/go-booklore/releases)
- Direct URL: `https://github.com/amalgamated-tools/go-booklore/releases/download/vX.Y.Z/go-booklore-vX.Y.Z-{OS}-{ARCH}`

Example:
```bash
wget https://github.com/amalgamated-tools/go-booklore/releases/download/v1.0.0/go-booklore-v1.0.0-linux-amd64
chmod +x go-booklore-v1.0.0-linux-amd64
```

## Troubleshooting

### Version validation failed
- Ensure the version follows the format: `vX.Y.Z`
- Both major, minor, and patch versions must be numeric

### Build failed
- Ensure Go 1.25.5 or later is installed
- Run `go mod tidy` and `go mod verify`
- Check that all dependencies are available

### Release not appearing on pkg.go.dev
- It can take a few minutes to index
- Ensure the tag points to a commit with a valid `go.mod` file
- Check https://pkg.go.dev/github.com/amalgamated-tools/go-booklore

## Rollback

If a release needs to be rolled back:

```bash
# Delete local tag
git tag -d v1.0.0

# Delete remote tag
git push origin :refs/tags/v1.0.0

# Delete GitHub release (via GitHub UI)
# Navigate to Releases, click on the release, and delete it
```

## Workflow Configuration

The release workflow is defined in `.github/workflows/release.yml` and:
- Triggers on git tag pushes matching `v*`
- Can be manually triggered via `workflow_dispatch`
- Builds for multiple platforms using `go build`
- Creates releases with GitHub Actions
- Notifies pkg.go.dev of new releases

For modifications to the workflow, edit `.github/workflows/release.yml`.
