# Cozy Prop Tech

A project for me to learn Fullstack on the property technology and management sector

[https://github.com/adhitamafikri/cozy-prop-tech](https://github.com/adhitamafikri/cozy-prop-tech)

---

## Project Structure

This project is basically a monorepo containing several distinct sub-projects with this structure

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

## Technology Stacks

### Frontend

| Category | Technology | Notes |
|----------|------------|-------|
| User Site | bun, Vite Vue 3, Pinia, TypeScript, TailwindCSS | Customer-facing |
| Admin Site | bun, Vite, Vue 3, Pinia, TypeScript, TailwindCSS | Admin-facing |
| Maps | Leaflet, OpenStreetMap | |
| Forms & Validation | VeeValidate | |
| Unit Testing | Vitest, @testing-library/vue | |
| E2E Testing | Playwright | |

### Backend

| Category | Technology | Notes |
|----------|------------|-------|
| Language | Go v1.2 | |
| Routing | Gin | github.com/gin-gonic/gin |
| Database | PostgreSQL v17, sqlx | |
| Migrations | Golang Migrate | Run with Docker |
| Caching | Redis v8, Go-Redis | |
| Logging | slog | |
| Auth | JWT | |
| Live Reload | Air | |
| Reverse Proxy | nginx | |
| Architecture | Hexagonal | |

### Infra CI/CD

| Category | Technology |
|----------|------------|
| Containerization | Docker, Docker Compose |
| Secrets Management | Infisical |
| CI/CD | GitHub Actions |

## Project Notes

**Documentation Naming Convention:**

- **ALWAYS use YYYY-MM-DD-slug.md format** for all documentation files
- Get current date from system: `date +%Y-%m-%d`
- Never assume or hardcode dates
- Use lowercase with hyphens for slugs (e.g., `2025-11-05-feature-implementation.md`)
- Examples:
  - ✅ `2025-11-07-docker-compose-setup.md`
  - ✅ `2025-10-22-ui-audit-results.md`
  - ❌ `setup.md` (missing date)
  - ❌ `IMPROVEMENTS.md` (missing date, uppercase)
  - ❌ `FINAL_AUDIT_2025-10-22.md` (date in wrong position)

**Documentation Lives Everywhere:**
Documentation is NOT only in `notes/` or `docs/` directories. Search comprehensively:

- **Workspace root**: `README.md`, `AGENTS.md`
- **Workspace notes**: `notes/*.md` (project-wide documentation including local development, setup, architecture)
- **Backend**: `backend/**/*.md`
- **Frontend**: `frontend/**/*.md`
- **In-code**: Comments, JSDoc, docstrings (implementation details)

**Before researching, search ALL locations:**

Find all markdown files excluding dependencies (node_modules, .nuxt).
Search across all documentation for relevant topics.

## Trust Code, Not Docs

Code is truth. Documentation may lag behind. When in doubt:

1. Read the actual implementation
2. Check git history
3. Test the behavior
4. Update docs if outdated

## Working Style

**You're a SENIOR ENGINEER with full autonomy**

- Research before implementing
- Follow existing patterns
- Test locally first

**When adding features**

- Check for similar implementations first
- Follow proper Go lang conventions for backend
- Follow proper Typescript and Vue 3 conventions for frontend
- Verify role-based access control

## Running Project Locally

We have multiple services here in this project. In order to know about which port the service is running on, consult to the table below:

| Services | Ports |
| -------- | ----- |
| web      | 5173  |
| admin    | 5174  |
| api      | 8082  |
