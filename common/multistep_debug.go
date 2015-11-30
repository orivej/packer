package common

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/orivej/packer/packer"
)

const actionRetry multistep.StepAction = multistep.ActionHalt + 1

const debugLocationAfterFailedRun multistep.DebugLocation = multistep.DebugLocationBeforeCleanup + 1

func debugFn(ui packer.Ui, loc multistep.DebugLocation, name string, state multistep.StateBag) multistep.StepAction {
	var locationString string
	switch loc {
	case multistep.DebugLocationAfterRun:
		locationString = "after run of"
	case debugLocationAfterFailedRun:
		locationString = "after failed"
	case multistep.DebugLocationBeforeCleanup:
		locationString = "before cleanup of"
	default:
		locationString = "at"
	}

	ui.Say(fmt.Sprintf("Pausing %s step '%s'.", locationString, name))

	result := make(chan multistep.StepAction, 1)
	go func() {
		result <- askForAction(ui)
	}()

	for {
		select {
		case action := <-result:
			return action
		case <-time.After(100 * time.Millisecond):
			if _, ok := state.GetOk(multistep.StateCancelled); ok {
				return multistep.ActionHalt
			}
		}
	}
}

func askForAction(ui packer.Ui) multistep.StepAction {
	for {
		line, err := ui.Ask("[A]bort, or [R]etry step (build may fail even if retry succeeds)? [ar]")
		if err != nil {
			log.Printf("Error asking for input: %s", err)
		}

		switch {
		case len(line) == 0 || line[0] == 'a':
			ui.Say("Abort.")
			return multistep.ActionHalt
		case line[0] == 'r':
			ui.Say("Retry.")
			return actionRetry
		}
		ui.Say(fmt.Sprintf("Incorrect input: %#v", line))
	}
}

// MultistepDebugFn will return a proper multistep.DebugPauseFn to
// use for debugging if you're using multistep in your builder.
func MultistepDebugFn(ui packer.Ui) multistep.DebugPauseFn {
	return func(loc multistep.DebugLocation, name string, state multistep.StateBag) {
		debugFn(ui, loc, name, state)
	}
}

func MultistepDebugSteps(steps []multistep.Step, ui packer.Ui) []multistep.Step {
	debugSteps := make([]multistep.Step, len(steps))
	for i, step := range steps {
		debugSteps[i] = &debugStep{step, ui}
	}
	return debugSteps
}

type debugStep struct {
	step multistep.Step
	ui   packer.Ui
}

func (s *debugStep) Cleanup(state multistep.StateBag) {
	s.step.Cleanup(state)
}

func (s *debugStep) Run(state multistep.StateBag) (action multistep.StepAction) {
	for {
		s.ui.Say("Running...")
		action = s.step.Run(state)

		if action != multistep.ActionHalt {
			return
		}

		s.ui.Say("Debugging...")
		action = debugFn(s.ui, debugLocationAfterFailedRun, typeName(s.step), state)

		if action != actionRetry {
			return
		}
	}
}

func typeName(i interface{}) string {
	return reflect.Indirect(reflect.ValueOf(i)).Type().Name()
}
