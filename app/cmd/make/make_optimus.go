package make

import (
	"fmt"
	"github.com/pjebs/optimus-go/generator"
	"github.com/spf13/cobra"
	"item-server/pkg/console"
)

var CmdMakeOptimus = &cobra.Command{
	Use:   "optimus",
	Short: "Crate optimus number, example: make optimus",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	o, err := generator.GenerateSeed()
	if err == nil {
		fmt.Printf(" prime:%v\n inverse:%v\n random:%v\n", o.Prime(), o.ModInverse(), o.Random())
	} else {
		console.Error(err.Error())
	}

}
