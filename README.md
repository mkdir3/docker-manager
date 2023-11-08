# Docker Manager 🐳

Welcome to Docker Manager, the CLI tool designed to make Docker container management a breeze. With a focus on user experience and simplicity, Docker Manager allows you to manage your Dockerized projects with ease.

## Features 🚀

- Start, stop, and manage Docker containers with simple commands.
- Add or remove project configurations with interactive prompts.
- List all Docker projects with paths and status.
- Load and save project configurations from a JSON file.

## Getting Started 🌟

### Prerequisites

- Docker installed on your machine.

### Commands 🛠️
- docker-manager start: Start a Docker project.
- docker-manager stop: Stop a Docker project.
- docker-manager list: List all Docker projects.
- docker-manager manage: Add or remove Docker projects.

### Configurations 📝
Docker Manager uses a projects.json file to store project configurations. If the file doesn't exist, it will prompt you to create one.

Example projects.json:
```json
{
  "ProjectName": "My Amazing Project,
  "AnotherProject": "/path/to/docker-compose.yml"
}
```
### License 📜
Distributed under the MIT License.
