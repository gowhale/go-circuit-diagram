# go-circuit-diagram
A repo to create circuit diagrams!

## Summary

So far I the main prog creates a a PNG of an LED which you can find in the images folder.

It also creates a board with multiple LED's painted onto it which you can find in the images folder.

### Example:

![image](https://user-images.githubusercontent.com/32711718/210105762-90de82bb-61ad-464f-ab40-6f26e62f6028.png)
## Actions created by this template:

### Testing

The pkg-cov workflow runs all go tests and ensures pkg coverage is above 80%.

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/pkg-cov.yml/badge.svg?event=push)

The pages workflow publishes a test coverage website everytime there is a push to the main branch. The website can be found here: https://gowhale.github.io/go-circuit-diagram/#file0

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/pages.yml/badge.svg?event=push)

### Linters

The revive workflow is executed to statically analsye go files: https://github.com/mgechev/revive

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/revive.yml/badge.svg?event=push)

The golangci-lint workflow runs the golangci-lint linter: https://github.com/golangci/golangci-lint

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/golangci-lint.yml/badge.svg?event=push)

### Project Management

The issue workflow adds a new issue to the projects Kanban board:

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/issue.yml/badge.svg?event=push)

The cut release workflow creates a binary executable everytime a release is published. The binary file is attached to the release.

![example event parameter](https://github.com/gowhale/go-circuit-diagram/actions/workflows/cut-release.yml/badge.svg?event=push)

