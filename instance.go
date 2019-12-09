package gaslight

// Instance is a structure containing information pertaining to a VM instance.
type Instance struct {
	Name string
}

// InstanceService is the interface that Gaslight uses to obtain information
// about VM instances.
type InstanceService interface {
	ManagerInstances() ([]*Instance, error)
}

// SwarmJoinService is the interface that Gaslight uses to attempt to join a
// Docker swarm.  The implementations of this service handle the specific
// strategy to remotely obtain a join token and then issue the docker swarm
// join command.
type SwarmJoinService interface {

	// Join attempts to join the VM instance that is running Gaslight to the
	// VM instance represented by the provided Instance structure.
	Join(*Instance) (bool, error)
}
