import { useState } from 'react'
import { useWebSocket } from './hooks/useWebSocket'
import Sidebar from './components/Sidebar'
import ChatWindow from './components/ChatWindow'
import './App.css'

function App() {
  const [selectedContact, setSelectedContact] = useState(null)
  
  // Mock contacts data - in a real app, this would come from an API
  const [contacts] = useState([
    { id: 1, name: 'Alice', lastMessage: 'Hey, how are you?' },
    { id: 2, name: 'Bob', lastMessage: 'See you later!' },
    { id: 3, name: 'Charlie', lastMessage: 'Thanks for the help' },
    { id: 4, name: 'Diana', lastMessage: 'Meeting at 3pm' },
  ])

  // WebSocket connection - adjust URL if your backend runs on a different port
  const wsUrl = 'ws://localhost:8080/ws'
  const { isConnected, messages, sendMessage, error } = useWebSocket(wsUrl)

  return (
    <div className="h-screen flex bg-gray-900">
      <Sidebar
        contacts={contacts}
        selectedContact={selectedContact}
        onSelectContact={setSelectedContact}
      />
      <ChatWindow
        selectedContact={selectedContact}
        messages={messages}
        sendMessage={sendMessage}
        isConnected={isConnected}
        error={error}
      />
    </div>
  )
}

export default App
