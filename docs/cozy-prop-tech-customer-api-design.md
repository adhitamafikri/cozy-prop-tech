# Customer-Facing API Design

**Base Path**: `/api/v1`

## Users

**Base Path**: `/api/v1/users`

### Get Current User

- Method: `GET`
- Path: `/users/me`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Response Body : 200 OK**

```typescript
type GetCurrentUserResponse = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
  updated_at: string;
  roles: { id: number; name: string }[] | null;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetCurrentUserError = ErrorResponse;
```

### Update Profile

- Method: `PUT`
- Path: `/users/me`
- Auth: Bearer (JWT access token)
- Security Aspects:
  - Need valid JWT access token
  - Robust request body validation
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute
  - Email cannot be changed

**Request Body**

```typescript
type UpdateProfileRequest = {
  name: string;
  phone: string;
};
```

**Response Body : 200 OK**

```typescript
type UpdateProfileResponse = Response<{
  id: number;
  name: string;
  email: string;
  phone: string;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type UpdateProfileError = ErrorResponse;
```

---

## Locations

**Base Path**: `/api/v1/locations`

### Get Locations

- Method: `GET`
- Path: `/locations`
- Auth: Bearer (guest JWT token or user JWT token)
- Security Aspects:
  - Accessible to guests (with guest JWT) and logged-in users
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type GetLocationsRequest = PaginatedQueryParams & {
  category?: "country" | "province" | "city" | "district" | "neighborhood";
};
```

**Response Body : 200 OK**

```typescript
type GetLocationsResponse = PaginatedResponse<{
  id: number;
  name: string;
  category: "country" | "province" | "city" | "district" | "neighborhood";
  description: string | null;
  latitude: number | null;
  longitude: number | null;
  parent_id: number | null;
  created_at: string;
  updated_at: string;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetLocationsError = ErrorResponse;
```

---

## Searches (search for listings)

**Base Path**: `/api/v1/search`

### Search Listings

- Method: `GET`
- Path: `/search`
- Auth: Bearer (guest JWT token or user JWT token)
- Security Aspects:
  - Accessible to guests (with guest JWT) and logged-in users
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Query Params**

```typescript
type SearchListingsRequest = PaginatedQueryParams & {
  city_id?: number;
  district_id?: number;
};
```

**Response Body : 200 OK**

```typescript
type SearchListingsResponse = PaginatedResponse<{
  id: number;
  title: string;
  description: string | null;
  base_price: number;
  status: "active" | "inactive";
  guest_capacity: number;
  minimum_reservation_nights: number;
  maximum_reservation_nights: number;
  cleaning_fee: number | null;
  extra_guest_capacity: number;
  extra_guest_fee: number | null;
  num_beds: number;
  num_bathrooms: number;
  num_bedrooms: number;
  amenities: Record<string, any> | null;
  created_at: string;
  updated_at: string;
  property: {
    id: number;
    name: string;
    address: string;
    latitude: number;
    longitude: number;
  } | null;
  image: { url: string } | null;
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type SearchListingsError = ErrorResponse;
```

---

## Listings

**Base Path**: `/api/v1/listings`

### Get Listing Details

- Method: `GET`
- Path: `/listings/:id`
- Auth: Bearer (guest JWT token or user JWT token)
- Security Aspects:
  - Accessible to guests (with guest JWT) and logged-in users
  - Rate Limiter: Per-IP limit, 60 requests per 1 minute

**Request Params**

```typescript
type RequestParams = {
  id: number; // listing ID
};
```

**Response Body : 200 OK**

```typescript
type GetListingDetailsResponse = Response<{
  id: number;
  title: string;
  description: string | null;
  base_price: number;
  status: "active" | "inactive";
  guest_capacity: number;
  minimum_reservation_nights: number;
  maximum_reservation_nights: number;
  cleaning_fee: number | null;
  extra_guest_capacity: number;
  extra_guest_fee: number | null;
  num_beds: number;
  num_bathrooms: number;
  num_bedrooms: number;
  amenities: Record<string, any> | null;
  created_at: string;
  updated_at: string;
  property: {
    id: number;
    name: string;
    address: string;
    latitude: number;
    longitude: number;
    location: { id: number; name: string } | null;
  } | null;
  images: {
    id: number;
    url: string;
    is_primary: boolean;
    sort_order: number;
  }[];
}>;
```

**Error Response Body : 4xx, 5xx**

```typescript
type GetListingDetailsError = ErrorResponse;
```

---

## Booking (P2 - Nice to Have)

**Base Path**: `/api/v1/bookings`
