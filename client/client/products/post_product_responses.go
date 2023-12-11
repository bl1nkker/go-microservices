// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostProductReader is a Reader for the PostProduct structure.
type PostProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPostProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPostProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /products] postProduct", response, response.Code())
	}
}

// NewPostProductOK creates a PostProductOK with default headers values
func NewPostProductOK() *PostProductOK {
	return &PostProductOK{}
}

/*
PostProductOK describes a response with status code 200, with default header values.

Product returns in the response
*/
type PostProductOK struct {
}

// IsSuccess returns true when this post product o k response has a 2xx status code
func (o *PostProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post product o k response has a 3xx status code
func (o *PostProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post product o k response has a 4xx status code
func (o *PostProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post product o k response has a 5xx status code
func (o *PostProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post product o k response a status code equal to that given
func (o *PostProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post product o k response
func (o *PostProductOK) Code() int {
	return 200
}

func (o *PostProductOK) Error() string {
	return fmt.Sprintf("[POST /products][%d] postProductOK ", 200)
}

func (o *PostProductOK) String() string {
	return fmt.Sprintf("[POST /products][%d] postProductOK ", 200)
}

func (o *PostProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProductUnprocessableEntity creates a PostProductUnprocessableEntity with default headers values
func NewPostProductUnprocessableEntity() *PostProductUnprocessableEntity {
	return &PostProductUnprocessableEntity{}
}

/*
PostProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type PostProductUnprocessableEntity struct {
}

// IsSuccess returns true when this post product unprocessable entity response has a 2xx status code
func (o *PostProductUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post product unprocessable entity response has a 3xx status code
func (o *PostProductUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post product unprocessable entity response has a 4xx status code
func (o *PostProductUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post product unprocessable entity response has a 5xx status code
func (o *PostProductUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post product unprocessable entity response a status code equal to that given
func (o *PostProductUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post product unprocessable entity response
func (o *PostProductUnprocessableEntity) Code() int {
	return 422
}

func (o *PostProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /products][%d] postProductUnprocessableEntity ", 422)
}

func (o *PostProductUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /products][%d] postProductUnprocessableEntity ", 422)
}

func (o *PostProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProductNotImplemented creates a PostProductNotImplemented with default headers values
func NewPostProductNotImplemented() *PostProductNotImplemented {
	return &PostProductNotImplemented{}
}

/*
PostProductNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type PostProductNotImplemented struct {
}

// IsSuccess returns true when this post product not implemented response has a 2xx status code
func (o *PostProductNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post product not implemented response has a 3xx status code
func (o *PostProductNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post product not implemented response has a 4xx status code
func (o *PostProductNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this post product not implemented response has a 5xx status code
func (o *PostProductNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this post product not implemented response a status code equal to that given
func (o *PostProductNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the post product not implemented response
func (o *PostProductNotImplemented) Code() int {
	return 501
}

func (o *PostProductNotImplemented) Error() string {
	return fmt.Sprintf("[POST /products][%d] postProductNotImplemented ", 501)
}

func (o *PostProductNotImplemented) String() string {
	return fmt.Sprintf("[POST /products][%d] postProductNotImplemented ", 501)
}

func (o *PostProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
