package arrows

import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_Node_normalize_meta_makes_caption_have_initial_upper_case(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	n := newNode("n0", "cap", []string{})
	n.normalize(ic, true)
	tt.CheckEqual("Cap", n.Caption)
	tt.CheckEqual(0, ic.Count())
}

func Test_Node_normalize_meta_sets_caption_to_first_label_if_caption_is_missing_and_removes_label(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	n := newNode("n0", "", []string{"cap"})
	n.normalize(ic, true)
	tt.CheckEqual("Cap", n.Caption)
	tt.CheckEqual(0, len(n.Labels))
	tt.CheckEqual(0, ic.Count())
}

func Test_Node_normalize_meta_capitalizes_labels_and_drops_empty_labels(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	n := newNode("n0", "cap", []string{"person", "åäö", ""})
	n.normalize(ic, true)
	tt.CheckEqual([]string{"Person", "Åäö"}, n.Labels)
	tt.CheckEqual(2, len(n.Labels))
	tt.CheckEqual(0, ic.Count())
}

func Test_Node_normalize_meta_drops_duplicate_labels(t *testing.T) {
	tt := testutils.NewTester((t))
	ic := validation.NewTerminatingIssueCollector()
	n := newNode("n0", "cap", []string{"LabelA", "LabelB", "LabelA"})
	n.normalize(ic, true)
	tt.CheckEqual([]string{"LabelA", "LabelB"}, n.Labels)
	tt.CheckEqual(2, len(n.Labels))
	tt.CheckEqual(0, ic.Count())
}
