package types

const (
	// ModuleName defines the module name
	ModuleName = "mynft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mynft"

    
)

var (
	ParamsKey = []byte("p_mynft")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
