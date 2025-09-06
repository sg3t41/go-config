package strategy

import (
	"fmt"
	"path/filepath"
)

var defaultManager = NewManager()

type Manager struct {
	strategies map[string]IStrategy
}

func NewManager() *Manager {
	return &Manager{
		strategies: make(map[string]IStrategy),
	}
}

func (m *Manager) Register(ext string, s IStrategy) {
	m.strategies[ext] = s
}

func (m *Manager) Get(path string) (IStrategy, error) {
	ext := filepath.Ext(path)
	if s, ok := m.strategies[ext]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("unsupported file type: %s", ext)
}

func Register(ext string, s IStrategy) {
	defaultManager.Register(ext, s)
}

func Get(path string) (IStrategy, error) {
	return defaultManager.Get(path)
}
