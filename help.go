package main

import "github.com/jroimartin/gocui"

type help struct {
	c        chan event
	contents string
	written  bool
}

func (self *help) deliver(e event) {
	self.c <- e
}

func (self *help) makechan() {
	self.c = make(chan event)
}

func (self *help) draw(v *gocui.View, refresh bool) {
	if refresh {
		self.written = false
	}

	if self.written {
		return
	}

	v.Write([]byte(self.contents))

	self.written = true
}

func (self *help) init() {
	self.contents = `
TAB switches tabs. PgUp/PgDown scrolls main window.

C-d                     Scroll down assembly
C-u                     Scroll up assembly
Left                    Center assembly on current instruction
Up                      Command history up
Down                    Command history down
C-x                     Enter hotkey mode, and then:
  S                     Start device
  s                     Stop/step device
  c                     Continue device
  R                     Restart device
C-b                     Bump stack

C-Q                     Log view
C-W                     Device output view
C-E                     Source view
C-R                     VM opcode view
C-T                     Memory view
C-Y                     Stack view
C-H                     Help

Debugger commands:

list <arg>              Center assembly on <arg> (addr/fn)
functions               List all known functions
functions <arg>         All functions matching regex
start                   Start device
step                    Stop/step device
cont                    Continue device
restart                 Restart device
break <arg>             Set breakpoint on <arg> (addr/fn)
clear <arg>             Clear breakpoint on <arg> (addr/fn)
runto <arg>             Execute instructions until <arg> 
stepover                If at CALL, run until that function returns
follow / nofollow       Assembly listing does / doesn't follow PC
dump <addr>             Load <addr> into memory dump

r/8 r24                 Display value of r24
r/16 r24:25             Display word in r24:25
r/s @r24:25             Read string in memory pointed to at r24:r25
r/m @r24                Dump memory pointed to by r24:r25
r/m @4000               Dump memory at 0x4000

load <file>             Load C source from <file>
compile                 Compile loaded C source
vmload                  Request device load VM from flash
vmexec                  Request device execute loaded code
flash                   Write compiled code to flash

echo foo or whatever    Is this thing on?
 
Everything is broken. Everything is awful. Have fun!
`
}

func (self *help) loop() {
	self.init()

	for {
		e := <-self.c
		switch e.kind {
		}
		redraw()
	}
}
