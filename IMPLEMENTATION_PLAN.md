# University CMS Backend Plan

## Goal

Finish a usable CMS backend for university information in 7 days by extending the current Go API from a single `posts` module into a small, focused admin/public content system.

This plan is based on the current repo state:

- Gin router in [internal/router/router.go](/home/student/Desktop/CIIP-api-go/internal/router/router.go:1)
- Handler/service/repository layering for `posts`
- PostgreSQL with migrations
- Minimal `posts` content model in [internal/models/post.go](/home/student/Desktop/CIIP-api-go/internal/models/post.go:1)

## Delivery Strategy

Do not try to build a full university ERP. Ship a narrow CMS MVP:

- Admin authentication
- University CRUD
- Faculty CRUD
- Program CRUD
- News/announcements CRUD
- Public read endpoints
- Basic file/image upload if time allows

Everything else is optional and should be cut if it threatens delivery.

## MVP Scope

### Must Have

- `POST /auth/login`
- Admin user model and password-based login
- CRUD for `universities`
- CRUD for `faculties`
- CRUD for `programs`
- CRUD for `posts` or rename `posts` to `news`
- Public endpoints for published data
- Input validation
- Migrations for all core tables
- Basic tests for critical flows

### Nice To Have

- Image upload
- `departments` table
- Filtering and pagination
- Swagger docs
- Seed script
- Soft deletes

### Out of Scope For This Week

- Complex role systems
- Rich text editor support on backend
- Full audit logs
- Notifications
- Search engine integration
- Multi-language support
- Analytics dashboard

## Recommended Data Model

Start small and relational:

- `users`
  - `id`, `email`, `password_hash`, `role`, `created_at`
- `universities`
  - `id`, `name`, `slug`, `description`, `address`, `phone`, `email`, `website`, `is_published`, `created_at`, `updated_at`
- `faculties`
  - `id`, `university_id`, `name`, `slug`, `description`, `is_published`, `created_at`, `updated_at`
- `programs`
  - `id`, `faculty_id`, `name`, `degree`, `duration_years`, `description`, `tuition_fee`, `is_published`, `created_at`, `updated_at`
- `posts`
  - `id`, `university_id`, `title`, `content`, `is_published`, `created_at`, `updated_at`

If you have extra time, add:

- `departments`
- `media`

## API Shape

Keep admin and public routes separate.

### Admin Routes

- `POST /auth/login`
- `GET /admin/universities`
- `POST /admin/universities`
- `GET /admin/universities/:id`
- `PUT /admin/universities/:id`
- `DELETE /admin/universities/:id`
- `GET /admin/faculties`
- `POST /admin/faculties`
- `PUT /admin/faculties/:id`
- `DELETE /admin/faculties/:id`
- `GET /admin/programs`
- `POST /admin/programs`
- `PUT /admin/programs/:id`
- `DELETE /admin/programs/:id`
- `GET /admin/posts`
- `POST /admin/posts`
- `PUT /admin/posts/:id`
- `DELETE /admin/posts/:id`

### Public Routes

- `GET /universities`
- `GET /universities/:id`
- `GET /universities/:id/faculties`
- `GET /faculties/:id/programs`
- `GET /posts`
- `GET /posts/:id`

## Code Structure Plan

Keep the current architecture and repeat it for each module:

- `internal/models`
- `internal/repository`
- `internal/service`
- `internal/handler`
- `internal/router`

Suggested additions:

- `internal/middleware/auth.go`
- `internal/models/user.go`
- `internal/handler/auth_handler.go`
- `internal/service/auth_service.go`
- `internal/repository/user_repository.go`

Build one module fully, then copy the pattern. Do not over-engineer shared abstractions this week.

## 7-Day Execution Plan

### Day 1: Freeze Scope and Fix Foundation

Deliverables:

- Finalize MVP entities and routes
- Create migrations for `users`, `universities`, `faculties`, `programs`
- Decide whether `posts` stays as `posts` or becomes `news`
- Fix current schema/code issues before adding features

Tasks:

- Review and correct the existing `posts` migration
- Align SQL field names with Go model fields
- Add timestamps and `is_published` fields where needed
- Standardize JSON binding and validation tags in models
- Create `.env.example`

