# PROJECT PROGRESS
This document is showing current project status and it's plan.

---

## Current Main Objectives
- [ ] Implemented authority domain.
- [ ] Implemented role domain.
- [ ] Implemented user_auth domain.

> [!NOTE]
> Cleared only if all layer has been completed. Included it's variant, 
> such as memory repository, sqlite repository, etc.

---

## Current Task
Specific task for each domain.

### BLOCKER
Need to added first.
- [x] Config and it's loader.

### Authority domain
- [x] Add domain error.
- [x] Entity

- Repository
    - Memory
        - [x] Implement GetByID method.
        - [x] Implement GetByIDs method.
        - [x] Implement Create method.
        - [x] Implement Update method.
        - [x] Implement DeleteByID method.
        - [x] Add test file
    - Postgres
        - [ ] Implement GetByID method.
        - [ ] Implement GetByIDs method.
        - [ ] Implement Create method.
        - [ ] Implement Update method.
        - [ ] Implement DeleteByID method.
        - [ ] Add test file
    - Sqlite
        - [ ] Implement GetByID method.
        - [ ] Implement GetByIDs method.
        - [ ] Implement Create method.
        - [ ] Implement Update method.
        - [ ] Implement DeleteByID method.
        - [ ] Add test file
    - [ ] Make DDL/Migrations structure

- Service
    - [x] Add GetByID method.
    - [x] Add GetByIDs method.
    - [x] Add Create method.
    - [x] Add Update method.
    - [x] Add DeleteByID method.

- HTTP Handler
    - [x] Add GetByID method.
    - [ ] Add GetByIDs method.
    - [x] Add Create method.
    - [ ] Add Update method.
    - [ ] Add DeleteByID method.

---

## Extra Task
This task is optionally, but recomended to be cleared,
due to performance and cleaner code.

- [ ] Consider returning copy for small data.
    It's mainly http request - response, normally the object lifetime 
    is short. So why returning reference for that(used heap)? 
    It's add job for GC later.
- [ ] Consider make handler for all error.
    Research about error that might causing panic.
    Location: handler.go:Pagination
