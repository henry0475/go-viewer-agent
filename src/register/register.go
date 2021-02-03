package register

import (
	"context"
	"log"
	"runtime"

	"github.com/henry0475/go-viewer-agent/src/reporter"
	collector "github.com/henry0475/go-viewer-agent/src/reporter/protos"
	"github.com/henry0475/go-viewer-agent/src/tools/ids"
)

// Register is ...
type Register struct {
	reporterHandler *reporter.Reporter
}

// NewRegister is the constructor
func NewRegister(reporterHandler *reporter.Reporter) *Register {
	var h = new(Register)
	h.reporterHandler = reporterHandler
	h.register()

	return h
}

func (r *Register) register() {
	reply, err := r.reporterHandler.Collector.RegisterNode(context.Background(), &collector.RegisterNodeRequest{
		Agent: &collector.AgentInfo{
			NodeID: ids.GetNodeID(),
			NumCPU: int64(runtime.NumCPU()),
		},
	})
	if err != nil {
		log.Println(err.Error())
	}

	if reply.GetStatus() != 0 {
		log.Println(reply.GetMsg())
	}
}
