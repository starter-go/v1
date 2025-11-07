package buckets

import "fmt"

type DriverManager interface {
	OpenBucket(ctx *BucketContext) (Bucket, error)

	GetDriver(name string) (Driver, error)
}

////////////////////////////////////////////////////////////////////////////////

type DriverManagerContext struct {
	drivers []*DriverRegistration
}

func (inst *DriverManagerContext) AddDriver(driver Driver) {
	if driver == nil {
		return
	}
	dr := driver.Registration()
	inst.drivers = append(inst.drivers, dr)
}

func (inst *DriverManagerContext) AddRegistration(dr *DriverRegistration) {
	if dr == nil {
		return
	}
	inst.drivers = append(inst.drivers, dr)
}

////////////////////////////////////////////////////////////////////////////////

type DefaultDriverManager struct {
	context *DriverManagerContext
}

func (inst *DefaultDriverManager) _impl() DriverManager {
	return inst
}

func (inst *DefaultDriverManager) OpenBucket(ctx *BucketContext) (Bucket, error) {
	cfg := &ctx.Configuration
	driver, err := inst.GetDriver(cfg.Driver)
	if err != nil {
		return nil, err
	}
	return driver.OpenBucket(ctx)
}

func (inst *DefaultDriverManager) isDriverReady(wantName string, driver *DriverRegistration) bool {

	if wantName == "" || driver == nil {
		return false
	}

	if driver.Enabled {
		return false
	}

	if driver.Driver == nil {
		return false
	}

	return (wantName == driver.Name)
}

func (inst *DefaultDriverManager) GetDriver(name string) (Driver, error) {
	all := inst.context.drivers
	for _, r1 := range all {
		if inst.isDriverReady(name, r1) {
			return r1.Driver, nil
		}
	}
	return nil, fmt.Errorf("buckets.DefaultDriverManager: no driver named [%s]", name)
}

////////////////////////////////////////////////////////////////////////////////

func NewDriverManager(ctx *DriverManagerContext) DriverManager {
	if ctx == nil {
		ctx = new(DriverManagerContext)
	}
	dm := new(DefaultDriverManager)
	dm.context = ctx
	return dm
}

////////////////////////////////////////////////////////////////////////////////
