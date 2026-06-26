# Contributing to Blinkfetch

Thank you for your interest in contributing to Blinkfetch! This is a learning project where I'm developing a neofetch clone in Go, and I'd love your help in making it better.

## Code of Conduct

Please be respectful and constructive in all interactions with other contributors.

## Getting Started

### Prerequisites

- **Go 1.26+** installed on your system
- Basic familiarity with Go and the command line
- Git for version control

### Setting Up Your Development Environment

1. **Fork the repository** on GitHub

2. **Clone your fork locally:**
   ```bash
   git clone https://github.com/YOUR_USERNAME/blinkfetch.git
   cd blinkfetch
   ```

3. **Add the upstream remote:**
   ```bash
   git remote add upstream https://github.com/nxkst/blinkfetch.git
   ```

4. **Install dependencies:**
   ```bash
   go mod download
   ```

5. **Build and test locally:**
   ```bash
   go build ./cmd/blinkfetch
   go test ./...
   ```

## How to Contribute

### Reporting Bugs

- Check the [existing issues](https://github.com/nxkst/blinkfetch/issues) to avoid duplicates
- Provide clear steps to reproduce
- Include your OS, Go version, and relevant system information
- Attach screenshots or error logs if applicable

### Suggesting Features

- Open a [new issue](https://github.com/nxkst/blinkfetch/issues/new) with a clear title and description
- Explain the use case and why the feature would be valuable
- Label it as an enhancement

### Submitting Code Changes

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes:**
   - Follow Go conventions and best practices
   - Keep code clean and well-documented
   - Add comments to explain complex logic

3. **Test your changes:**
   ```bash
   go test ./...
   go build ./cmd/blinkfetch
   ```

4. **Commit with clear messages:**
   ```bash
   git commit -m "Brief description of changes"
   ```
   - Use present tense ("Add feature" not "Added feature")
   - Reference related issues if applicable

5. **Push to your fork:**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Open a Pull Request:**
   - Provide a descriptive title and summary
   - Link related issues using `#issue-number`
   - Explain what the PR does and why
   - Include any relevant testing details

## Code Style

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code:
  ```bash
  go fmt ./...
  ```
- Use meaningful variable and function names
- Keep functions focused and single-purpose
- Write comments for exported functions

## Project Structure

```
blinkfetch/
├── cmd/
│   └── blinkfetch/        # Main entry point
├── internal/              # Internal packages
├── screenshots/           # Project screenshots
├── go.mod & go.sum       # Go module files
├── README.md             # Project overview
└── LICENSE               # MIT License
```

- `cmd/` - Command-line applications
- `internal/` - Private packages not exposed as part of the public API

## Dependencies

We currently use:
- **go-distro** - For Linux distribution detection
- **gopsutil** - For system information gathering

Please discuss any new dependencies with the maintainer before adding them.

## Testing

- Write tests for new features and bug fixes
- Ensure all tests pass before submitting a PR:
  ```bash
  go test -v ./...
  ```

## License

By contributing to Blinkfetch, you agree that your contributions will be licensed under the MIT License.

## Questions?

Feel free to:
- Open an issue for clarification
- Ask in pull request comments
- Check existing issues and discussions

Happy coding! 🚀
