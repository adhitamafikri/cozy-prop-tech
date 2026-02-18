---
title: Cozy Prop API Design
description: Contains the documentation API endpoints along with their request and response schema
updatedDate: 2026-02-17
---

# Overview

This _Cozy Prop Tech_ project is a minimal-bootleg replica of the original Kozystay platform. This project is created to emulate the functionality of the original Kozystay platform, from managing properties, listings, to bookings. There will be 2 kind of users using this platform, _admins_ and _customers_ (owners, guests).

## API Endpoints Design

> This section contains the full endpoint along with the request and response schema design

There will be 2 types of endpoints, _user-facing_ and _admin_ endpoints

- **Base URL**: `http://api.cozy-prop.local:8082`
- **Base Path**: `/api/v1`
- **Base Path (Admin)**: `/api/v1/admin`
- **Base Path (Auth)**: `/api/v1/auth`

### Base Schemas

> This section contains base schemas that could be reused across the endpoints

#### Standard Paginated Request Query Params

```typescript
type PaginatedQueryParams = {
  q?: string;
  page?: number;
  per_page?: number;
  sort_by?: string;
  order?: "asc" | "desc";
};
```

#### Standard Response Body Schema

```typescript
type Response<T> = {
  data: T;
};
```

#### Paginated Response Body Schema

```typescript
type PaginatedResponse<T> = {
  data: T[];
  meta: {
    current_page: number;
    per_page: number;
    total: number;
    last_page: number;
    from: number;
    to: number;
  };
};
```

#### Error Response Body Schema

```typescript
type ErrorResponse = {
  error: string;
};
```

### Detailed API Design Documents

> We have extensive list of APIs, ranging from auth, admin-facing APIs, and customer-facing APIs. We split them into these 3 documents

1. **Auth API Design**: [Auth API Design](./cozy-prop-tech-auth-api-design.md)
2. **Admin API Design**: [Admin API Design](./cozy-prop-tech-admin-api-design.md)
3. **Customer API Design**: [Customer API Design](./cozy-prop-tech-customer-api-design.md)

### Security

- All endpoints should use JWT for authorizing user requests
- All the endpoints should be guarded by RBAC middleware
