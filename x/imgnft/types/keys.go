package types

const (
	// ModuleName defines the module name
	ModuleName = "imgnft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_imgnft"
)

var (
	ParamsKey = []byte("p_imgnft")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
