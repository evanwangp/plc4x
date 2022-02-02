/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model


import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// The data-structure of this message
type MQTT_Property struct {
	PropertyType MQTT_PropertyType
	Child IMQTT_PropertyChild
}

// The corresponding interface
type IMQTT_Property interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IMQTT_PropertyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IMQTT_Property, serializeChildFunction func() error) error
	GetTypeName() string
}

type IMQTT_PropertyChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *MQTT_Property, propertyType MQTT_PropertyType)
	GetTypeName() string
	IMQTT_Property
}

func NewMQTT_Property(propertyType MQTT_PropertyType) *MQTT_Property {
	return &MQTT_Property{PropertyType: propertyType}
}

func CastMQTT_Property(structType interface{}) *MQTT_Property {
	castFunc := func(typ interface{}) *MQTT_Property {
		if casted, ok := typ.(MQTT_Property); ok {
			return &casted
		}
		if casted, ok := typ.(*MQTT_Property); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MQTT_Property) GetTypeName() string {
	return "MQTT_Property"
}

func (m *MQTT_Property) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MQTT_Property) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *MQTT_Property) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyType)
	lengthInBits += 8

	return lengthInBits
}

func (m *MQTT_Property) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_PropertyParse(readBuffer utils.ReadBuffer) (*MQTT_Property, error) {
	if pullErr := readBuffer.PullContext("MQTT_Property"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (propertyType)
	if pullErr := readBuffer.PullContext("propertyType"); pullErr != nil {
		return nil, pullErr
	}
_propertyType, _propertyTypeErr := MQTT_PropertyTypeParse(readBuffer)
	if _propertyTypeErr != nil {
		return nil, errors.Wrap(_propertyTypeErr, "Error parsing 'propertyType' field")
	}
	propertyType := _propertyType
	if closeErr := readBuffer.CloseContext("propertyType"); closeErr != nil {
		return nil, closeErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *MQTT_Property
	var typeSwitchError error
	switch {
case propertyType == MQTT_PropertyType_PAYLOAD_FORMAT_INDICATOR : // MQTT_Property_PAYLOAD_FORMAT_INDICATOR
		_parent, typeSwitchError = MQTT_Property_PAYLOAD_FORMAT_INDICATORParse(readBuffer, )
case propertyType == MQTT_PropertyType_MESSAGE_EXPIRY_INTERVAL : // MQTT_Property_MESSAGE_EXPIRY_INTERVAL
		_parent, typeSwitchError = MQTT_Property_MESSAGE_EXPIRY_INTERVALParse(readBuffer, )
case propertyType == MQTT_PropertyType_CONTENT_TYPE : // MQTT_Property_CONTENT_TYPE
		_parent, typeSwitchError = MQTT_Property_CONTENT_TYPEParse(readBuffer, )
case propertyType == MQTT_PropertyType_RESPONSE_TOPIC : // MQTT_Property_RESPONSE_TOPIC
		_parent, typeSwitchError = MQTT_Property_RESPONSE_TOPICParse(readBuffer, )
case propertyType == MQTT_PropertyType_CORRELATION_DATA : // MQTT_Property_CORRELATION_DATA
		_parent, typeSwitchError = MQTT_Property_CORRELATION_DATAParse(readBuffer, )
case propertyType == MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER : // MQTT_Property_SUBSCRIPTION_IDENTIFIER
		_parent, typeSwitchError = MQTT_Property_SUBSCRIPTION_IDENTIFIERParse(readBuffer, )
case propertyType == MQTT_PropertyType_SESSION_EXPIRY_INTERVAL : // MQTT_Property_EXPIRY_INTERVAL
		_parent, typeSwitchError = MQTT_Property_EXPIRY_INTERVALParse(readBuffer, )
case propertyType == MQTT_PropertyType_ASSIGNED_CLIENT_IDENTIFIER : // MQTT_Property_ASSIGNED_CLIENT_IDENTIFIER
		_parent, typeSwitchError = MQTT_Property_ASSIGNED_CLIENT_IDENTIFIERParse(readBuffer, )
case propertyType == MQTT_PropertyType_SERVER_KEEP_ALIVE : // MQTT_Property_SERVER_KEEP_ALIVE
		_parent, typeSwitchError = MQTT_Property_SERVER_KEEP_ALIVEParse(readBuffer, )
case propertyType == MQTT_PropertyType_AUTHENTICATION_METHOD : // MQTT_Property_AUTHENTICATION_METHOD
		_parent, typeSwitchError = MQTT_Property_AUTHENTICATION_METHODParse(readBuffer, )
case propertyType == MQTT_PropertyType_AUTHENTICATION_DATA : // MQTT_Property_AUTHENTICATION_DATA
		_parent, typeSwitchError = MQTT_Property_AUTHENTICATION_DATAParse(readBuffer, )
case propertyType == MQTT_PropertyType_REQUEST_PROBLEM_INFORMATION : // MQTT_Property_REQUEST_PROBLEM_INFORMATION
		_parent, typeSwitchError = MQTT_Property_REQUEST_PROBLEM_INFORMATIONParse(readBuffer, )
case propertyType == MQTT_PropertyType_WILL_DELAY_INTERVAL : // MQTT_Property_WILL_DELAY_INTERVAL
		_parent, typeSwitchError = MQTT_Property_WILL_DELAY_INTERVALParse(readBuffer, )
case propertyType == MQTT_PropertyType_REQUEST_RESPONSE_INFORMATION : // MQTT_Property_REQUEST_RESPONSE_INFORMATION
		_parent, typeSwitchError = MQTT_Property_REQUEST_RESPONSE_INFORMATIONParse(readBuffer, )
case propertyType == MQTT_PropertyType_RESPONSE_INFORMATION : // MQTT_Property_RESPONSE_INFORMATION
		_parent, typeSwitchError = MQTT_Property_RESPONSE_INFORMATIONParse(readBuffer, )
case propertyType == MQTT_PropertyType_SERVER_REFERENCE : // MQTT_Property_SERVER_REFERENCE
		_parent, typeSwitchError = MQTT_Property_SERVER_REFERENCEParse(readBuffer, )
case propertyType == MQTT_PropertyType_REASON_STRING : // MQTT_Property_REASON_STRING
		_parent, typeSwitchError = MQTT_Property_REASON_STRINGParse(readBuffer, )
case propertyType == MQTT_PropertyType_RECEIVE_MAXIMUM : // MQTT_Property_RECEIVE_MAXIMUM
		_parent, typeSwitchError = MQTT_Property_RECEIVE_MAXIMUMParse(readBuffer, )
case propertyType == MQTT_PropertyType_TOPIC_ALIAS_MAXIMUM : // MQTT_Property_TOPIC_ALIAS_MAXIMUM
		_parent, typeSwitchError = MQTT_Property_TOPIC_ALIAS_MAXIMUMParse(readBuffer, )
case propertyType == MQTT_PropertyType_TOPIC_ALIAS : // MQTT_Property_TOPIC_ALIAS
		_parent, typeSwitchError = MQTT_Property_TOPIC_ALIASParse(readBuffer, )
case propertyType == MQTT_PropertyType_MAXIMUM_QOS : // MQTT_Property_MAXIMUM_QOS
		_parent, typeSwitchError = MQTT_Property_MAXIMUM_QOSParse(readBuffer, )
case propertyType == MQTT_PropertyType_RETAIN_AVAILABLE : // MQTT_Property_RETAIN_AVAILABLE
		_parent, typeSwitchError = MQTT_Property_RETAIN_AVAILABLEParse(readBuffer, )
case propertyType == MQTT_PropertyType_USER_PROPERTY : // MQTT_Property_USER_PROPERTY
		_parent, typeSwitchError = MQTT_Property_USER_PROPERTYParse(readBuffer, )
case propertyType == MQTT_PropertyType_MAXIMUM_PACKET_SIZE : // MQTT_Property_MAXIMUM_PACKET_SIZE
		_parent, typeSwitchError = MQTT_Property_MAXIMUM_PACKET_SIZEParse(readBuffer, )
case propertyType == MQTT_PropertyType_WILDCARD_SUBSCRIPTION_AVAILABLE : // MQTT_Property_WILDCARD_SUBSCRIPTION_AVAILABLE
		_parent, typeSwitchError = MQTT_Property_WILDCARD_SUBSCRIPTION_AVAILABLEParse(readBuffer, )
case propertyType == MQTT_PropertyType_SUBSCRIPTION_IDENTIFIER_AVAILABLE : // MQTT_Property_SUBSCRIPTION_IDENTIFIER_AVAILABLE
		_parent, typeSwitchError = MQTT_Property_SUBSCRIPTION_IDENTIFIER_AVAILABLEParse(readBuffer, )
case propertyType == MQTT_PropertyType_SHARED_SUBSCRIPTION_AVAILABLE : // MQTT_Property_SHARED_SUBSCRIPTION_AVAILABLE
		_parent, typeSwitchError = MQTT_Property_SHARED_SUBSCRIPTION_AVAILABLEParse(readBuffer, )
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("MQTT_Property"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, propertyType)
	return _parent, nil
}

func (m *MQTT_Property) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *MQTT_Property) SerializeParent(writeBuffer utils.WriteBuffer, child IMQTT_Property, serializeChildFunction func() error) error {
	if pushErr :=writeBuffer.PushContext("MQTT_Property"); pushErr != nil {
		return pushErr
	}

	// Simple Field (propertyType)
	if pushErr := writeBuffer.PushContext("propertyType"); pushErr != nil {
		return pushErr
	}
	_propertyTypeErr := m.PropertyType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyType"); popErr != nil {
		return popErr
	}
	if _propertyTypeErr != nil {
		return errors.Wrap(_propertyTypeErr, "Error serializing 'propertyType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MQTT_Property"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *MQTT_Property) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}


