package flowmanager

import (
	"strconv"
	"testing"

	"github.com/coreos/ksched/pkg/types"
	"github.com/coreos/ksched/pkg/util"
	pb "github.com/coreos/ksched/proto"
	"github.com/coreos/ksched/scheduling/flow/costmodel"
	"github.com/coreos/ksched/scheduling/flow/dimacs"
)

func TestAddResourceNode(t *testing.T) {
	//TODO
}

// Create a Graph Manager using the trivial cost model
func createGMTrivial() *graphManager {
	resourceMap := types.NewResourceMap()
	taskMap := types.NewTaskMap()
	leafResourceIDs := make(map[types.ResourceID]struct{})
	dimacsStats := &dimacs.ChangeStats{}
	costModeler := costmodel.NewTrivial(resourceMap, taskMap, leafResourceIDs)
	gm := NewGraphManager(costModeler, leafResourceIDs, dimacsStats)
	return gm
}

// TODO: Helper functions that may just be duplicated into each unit test later
func createMachine(rtnd *pb.ResourceTopologyNodeDescriptor, machineName string) *pb.ResourceDescriptor {
	util.SeedStringRng(machineName)
	rID := util.GenerateResourceID()
	rd := rtnd.ResourceDesc
	rd.Uuid = strconv.FormatUint(uint64(rID), 10)
	rd.Type = pb.ResourceDescriptor_ResourceMachine
	return rd
}

func createTask(jd pb.JobDescriptor, jobIDSeed uint64) {
	util.SeedIntRng(int64(jobIDSeed))
	jobID := util.GenerateJobID()
	jd.Uuid = strconv.FormatUint(uint64(jobID), 10)
}
