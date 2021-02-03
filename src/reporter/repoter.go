package reporter

import (
	"fmt"
	"log"

	"github.com/henry0475/go-viewer-agent/src/options"
	collector "github.com/henry0475/go-viewer-agent/src/reporter/protos"
	"google.golang.org/grpc"
)

// Reporter is ...
type Reporter struct {
	conn      *grpc.ClientConn
	Collector collector.CollectorClient
}

var reporterHandler *Reporter

// NewReporter is the constructor
func NewReporter(gRPCOpt *options.GRPCConfigs) *Reporter {
	reporterHandler = new(Reporter)

	// altsTC := alts.NewClientCreds(alts.DefaultClientOptions())
	// conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCOpt.Address, gRPCOpt.Port), grpc.WithTransportCredentials(altsTC))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCOpt.Address, gRPCOpt.Port), grpc.WithInsecure())
	reporterHandler.conn = conn
	if err != nil {
		log.Println("Cannot connect: " + err.Error())
		return nil
	}
	reporterHandler.Collector = collector.NewCollectorClient(reporterHandler.conn)

	return reporterHandler
}

// GetRepoter returns the handler of the repoter
func GetRepoter() *Reporter {
	return reporterHandler
}
