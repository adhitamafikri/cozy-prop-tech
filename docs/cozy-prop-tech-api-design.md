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

### Auth Endpoints

**Base Path (Auth)**: `/api/v1/auth`

#### Guest Token Generation

- Method: `POST`
- Path: `/guest`
- Security Aspects:
  - Token Lifetime: 15 minutes
  - Rate Limiter: Per-IP limit, 4 requests per hour

**Response Body : 200 OK**

```typescript
type GuestTokenResponse = Response<{ token: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GuestTokenErrorResponse = ErrorResponse;
```

#### User Registration

- Method: `POST`
- Path: `/register`
- Auth: Bearer (guest JWT token)
- Security Aspects:
  - Need valid guest JWT token
  - Rate Limiter: Per-IP limit, 1 requests for 5 minutes

**Request Body**

```typescript
type RegistrationRequest = {
  name: string;
  email: string;
  phone: string;
  password: string;
};
```

**Response Body : 201 OK**

```typescript
type RegistrationResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type RegistrationErrorResponse = ErrorResponse;
```

#### User Login

- Method: `POST`
- Path: `/login`
- Auth: Bearer (guest JWT token)
- Security Aspects:
  - Need valid guest JWT token
  - Rate Limiter: Per-IP limit, 1 requests for 5 minutes
  - Access Token: 15 minutes expiry
  - Refresh Token: 24 hours expiry

**Request Body**

```typescript
type LoginRequest = {
  email: string;
  password: string;
};
```

**Response Header : 201 OK**

```typescript
type LoginResponseHeader = {
  access_token: string;
  refresh_token: string;
};
```

**Response Body : 201 OK**

```typescript
type LoginResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type LoginErrorResponse = ErrorResponse;
```

### Admin-Facing Endpoints

**Base Path**: `/api/v1/admin`

#### User Management

**Base Path**: `/api/v1/admin/users`

##### GET Users

- Method: `GET`
- Path: `/users`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type GetUsersRequest = PaginatedQueryParams;
```

**Response Body : 200 OK**

```typescript
type GetUsersRequest = PaginatedResponse<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetUsersError = ErrorResponse;
```

##### GET User By ID

- Method: `GET`
- Path: `/users/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of user
};
```

**Response Body : 200 OK**

```typescript
type GetUserById = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
  roles: {
    name: string;
    permissions: string[];
  }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetUserByIDError = ErrorResponse;
```

##### Create User

- Method: `POST`
- Path: `/users`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Robust request body validation
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Body**

```typescript
type CreateUserRequest = {
  name: string;
  email: string;
  phone: string;
  roles: {
    name: string;
    permissions: string[];
  }[];
};
```

**Response Body : 201 OK**

```typescript
type CreateUserResponse = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
  roles: {
    name: string;
    permissions: string[];
  }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type CreateUserError = ErrorResponse;
```

##### Update User

- Method: `PUT`
- Path: `/users/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Robust request body validation
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Body**

```typescript
type UpdateUserRequest = {
  name: string;
  email: string;
  phone: string;
  roles: {
    name: string;
    permissions: string[];
  }[];
};
```

**Response Body : 200 OK**

```typescript
type UpdateUserResponse = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
  roles: {
    name: string;
    permissions: string[];
  }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type UpdateUserError = ErrorResponse;
```

##### Delete User

- Method: `DELETE`
- Path: `/users/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**
```typescript
type RequestParams = {
  id: number
}
```

**Response Body : 200 OK**

```typescript
type DeleteUserResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type DeleteUserError = ErrorResponse;
```

#### Location Management

#### Properties Management

#### Listings Management

#### Listing Availability Management

#### Booking Management

### Customer-Facing Endpoints

### Security

- All endpoints should use JWT for authorizing user requests
- All the endpoints should be guarded by RBAC middleware
