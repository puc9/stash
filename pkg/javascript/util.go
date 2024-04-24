package javascript

import (
	"fmt"
	"os"
	"time"

	"github.com/stashapp/stash/pkg/logger"
)

type Util struct{}

func (u *Util) sleepFunc(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func (u *Util) loadTextFile(filePath string) string {
    // TODO: Add relative path logic
	data, err := os.ReadFile(filePath)
	if err != nil {
        logger.Debugf("Unable to read text file: %w", err)
		return ""
	}

    return string(data)
}

func (u *Util) AddToVM(globalName string, vm *VM) error {
	util := vm.NewObject()
	if err := SetAll(util,
		ObjectValueDef{"Sleep", u.sleepFunc},
		ObjectValueDef{"LoadTextFile", u.loadTextFile},
	); err != nil {
		return err
	}

	if err := vm.Set(globalName, util); err != nil {
		return fmt.Errorf("unable to set util: %w", err)
	}

	return nil
}
