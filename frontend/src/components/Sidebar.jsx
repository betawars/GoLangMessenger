function Sidebar({ contacts = [], selectedContact, onSelectContact }) {
  return (
    <div className="w-64 bg-gray-800 text-white flex flex-col border-r border-gray-700">
      <div className="p-4 border-b border-gray-700">
        <h2 className="text-xl font-bold">Contacts</h2>
      </div>
      <div className="flex-1 overflow-y-auto">
        {contacts.length === 0 ? (
          <div className="p-4 text-gray-400 text-sm">No contacts available</div>
        ) : (
          contacts.map((contact) => (
          <div
            key={contact.id}
            onClick={() => onSelectContact(contact)}
            className={`p-4 cursor-pointer hover:bg-gray-700 transition-colors ${
              selectedContact?.id === contact.id ? 'bg-gray-700 border-l-4 border-blue-500' : ''
            }`}
          >
            <div className="flex items-center space-x-3">
              <div className="w-10 h-10 rounded-full bg-gray-600 flex items-center justify-center text-sm font-semibold">
                {contact.name.charAt(0).toUpperCase()}
              </div>
              <div className="flex-1 min-w-0">
                <p className="font-medium truncate">{contact.name}</p>
                {contact.lastMessage && (
                  <p className="text-sm text-gray-400 truncate">{contact.lastMessage}</p>
                )}
              </div>
            </div>
          </div>
        ))
        )}
      </div>
    </div>
  )
}

export default Sidebar

