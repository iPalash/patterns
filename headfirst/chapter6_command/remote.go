package chapter6_command

type RemoteControl interface {
	Assign(slot int, on Command, off Command)
	On(slot int)
	Off(slot int)
	Undo()
}

type RemoteControlImpl struct {
	slots       int
	undo        Command
	onCommands  []Command
	offCommands []Command
}

func NewRemoteControl(slots int) RemoteControl {
	onCommands := make([]Command, slots)
	offCommands := make([]Command, slots)
	for i := range onCommands {
		onCommands[i] = EmptyCommand()
		offCommands[i] = EmptyCommand()
	}
	rc := &RemoteControlImpl{slots: slots, onCommands: onCommands, offCommands: offCommands}
	return rc
}

func (rc *RemoteControlImpl) Assign(slot int, on Command, off Command) {
	rc.onCommands[slot] = on
	rc.offCommands[slot] = off
}

func (rc *RemoteControlImpl) On(slot int) {
	rc.onCommands[slot].execute()
	rc.undo = rc.onCommands[slot]
}

func (rc *RemoteControlImpl) Off(slot int) {
	rc.offCommands[slot].execute()
	rc.undo = rc.offCommands[slot]
}

func (rc *RemoteControlImpl) Undo() {
	rc.undo.undo()
}
