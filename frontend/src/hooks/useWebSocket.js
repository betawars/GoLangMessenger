import { useEffect, useRef, useState, useCallback } from 'react'

export function useWebSocket(url) {
  const [isConnected, setIsConnected] = useState(false)
  const [messages, setMessages] = useState([])
  const [error, setError] = useState(null)
  const wsRef = useRef(null)
  const reconnectTimeoutRef = useRef(null)

  const connect = useCallback(() => {
    try {
      const ws = new WebSocket(url)
      wsRef.current = ws

      ws.onopen = () => {
        setIsConnected(true)
        setError(null)
        console.log('WebSocket connected')
      }

      ws.onmessage = (event) => {
        try {
          const message = event.data
          setMessages((prev) => [...prev, { text: message, timestamp: new Date() }])
        } catch (err) {
          console.error('Error parsing message:', err)
        }
      }

      ws.onerror = (error) => {
        console.error('WebSocket error:', error)
        setError('WebSocket connection error')
      }

      ws.onclose = () => {
        setIsConnected(false)
        console.log('WebSocket disconnected')
        
        // Attempt to reconnect after 3 seconds
        reconnectTimeoutRef.current = setTimeout(() => {
          if (wsRef.current?.readyState === WebSocket.CLOSED) {
            connect()
          }
        }, 3000)
      }
    } catch (err) {
      setError('Failed to create WebSocket connection')
      console.error('WebSocket connection error:', err)
    }
  }, [url])

  const sendMessage = useCallback((message) => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(message)
      return true
    } else {
      console.warn('WebSocket is not connected')
      return false
    }
  }, [])

  const disconnect = useCallback(() => {
    if (reconnectTimeoutRef.current) {
      clearTimeout(reconnectTimeoutRef.current)
      reconnectTimeoutRef.current = null
    }
    if (wsRef.current) {
      wsRef.current.close()
      wsRef.current = null
    }
  }, [])

  useEffect(() => {
    connect()

    return () => {
      disconnect()
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url]) // Only reconnect if URL changes

  return {
    isConnected,
    messages,
    error,
    sendMessage,
    disconnect,
    reconnect: connect,
  }
}

