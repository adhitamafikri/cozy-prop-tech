# Cozy Prop Tech

This is the learning project for me to onboard as Fullstack Engineer in a leading High-End property management company in Indonesia

## Prerequisites

Add these to your `/etc/hosts`

```
127.0.0.1 cozy-prop.local
127.0.0.1 admin.cozy-prop.local
127.0.0.1 api.cozy-prop.local
```

You need to have these things installed on your machine, in order to run this project locally

- Docker
- Bun
- Go (programming language)

## AI Assisted Development

> This project supports AI-assisted engineering with **Opencode**

This is one of the interesting part in this project. This project demonstrates AI-Assisted development using multiple agents. Each of these agents are called a `subagent`, which operates in a _Role-isolated execution context_.

### Implementation Agents

> Focuses feature implementation while maintaining the code quality through unit testing

1. Frontend Agent
2. Backend Agent

### Review Agents

1. Frontend Reviewer Agent
2. Backend Reviewer Agent

### Audit and E2E Testing Agents

1. Frontend Accessibility Agent
2. Frontend E2E Testing Agent

## Projects

This repository contains several sub-projects

### Frontend

- **web**: customer-facing site
- **admin**: admin-facing site

### Backend

Our backend **api** is responsible for handling these domains:

- Auth (RBAC)
- Users
- Properties
- Listings
- Availability
- Booking

This backend project uses _Hexagonal Architecture_

## Running the Project Locally

**Copy .env and Install Dependencies for All Projects**

```bash
make prepare
```

**Run the projects locally**

```bash
make up
```

**Stops the projects locally**

```bash
make down
```

You will have these services running on your machine:

| Services | Ports |
| -------- | ----- |
| web      | 5173  |
| admin    | 5174  |
| api      | 8082  |

## Reading Links

- [https://github.com/khannedy/golang-clean-architecture/blob/main](https://github.com/khannedy/golang-clean-architecture/blob/main)
- [https://betterstack.com/community/guides/logging/logging-in-go/](https://betterstack.com/community/guides/logging/logging-in-go/)
- [https://opencode.ai/docs/agents](https://opencode.ai/docs/agents)
- [https://code.claude.com/docs/en/sub-agents](https://code.claude.com/docs/en/sub-agents)
