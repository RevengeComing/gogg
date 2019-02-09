package gogg

type Node struct {
	Host string
	Port int
}

// DistributedGoroutine Api must be something like:
// channel := make(chan interface{}, 10) // or type of the thing you want to transfer
// dg := DistributedGoroutine{
// 		name: "something",
// 		channel: channel,
// }
// go goroutinedFunction(args ... interface{}, channel)
// and Add it to a group or groups.
type DistributedGoroutine struct {
	name    string
	isLocal bool
	node    Node
	channel chan interface{}
}

type GoroutineGroupTable interface {
	// Run pull server
	Run()

	// Returns all nodes
	GetNodes() []*Node
	// Link local node to destiny node recieving all nodes linked to destiny node and add them to all nodes
	LinkNode(node Node) error

	// Create a gg (goroutine group) and notify all linked nodes
	Create(groupName string)
	// Delete a gg and notify all linked nodes
	Delete(groupName string)

	// Join a goroutine to a gg and notify all linked nodes
	Join(goroutineName string, goroutineGroup string) error
	// Leave a goroutine from a gg and notify all linked nodes
	Leave(goroutineName string, goroutineGroup string) error

	// Get all members of a gg
	GetMembers(groupName string) ([]*DistributedGoroutine, error)
	GetLocalMembers(groupName string) ([]*DistributedGoroutine, error)
	// Return all ggs
	WhichGroups() []string

	// Transfer a message to a gg and all its goroutines's channel
	TransferToGroup(message interface{}, groupName string) error
}
