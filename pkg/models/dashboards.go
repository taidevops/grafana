package models

// DashboardErr represents a dashboard error.
type DashboardErr struct {
	StatusCode int
	Status     string
	Reason     string
}

// Equal returns whether equal to another DashboardErr.
func (e DashboardErr) Equal(o DashboardErr) bool {
	return o.StatusCode == e.StatusCode && o.Status == e.Status && o.Reason == e.Reason
}

func (e DashboardErr) Error() string {
	if e.Reason != "" {
		return e.Reason
	}
	return "Dashboard Error"
}


