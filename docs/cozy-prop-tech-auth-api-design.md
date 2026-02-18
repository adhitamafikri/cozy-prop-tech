# Auth Endpoints

**Base Path (Auth)**: `/api/v1/auth`

### Guest Token Generation

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

### User Registration

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

### User Login

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

