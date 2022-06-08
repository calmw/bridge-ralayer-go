package message

import "sync"

type Message struct {
	sync.RWMutex
	MessageId map[[32]byte][]byte
}

var AllMessage *Message

func (m *Message) Save(msgId [32]byte, data []byte) {
	m.Lock()
	defer m.Unlock()
	m.MessageId[msgId] = data
}

func (m *Message) Get(msgId [32]byte) (bool, []byte) {
	m.RLock()
	defer m.RUnlock()
	if data, ok := m.MessageId[msgId]; ok {
		return true, data
	}

	return false, nil
}
