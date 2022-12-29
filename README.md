# golang-repo-template
A template for all my future Go projects!

## Actions created by this template:


### Testing

The pkg-cov workflow runs all go tests and ensures pkg coverage is above 80%.

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/pkg-cov.yml/badge.svg?event=push)

The pages workflow publishes a test coverage website everytime there is a push to the main branch. The website can be found here: https://gowhale.github.io/go-shopping-list/#file0

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/pages.yml/badge.svg?event=push)

### Linters

The revive workflow is executed to statically analsye go files: https://github.com/mgechev/revive

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/revive.yml/badge.svg?event=push)

The golangci-lint workflow runs the golangci-lint linter: https://github.com/golangci/golangci-lint

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/golangci-lint.yml/badge.svg?event=push)

### Project Management

The issue workflow adds a new issue to the projects Kanban board:

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/issue.yml/badge.svg?event=push)

The cut release workflow creates a binary executable everytime a release is published. The binary file is attached to the release.

![example event parameter](https://github.com/gowhale/go-shopping-list/actions/workflows/cut-release.yml/badge.svg?event=push)

