# AI Integration Guide

## Supported AI Providers

### 1. Ollama (Recommended for Local Development)

**Setup:**
```bash
# Install Ollama from https://ollama.ai
# Start Ollama
ollama serve

# Pull a model
ollama pull qwen2.5-coder
```

**Configuration:**
```json
{
  "aiProvider": "ollama",
  "ollamaHost": "http://localhost:11434",
  "ollamaModel": "qwen2.5-coder"
}
```

**Benefits:**
- ✅ Runs locally - no API calls
- ✅ Private - data stays on device
- ✅ Free - open source models
- ✅ Fast - optimized for CPU/GPU

### 2. Google Gemini API

**Setup:**
```bash
# Get API key from https://makersuite.google.com/app/apikey
export GOOGLE_API_KEY="your-key-here"
```

**Configuration:**
```json
{
  "aiProvider": "gemini",
  "geminiApiKey": "sk-..."
}
```

### 3. OpenAI API

**Setup:**
```bash
export OPENAI_API_KEY="sk-..."
```

**Configuration:**
```json
{
  "aiProvider": "openai",
  "openaiApiKey": "sk-..."
}
```

### 4. Groq API

**Setup:**
```bash
export GROQ_API_KEY="gsk_..."
```

**Configuration:**
```json
{
  "aiProvider": "groq",
  "groqApiKey": "gsk-..."
}
```

### 5. LMStudio (Local)

**Setup:**
```bash
# Download LMStudio from https://lmstudio.ai
# Load a model and start server
```

**Configuration:**
```json
{
  "aiProvider": "lmstudio",
  "lmstudioHost": "http://localhost:1234"
}
```

## Streaming Responses

All AI providers support streaming for real-time responses:

```go
provider := ai.NewOllamaProvider("http://localhost:11434")

req := ai.ChatRequest{
    Text: "Explain Go concurrency",
    Provider: "ollama",
}

stream, err := provider.Stream(ctx, req)
if err != nil {
    log.Fatal(err)
}

for chunk := range stream {
    fmt.Print(chunk.Text)
    if chunk.Complete {
        break
    }
}
```

## Context Management

Maintain conversation history:

```go
type ChatSession struct {
    ID       string
    Messages []ChatMessage
}

// Add message to session
session.AddMessage(ChatMessage{
    Role:    "user",
    Content: "Hello",
})

// Send with context
response := provider.StreamWithContext(ctx, req, session.GetContext())
```

## Rate Limiting

Implement quotas per provider:

```go
limiter := ai.NewRateLimiter()

// Check quota
if !limiter.CanUse("gemini", 1000) {
    return errors.New("quota exceeded")
}

// Use tokens
limiter.Use("gemini", 1000)
```

## Model Selection

Choose the right model for your task:

### Code Understanding
- qwen2.5-coder (Ollama)
- GPT-4 (OpenAI)
- Claude (Anthropic)

### General Chat
- llama2 (Ollama)
- GPT-3.5 (OpenAI)
- Gemini (Google)

### Fast Responses
- qwen2.5-0.5b (Ollama - 500M)
- GPT-3.5 Turbo (OpenAI)
- Groq (fastest inference)

## Error Handling

Handle AI provider errors gracefully:

```go
response, err := provider.Stream(ctx, req)

switch {
case errors.Is(err, context.DeadlineExceeded):
    // Timeout handling
    return "Request timed out"
case errors.Is(err, ErrQuotaExceeded):
    // Quota handling
    return "Daily quota exceeded"
case err != nil:
    // Generic error
    return fmt.Sprintf("Error: %v", err)
}
```

## Best Practices

1. **Always use context for cancellation**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()
   ```

2. **Stream responses instead of waiting for completion**
   ```go
   // Good: Stream
   stream := provider.Stream(ctx, req)
   
   // Avoid: Blocking
   response := provider.Complete(ctx, req)
   ```

3. **Implement retry logic for flaky providers**
   ```go
   for attempt := 0; attempt < 3; attempt++ {
       if err := tryRequest(); err == nil {
           return nil
       }
       time.Sleep(exponentialBackoff(attempt))
   }
   ```

4. **Cache responses when appropriate**
   ```go
   response, cached := cache.Get(req.Hash())
   if cached {
       return response
   }
   ```

5. **Monitor usage and costs**
   ```go
   metrics.RecordAICall(provider, tokens, latency, cost)
   ```
