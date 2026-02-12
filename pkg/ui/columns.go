package ui

type columnWidth struct {
	Type int
	Repo int
	Time int
}

func computeWidths(rows []EventRow) columnWidth {
	maxType := len("TYPE")
	maxRepo := len("REPO")
	maxTime := len("WHEN")

	for _, r := range rows {
		if len(r.Type) > maxType {
			maxType = len(r.Type)
		}
		if len(r.Repo) > maxRepo {
			maxRepo = len(r.Repo)
		}
		if len(r.Time) > maxTime {
			maxTime = len(r.Time)
		}
	}

	return columnWidth{
		Type: maxType,
		Repo: maxRepo,
		Time: maxTime,
	}
}
