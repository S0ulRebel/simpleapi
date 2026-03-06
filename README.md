Overview

This project is a small REST API built with Go and Gin.
It was originally created as a technical showcase during a hiring process.

Since most of the production code I have written is under NDA, this repository serves as a public example of how I structure backend services and APIs.

Although the scope is intentionally simple, the project demonstrates:

layered architecture

repository pattern for database abstraction

authentication flow

configurable database backend

basic unit testing

Technical Notes

This was my first project using Go and the Gin framework, built under a short deadline as part of a technical assignment. Despite the small scope, I aimed to structure the code as I would a production backend.

Key architectural decisions:

Repository Pattern

The API uses a repository abstraction to decouple business logic from the database implementation.

This allows switching the database engine via configuration.

Supported databases:

PostgreSQL

MongoDB

The database type can be configured in the .env file.

Authentication Flow

Two endpoints are public:

CreateUser – user registration

Login

All other endpoints require authentication.

After authentication, users can:

fetch users

update users

delete users

create posts

For simplicity (since this is a technical test), the project does not implement roles or complex permission systems.

Testing

A small set of unit tests is included for one service to demonstrate testing structure.
Since this is a demo project rather than production code, test coverage is intentionally limited.

API Testing

A Postman collection is included in the project root for easier API testing.
