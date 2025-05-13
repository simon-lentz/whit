package tc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
)

// DefaultUUIDChecker is a TypeChecker for a UUID in standard UUIDv4 format, or yammm
// special base57 encoded short UUUID, or a local ID to be transformed into a UUID, or
// a transformed UUID with local ID included.
var DefaultUUIDChecker = &uuidChecker{}

type uuidChecker struct {
	// baseChecker
}

func (c *uuidChecker) SyntaxString() string {
	return "UUID"
}
func (c *uuidChecker) TypeString() string {
	return c.SyntaxString()
}

func (c *uuidChecker) BaseType() BaseType {
	return valueType(StringKind)
}

var rLocalUUID = regexp.MustCompile(`\$\$\w+`)
var rGlobalUUID = regexp.MustCompile(`\$\$:\w+:(\w+)`)

// Check checks value is a valid UUID.
// It Accepts: Base57 encoded v4 UUID (22 characters),
// Yammm special reference $$:<22chars base57>:<local id string>,
// and Standard v4 UUID (36 hex chars including - separators) as well as all formats
// supported by google UUID Parse().
func (c *uuidChecker) Check(v any) (bool, string) {
	_, err := c.GetUUID(v)
	if errors.Is(err, ErrRequiresResolution) || err == nil {
		return true, ""
	}
	return false, err.Error()
}

// ErrRequiresResolution is an error describing that a uuid string requires resolution.
var ErrRequiresResolution = errors.New("$$xxx uuid reference requires resolution")

// ErrNotUUIDTypeChecker is an error describing that GetUUID was called on a non UUID type checker.
var ErrNotUUIDTypeChecker = errors.New("the typechecker is not of UUID type")

var nullUUID uuid.UUID

// GetUUID returns a UUID for the given v if it is a string on one of the accepted forms
// of the yammm UUID ($$:shortuid:local, shortuuid (22 chars), or uuid pareable by uuid.UUID.Parse()).
// If the uuid string v is a $$local uuid string the error RequiresResolution is returned.
// Other errors will be return if the string is neither a shortuuid nor parseable by uuid.UUID.Parse().
func (c *uuidChecker) GetUUID(v any) (uid uuid.UUID, err error) {
	if v == nil {
		return uid, fmt.Errorf("nil value is not an UUID")
	}
	if s, ok := v.(string); ok {
		if strings.HasPrefix(s, "$$") {
			if rLocalUUID.MatchString(s) {
				return nullUUID, ErrRequiresResolution // To be replaced by ID replacer
			}
			if rGlobalUUID.MatchString(s) {
				parts := strings.Split(s, ":")
				if len(parts[1]) != 22 {
					return nullUUID, fmt.Errorf("not a 22 chars long base57 encoded UUID")
				}
				if uid, err = shortuuid.DefaultEncoder.Decode(parts[1]); err != nil {
					return nullUUID, err
				}
				return uid, nil
			}
			return nullUUID, fmt.Errorf("$$ prefixed UUID is neither a local nor global UUID")
		}
		if len(s) == 22 {
			if uid, err = shortuuid.DefaultEncoder.Decode(s); err != nil {
				return nullUUID, err
			}
			return uid, nil
		}
		if uid, err = uuid.Parse(s); err != nil {
			return uid, err
		}

		return uid, nil
	}
	return nullUUID, fmt.Errorf("value is not a string based UUID")
}
func (c *uuidChecker) Refine(instr []any) TypeChecker {
	if len(instr) == 0 {
		return c
	}
	return nil
}
