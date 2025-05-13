// Package yammm (Yet Another Meta-Meta Model) contains a simple meta-meta model for describing data with relations.
package yammm

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/wyrth-io/whit/internal/validation"
)

var ucName = regexp.MustCompile(`^[A-Z]\w*$`)
var lcName = regexp.MustCompile(`^[a-z]\w*$`)

// IsUCName returns true if the given name is ok as an initially upper cased name.
func IsUCName(s string) bool {
	return ucName.MatchString(s)
}

// IsLCName returns true if the given name is ok as an initially lower cased name.
func IsLCName(s string) bool {
	return lcName.MatchString(s)
}

// WhitNamespace is a UUID v5 namespace for SHA1 based deterministic UUIDs.
var WhitNamespace uuid.UUID

// init creates the UUID v5 namespace UUID based on the URL to the whit repo. This namespace
// is used for deterministic UUIDs of instances of whit/yammm types.
func init() {
	WhitNamespace = uuid.NewSHA1(uuid.NameSpaceURL, []byte("https://github.com/wyrth-io/whit"))
}

// ExternalValidator is an interface for an implementation that can be registered to be part of validation.
// It will be called back for each type, property, association, association property and composition.
// See [AddValidator] for registration.
type ExternalValidator interface {
	ValidateType(t *Type, ic validation.IssueCollector)
	ValidateProperty(t *Type, p *Property, ic validation.IssueCollector)
	ValidateAssociation(t *Type, a *Association, ic validation.IssueCollector)
	ValidateAssociationProperty(t *Type, a *Association, p *Property, ic validation.IssueCollector)
	ValidateComposition(t *Type, c *Composition, ic validation.IssueCollector)
}

var externalValidators []ExternalValidator

// RegisterValidator registers an external validator with this validator. All registered validators will
// receive callbacks to all relevant methods during validation.
func RegisterValidator(ev ExternalValidator) {
	externalValidators = append(externalValidators, ev)
}
