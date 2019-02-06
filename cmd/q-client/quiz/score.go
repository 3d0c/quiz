// Code generated by go generate;
package quiz

import (
	"flag"

	"github.com/3d0c/cli/pkg"
	"github.com/3d0c/quiz/pkg/rpc"
)

type score struct {
	*cli.General

	addr     string
	endpoint string
}

func init() {
	cli.Register("quiz.score", &score{General: &cli.General{}})
}

func (cmd *score) Register(f *flag.FlagSet) {
	cmd.General.Register(f)
	f.StringVar(&cmd.addr, "addr", "127.0.0.1:5560", "quiz server location")
}

func (cmd *score) Description() string {
	return `
Description:
	Show QUIZ score.

Examples:
	; show quiz score
	q-client quiz.score

	; optionally define quiz server
	q-client quiz.score -addr=192.168.0.10:5560
	`
}

func (cmd *score) Process() error {
	cmd.endpoint = "http://" + cmd.addr + "/score" + "/" + key()
	return nil
}

func (cmd *score) Run(f *flag.FlagSet) ([]byte, error) {
	return rpc.Get(cmd.endpoint, nil)
}
