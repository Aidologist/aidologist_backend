package models

type Project struct {
	Id			int
	Name		string
	Image		string
	Desc		string
	// Time records
	StartTime	time.Time
	UpdateTIme	time.Time
	ExpectEnd[]	time.Time
	// counters
}

type ProjectType struct {
	Id			int
	Name		string
	Father		ProjectType
	Children	ProjectType
}

func (this *Project) AddProject() (int64, bool) {
	success, o := databases.ConnectOrm()
	if success {
		index, err := o.Insert(this)
		if err == nil {
			return index, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

