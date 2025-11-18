package platforms

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type Platform interface {
	fmt.Stringer

	Arch() ArchName

	OS() OST

	OSVersion() OperatingSystemVersion

	OSRevision() OperatingSystemRevision
}

type Info = Platform

////////////////////////////////////////////////////////////////////////////////

type InfoBuilder struct {
	Arch ArchName
	OST  OperatingSystemType
	OSV  OperatingSystemVersion
	OSR  OperatingSystemRevision
}

func (inst *InfoBuilder) Info() Info {

	info := new(innerPlatformInfo)

	info.arch = inst.Arch
	info.ost = inst.OST
	info.osr = inst.OSR
	info.osv = inst.OSV

	return info
}

////////////////////////////////////////////////////////////////////////////////

type innerPlatformInfo struct {
	arch ArchName
	ost  OperatingSystemType
	osv  OperatingSystemVersion
	osr  OperatingSystemRevision
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
	panic("unimplemented")
}

// OSVersion implements Platform.
func (i *innerPlatformInfo) OSVersion() OperatingSystemVersion {
	panic("unimplemented")
}

func (i *innerPlatformInfo) String() string {

	ver := i.osv
	rev := i.osr
	ost := i.ost
	ar := i.arch

	b := new(strings.Builder)

	b.WriteString("platform:")
	b.WriteString(ost.Normalize().String())
	b.WriteString(":")
	b.WriteString(rev.String())
	b.WriteString("/")
	b.WriteString(ver.Normalize().String())
	b.WriteString("?arch=")
	b.WriteString(ar.String())

	return b.String()

}

////////////////////////////////////////////////////////////////////////////////
// current

var theCurrentPlatformInfo Platform

func Current() Platform {
	info := theCurrentPlatformInfo
	if info == nil {
		loader := new(innerCurrentPlatformInfoLoader)
		info = loader.load()
		theCurrentPlatformInfo = info
	}
	return info
}

////////////////////////////////////////////////////////////////////////////////
// EOF
