import { useState } from "react";
import { Plus } from "lucide-react";

import Sidebar from "./components/Sidebar";
import ChatWindow from "./components/Chatwindow";

import "./index.css";

export default function App() {
  const [messages, setMessages] = useState([]);
  const [collapsed, setCollapsed] = useState(false);
  const [loading, setLoading] = useState(false);

  function handleNewChat() {
    if (loading) return;
    setMessages([]);
  }

  async function handleSend(text) {
    const message = text.trim();

    if (!message || loading) return;

    // Keep the previous conversation for AI context.
    const history = messages
      .filter((msg) => msg.role === "user" || msg.role === "assistant")
      .map(({ role, content }) => ({
        role,
        content,
      }));

    // Show the user's message immediately.
    setMessages((previous) => [
      ...previous,
      {
        id: crypto.randomUUID(),
        role: "user",
        content: message,
      },
    ]);

    setLoading(true);

    try {
      const response = await fetch("/api/chat", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          message,
          history,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`API error (${response.status}): ${errorText}`);
      }

      const data = await response.json();

      if (typeof data.reply !== "string") {
        throw new Error("Invalid response from Go backend");
      }

      setMessages((previous) => [
        ...previous,
        {
          id: crypto.randomUUID(),
          role: "assistant",
          content: data.reply,
        },
      ]);
    } catch (error) {
      console.error("Chat API error:", error);

      setMessages((previous) => [
        ...previous,
        {
          id: crypto.randomUUID(),
          role: "assistant",
          content:
            "Sorry, I couldn't get an AI response. Please check your Go backend, API key, and terminal logs.",
        },
      ]);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="app-shell">
      <Sidebar
        onNewChat={handleNewChat}
        collapsed={collapsed}
        onToggle={() => setCollapsed((previous) => !previous)}
      />

      <ChatWindow messages={messages} onSend={handleSend} loading={loading} />

      <button
        className="mobile-new-chat"
        onClick={handleNewChat}
        aria-label="New chat"
        disabled={loading}
      >
        <Plus size={20} />
      </button>
    </div>
  );
}
