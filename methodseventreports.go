package objectsmispformat

func NewEventReports() *EventReports {
	return &EventReports{}
}

// SetName устанавливает значение для Name
func (r *EventReports) SetName(v string) {
	r.Name = v
}

// GetName возвращает значение Name
func (r *EventReports) GetName() string {
	return r.Name
}

// SetContent устанавливает значение для Content
func (r *EventReports) SetContent(v string) {
	r.Content = v
}

// GetContent возвращает значение Content
func (r *EventReports) GetContent() string {
	return r.Content
}

// SetDistribution устанавливает значение для Distribution
func (r *EventReports) SetDistribution(v string) {
	r.Distribution = v
}

// GetDistribution возвращает значение Distribution
func (r *EventReports) GetDistribution() string {
	return r.Distribution
}

// Comparison выполняет сравнение двух объектов типа EventReports
func (r *EventReports) Comparison(newReport *EventReports) bool {
	if r.Name != newReport.Name {
		return false
	}

	if r.Content != newReport.Content {
		return false
	}

	if r.Distribution != newReport.Distribution {
		return false
	}

	return true
}

// MatchingAndReplacement выполняет сопоставление элементов объекта и их замену
func (r *EventReports) MatchingAndReplacement(newReport EventReports) EventReports {
	if r.Name != "" && r.Name != newReport.Name {
		newReport.Name = r.Name
	}

	if r.Content != "" && r.Content != newReport.Content {
		newReport.Content = r.Content
	}

	if r.Distribution != "" && r.Distribution != newReport.Distribution {
		newReport.Distribution = r.Distribution
	}

	return newReport
}
