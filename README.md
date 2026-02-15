# Cozy Prop Tech

This is the learning project for me to onboard as Fullstack Engineer in a leading High-End property management company in Indonesia

## Prerequisites

Add these to your `/etc/hosts`

```
127.0.0.1 cozy-prop.local
127.0.0.1 admin.cozy-prop.local
127.0.0.1 api.cozy-prop.local
127.0.0.1 api.cozy-prop-auth.local
127.0.0.1 api.cozy-prop-listings.local
127.0.0.1 api.cozy-prop-customer-support.local
```

You need to have these things installed on your machine, in order to run this project locally

- Docker
- Bun
- Go (programming language)

## Development Process

This project supports AI-assisted engineering with **Opencode** and **Claude Code**

## Projects

This repository contains several sub-projects

### Frontend

- **web**: customer-facing site
- **admin**: admin-facing site

### Backend

- **api-gateway**: client request entry point, request routing, auth checking
- **auth-service**: user registration, user login
- **listing-service**: property listing service
- **customer-support-service**: customer support service, handles simple complaint management and live chat

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
|---|---|
| web | 5173 |
| admin | 5174 |
| api-gateway | 8080 |
| auth-service | 8081 |
| listings-service | 8082 |
| chat-service | 8083 |