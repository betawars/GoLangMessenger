import { useState, useRef, useEffect } from 'react'

function ChatWindow({ selectedContact, messages = [], sendMessage, isConnected, error }) {
  const [inputMessage, setInputMessage] = useState('')
  const messagesEndRef = useRef(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const handleSend = (e) => {
    e.preventDefault()
    if (inputMessage.trim() && sendMessage(inputMessage.trim())) {
      setInputMessage('')
    }
  }

  if (!selectedContact) {
    return (
      <div className="flex-1 flex items-center justify-center bg-gray-900 text-gray-400">
        <div className="text-center">
          <p className="text-xl">Select a contact to start chatting</p>
        </div>
      </div>
    )
  }

  return (
    <div className="flex-1 flex flex-col bg-gray-900">
      {/* Header */}
      <div className="bg-gray-800 p-4 border-b border-gray-700">
        <div className="flex items-center space-x-3">
          <div className="w-10 h-10 rounded-full bg-gray-600 flex items-center justify-center text-sm font-semibold">
            {selectedContact.name.charAt(0).toUpperCase()}
          </div>
          <div>
            <h3 className="font-semibold text-white">{selectedContact.name}</h3>
            <p className="text-sm text-gray-400">
              {isConnected ? 'Online' : 'Connecting...'}
            </p>
          </div>
        </div>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {error && (
          <div className="bg-red-900/20 border border-red-500 text-red-400 rounded-lg px-4 py-2 text-sm mb-4">
            Connection error: {error}
          </div>
        )}
        {messages.length === 0 ? (
          <div className="flex items-center justify-center h-full text-gray-500">
            <p>No messages yet. Start the conversation!</p>
          </div>
        ) : (
          messages.map((message, index) => (
          <div
            key={index}
            className="flex justify-start"
          >
            <div className="bg-gray-800 rounded-lg px-4 py-2 max-w-xs lg:max-w-md">
              <p className="text-white text-sm">{message.text}</p>
              <p className="text-gray-400 text-xs mt-1">
                {new Date(message.timestamp).toLocaleTimeString()}
              </p>
            </div>
          </div>
        ))
        )}
        <div ref={messagesEndRef} />
      </div>

      {/* Input */}
      <div className="bg-gray-800 p-4 border-t border-gray-700">
        <form onSubmit={handleSend} className="flex space-x-2">
          <input
            type="text"
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 bg-gray-700 text-white rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            disabled={!isConnected}
          />
          <button
            type="submit"
            disabled={!isConnected || !inputMessage.trim()}
            className="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-600 disabled:cursor-not-allowed text-white px-6 py-2 rounded-lg transition-colors"
          >
            Send
          </button>
        </form>
      </div>
    </div>
  )
}

export default ChatWindow

