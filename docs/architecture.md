# Architecture Overview

## System Design

The AI Desktop Assistant follows a **clean architecture** pattern with clear separation of concerns:

```
┌─────────────────────────────────────┐
│         UI Layer (React)            │
│  Components, Pages, State (Zustand) │
└──────────────┬──────────────────────┘
               │ WebSocket / HTTP
┌──────────────▼──────────────────────┐
│     API Gateway Layer (Wails)       │
│    WebSocket Handlers, Routes       │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Service Layer (Business Logic)  │
│  AI, Audio, OCR, Capture Services   │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│    Repository Layer (Data Access)   │
│       Storage, Events, Config       │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│   Infrastructure Layer               │
│  Database, External APIs, Logging    │
└─────────────────────────────────────┘
```

## Module Organization

### Backend Structure

#### `internal/ai/`
- **Purpose**: AI provider abstraction
- **Key Components**:
  - `provider.go`: Interface for AI providers
  - `ollama.go`: Ollama integration
  - `rate_limiter.go`: Rate limiting logic
- **Responsibility**: Handle AI inference and streaming

#### `internal/audio/`
- **Purpose**: Audio capture and processing
- **Key Components**:
  - `capture.go`: Microphone capture
  - `processor.go`: Audio format conversion
  - `vad.go`: Voice activity detection
- **Responsibility**: Capture and preprocess audio

#### `internal/capture/`
- **Purpose**: Screenshot and OCR
- **Key Components**:
  - `screenshot.go`: Capture desktop/region
  - `ocr.go`: OCR extraction
  - `preprocessor.go`: Image preprocessing
- **Responsibility**: Extract text from images

#### `internal/overlay/`
- **Purpose**: Window management
- **Key Components**:
  - `manager.go`: Window lifecycle
  - `input.go`: Input handling
  - `renderer.go`: Rendering pipeline
- **Responsibility**: Manage overlay window

#### `internal/websocket/`
- **Purpose**: Real-time communication
- **Key Components**:
  - `hub.go`: Connection management
  - `handlers.go`: Message handlers
  - `broadcaster.go`: Event broadcasting
- **Responsibility**: WebSocket server and messaging

#### `internal/workers/`
- **Purpose**: Async task processing
- **Key Components**:
  - `pool.go`: Worker pool
  - `task.go`: Task definition
  - `executor.go`: Task execution
- **Responsibility**: Execute long-running tasks

#### `internal/events/`
- **Purpose**: Event bus (pub/sub)
- **Key Components**:
  - `bus.go`: Event bus implementation
  - `events.go`: Event type definitions
- **Responsibility**: Decouple components via events

#### `internal/storage/`
- **Purpose**: Data persistence
- **Key Components**:
  - `db.go`: Database initialization
  - `repositories/`: CRUD operations
  - `migrations/`: Schema versioning
- **Responsibility**: Data access layer

#### `internal/logging/`
- **Purpose**: Structured logging
- **Key Components**:
  - `logger.go`: Logger configuration
  - `middleware.go`: Logging middleware
- **Responsibility**: Consistent, structured logging

#### `internal/config/`
- **Purpose**: Configuration management
- **Key Components**:
  - `config.go`: Config loading
  - `env.go`: Environment variables
- **Responsibility**: Application configuration

#### `internal/security/`
- **Purpose**: Security utilities
- **Key Components**:
  - `encryption.go`: Credential encryption
  - `validator.go`: Input validation
- **Responsibility**: Security operations

### Frontend Structure

#### `src/components/`
Reusable React components using shadcn/ui

#### `src/hooks/`
Custom React hooks for business logic

#### `src/pages/`
Page/layout components

#### `src/services/`
- `api.ts`: HTTP client
- `websocket.ts`: WebSocket client
- `ai.ts`: AI service

#### `src/store/`
Zustand stores for state management

#### `src/types/`
Shared TypeScript types

## Data Flow

### AI Chat Flow

```
User Input
    ↓
React Component → Send message via WebSocket
    ↓
WebSocket Handler (Backend)
    ↓
AI Service → Call AI Provider
    ↓
Stream chunks via WebSocket
    ↓
React receives chunks → Display in UI
```

### Screenshot + OCR Flow

```
User captures screenshot
    ↓
Capture Service → Screenshot region
    ↓
Worker Pool → Process image
    ↓
OCR Service → Extract text
    ↓
Event Bus → Publish OCR complete
    ↓
React listens → Display results
```

### Audio Transcription Flow

```
Microphone input
    ↓
Audio Capture → Record PCM
    ↓
VAD Detection → Segment voice
    ↓
Worker Pool → Transcribe with whisper.cpp
    ↓
Event Bus → Publish transcription
    ↓
React listens → Display transcription
```

## Design Patterns

### 1. Provider Pattern
```go
type AIProvider interface {
    Stream(ctx context.Context, req ChatRequest) (<-chan StreamChunk, error)
    GetModel() string
}
```
Allows swapping AI backends (Ollama, Gemini, OpenAI, Groq)

### 2. Event Bus Pattern
Components communicate via events instead of direct calls.
Enables loose coupling and scalability.

### 3. Worker Pool Pattern
Async processing for CPU-intensive tasks (OCR, transcription, AI).
Avoids blocking the UI.

### 4. Repository Pattern
Abstraction for data access.
Enables easy testing and swapping storage backends.

### 5. Dependency Injection
Services receive dependencies via constructors.
Facilitates testing and modularity.

## Concurrency Model

- **Goroutines**: Used for async operations
- **Channels**: For communication between goroutines
- **Context**: For cancellation and timeouts
- **Worker Pool**: For limiting concurrent tasks

## Testing Strategy

### Backend (Go)
- **Unit Tests**: Individual functions/methods
- **Integration Tests**: Service interactions
- **Mock Objects**: For external dependencies

### Frontend (React)
- **Component Tests**: React Testing Library
- **Hook Tests**: Test custom hooks
- **E2E Tests**: User workflows

## Performance Considerations

1. **Streaming**: AI responses stream in real-time
2. **Async Processing**: Heavy tasks don't block UI
3. **Memory**: Efficient buffer management
4. **Goroutines**: Limited by worker pool
5. **Rendering**: React optimizations (memo, useMemo)

## Security Considerations

1. **Input Validation**: All inputs sanitized
2. **Encryption**: Credentials encrypted at rest
3. **Logging**: No sensitive data in logs
4. **Panic Recovery**: Graceful error handling
5. **Permissions**: Respect OS permissions
