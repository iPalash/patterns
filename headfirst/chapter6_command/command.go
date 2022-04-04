package chapter6_command

import (
	"headfirst/chapter6_command/appliances"
)

type Command interface {
	execute()
	undo()
}

type NoOp struct {
}

func EmptyCommand() Command {
	return &NoOp{}
}

func (n *NoOp) execute() {
}

func (n *NoOp) undo() {
}

type TVOnCommand struct {
	*appliances.TV
}

func NewTVOnCommand(t *appliances.TV) Command {
	return &TVOnCommand{t}
}

func (t *TVOnCommand) execute() {
	t.TV.On()
}

func (t *TVOnCommand) undo() {
	t.TV.Off()
}

type TVOffCommand struct {
	*appliances.TV
}

func NewTVOffCommand(t *appliances.TV) Command {
	return &TVOffCommand{t}
}

func (t *TVOffCommand) execute() {
	t.TV.Off()
}

func (t *TVOffCommand) undo() {
	t.TV.On()
}

type LightOnCommand struct {
	*appliances.Light
}

func NewLightOnCommand(t *appliances.Light) Command {
	return &LightOnCommand{t}
}

func (t *LightOnCommand) execute() {
	t.Light.On()
}

func (t *LightOnCommand) undo() {
	t.Light.Off()
}

type LightOffCommand struct {
	*appliances.Light
}

func NewLightOffCommand(t *appliances.Light) Command {
	return &LightOffCommand{t}
}

func (t *LightOffCommand) execute() {
	t.Light.Off()
}

func (t *LightOffCommand) undo() {
	t.Light.On()
}

type MacroCommand struct {
	cmds []Command
}

func NewMacroCommand(cmds ...Command) Command {
	return &MacroCommand{
		cmds: cmds,
	}
}

func (m *MacroCommand) execute() {
	for i := range m.cmds {
		m.cmds[i].execute()
	}
}

func (m *MacroCommand) undo() {
	for i := len(m.cmds) - 1; i >= 0; i-- {
		m.cmds[i].undo()
	}
}
