package meritop

// Task is a logic repersentation of a computing unit.
// Each task contain at least one Node.
// Each task has exact one master Node and might have multiple salve Nodes.
type Task interface {
	// This is useful to bring the task up to speed from scratch or if it recovers.
	Init(config Config)

	// Task need to finish up for exit, last chance to save work?
	Exit()

	// Nodes returns the IDs of the node associated with this task.
	Nodes() []Node
	Master() Node
	Slaves() []Node

	// These are called by framework implementation so that task implementation can
	// reacts to parent or children restart.
	ParentRestart(taskID uint64)
	ChildRestart(taskID uint64)

	ParentDie(taskID uint64)
	ChildDie(taskID uint64)

	// Ideally, we should also have the following:
	ParentReady(taskID uint64, data []byte)
	ClientReady(taskID uint64, data []byte)

	// This give the task an opportunity to cleanup and regroup.
	SetEpoch(epochID uint64)

	// Some hooks that need for master slave etc.
	BecameMaster(nodeID uint64)
	BecameSlave(nodeID uint64)

	// This method make framework available to task implementation.
	SetFramework(framework Framework)
}