# Hexagonal Architecture Directory Structure Analysis

**Date:** 2026-02-17

## Overview

This document analyzes the current hexagonal architecture scaffolding in `backend/api/internal/` and compares it with professional-production-grade Go projects.

---

## Current Structure

```
internal/
├── config/       ✅ (infrastructure)
├── delivery/     ✅ (adapters)
├── entity/       ✅ (domain/core)
├── repository/   ✅ (ports)
└── usecase/      ✅ (application)
```

**Assessment:** Good starting point, but missing key elements for production-grade code.

---

## What is Missing

| Aspect | Current | Professional |
|--------|---------|--------------|
| Domain grouping | Flat `entity/` | Organized by domain: `domain/user/`, `domain/property/` |
| Ports/Interfaces | In `repository/` | Explicit `repository/port/` for interfaces |
| DTOs/Models | Mixed in delivery | Separate `delivery/dto/`, `delivery/request/`, `delivery/response/` |
| Domain Services | Missing | `domain/service/` for cross-entity logic |
| Value Objects | Missing | `domain/valueobject/` for validated types |
| Dependency Injection | Manual in main.go | `di/` container or wire |
| Error Handling | Likely inline | `errors/` package with custom error types |

---

## Recommended Professional Structure

```
internal/
├── config/                    # Infrastructure config (DB, Redis, Logger)
├── delivery/
│   ├── dto/                  # Data Transfer Objects
│   ├── handler/              # HTTP handlers/controllers (Gin)
│   ├── middleware/           # Gin middleware (auth, logging)
│   └── presenter/            # Response formatters
├── di/                       # Dependency injection container
├── domain/
│   ├── user/
│   │   ├── entity.go         # User struct
│   │   ├── valueobject/     # Email, Password, etc
│   │   ├── repository.go    # Interface (port)
│   │   └── service.go        # Domain logic
│   ├── property/
│   ├── listing/
│   └── booking/
├── repository/
│   ├── postgres/             # sqlx implementations
│   ├── cache/                # Redis implementations
│   └── port/                 # Explicit interfaces
└── usecase/
    ├── user/
    ├── property/
    └── booking/
```

---

## Key Concepts from Clean Architecture

### The Dependency Rule
- Source code dependencies must point **inwards**
- Inner layers (domain/usecase) know nothing about outer layers (delivery/repository)
- Use interfaces (ports) to invert dependencies

### Layers Explained

1. **Domain ( innermost)** - Entities, Value Objects, Domain Services
   - Pure business logic, no external dependencies
   - Most stable, least likely to change

2. **Use Case (Application)** - Orchestrates domain objects
   - Implements application-specific business rules
   - Coordinates flow of data

3. **Delivery (Adapters)** - HTTP, gRPC, CLI handlers
   - Converts external requests to internal formats
   - Converts internal responses to external formats

4. **Repository (Ports)** - Interfaces for data access
   - Defines contracts, not implementations
   - Allows swapping data sources

---

## Quality References

### Articles
- **Uncle Bob - The Clean Architecture**: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- **Alistair Cockburn - Hexagonal Architecture**: https://alistair.cockburn.us/hexagonal-architecture/
- **Package-Oriented Design**: https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md

### Videos
- **GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps**: https://www.youtube.com/watch?v=oL6JBUk6tj0
- **GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming**: https://www.youtube.com/watch?v=PTE4VJIdHPg

### Project References
- **go-kit**: https://github.com/go-kit/kit
- **ent**: https://github.com/ent/ent

---

## Comparison with khannedy/golang-clean-architecture

This repository follows textbook Clean Architecture:

| Aspect | khannedy/golang-clean-architecture | Notes |
|--------|-----------------------------------|-------|
| Layer naming | Delivery, UseCase, Repository, Entity | Standard terminology |
| DDD Elements | Basic (Entity only) | No Ubiquitous Language, Bounded Contexts, Aggregates |
| Focus | Dependency separation | Not full DDD |

**Verdict:** This is Clean Architecture (Hexagonal), NOT Domain Driven Design.

---

## Next Steps

1. Consider restructuring by domain (user/, property/, listing/)
2. Add explicit port interfaces in domain layer
3. Add DTOs for clean delivery/domain separation
4. Implement dependency injection for testability

---

## Visual Diagram

See `hexagonal-architecture.svg` for the visual representation of this architecture.
