Client / User always interacts with the Presentation Layer

Controllers in the presentation layer for REST API structure
The business logic for cl

3 tier architecture (Top layers can access the ones below them - constraint of layered architecture)
Presentation (Controllers)
    ↓
Business Logic (Semantics)
    ↓
Data Access Layer (Database Interactions)
    ↓
Database

Problems:
1. Creates a transitive dependency between Presenation layer and the DAL and some DAL logic might end up in the layers above.
2. Business logic becomes convolutes with the DAL with time.

Clean Architecture
Presentation
    ↓
Application (Use Cases / Features of the system) → |
    ↓          | →  Business Logic
  Domain     → |