Important current issues to fix first:

- [migrations/000001_create_posts_table.up.sql](/home/student/Desktop/CIIP-api-go/migrations/000001_create_posts_table.up.sql:1) has a trailing comma after `content TEXT NOT NULL`
- [internal/repository/post_repository.go](/home/student/Desktop/CIIP-api-go/internal/repository/post_repository.go:1) inserts into `context` but selects `content`
- [internal/repository/post_repository.go](/home/student/Desktop/CIIP-api-go/internal/repository/post_repository.go:1) does not check `rows.Scan` errors

### Day 2: Authentication

Deliverables:

- Admin login working
- Auth middleware protecting admin routes

Tasks:

- Add `users` table migration
- Add password hashing with `bcrypt`
- Add JWT login flow or simple token auth
- Create seeded admin user
- Protect `/admin/*` routes with middleware

Minimum acceptable result:

- One admin can log in and use protected CRUD endpoints

### Day 3: Universities Module

Deliverables:

- Full university CRUD

Tasks:

- Add `University` model
- Add repository methods: create, list, get by id, update, delete
- Add service validation
- Add admin handlers and routes
- Add public list/detail endpoints for published universities

Definition of done:

- You can create, edit, list, and fetch universities from Postman

### Day 4: Faculties and Programs

Deliverables:

- Full faculty CRUD
- Full program CRUD

Tasks:

- Add foreign keys
- Build faculty module
- Build program module
- Add list filters by `university_id` and `faculty_id`
- Return clean API errors for missing parent records

Definition of done:

- One university can have many faculties
- One faculty can have many programs

### Day 5: Posts/News and Public Read API

Deliverables:

- News/announcement management
- Public content endpoints usable by frontend

Tasks:

- Upgrade current `posts` module to include publication fields
- Attach posts to universities if needed
- Add public-only published filters
- Add consistent response structure
- Add pagination to list endpoints if easy

If time remains:

- Add upload endpoint for thumbnails or attachments

### Day 6: Testing and Stabilization

Deliverables:

- Critical paths tested
- Main bugs fixed

Tasks:

- Add handler tests for auth and one CRUD resource
- Add service tests for validation rules
- Run through manual test checklist
- Validate bad input, unauthorized access, and empty states
- Verify migration flow on a clean database

Minimum test checklist:

- Login works
- Invalid login fails
- Create university works
- Create faculty with bad `university_id` fails
- Create program works
- Public list returns only published records

### Day 7: Deployment and Buffer

Deliverables:

- Deployable API
- Final bugfix buffer

Tasks:

- Prepare production env vars
- Run migrations in target environment
- Smoke test live endpoints
- Write brief API usage notes
- Fix only high-priority issues

Do not add new features on the last day.

## Work Order

When time is tight, build in this exact order:

1. Fix current `posts` implementation
2. Authentication
3. Universities
4. Faculties
5. Programs
6. Public read endpoints
7. Posts/news improvements
8. Uploads
9. Extra polish

## Daily Output Checklist

Each day should end with something runnable:

- Day 1: migrations and baseline cleanup merged
- Day 2: admin login works
- Day 3: university CRUD works
- Day 4: faculty and program CRUD work
- Day 5: public API works
- Day 6: tests and bug fixes are done
- Day 7: deployed or deployment-ready build

## Practical Rules

- Reuse the current handler/service/repository pattern
- Keep one module shape for all entities
- Prefer simple SQL over generic repository abstractions
- Validate all input at handler and service boundaries
- Keep admin routes separate from public routes
- Do not spend time on perfect architecture this week
- Do not add optional tables until core CRUD is stable

## Suggested Immediate Next Steps

Start with these tasks in the repo:

1. Fix the `posts` migration and repository column mismatch
2. Add `users` migration and auth module
3. Create `universities` model, migration, repository, service, and handler
4. Add `/admin` and public route groups in the router

## Definition Of Success

By the end of the week, success means:

- Admin can log in
- Admin can manage university, faculty, program, and news data
- Public clients can fetch published university information
- The API starts cleanly, migrations run cleanly, and core flows are tested
