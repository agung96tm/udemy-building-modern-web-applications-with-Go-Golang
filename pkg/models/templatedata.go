package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	BoolMap   map[string]bool
	Data      map[string]any
	CSRFToken string
	Flash     string
	Error     string
}
