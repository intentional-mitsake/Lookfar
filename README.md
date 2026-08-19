# Lookfar

A community-driven city map platform where verified residents share local knowledge — festivals, landmarks, routes, and on-the-ground information that maps alone don't carry.

---

## Features

### Resident Verification
- New accounts start unverified and gain contribution access only after approval
- Submit a proof-of-residence document to enter the verification queue
- Admins review and approve or reject submissions with a recorded reason
- Every admin action is written to an audit log
- Unverified users can browse everything; only verified residents can contribute

### Landmarks
- Verified residents can add landmarks (temples, monuments, markets, parks, and more) with a name, description, category, and precise coordinates
- Query landmarks by radius around a point or by the bounding box of a map viewport
- Each landmark aggregates a live average rating from community reviews

### Festivals
- Add festivals with a name, description, location, and a start and end time
- Filter by currently active festivals, upcoming festivals, or by proximity
- A dedicated endpoint returns all active festivals as a GeoJSON FeatureCollection — ready to drop into any map library

### Reviews
- Verified residents can leave a rating and written review on any landmark or festival
- One review per resident per target per 30-day window, enforced at the database level
- Ratings are aggregated and returned with every landmark and festival response

### Community Information Posts
- Verified residents can submit free-form information posts attached to a landmark or festival (road closures, crowd conditions, parking notes, anything useful)
- Posts go into a moderation queue and are only publicly visible after admin approval

### Community Routes
- Verified residents can contribute walking or travel routes to active festivals, stored as geographic line geometries
- Routes are associated with a specific festival and validated to start near it
- Other residents can upvote routes; routes are returned sorted by upvotes
- Double-voting is prevented at the database level

### Real-Time Festival Updates
- Connect over WebSocket to receive live events when a festival starts or ends
- The server checks active festival windows every 60 seconds and broadcasts typed events to all connected clients
- Dead connections are cleaned up automatically via ping/pong

### Rate Limiting & Caching
- Per-IP and per-user token bucket rate limiting on all write endpoints, implemented with a Lua script in Redis
- Active festival and landmark data cached in Redis with TTL-based expiry and update-time invalidation

### Observability
- Structured request logging with IP and user ID tracking
- Health endpoint that checks database and Redis connectivity
- Metrics endpoint exposing active WebSocket connections and cache hit counters

---

## Tech Stack

| | |
|---|---|
| Language | Go |
| Database | PostgreSQL 16 + PostGIS 3.4 |
| Cache / Rate Limiting | Redis 7 |
| Real-time | WebSocket (gorilla/websocket) |
| Auth | JWT access tokens + rotating refresh tokens |
| Containerization | Docker + Docker Compose |

---

## Project Status

Early development. Features are being built in phases — verification pipeline first, then landmarks and festivals, then reviews, routes, and real-time.
