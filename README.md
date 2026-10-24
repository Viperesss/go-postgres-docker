# Go PostgreSQL Docker

An educational project for learning backend development in Go and working with the tools commonly used in backend applications.

The project focuses on database integration, development environment configuration, containerization, and application lifecycle management.

## Technologies

- **Go** — application development
- **PostgreSQL** — relational database
- **Docker** — containerization
- **Docker Compose** — managing application services
- **SQL migrations** — managing database schema changes
- **Environment variables (`.env`)** — application configuration
- **Makefile** — automating common development commands
- **Git** — version control and branch-based development
- **OS signals** — graceful application shutdown

## What I Learned

- Connecting a Go application to PostgreSQL.
- Creating and applying database migrations.
- Managing environment variables and application configuration.
- Running services using Docker and Docker Compose.
- Configuring persistent database storage with Docker volumes.
- Automating development tasks with Makefile.
- Working with Git branches and organizing changes into separate development workflows.
- Handling operating system signals and implementing graceful shutdown.

## Project Structure

The project structure depends on the current implementation. The main components include the Go application, database configuration, migration files, and containerization setup.

## Getting Started

### Prerequisites

Make sure the following tools are installed:

- Go
- Docker
- Docker Compose

### Configuration

Create a `.env` file in the project root and specify the environment variables required by the application.

Do not commit `.env` files containing passwords, credentials, or other secrets to version control.

### Running the Project

Use Docker Compose to build and start the configured services:

```bash
docker compose up --build
```

To stop the services:

```bash
docker compose down
```

To stop the services and remove named volumes:

```bash
docker compose down -v
```

**Warning:** The last command deletes the Compose project's named volumes, which may contain your PostgreSQL data.

Refer to the project's Makefile and configuration files for any additional commands required to run the application and database migrations.

## Purpose

This repository serves as a hands-on learning project for understanding the development workflow of a Go backend application, from database setup and schema migrations to containerized environments and graceful shutdown.

## Status

Educational project. The repository may evolve as new backend development concepts and tools are explored.
