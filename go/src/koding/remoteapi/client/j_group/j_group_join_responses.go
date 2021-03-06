package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JGroupJoinReader is a Reader for the JGroupJoin structure.
type JGroupJoinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupJoinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupJoinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupJoinOK creates a JGroupJoinOK with default headers values
func NewJGroupJoinOK() *JGroupJoinOK {
	return &JGroupJoinOK{}
}

/*JGroupJoinOK handles this case with default header values.

OK
*/
type JGroupJoinOK struct {
	Payload JGroupJoinOKBody
}

func (o *JGroupJoinOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.join/{id}][%d] jGroupJoinOK  %+v", 200, o.Payload)
}

func (o *JGroupJoinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupJoinOKBody j group join o k body
swagger:model JGroupJoinOKBody
*/
type JGroupJoinOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupJoinOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupJoinOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupJoinOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupJoinOKBodyAO0

	var jGroupJoinOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupJoinOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupJoinOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupJoinOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupJoinOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupJoinOKBodyAO0)

	jGroupJoinOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupJoinOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group join o k body
func (o *JGroupJoinOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
