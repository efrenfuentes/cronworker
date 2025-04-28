# CronWorker

A CLI application for managing cron jobs, built with Go.

## Installation

### Using Go

```bash
go install github.com/efrenfuentes/cronworker
```

### Using Docker

```bash
# Build the Docker image
docker build -t cronworker .

# Run the container
docker run -v $(pwd)/config.yaml:/app/config.yaml cronworker

# Run with custom config file
docker run -v /path/to/your/config.yaml:/app/config.yaml cronworker
```

## Configuration

CronWorker uses a YAML configuration file. By default, it looks for a `cronworker.yaml` file in your home directory. You can also specify a custom config file using the `--config` flag.

Example configuration file (`config.yaml`):

```yaml
jobs:
  - name: "example-job"
    schedule: "*/2 * * * *"  # Run every 2 minutes
    command: "echo 'Hello from cronworker'"
    enabled: true
```

## Usage

### Basic Commands

```bash
# Show help
cronworker --help

# Show version
cronworker version

# Use custom config file
cronworker --config /path/to/config.yaml
```

## Development

To build the project:

```bash
go build
```

To run tests:

```bash
go test ./...
```

## License

MIT 