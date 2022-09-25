package ext

import (
	"github.com/fairytale5571/secExt/pkg/errs"
	"github.com/fairytale5571/secExt/pkg/logger"
	"strings"
)

var (
	dataBus chan extData
)

type extData struct {
	FuncName string
	Args     []string
	logger   *logger.Wrapper
}

func RVExtensionHandle(funcName string, args []string) string {
	dataBus <- extData{
		FuncName: funcName,
		Args:     args,
	}
	return ""
}

func handleLoop() {
	var err error
	for {
		select {
		case data, ok := <-dataBus:
			if !ok {
				// channel is busto lets exit
				logger.Print("Channel busted, exiting\n")
				return
			}
			funcName, args := data.FuncName, data.Args
			switch funcName {
			case ":NEW:UNIT:":
				err = rvNewUnitHandler(args)
			case ":NEW:VEH:":
				err = rvNewVehicleHandler(args)
			case ":EVENT:":
				err = rvEventHandler(args)
			case ":UPDATE:UNIT:":
				err = rvUpdateUnitHandler(args)
			case ":UPDATE:VEH:":
				err = rvUpdateVehicleHandler(args)
			case ":SAVE:":
				err = rvSaveHandler(args)
			case ":LOG:":
			case ":START:":
				err = rvStartHandler(args)
			case ":FIRED:":
				err = rvFiredHandler(args)
			default:
				err = errs.ErrorNotHandled
			}

			if err != nil {
				logger.Printf("ERR: %s, Func: %s, Args: %s\n", err, funcName, strings.Join(args, ","))
			}
			logger.Printf("Func: %s, Args: %s\n", funcName, strings.Join(args, ","))
		}
	}
}
