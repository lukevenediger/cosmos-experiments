package types

const (
	// ModuleName defines the module name
	ModuleName = "experi"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_experi"
)

var (
	ParamsKey = []byte("p_experi")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
