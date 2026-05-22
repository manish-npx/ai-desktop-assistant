# AI Desktop Assistant

> A lightweight, native desktop AI assistant built with **Wails**, **Go**, and **React 19**

## Overview

This project is an **educational reference implementation** for building production-grade desktop applications with:

- **Backend**: Go with clean architecture
- **Frontend**: React 19 + TypeScript + TailwindCSS + shadcn/ui
- **Desktop**: Wails v2 (lightweight alternative to Electron)
- **AI**: Local Ollama + streaming responses
- **Audio**: Real-time transcription with whisper.cpp
- **OCR**: Image text extraction with preprocessing
- **Overlay**: Transparent desktop overlay window

## Project Goals

✅ Learning architecture patterns for desktop applications  
✅ Implementing real-time AI streaming  
✅ Building audio/image processing pipelines  
✅ Writing maintainable, testable code  
✅ Following SOLID principles and clean architecture  
✅ Creating production-ready engineering practices  
✅ Demonstrating ethical, privacy-first design  

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Wails v2
- **Database**: SQLite with migrations
- **Logging**: Uber Zap (structured logging)
- **WebSocket**: Gorilla WebSocket

### Frontend
- **Framework**: React 19
- **Language**: TypeScript (strict mode)
- **Styling**: TailwindCSS v4
- **Components**: shadcn/ui
- **State**: Zustand
- **HTTP**: TanStack Query + Axios
- **Icons**: Lucide React
- **Hotkeys**: react-hotkeys-hook

### AI & Processing
- **LLM**: Ollama (local inference)
- **Speech**: whisper.cpp (local transcription)
- **OCR**: Tesseract via gosseract
- **Screenshots**: robotgo + image processing

## Project Structure

```
ai-desktop-assistant/
├── backend/
│   ├── cmd/app/main.go              # Entry point
│   ├── internal/
│   │   ├── ai/                      # AI provider abstraction
│   │   ├── audio/                   # Audio capture & processing
│   │   ├── capture/                 # Screenshots & OCR
│   │   ├── overlay/                 # Window management
│   │   ├── websocket/               # WebSocket handlers
│   │   ├── workers/                 # Async task workers
│   │   ├── events/                  # Event bus/pub-sub
│   │   ├── storage/                 # SQLite repositories
│   │   ├── logging/                 # Structured logging
│   │   ├── config/                  # Configuration management
│   │   └── security/                # Security utilities
│   ├── pkg/
│   │   ├── models/                  # Shared data models
│   │   └── utils/                   # Utilities
│   └── tests/                       # Integration tests
├── frontend/
│   ├── src/
│   │   ├── components/              # React components
│   │   ├── hooks/                   # Custom hooks
│   │   ├── pages/                   # Page layouts
│   │   ├── services/                # API/WebSocket clients
│   │   ├── store/                   # Zustand stores
│   │   ├── types/                   # TypeScript types
│   │   └── styles/                  # Global styles
│   └── public/                      # Static assets
├── storage/
│   └── migrations/                  # Database migrations
├── docs/
│   ├── architecture.md              # System design
│   ├── tdd-workflow.md              # Testing approach
│   ├── ai-usage.md                  # AI integration guide
│   └── deployment.md                # Release process
└── scripts/
    ├── build.sh                     # Build script
    ├── dev.sh                       # Development startup
    └── migrate.sh                   # Database migrations
```

## Prerequisites

- **Go**: 1.21 or higher
- **Node.js**: 18 or higher
- **Wails**: Latest version
- **Ollama**: For local AI (optional, for development)

```bash
# Install Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Install Node dependencies
cd frontend
npm install
cd ..
```

## Development

### Start Development Server

```bash
# Install dependencies
go mod download
cd frontend && npm install && cd ..

# Run development mode (hot reload)
wails dev
```

### Run Tests

```bash
# Go tests
go test ./...

# React tests
cd frontend && npm test
```

### Build Production

```bash
# Windows
wails build -platform windows/amd64

# macOS
wails build -platform darwin/universal

# Linux
wails build -platform linux/amd64
```

## Architecture Highlights

### Clean Architecture Layers

```
UI Layer (React)
      ↓
API/WebSocket Gateway
      ↓
Service Layer (Use Cases)
      ↓
Repository Layer (Data Access)
      ↓
Database (SQLite)
```

### Event-Driven Architecture

- **Event Bus**: Publish/Subscribe pattern for loose coupling
- **Worker Pool**: Async processing for OCR, transcription, AI
- **WebSocket Gateway**: Real-time bidirectional communication

### Provider Pattern

```go
type AIProvider interface {
    Stream(ctx context.Context, req ChatRequest) (<-chan StreamChunk, error)
    GetModel() string
}

// Implementations: Ollama, Gemini, OpenAI, Groq, LMStudio
```

## Security & Privacy

✅ **Local-first**: All processing can run locally  
✅ **No telemetry**: No tracking or analytics  
✅ **Encrypted storage**: Credentials encrypted at rest  
✅ **Input validation**: All inputs sanitized  
✅ **Graceful errors**: No sensitive info in logs  

## Documentation

See the `docs/` directory for detailed documentation:

- **[Architecture](docs/architecture.md)**: System design and patterns
- **[TDD Workflow](docs/tdd-workflow.md)**: Testing approach
- **[AI Integration](docs/ai-usage.md)**: Using different AI providers
- **[Deployment](docs/deployment.md)**: Building and packaging

## Contributing

This is an educational project. Contributions following the project guidelines are welcome!

## License

MIT License - see LICENSE file for details

## Learning Resources

- [Wails Documentation](https://wails.io)
- [Go Clean Architecture](https://github.com/golang-standards/project-layout)
- [React 19 Docs](https://react.dev)
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID)

---

**Built with ❤️ for learning and reference**
