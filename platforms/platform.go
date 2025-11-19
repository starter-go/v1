package platforms

import (
	"encoding/json"
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type Platform interface {
	fmt.Stringer

	Arch() ArchName

	OS() OST

	OSName() OperatingSystemName

	OSVersion() OperatingSystemVersion

	OSRevision() OperatingSystemRevision

	GetProperties(dst map[string]string) map[string]string
}

type Info = Platform

////////////////////////////////////////////////////////////////////////////////

type InfoBuilder struct {
	Arch ArchName

	OST OperatingSystemType
	OSN OperatingSystemName
	OSV OperatingSystemVersion
	OSR OperatingSystemRevision

	props map[string]string
}

func (inst *InfoBuilder) Info() Info {

	info := new(innerPlatformInfo)

	info.arch = inst.Arch.Normalize()
	info.ost = inst.OST.Normalize()
	info.osn = inst.OSN.Normalize()
	info.osr = inst.OSR
	info.osv = inst.OSV.Normalize()
	info.props = inst.props

	return info
}

func (inst *InfoBuilder) String() string {
	const (
		prefix = "  "
		indent = "\t"
	)
	bin, err := json.MarshalIndent(inst, prefix, indent)
	if err != nil {
		return err.Error()
	}
	return string(bin)
}

////////////////////////////////////////////////////////////////////////////////

type InfoLoader interface {
	Load() Info

	OnLoad(ib *InfoBuilder) error
}

////////////////////////////////////////////////////////////////////////////////

type innerPlatformInfo struct {
	arch ArchName

	ost OperatingSystemType
	osn OperatingSystemName
	osv OperatingSystemVersion
	osr OperatingSystemRevision

	props map[string]string
}

// GetProperties implements Platform.
func (inst *innerPlatformInfo) GetProperties(dst map[string]string) map[string]string {
	src := inst.props
	if src == nil {
		src = make(map[string]string)
	}
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// OSName implements Platform.
func (i *innerPlatformInfo) OSName() OperatingSystemName {
	return i.osn
}

// Arch implements Platform.
func (i *innerPlatformInfo) Arch() ArchName {
	return i.arch
}

// OS implements Platform.
func (i *innerPlatformInfo) OS() OST {
	return i.ost
}

// OSRevision implements Platform.
func (i *innerPlatformInfo) OSRevision() OperatingSystemRevision {
	return i.osr
}

// OSVersion implements Platform.
func (i *innerPlatformInfo) OSVersion() OperatingSystemVersion {
	return i.osv
}

func (i *innerPlatformInfo) String() string {

	ver := i.osv
	rev := i.osr
	ost := i.ost
	osn := i.osn
	ar := i.arch

	b := new(strings.Builder)

	b.WriteString("platform:")
	b.WriteString(ost.Normalize().String())
	b.WriteString("/")
	b.WriteString(osn.String())
	b.WriteString(":")
	b.WriteString(rev.String())
	b.WriteString("/")
	b.WriteString(ver.Normalize().String())
	b.WriteString("?arch=")
	b.WriteString(ar.String())

	return b.String()
}

////////////////////////////////////////////////////////////////////////////////

func innerLoadWithLoader(loader InfoLoader) Info {
	ib := new(InfoBuilder)
	if loader != nil {
		loader.OnLoad(ib)
	}
	return ib.Info()
}

////////////////////////////////////////////////////////////////////////////////
// current

var theCurrentPlatformInfo Platform

func Current() Platform {
	info := theCurrentPlatformInfo
	if info == nil {
		loader1 := new(innerCurrentPlatformInfoLoader)
		loader2 := loader1.tryInit()
		info = loader2.Load()
		theCurrentPlatformInfo = info
	}
	return info
}

////////////////////////////////////////////////////////////////////////////////
// EOF
