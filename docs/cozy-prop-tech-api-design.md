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

---

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

  **Request Params**

```typescript
type RequestParams = {
  id: number;
};
```

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
  id: number;
};
```

**Response Body : 200 OK**

```typescript
type DeleteUserResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type DeleteUserError = ErrorResponse;
```

#### Assign Role to User

- Method: `POST`
- Path: `/users/:id/roles`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of role
};
```

**Request Body**

```typescript
type AssignRoleToUserRequest = {
  role_ids: number[];
};
```

**Response Body : 200 OK**

```typescript
type AssignRoleToUserResponse = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: timestamp;
  updated_at: timestamp;
  roles: { id: number; name: string }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type AssignRoleToUserError = ErrorResponse;
```

##### GET Me (currently login admin data)

- Method: `GET`
- Path: `/users/me`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Response Body : 200 OK**

```typescript
type GetMeResponse = Response<{
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
type GetMeError = ErrorResponse;
```

---

#### Role Management

**Base Path**: `/api/v1/admin/roles`

##### Get Roles

- Method: `GET`
- Path: `/roles`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type GetRolesRequest = PaginatedQueryParams;
```

**Response Body : 200 OK**

```typescript
type GetRolesResponse = PaginatedResponse<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetRolesError = ErrorResponse;
```

##### Get Role By ID

- Method: `GET`
- Path: `/roles/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of role
};
```

**Response Body : 200 OK**

```typescript
type GetRoleByIDResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetRoleByIDError = ErrorResponse;
```

##### Create Role

- Method: `POST`
- Path: `/roles`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Body**

```typescript
type CreateRoleRequest = {
  name: string;
  description: string;
};
```

**Response Body : 200 OK**

```typescript
type CreateRoleResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type CreateRoleError = ErrorResponse;
```

##### Update Role

- Method: `PUT`
- Path: `/roles/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of role
};
```

**Request Body**

```typescript
type UpdateRoleRequest = {
  name: string;
  description: string;
};
```

**Response Body : 200 OK**

```typescript
type UpdateRoleResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type UpdateRoleError = ErrorResponse;
```

##### Delete Role

- Method: `DELETE`
- Path: `/roles/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute
  - Should return error if the role is used

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of role
};
```

**Response Body : 200 OK**

```typescript
type DeleteRoleResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type DeleteRoleError = ErrorResponse;
```

#### Assign Permission to Role

- Method: `POST`
- Path: `/roles/:id/permissions`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of role
};
```

**Request Body**

```typescript
type AssignPermissionToRoleRequest = {
  permission_ids: number[];
};
```

**Response Body : 200 OK**

```typescript
type AssignPermissionToRoleResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
  permissions: { id: number; name: string }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type AssignPermissionToRoleError = ErrorResponse;
```

---

#### Permission Management

**Base Path**: `/api/v1/admin/permissions`

##### Get Permission List

- Method: `GET`
- Path: `/permissions`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type GetPermissionRequest = PaginatedQueryParams;
```

**Response Body : 200 OK**

```typescript
type GetPermissionResponse = PaginatedResponse<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetPermissionError = ErrorResponse;
```

##### Get Permission By ID

- Method: `GET`
- Path: `/permissions/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of permission
};
```

**Response Body : 200 OK**

```typescript
type GetPermissionByIDResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetPermissionByIDError = ErrorResponse;
```

##### Create Permission

- Method: `POST`
- Path: `/permissions`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Body**

```typescript
type CreatePermissionRequest = {
  name: string;
  description: string;
};
```

**Response Body : 200 OK**

```typescript
type CreatePermissionResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type CreatePermissionError = ErrorResponse;
```

##### Update Permission

- Method: `PUT`
- Path: `/permissions/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of permisssion
};
```

**Request Body**

```typescript
type UpdatePermissionRequest = {
  name: string;
  description: string;
};
```

**Response Body : 200 OK**

```typescript
type UpdatePermissionResponse = Response<{
  id: number;
  name: string;
  description: string;
  created_at: timestamp;
  updated_at: timestamp;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type UpdatePermissionError = ErrorResponse;
```

##### Delete Permission

- Method: `DELETE`
- Path: `/permissions/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute
  - Should return error if the permission is used

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of permission
};
```

**Response Body : 200 OK**

```typescript
type DeletePermissionResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type DeletePermissionError = ErrorResponse;
```

---

#### Location Management

**Base Path**: `/api/v1/admin/locations`

##### Get Locations

- Method: `GET`
- Path: `/locations`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type GetLocationsRequest = PaginatedQueryParams;
```

**Response Body : 200 OK**

```typescript
type GetLocationsResponse = PaginatedResponse<{
  id: number;
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetLocationsError = ErrorResponse;
```

##### Get Location By ID

- Method: `GET`
- Path: `/locations/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of location
};
```

**Response Body : 200 OK**

```typescript
type GetLocationByIDResponse = Response<{
  id: number;
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetLocationByIDError = ErrorResponse;
```

##### Create Location

- Method: `POST`
- Path: `/locations`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Robust request body validation
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Body**

```typescript
type CreateLocationRequest = {
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
};
```

**Response Body : 201 OK**

```typescript
type CreateLocationResponse = Response<{
  id: number;
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type CreateLocationError = ErrorResponse;
```

##### Update Location

- Method: `PUT`
- Path: `/locations/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Robust request body validation
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of location
};
```

**Request Body**

```typescript
type UpdateLocationRequest = {
  name?: string;
  category?: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
};
```

**Response Body : 200 OK**

```typescript
type UpdateLocationResponse = Response<{
  id: number;
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description?: string;
  latitude?: number;
  longitude?: number;
  parent_id?: number;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type UpdateLocationError = ErrorResponse;
```

##### Delete Location

- Method: `DELETE`
- Path: `/locations/:id`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute
  - Should return error if the location is used by other entities (properties, listings)

**Request Params**

```typescript
type RequestParams = {
  id: number; // id of location
};
```

**Response Body : 200 OK**

```typescript
type DeleteLocationResponse = Response<{ message: string }>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type DeleteLocationError = ErrorResponse;
```

---

#### Properties Management

**Base Path**: `/api/v1/admin/properties`

---

#### Listings Management

**Base Path**: `/api/v1/admin/listings`

---

#### Listing Availability Management

**Base Path**: `/api/v1/admin/listing-availability`

---

#### Booking Management

**Base Path**: `/api/v1/admin/bookings`

---

### Customer-Facing Endpoints

**Base Path**: `/api/v1`

#### Users

**Base Path**: `/api/v1/users`

#### Locations

**Base Path**: `/api/v1/locations`

#### Searches (search for listings)

**Base Path**: `/api/v1/search`

#### Listings

**Base Path**: `/api/v1/listings`

#### Bookings

**Base Path**: `/api/v1/bookings`

---

### Security

- All endpoints should use JWT for authorizing user requests
- All the endpoints should be guarded by RBAC middleware

---
