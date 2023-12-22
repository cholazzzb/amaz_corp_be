```mermaid
---
title: List all member in group
---

sequenceDiagram
    FE ->> BE: GET /api/v1/rooms/{roomID}/members
    BE ->> FE: success/error

```
