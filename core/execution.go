package core

import (
	"log"
)

func Execute(w *Workspace, p *Project, t *Task) {
	log.Printf("Running: %v/%v", p.Name, t.Name)
}
