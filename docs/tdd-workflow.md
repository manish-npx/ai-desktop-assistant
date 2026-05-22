# TDD Workflow

## Philosophy

Test-Driven Development is core to this project:

1. **Write failing test** → Red
2. **Implement minimum code** → Green
3. **Refactor** → Refactor
4. **Commit** → Clean history

## Backend Testing (Go)

### Test Structure

```go
// {module}_test.go
package {module}_test

import (
	"context"
	"testing"

	"github.com/manish-npx/ai-desktop-assistant/backend/internal/{module}"
)

func TestFeature(t *testing.T) {
	// Arrange
	svc := module.NewService()

	// Act
	result := svc.DoSomething()

	// Assert
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific test
go test -run TestFeature ./internal/{module}

# Run with coverage
go test -cover ./...
```

### Test Coverage Goals

- **Critical paths**: 100% coverage
- **Services**: 80%+ coverage
- **Utilities**: 70%+ coverage

## Frontend Testing (React)

### Test Structure

```typescript
// Component.test.tsx
import { render, screen } from '@testing-library/react';
import Component from './Component';

describe('Component', () => {
  it('should render with correct text', () => {
    render(<Component />);
    expect(screen.getByText('text')).toBeInTheDocument();
  });
});
```

### Running Tests

```bash
# Run all tests
cd frontend && npm test

# Run in watch mode
npm test -- --watch

# Run with coverage
npm test -- --coverage
```

## Git Workflow

### After Each TDD Cycle

1. **Tests pass**
   ```bash
   go test ./...
   cd frontend && npm test -- --run
   ```

2. **Format code**
   ```bash
   go fmt ./...
   cd frontend && npm run format
   ```

3. **Commit with meaningful message**
   ```bash
   git add .
   git commit -m "feat: add user authentication"
   git push origin main
   ```

## Commit Message Format

```
<type>: <subject>

<body>

<footer>
```

### Types

- **feat**: New feature
- **fix**: Bug fix
- **refactor**: Code refactoring
- **test**: Test additions or modifications
- **chore**: Maintenance tasks
- **docs**: Documentation updates

### Examples

```
feat: implement websocket streaming gateway

Add WebSocket server for real-time AI response streaming.
Uses Gorilla WebSocket for connection management.
Supports concurrent connections with graceful shutdown.

Closes #42

---

fix: resolve websocket reconnection timeout

Increase reconnection timeout from 5s to 30s.
Add exponential backoff for retry attempts.

Fixes #123

---

refactor: extract ai provider interfaces

Create AIProvider interface for cleaner abstraction.
Move common logic to base implementation.
Reduce code duplication across providers.
```

## Feature Development Checklist

- [ ] Create failing test
- [ ] Implement feature
- [ ] Test passes
- [ ] Refactor if needed
- [ ] Code review
- [ ] Commit with meaningful message
- [ ] Push to main
- [ ] Update documentation

## Common Test Patterns

### Testing AI Streaming

```go
func TestAIStreamingResponse(t *testing.T) {
	provider := &MockAIProvider{}
	chan := provider.Stream(context.Background(), ChatRequest{})

	chunks := []string{}
	for chunk := range chan {
		chunks = append(chunks, chunk.Text)
	}

	if len(chunks) != expectedCount {
		t.Errorf("expected %d chunks, got %d", expectedCount, len(chunks))
	}
}
```

### Testing WebSocket

```go
func TestWebSocketConnection(t *testing.T) {
	hub := events.New(logger)
	
	// Test publish/subscribe
	var received Event
	hub.Subscribe("test", func(ctx context.Context, e Event) error {
		received = e
		return nil
	})
	
	hub.Publish(context.Background(), "test", testEvent)
	
	if received != testEvent {
		t.Error("event not received")
	}
}
```

### Testing React Hooks

```typescript
import { renderHook, act } from '@testing-library/react';
import useChat from './useChat';

it('should add message to chat', () => {
  const { result } = renderHook(() => useChat());

  act(() => {
    result.current.addMessage('Hello');
  });

  expect(result.current.messages).toHaveLength(1);
});
```

## Continuous Integration

Ensure tests pass before committing:

```bash
# Pre-commit hook
#!/bin/bash
set -e
go test ./...
cd frontend && npm test -- --run && cd ..
echo "All tests passed!"
```

## Coverage Reports

Generate coverage reports for quality metrics:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```
