import {
  Bot,
  Plus,
  MessageSquare,
  Settings,
  PanelLeftClose,
  Code2,
} from "lucide-react";

export default function Sidebar({ onNewChat, collapsed, onToggle }) {
  return (
    <aside className={`sidebar ${collapsed ? "collapsed" : ""}`}>
      <div className="sidebar-top">
        <div className="brand">
          <div className="brand-icon">
            <Bot size={22} />
          </div>

          {!collapsed && (
            <div className="brand-text">
              <strong>Go AI Agent</strong>
              <span>Powered by OpenRouter</span>
            </div>
          )}
        </div>

        <button
          className="icon-button sidebar-toggle"
          onClick={onToggle}
          title="Toggle sidebar"
        >
          <PanelLeftClose size={19} />
        </button>
      </div>

      <button className="new-chat-button" onClick={onNewChat}>
        <Plus size={19} />
        {!collapsed && <span>New chat</span>}
      </button>

      {!collapsed && (
        <>
          <div className="sidebar-section-label">YOUR WORKSPACE</div>

          <div className="sidebar-nav-item active">
            <MessageSquare size={18} />
            <span>AI Assistant</span>
          </div>

          <div className="sidebar-section-label recent-label">RECENT CHATS</div>

          <div className="empty-history">
            <MessageSquare size={18} />
            <p>Your conversations will appear here.</p>
          </div>
        </>
      )}

      <div className="sidebar-bottom">
        <div className="connection-status">
          <span className="status-dot" />
          {!collapsed && <span>AI workspace</span>}
        </div>

        <button className="sidebar-footer-item">
          <Settings size={18} />
          {!collapsed && <span>Settings</span>}
        </button>

        <a
          className="sidebar-footer-item"
          href="https://github.com/"
          target="_blank"
          rel="noreferrer"
        >
          <Code2 size={16} />
          {!collapsed && <span>GitHub</span>}
        </a>
      </div>
    </aside>
  );
}
