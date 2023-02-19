/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// UnConnectedDataItem is the corresponding interface of UnConnectedDataItem
type UnConnectedDataItem interface {
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetService returns Service (property field)
	GetService() CipService
}

// UnConnectedDataItemExactly can be used when we want exactly this type and not a type which fulfills UnConnectedDataItem.
// This is useful for switch cases.
type UnConnectedDataItemExactly interface {
	UnConnectedDataItem
	isUnConnectedDataItem() bool
}

// _UnConnectedDataItem is the data-structure of this message
type _UnConnectedDataItem struct {
	*_TypeId
	Service CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnConnectedDataItem) GetId() uint16 {
	return 0x00B2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnConnectedDataItem) InitializeParent(parent TypeId) {}

func (m *_UnConnectedDataItem) GetParent() TypeId {
	return m._TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnConnectedDataItem) GetService() CipService {
	return m.Service
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnConnectedDataItem factory function for _UnConnectedDataItem
func NewUnConnectedDataItem(service CipService, order IntegerEncoding) *_UnConnectedDataItem {
	_result := &_UnConnectedDataItem{
		Service: service,
		_TypeId: NewTypeId(order),
	}
	_result._TypeId._TypeIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnConnectedDataItem(structType interface{}) UnConnectedDataItem {
	if casted, ok := structType.(UnConnectedDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*UnConnectedDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_UnConnectedDataItem) GetTypeName() string {
	return "UnConnectedDataItem"
}

func (m *_UnConnectedDataItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (packetSize)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UnConnectedDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnConnectedDataItemParse(theBytes []byte, order IntegerEncoding) (UnConnectedDataItem, error) {
	return UnConnectedDataItemParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased((utils.InlineIf(bool((order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder))), order)
}

func UnConnectedDataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, order IntegerEncoding) (UnConnectedDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnConnectedDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnConnectedDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (packetSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	packetSize, _packetSizeErr := readBuffer.ReadUint16("packetSize", 16)
	_ = packetSize
	if _packetSizeErr != nil {
		return nil, errors.Wrap(_packetSizeErr, "Error parsing 'packetSize' field of UnConnectedDataItem")
	}

	// Simple Field (service)
	if pullErr := readBuffer.PullContext("service"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for service")
	}
	_service, _serviceErr := CipServiceParseWithBuffer(ctx, readBuffer, bool(bool(false)), uint16(packetSize), IntegerEncoding(order))
	if _serviceErr != nil {
		return nil, errors.Wrap(_serviceErr, "Error parsing 'service' field of UnConnectedDataItem")
	}
	service := _service.(CipService)
	if closeErr := readBuffer.CloseContext("service"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for service")
	}

	if closeErr := readBuffer.CloseContext("UnConnectedDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnConnectedDataItem")
	}

	// Create a partially initialized instance
	_child := &_UnConnectedDataItem{
		_TypeId: &_TypeId{
			Order: order,
		},
		Service: service,
	}
	_child._TypeId._TypeIdChildRequirements = _child
	return _child, nil
}

func (m *_UnConnectedDataItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer((utils.InlineIf(bool((m.Order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder)))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnConnectedDataItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnConnectedDataItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnConnectedDataItem")
		}

		// Implicit Field (packetSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		packetSize := uint16(m.GetService().GetLengthInBytes(ctx))
		_packetSizeErr := writeBuffer.WriteUint16("packetSize", 16, (packetSize))
		if _packetSizeErr != nil {
			return errors.Wrap(_packetSizeErr, "Error serializing 'packetSize' field")
		}

		// Simple Field (service)
		if pushErr := writeBuffer.PushContext("service"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for service")
		}
		_serviceErr := writeBuffer.WriteSerializable(ctx, m.GetService())
		if popErr := writeBuffer.PopContext("service"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for service")
		}
		if _serviceErr != nil {
			return errors.Wrap(_serviceErr, "Error serializing 'service' field")
		}

		if popErr := writeBuffer.PopContext("UnConnectedDataItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnConnectedDataItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UnConnectedDataItem) isUnConnectedDataItem() bool {
	return true
}

func (m *_UnConnectedDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}