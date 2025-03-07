// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type execsnoopWithCwdBufT struct{ Buf [32768]uint8 }

type execsnoopWithCwdEvent struct {
	MntnsId   uint64
	Timestamp uint64
	Pid       uint32
	Ppid      uint32
	Uid       uint32
	Gid       uint32
	Loginuid  uint32
	Sessionid uint32
	Retval    int32
	ArgsCount int32
	ArgsSize  uint32
	Comm      [16]uint8
	Cwd       [4096]uint8
	Args      [7680]uint8
	_         [4]byte
}

// loadExecsnoopWithCwd returns the embedded CollectionSpec for execsnoopWithCwd.
func loadExecsnoopWithCwd() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ExecsnoopWithCwdBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load execsnoopWithCwd: %w", err)
	}

	return spec, err
}

// loadExecsnoopWithCwdObjects loads execsnoopWithCwd and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*execsnoopWithCwdObjects
//	*execsnoopWithCwdPrograms
//	*execsnoopWithCwdMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadExecsnoopWithCwdObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadExecsnoopWithCwd()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// execsnoopWithCwdSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopWithCwdSpecs struct {
	execsnoopWithCwdProgramSpecs
	execsnoopWithCwdMapSpecs
}

// execsnoopWithCwdSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopWithCwdProgramSpecs struct {
	IgExecveE *ebpf.ProgramSpec `ebpf:"ig_execve_e"`
	IgExecveX *ebpf.ProgramSpec `ebpf:"ig_execve_x"`
}

// execsnoopWithCwdMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopWithCwdMapSpecs struct {
	Bufs                 *ebpf.MapSpec `ebpf:"bufs"`
	Events               *ebpf.MapSpec `ebpf:"events"`
	Execs                *ebpf.MapSpec `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.MapSpec `ebpf:"gadget_mntns_filter_map"`
}

// execsnoopWithCwdObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopWithCwdObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopWithCwdObjects struct {
	execsnoopWithCwdPrograms
	execsnoopWithCwdMaps
}

func (o *execsnoopWithCwdObjects) Close() error {
	return _ExecsnoopWithCwdClose(
		&o.execsnoopWithCwdPrograms,
		&o.execsnoopWithCwdMaps,
	)
}

// execsnoopWithCwdMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopWithCwdObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopWithCwdMaps struct {
	Bufs                 *ebpf.Map `ebpf:"bufs"`
	Events               *ebpf.Map `ebpf:"events"`
	Execs                *ebpf.Map `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.Map `ebpf:"gadget_mntns_filter_map"`
}

func (m *execsnoopWithCwdMaps) Close() error {
	return _ExecsnoopWithCwdClose(
		m.Bufs,
		m.Events,
		m.Execs,
		m.GadgetMntnsFilterMap,
	)
}

// execsnoopWithCwdPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopWithCwdObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopWithCwdPrograms struct {
	IgExecveE *ebpf.Program `ebpf:"ig_execve_e"`
	IgExecveX *ebpf.Program `ebpf:"ig_execve_x"`
}

func (p *execsnoopWithCwdPrograms) Close() error {
	return _ExecsnoopWithCwdClose(
		p.IgExecveE,
		p.IgExecveX,
	)
}

func _ExecsnoopWithCwdClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed execsnoopwithcwd_bpfel.o
var _ExecsnoopWithCwdBytes []byte
