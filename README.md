# Event-Driven Realtime System

## Overview
This project is a **simple event-driven backend system** that demonstrates how modern backend applications handle tasks **asynchronously** and deliver **realtime updates** to clients.

Instead of processing everything directly inside an API request, this system uses:
- Events
- Message Queue
- Background Worker
- Realtime communication

This approach is commonly used in **production systems** such as notification services, order processing, monitoring systems, and SaaS platforms.

---

## Why This Project Exists
Traditional CRUD-based backends process requests synchronously:

```
Client → API → Process → Response
```

This approach can become slow and hard to scale when:
- Tasks are heavy
- Events happen frequently
- Realtime updates are required

This project demonstrates an alternative approach:

```
Client → API → Event → Queue → Worker → Realtime Update
```


The API stays **fast**, while background workers handle the heavy processing.

---

## ⚙️ How the System Works

### 1️⃣ Client Sends an Event
The client sends an event to the API:

```http
POST /events
{
  "type": "USER_ACTION",
  "message": "User performed an action"
}
```

