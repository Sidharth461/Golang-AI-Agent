
import { ArrowUp, Paperclip } from "lucide-react";
import { useState } from "react";

export default function MessageInput({ onSend, loading }) {
  const [message, setMessage] = useState("");

  function handleSubmit(event) {
    event.preventDefault();

    const text = message.trim();

    if (!text || loading) return;

    onSend(text);
    setMessage("");
  }

  function handleKeyDown(event) {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      event.currentTarget.form?.requestSubmit();
    }
  }

  return (
    <form className="composer" onSubmit={handleSubmit}>
      <textarea
        value={message}
        onChange={(event) => setMessage(event.target.value)}
        onKeyDown={handleKeyDown}
        placeholder="Message your AI assistant..."
        rows={1}
        aria-label="Your message"
      />

      <div className="composer-footer">
        <div className="composer-hint">
          <Paperclip size={16} />
          <span>Go AI Agent</span>
        </div>

        <button
          type="submit"
          className="send-button"
          disabled={!message.trim() || loading}
          aria-label="Send message"
        >
          <ArrowUp size={19} />
        </button>
      </div>
    </form>
  );
}