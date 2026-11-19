# Go RBAC Authorization System

A production-oriented **Role-Based Access Control (RBAC) system** built with **Go, Fiber v3, GORM, PostgreSQL, and JWT**.

The project provides a complete authentication and authorization foundation for Go applications, including user management, JWT authentication, roles, permissions, role-based access control, authorization middleware, secure password handling, and protected REST APIs.

The architecture is designed to be modular, testable, and reusable across different backend applications.

## Core Features

- User registration and authentication
- Secure password hashing and verification
- JWT access and refresh tokens
- JWT validation and expiration handling
- User management
- Role management
- Permission management
- User-to-role relationships
- Role-to-permission relationships
- Role-based authorization
- Permission-based authorization
- Fiber v3 authentication middleware
- Fiber v3 authorization middleware
- Protected REST API endpoints
- GORM-based database persistence
- PostgreSQL support
- Database migrations
- Transaction-safe operations
- Request validation
- Centralized error handling
- Unit and integration tests
- Docker support
- Configuration through environment variables

## Authorization Model

The system follows a standard RBAC model:

```text
User
  |
  | N:M
  v
Role
  |
  | N:M
  v
Permission
```

A user can have multiple roles, and each role can contain multiple permissions.

Example:

```text
Mohammad
   |
   +── Admin
   |     ├── user.read
   |     ├── user.create
   |     ├── user.update
   |     └── user.delete
   |
   +── Developer
         ├── project.read
         └── project.update
```

## Authentication Flow

```text
Client
  |
  | POST /auth/login
  v
Fiber v3
  |
  v
Auth Service
  |
  +── Verify Password
  |
  +── Generate JWT
  |
  v
Access Token + Refresh Token
```

Authenticated requests then follow:

```text
Client
  |
  | Authorization: Bearer <JWT>
  v
Fiber v3
  |
  v
JWT Middleware
  |
  +── Validate Signature
  +── Validate Expiration
  +── Extract User ID
  |
  v
Authorization Middleware
  |
  +── Resolve Roles
  +── Resolve Permissions
  |
  v
Protected Handler
```

## Example

An endpoint can require the `user.read` permission:

```text
GET /api/users

Required Permission:
    user.read
```

A user with the following role:

```text
Role: Admin

Permissions:
    user.read
    user.create
    user.update
    user.delete
```

is authorized to access the endpoint.

A user without the required permission receives:

```text
403 Forbidden
```

## Technology Stack

| Component | Technology |
|---|---|
| Language | Go |
| HTTP Framework | Fiber v3 |
| ORM | GORM |
| Database | PostgreSQL |
| Authentication | JWT |
| API | REST |
| Containerization | Docker |
| Testing | Go testing |

## Architecture

The application separates HTTP handling, authentication, authorization, business logic, and persistence.

```text
                    ┌─────────────────┐
                    │     Client      │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │   Fiber v3      │
                    │    Router       │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │ Authentication  │
                    │    Middleware   │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │ JWT Validation  │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │ Authorization   │
                    │    Middleware   │
                    └────────┬────────┘
                             │
                    ┌────────┴────────┐
                    │                 │
                    v                 v
                  Roles          Permissions
                    │                 │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │    Handler      │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │    Service      │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │   Repository    │
                    │      GORM       │
                    └────────┬────────┘
                             │
                             v
                    ┌─────────────────┐
                    │   PostgreSQL    │
                    └─────────────────┘
```

## Project Goals

The main goal is to provide a clean and reusable implementation of authentication and authorization for Go applications.

The project focuses on:

- Clean architecture
- Separation of authentication and authorization
- Reusable Fiber middleware
- Explicit RBAC relationships
- Secure credential handling
- JWT-based stateless authentication
- Database consistency
- Transaction safety
- Testability
- Maintainability
- Extensibility

## Potential Use Cases

This system can be used as a foundation for:

- SaaS applications
- Admin dashboards
- Enterprise applications
- Banking and financial systems
- Internal company platforms
- Management systems
- REST APIs
- Microservices
- Multi-user applications

## Future Extensions

The architecture can be extended with:

- Multi-tenancy
- Organization-based authorization
- Resource-level permissions
- Attribute-Based Access Control (ABAC)
- OAuth2 / OpenID Connect
- API keys
- Token revocation
- Session management
- Audit logging
- Login history
- Account lockout
- Two-factor authentication
- Policy-based authorization
- Casbin integration

## License

Open-source project intended as a reusable reference implementation for building secure authentication and RBAC authorization systems with Go and Fiber v3.
