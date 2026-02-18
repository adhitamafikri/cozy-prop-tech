---
title: Cozy Prop Tech PRD (Product Design Documentation)
description: This is the comprehensive documentation on the functionalities that should be possessed by this platform
createdDate: 2026-02-17
updatedDate: 2026-02-17
version: 1.0.0
---

# Overview

This _Cozy Prop Tech_ project is a minimal-bootleg replica of the original Kozystay platform. This project is created to emulate the functionality of the original Kozystay platform, from managing properties, listings, to bookings. There will be 2 kind of users using this platform, _admins_ and _customers_ (owners, guests).

## Problem Statement

We have so much potential for the high-end property market in Indonesia, especially in the big cities such as Jakarta, Bandung, and Bali. There are also so many unused apartments bought by the people, but they are not actually live there.

The big idea is, we will manage this high-end properties end to end. From acquiring the properties, manage listings, implement pricing, booking mechanism, digital concierge, room services, maintenance, and maintaining the relationship with the cusstomers (owners and guests).

We have a problem though. In order to do all that things, we need to build a robust platform that is capable of handling those duties. The platform that we are going to build will mimic the largest high-end property management company in Indonesia, _Kozystay_. The first version of our product would be the "lite" version of it.

## User Personas

We have several type of users that would be involved in this product:

### Customers

There are 2 types of customers:

- **guests**: Users that will browse our listings and book for stays in one of our listings
- **owners**: Users that will partner with us in managing their property and listings end-to-end

### Admins

The operations would be in chaos if we have no people that oversees the operations in our platform. Managing customers, properties, listings, bookings, would be impossible without admins.

## Goals

> This section explains the boundaries of our projects to avoid scope-creep and limiting the expectations

The first version of this product would have a limited amount of functionalities, explained below

**Admin**

- Able to manage users
- Able to manage properties
- Able to manage listings
- Able to manage listing availability and pricings
- Able to manage booking

**Customers**

- Able to register and login to our platform
- Able to browse our listings
- Able to book for a stay in one of our listings at a time

## User Stories/Usecases

> This section contains the use cases that have to be fulfilled by the current version of our platform. Each user stories would be group by personas (customers and admins). Each of the user stories are also marked by priority codes -> P0: must have, P1: must have, P2: nice to have

### Admins

1. **[P0]** As an admin, I need to be able to login to the admin platform
2. **[P0]** As an admin, I need to be able to manage registered users data in our platform
3. **[P0]** As an admin, I need to be able to manage locations data to be used in the platform
4. **[P0]** As an admin, I need to be able to manage properties in our platform
5. **[P0]** As an admin, I need to be able to manage listings in our platform
6. **[P0]** As an admin, I need to be able to manage listing availability in our platform
7. **[P2]** As an admin, I need to be able to manage bookings made by our customers in the platform

### Customers (owners and customers)

1. **[P0]** As a customer, I need to be able to register to the platform
2. **[P0]** As a customer, I need to be able to login to the platform
3. **[P0]** As a customer, I need to be able to access the home page of the platform
4. **[P0]** As a customer, I need to be able search for listings from the home page
5. **[P0]** As a customer, I need to be able search for listings from the search page
6. **[P0]** As a customer, I need to be able see the selected listing details on the listing details page
7. **[P2]** As a customer, I need to be able to make a booking from the listing details page

## Related Documents

- [Rough System Design](./system-design/rough-system-design.png)
- [Rough ERD](./system-design/erd.png)
- [API Design Doc](./cozy-prop-tech-api-design.md)
