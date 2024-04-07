# Project Layout

This document describes the layout of the project.

## Directory Structure

The project is structured as follows:

### adapter

This directory typically contains code that adapts the core business logic of your application to various interfaces
such as HTTP, gRPC, or a database. For example, it might contain your HTTP handlers or gRPC server definitions.

### app

This directory usually contains the core business logic of your application. It might contain domain models, use cases,
and interfaces that define the operations your application can perform.

### docs

This directory contains documentation for your project. This could include design documents, and user guides.

### entity

This directory typically contains the entities or domain models of your application. These are the objects that your
application manipulates. They encapsulate the most important business rules.

### pkg

This directory is often used in larger applications and open-source projects to contain public code that other
applications or services might want to use. It's a way of explicitly sharing code across multiple projects.

### test

This directory contains all the test files. It can include unit tests, integration tests, end-to-end tests, etc.
