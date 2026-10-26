import {
  Bot,
  Sparkles,
  Code2,
  Lightbulb,
  MessageCircle,
  Copy,
  Check,
} from "lucide-react";
import { useState } from "react";
import MessageInput from "./MessageInput";

const suggestions = [
  {
    icon: Code2,
    title: "Explain Go concepts",
    description: "Understand goroutines, channels, and interfaces.",
    prompt: "Explain goroutines in Go with a simple example.",
  },
  {
    icon: Lightbulb,
    title: "Help me build",
    description: "Plan and develop a backend project.",
    prompt: "Help me plan a clean Go backend project architecture.",
  },
  {
    icon: MessageCircle,
    title: "Explore AI development",
    description: "Learn how AI APIs work in Go.",
    prompt: "Explain how an AI API works in a Go application.",
  },
];

function WelcomeScreen({ onSuggestion }) {
  return (
    <div className="welcome-screen">
      <div className="welcome-icon">
        <Bot size={32} />
        <span className="welcome-sparkle">
          <Sparkles size={15} />
        </span>
      </div>

      <p className="eyebrow">YOUR PERSONAL AI WORKSPACE</p>

      <h1>
        What can I help
        <br />
        <span>you build today?</span>
      </h1>

      <p className="welcome-description">
        Ask questions, explore ideas, or write code.
        <br />
        Let's make something useful together.
      </p>

      <div className="suggestion-grid">
        {suggestions.map((item) => {
          const Icon = item.icon;

          return (
            <button
              className="suggestion-card"
              key={item.title}
              onClick={() => onSuggestion(item.prompt)}
            >
              <Icon size={20} />
              <strong>{item.title}</strong>
              <span>{item.description}</span>
            </button>
          );
        })}
      </div>
    </div>
  );
}

function MessageBubble({ message }) {
  const [copied, setCopied] = useState(false);

  async function copyMessage() {
    try {
      await navigator.clipboard.writeText(message.content);
      setCopied(true);
      setTimeout(() => setCopied(false), 1500);
    } catch {
      // Clipboard may be unavailable in some browser contexts.
    }
  }

  return (
    <div className={`message-row ${message.role}`}>
      {message.role === "assistant" && (
        <div className="message-avatar">
          <Bot size={18} />
        </div>
      )}

      <div className="message-content">
        <div className="message-author">
          {message.role === "user" ? "You" : "Go AI Agent"}
        </div>

        <div className="message-text">{message.content}</div>

        {message.role === "assistant" && (
          <button
            className="copy-message-button"
            onClick={copyMessage}
            title="Copy response"
          >
            {copied ? <Check size={14} /> : <Copy size={14} />}
            {copied ? "Copied" : "Copy"}
          </button>
        )}
      </div>
    </div>
  );
}

export default function ChatWindow({ messages, loading, onSend }) {
  return (
    <main className="chat-main">
      <header className="chat-header">
        <div>
          <span className="header-title">AI Assistant</span>
          <span className="model-label">
            <span className="status-dot" />
            OpenRouter · Your selected model
          </span>
        </div>

        <div className="header-badge">
          <Sparkles size={15} />
          AI Workspace
        </div>
      </header>

      <section className="conversation">
        {messages.length === 0 ? (
          <WelcomeScreen onSuggestion={onSend} />
        ) : (
          <div className="message-list">
            {messages.map((message) => (
              <MessageBubble key={message.id} message={message} />
            ))}

            {loading && (
              <div className="message-row assistant">
                <div className="message-avatar">
                  <Bot size={18} />
                </div>

                <div className="message-content">
                  <div className="message-author">Go AI Agent</div>
                  <div className="typing-indicator">
                    <span />
                    <span />
                    <span />
                  </div>
                </div>
              </div>
            )}
          </div>
        )}
      </section>

      <footer className="chat-footer">
        <MessageInput onSend={onSend} loading={loading} />
        <p className="disclaimer">
          AI can make mistakes. Verify important information.
        </p>
      </footer>
    </main>
  );
}
