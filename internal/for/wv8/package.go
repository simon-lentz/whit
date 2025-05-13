// Package wv8 contains the whit/yammm support for the Weaviate vector/structured store.
package wv8

import (
	"fmt"

	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// EdgeName produces the name used in Weaviate for a class representing an association edge between
// the given type t, and the type referenced by the given association.
func EdgeName(t *yammm.Type, a *yammm.Association) string {
	return fmt.Sprintf("EDGE_%s_%s_%s", t.Name, a.Name, a.To)
}

// FormatPropName returns a de-capitalized name for a yammm name (that may be with initial capital letter).
func FormatPropName(yammmName string) string {
	return utils.DeCapitalize(yammmName)
}
