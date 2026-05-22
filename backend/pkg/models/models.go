package models

import "time"

// ChatMessage represents a message in a conversation
type ChatMessage struct {
	ID        string    `json:"id"`
	SessionID string    `json:"sessionId"`
	Role      string    `json:"role"` // "user" or "assistant"
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// ChatRequest is sent by frontend to request AI response
type ChatRequest struct {
	Text     string `json:"text"`
	Image    string `json:"image,omitempty"` // base64 encoded
	Provider string `json:"provider"`        // "ollama", "gemini", etc.
}

// StreamChunk represents a chunk of streaming response
type StreamChunk struct {
	ID        string `json:"id"`
	Chunk     string `json:"chunk"`
	Complete bool   `json:"complete"`
	Error     string `json:"error,omitempty"`
}

// Session represents a conversation session
type Session struct {
	ID        string       `json:"id"`
	Title     string       `json:"title"`
	Messages  []ChatMessage `json:"messages"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

// Config represents application configuration
type Config struct {
	APIKey        string `json:"apiKey"`
	OllamaHost    string `json:"ollamaHost"`
	OllamaModel   string `json:"ollamaModel"`
	Theme         string `json:"theme"`
	FontSize      string `json:"fontSize"`
	AutoStartChat bool   `json:"autoStartChat"`
}

// Screenshot represents a captured screenshot
type Screenshot struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Data      []byte    `json:"data"` // PNG encoded
	Width     int       `json:"width"`
	Height    int       `json:"height"`
}

// OCRResult represents the result of OCR processing
type OCRResult struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	Confidence float64  `json:"confidence"`
}

// TranscriptionResult represents speech-to-text result
type TranscriptionResult struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Language  string    `json:"language"`
	Timestamp time.Time `json:"timestamp"`
	Duration  int       `json:"duration"` // milliseconds
}
