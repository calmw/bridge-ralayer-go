package message

import "sync"

type Details struct {
	State bool
	Data  []byte
}

type Message struct {
	sync.RWMutex
	MessageId map[[32]byte]Details
}

var AllMessage *Message

func init() {
	AllMessage = &Message{
		MessageId: map[[32]byte]Details{},
	}
}

func (m *Message) Save(msgId [32]byte, data []byte, state bool) {
	m.Lock()
	defer m.Unlock()
	m.MessageId[msgId] = Details{
		state,
		data,
	}
	m.ChangeRecord(msgId, data, state)
}

func (m *Message) Get(msgId [32]byte) (bool, Details) {
	m.RLock()
	defer m.RUnlock()
	if data, ok := m.MessageId[msgId]; ok {
		return true, data
	}

	return false, Details{}
}

func (m *Message) ChangeRecord(msgId [32]byte, data []byte, state bool) {

	// Update Record

	// Delete
	if state == true {
		delete(m.MessageId, msgId)
	}
}
