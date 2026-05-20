# Widyanaya API

Widyanaya API is a backend REST API built with Go to help researchers better understand and manage the publication process.

The project is designed to support research and publication workflows by providing a structured backend system for handling publication-related data and researcher interactions.

---

## Purpose

Widyanaya API aims to help researchers by:

* Simplifying publication management
* Organizing research-related data
* Supporting publication workflows
* Providing a backend system for research platforms

---

## Tech Stack

* Go (Golang)
* Supabase

---

## Concepts Used

* Backend API development
* CRUD operations
* Layered architecture
* Modular backend structure
* Database integration
* Environment configuration

---

## Main Features

* Research publication management
* User management
* Publication-related operations
* Structured API endpoints
* Database-backed storage

---

## Project Structure

```txt
cmd/         → application entry point
config/      → application configuration
database/    → database connection setup
helper/      → helper utilities
internal/    → main application modules
server/      → server initialization
```

---

## Running The Project

### Install dependencies

```bash
go mod tidy
```

### Run the server

```bash
go run cmd/widyanaya/main.go
```

---
