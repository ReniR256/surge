package surge

import "fmt"

// ErrUnsupportedMarshalType is returned when the an unsupported type is
// encountered during marshaling.
type ErrUnsupportedMarshalType struct {
	error
}

func newErrUnsupportedMarshalType(v interface{}) error {
	return ErrUnsupportedMarshalType{error: fmt.Errorf("marshal error: unsupported type %T", v)}
}

// ErrUnsupportedUnmarshalType is returned when the an unsupported type is
// encountered during unmarshaling.
type ErrUnsupportedUnmarshalType struct {
	error
}

func newErrUnsupportedUnmarshalType(v interface{}) error {
	return ErrUnsupportedUnmarshalType{error: fmt.Errorf("unmarshal error: unsupported type %T", v)}
}

// ErrTagNotFound is returned when unmarshaling encounters a struct tag that is
// not supported by the destination struct.
type ErrTagNotFound struct {
	error
}

func newErrTagNotFound(v uint8) error {
	return ErrTagNotFound{error: fmt.Errorf("tag error: \"%v\" not found", v)}
}

// ErrTagOutOfRange is returned when a struct tag is encountered that is not the
// supported range `[0 .. 255)`.
type ErrTagOutOfRange struct {
	error
}

func newErrTagOutOfRange(tag uint8) error {
	return ErrTagOutOfRange{error: fmt.Errorf("tag error: \"%v\" is out of range (must be 0 to 254)", tag)}
}

// ErrTagDuplicate is returned when the a struct tag is appears more than once
// in a struct.
type ErrTagDuplicate struct {
	error
}

func newErrTagDuplicate(tag uint8) error {
	return ErrTagDuplicate{error: fmt.Errorf("tag error: \"%v\" appears more than once", tag)}
}
