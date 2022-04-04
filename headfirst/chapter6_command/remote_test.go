package chapter6_command

import (
	"headfirst/chapter6_command/appliances"
	"testing"
)

func TestNewRemoteControl(t *testing.T) {
	rc := NewRemoteControl(3)
	t.Run("NoOp", func(t *testing.T) {
		rc.On(1)
		rc.Off(1)
	})
	tv := appliances.NewTV()
	rc.Assign(1, NewTVOnCommand(tv), NewTVOffCommand(tv))
	t.Run("TV On/off", func(t *testing.T) {
		rc.On(1)
		rc.Off(1)
		rc.Undo()
	})

	light := appliances.NewLight()
	rc.Assign(0, NewLightOnCommand(light), NewLightOffCommand(light))
	t.Run("Light On/off", func(t *testing.T) {
		rc.On(0)
		rc.Off(0)
		rc.Undo()
	})

	enterRoomMacro := NewMacroCommand(NewLightOnCommand(light), NewTVOnCommand(tv))
	exitRoomMacro := NewMacroCommand(NewLightOffCommand(light), NewTVOffCommand(tv))
	rc.Assign(2, enterRoomMacro, exitRoomMacro)
	t.Run("Room enter/exit macro", func(t *testing.T) {
		rc.On(2)
		rc.Off(2)
		rc.Undo()
	})
}
