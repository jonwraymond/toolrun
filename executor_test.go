package toolrun

import "testing"

func TestMockMCPExecutor_ImplementsInterface(t *testing.T) {
	var _ MCPExecutor = newMockMCPExecutor()
}

func TestMockProviderExecutor_ImplementsInterface(t *testing.T) {
	var _ ProviderExecutor = newMockProviderExecutor()
}

func TestMockLocalRegistry_ImplementsInterface(t *testing.T) {
	var _ LocalRegistry = newMockLocalRegistry()
}
