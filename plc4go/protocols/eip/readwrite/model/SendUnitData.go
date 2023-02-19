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
	"fmt"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const SendUnitData_INTERFACEHANDLE uint32 = 0x00000000

// SendUnitData is the corresponding interface of SendUnitData
type SendUnitData interface {
	utils.LengthAware
	utils.Serializable
	EipPacket
	// GetTimeout returns Timeout (property field)
	GetTimeout() uint16
	// GetItemCount returns ItemCount (property field)
	GetItemCount() uint16
	// GetTypeId returns TypeId (property field)
	GetTypeId() []TypeId
}

// SendUnitDataExactly can be used when we want exactly this type and not a type which fulfills SendUnitData.
// This is useful for switch cases.
type SendUnitDataExactly interface {
	SendUnitData
	isSendUnitData() bool
}

// _SendUnitData is the data-structure of this message
type _SendUnitData struct {
	*_EipPacket
	Timeout   uint16
	ItemCount uint16
	TypeId    []TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SendUnitData) GetCommand() uint16 {
	return 0x0070
}

func (m *_SendUnitData) GetResponse() bool {
	return false
}

func (m *_SendUnitData) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SendUnitData) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_SendUnitData) GetParent() EipPacket {
	return m._EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SendUnitData) GetTimeout() uint16 {
	return m.Timeout
}

func (m *_SendUnitData) GetItemCount() uint16 {
	return m.ItemCount
}

func (m *_SendUnitData) GetTypeId() []TypeId {
	return m.TypeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_SendUnitData) GetInterfaceHandle() uint32 {
	return SendUnitData_INTERFACEHANDLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSendUnitData factory function for _SendUnitData
func NewSendUnitData(timeout uint16, itemCount uint16, typeId []TypeId, sessionHandle uint32, status uint32, senderContext []byte, options uint32, order IntegerEncoding) *_SendUnitData {
	_result := &_SendUnitData{
		Timeout:    timeout,
		ItemCount:  itemCount,
		TypeId:     typeId,
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options, order),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSendUnitData(structType interface{}) SendUnitData {
	if casted, ok := structType.(SendUnitData); ok {
		return casted
	}
	if casted, ok := structType.(*SendUnitData); ok {
		return *casted
	}
	return nil
}

func (m *_SendUnitData) GetTypeName() string {
	return "SendUnitData"
}

func (m *_SendUnitData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (interfaceHandle)
	lengthInBits += 32

	// Simple field (timeout)
	lengthInBits += 16

	// Simple field (itemCount)
	lengthInBits += 16

	// Array field
	if len(m.TypeId) > 0 {
		for _curItem, element := range m.TypeId {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.TypeId), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SendUnitData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SendUnitDataParse(theBytes []byte, order IntegerEncoding, response bool) (SendUnitData, error) {
	return SendUnitDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased((utils.InlineIf(bool((order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder))), order, response)
}

func SendUnitDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, order IntegerEncoding, response bool) (SendUnitData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SendUnitData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SendUnitData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (interfaceHandle)
	interfaceHandle, _interfaceHandleErr := readBuffer.ReadUint32("interfaceHandle", 32)
	if _interfaceHandleErr != nil {
		return nil, errors.Wrap(_interfaceHandleErr, "Error parsing 'interfaceHandle' field of SendUnitData")
	}
	if interfaceHandle != SendUnitData_INTERFACEHANDLE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", SendUnitData_INTERFACEHANDLE) + " but got " + fmt.Sprintf("%d", interfaceHandle))
	}

	// Simple Field (timeout)
	_timeout, _timeoutErr := readBuffer.ReadUint16("timeout", 16)
	if _timeoutErr != nil {
		return nil, errors.Wrap(_timeoutErr, "Error parsing 'timeout' field of SendUnitData")
	}
	timeout := _timeout

	// Simple Field (itemCount)
	_itemCount, _itemCountErr := readBuffer.ReadUint16("itemCount", 16)
	if _itemCountErr != nil {
		return nil, errors.Wrap(_itemCountErr, "Error parsing 'itemCount' field of SendUnitData")
	}
	itemCount := _itemCount

	// Array field (typeId)
	if pullErr := readBuffer.PullContext("typeId", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeId")
	}
	// Count array
	typeId := make([]TypeId, itemCount)
	// This happens when the size is set conditional to 0
	if len(typeId) == 0 {
		typeId = nil
	}
	{
		_numItems := uint16(itemCount)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := TypeIdParseWithBuffer(arrayCtx, readBuffer, order)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'typeId' field of SendUnitData")
			}
			typeId[_curItem] = _item.(TypeId)
		}
	}
	if closeErr := readBuffer.CloseContext("typeId", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeId")
	}

	if closeErr := readBuffer.CloseContext("SendUnitData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SendUnitData")
	}

	// Create a partially initialized instance
	_child := &_SendUnitData{
		_EipPacket: &_EipPacket{
			Order: order,
		},
		Timeout:   timeout,
		ItemCount: itemCount,
		TypeId:    typeId,
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_SendUnitData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer((utils.InlineIf(bool((m.Order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder)))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SendUnitData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SendUnitData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SendUnitData")
		}

		// Const Field (interfaceHandle)
		_interfaceHandleErr := writeBuffer.WriteUint32("interfaceHandle", 32, 0x00000000)
		if _interfaceHandleErr != nil {
			return errors.Wrap(_interfaceHandleErr, "Error serializing 'interfaceHandle' field")
		}

		// Simple Field (timeout)
		timeout := uint16(m.GetTimeout())
		_timeoutErr := writeBuffer.WriteUint16("timeout", 16, (timeout))
		if _timeoutErr != nil {
			return errors.Wrap(_timeoutErr, "Error serializing 'timeout' field")
		}

		// Simple Field (itemCount)
		itemCount := uint16(m.GetItemCount())
		_itemCountErr := writeBuffer.WriteUint16("itemCount", 16, (itemCount))
		if _itemCountErr != nil {
			return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
		}

		// Array Field (typeId)
		if pushErr := writeBuffer.PushContext("typeId", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for typeId")
		}
		for _curItem, _element := range m.GetTypeId() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetTypeId()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'typeId' field")
			}
		}
		if popErr := writeBuffer.PopContext("typeId", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for typeId")
		}

		if popErr := writeBuffer.PopContext("SendUnitData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SendUnitData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SendUnitData) isSendUnitData() bool {
	return true
}

func (m *_SendUnitData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}