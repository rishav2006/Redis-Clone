# GoCacheDB

A Redis-inspired in-memory key-value database built in Go, designed to explore systems programming concepts such as TCP networking, concurrent client handling, synchronization primitives, persistence, key expiration, and publish-subscribe messaging.

## Features
• TCP-based server supporting multiple client connections  
• Concurrent request handling using Goroutines  
• Thread-safe in-memory storage using sync.RWMutex  
• Key-value operations: SET, GET, EXISTS  
• Time-To-Live (TTL) based key expiration  
• Snapshot persistence using JSON serialization  
• Automatic data recovery on server startup  
• Publish/Subscribe (Pub/Sub) messaging system  
• Graceful handling of invalid commands and malformed requests

## Architecture
![Logo Description](https://i.ibb.co/sdDVX1yV/diagram-export-6-18-2026-1-43-58-PM.png)

## Implemented Commands
```markdown

| Command                   | Description                    |
|---------------------------|--------------------------------|
| `SET key value`           | Store a value                  |
| `GET key`                 | Retrieve a value               |
| `EXISTS key`              | Check if a key exists          |
| `SETEX key ttl value`     | Store a value with expiration  |
| `SUBSCRIBE channel`       | Subscribe to a channel         |
| `PUBLISH channel message` | Publish a message              |
```

## Project Evolution

### V1 – TCP Server Foundation
- Built a TCP server listening on port `6379`
- Added support for client connections
- Implemented basic command parsing

### V2 – Concurrent Client Handling
- Added Goroutines for serving multiple clients simultaneously
- Prevented server blocking caused by individual client requests

### V3 – Thread-Safe Storage
- Replaced initial data structures with a hash map
- Introduced `sync.RWMutex` for safe concurrent access
- Eliminated race conditions during reads and writes

### V4 – Key Expiration (TTL)
- Added support for expiring keys automatically
- Implemented background expiration management

### V5 – Persistence
- Implemented snapshot-based persistence using JSON
- Added startup recovery mechanism to restore database state

### V6 – Pub/Sub Messaging
- Added channel-based publish-subscribe functionality
- Enabled real-time message delivery to subscribed clients

## Running the Project
### Clone the Repository
```markdown
git clone https://github.com/rishav2006/GoCacheDB.git
cd GoCacheDB
```

### Start Main Server
```markdown
cd cmd/server
go run main.go
```

### Start Client
```markdown
cd cmd/client
go run .
```

### Connect using Netcat
```markdown
nc localhost 6379
```

### Example
```markdown
SET name Rishav
Okay

GET name
Rishav

EXISTS name
YES
```

## Key Learnings
Through this project, I gained hands-on experience with:

• TCP socket programming  
• Concurrent systems development in Go  
• GoRoutines and synchronization primitives  
• Thread-safe data structures  
• Persistance and recovery mechanisms  
• Publish-Subscribe architecture  
• Designing Redis-inspired database internals  

## Future Improvements
• Redis RESP protocol support  
• AOF (Append Only File) persistence  
• Master-Replica replication  
• Distributed clustering  
• LRU/LFU eviction policies  
• Benchmarking and performance profiling  

## Tech Stack
• Go  
• TCP Networking  
• Goroutines  
• sync.RWMutex  
• JSON Persistance  
• Channels  
• Concurrent Data Structures  
