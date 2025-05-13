package arrows

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	yammmcue "github.com/wyrth-io/whit/internal/for/cue"
	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/validation"
)

// Arrows represents neo4j arrows content.
type Arrows struct {
	file string
	// contents map[string]interface{}
	Parsed *Graph `json:"graph"`
}

// Parse parses the arrows contents in the file set in the receiver.
// The unmarshalled information is stored in the receiver.
// This function will panic with a catch.Error if the file cannot be
// opened or unmarshalled as JSON.
// Will possibly panic in unknown location with an error if not given an IssueCollector that terminates on fatal issue.
func (arrows *Arrows) Parse(ic validation.IssueCollector) {
	// TODO: Add method in collector to ensure that collector terminates on fatal.
	file, err := os.Open(arrows.file)
	ic.CollectFatalIfErrorf("Error opening arrows json file: %s", err)
	defer pio.Close(file)

	data, err := io.ReadAll(file)
	ic.CollectFatalIfErrorf("error reading json data: %s", err)

	err = json.Unmarshal(data, arrows)
	ic.CollectFatalIfErrorf("Error unmarshalling arrows json data: %s", err)

	// Arrows uses two formats, one with an extra "graph" node (the saved format), and one without (the UI export format),
	// if the attempt above did not find  a "graph" element, parse again starting inside instead of outside the graph.
	if arrows.Parsed == nil {
		var theGraph Graph
		err = json.Unmarshal(data, &theGraph)
		ic.CollectFatalIfErrorf("Error unmarshalling arrows json data: %s", err)
		arrows.Parsed = &theGraph
	}
}

// New creates a new Arrows context represengint a filename and an empty (to be parsed graph).
func New(file string) Arrows {
	return Arrows{file: file}
}

// CueMarshalMeta interprets the graph as a meta graph and outputs to a cue schema (on stdout in first version).
func (arrows *Arrows) CueMarshalMeta(w io.Writer, ic validation.IssueCollector) {
	name := strings.TrimSuffix(arrows.file, filepath.Ext(arrows.file))
	ctx := arrows.Parsed.buildYammmModel(name, ic)
	if ic.HasErrors() || ic.HasFatal() {
		ic.Collectf(validation.Warning, "No output generated due to earlier errors")
		return
	}
	out := pio.WriterOn(w)
	yammmcue.Marshal(ctx, out)
}

// MarshalYammm interprets the graph as a meta graph and outputs a yammm schema in JSON.
// If the given name is empty, the name of the arrows file without file extension will be used as the
// name of the model. This name is important when generating go code from the model.
func (arrows *Arrows) MarshalYammm(name string, w io.Writer, ic validation.IssueCollector) {
	if len(name) == 0 {
		name = strings.TrimSuffix(arrows.file, filepath.Ext(arrows.file))
	}
	ctx := arrows.Parsed.buildYammmModel(name, ic)
	if ic.HasErrors() || ic.HasFatal() {
		ic.Collectf(validation.Warning, "No output generated due to earlier errors")
		return
	}
	if err := ctx.WriteModelAsJSON(w); err != nil {
		ic.Collectf(validation.Error, err.Error())
	}
}

// CueMarshalInstance interprets the graph as an instance graph and outputs to a cue schema (on stdout in first version).
// The intention is for the instance graph to be compliant with an already created schema.
// TODO: There needs to be validation, but this simply outputs content for the time being.
func (arrows *Arrows) CueMarshalInstance(ic validation.IssueCollector) {
	arrows.Parsed.CueMarshalInstance(ic)
}
