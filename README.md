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

## Development Process

This project supports AI-assisted engineering with **Opencode** and **Claude Code**

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

| Services                 | Ports |
| ------------------------ | ----- |
| web                      | 5173  |
| admin                    | 5174  |
| api                      | 8082  |

## Reading Links

- [https://github.com/khannedy/golang-clean-architecture/blob/main](https://github.com/khannedy/golang-clean-architecture/blob/main)
- [https://betterstack.com/community/guides/logging/logging-in-go/](https://betterstack.com/community/guides/logging/logging-in-go/)
