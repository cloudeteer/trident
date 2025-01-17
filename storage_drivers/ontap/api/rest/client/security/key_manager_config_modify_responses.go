// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// KeyManagerConfigModifyReader is a Reader for the KeyManagerConfigModify structure.
type KeyManagerConfigModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyManagerConfigModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyManagerConfigModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKeyManagerConfigModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyManagerConfigModifyOK creates a KeyManagerConfigModifyOK with default headers values
func NewKeyManagerConfigModifyOK() *KeyManagerConfigModifyOK {
	return &KeyManagerConfigModifyOK{}
}

/* KeyManagerConfigModifyOK describes a response with status code 200, with default header values.

OK
*/
type KeyManagerConfigModifyOK struct {
}

func (o *KeyManagerConfigModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/key-manager-configs][%d] keyManagerConfigModifyOK ", 200)
}

func (o *KeyManagerConfigModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKeyManagerConfigModifyDefault creates a KeyManagerConfigModifyDefault with default headers values
func NewKeyManagerConfigModifyDefault(code int) *KeyManagerConfigModifyDefault {
	return &KeyManagerConfigModifyDefault{
		_statusCode: code,
	}
}

/* KeyManagerConfigModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 65536139 | Cluster-wide passphrase is incorrect. |
| 65536805 | Common Criteria Mode requires an effective cluster version of ONTAP 9.4 or later. |
| 65536806 | Passphrase length error. |
| 65536807 | MetroCluster cannot be configured while in Common Criteria mode. |
| 65536809 | Common Criteria mode is disabled on the cluster. Contact technical support for assistance in enabling Common Criteria mode. |
| 65537302 | The passphrase field is required when changing cc_mode_enabled to true. |
| 65537304 | Unable to modify polling period because no external key management is configured on the cluster. |

*/
type KeyManagerConfigModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the key manager config modify default response
func (o *KeyManagerConfigModifyDefault) Code() int {
	return o._statusCode
}

func (o *KeyManagerConfigModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/key-manager-configs][%d] key_manager_config_modify default  %+v", o._statusCode, o.Payload)
}
func (o *KeyManagerConfigModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KeyManagerConfigModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
