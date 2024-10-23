package types

const (
	// ModuleName defines the module name
	ModuleName = "chain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chain"
)

var (
	ParamsKey = []byte("p_chain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
