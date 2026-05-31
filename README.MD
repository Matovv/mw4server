Magic Wars 4 - Server

Network:
Join workflow:
Client Connects
      │
      ▼
Send JoinSessionRequest
      │
      ▼
SessionManager.Get()
      │
      ▼
AssignPlayer()
      │
      ▼
Send JoinSessionResponse
      │
      ▼
Receive World Snapshot
      │
      ▼
Start Sending Commands