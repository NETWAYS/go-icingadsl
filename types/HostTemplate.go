package icingadsl

type HostTemplate struct {
	Name string
}

func (ht *HostTemplate) GetName() string {
	return ht.Name
}
