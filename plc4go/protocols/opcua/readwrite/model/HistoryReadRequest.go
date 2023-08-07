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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// HistoryReadRequest is the corresponding interface of HistoryReadRequest
type HistoryReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetHistoryReadDetails returns HistoryReadDetails (property field)
	GetHistoryReadDetails() ExtensionObject
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetReleaseContinuationPoints returns ReleaseContinuationPoints (property field)
	GetReleaseContinuationPoints() bool
	// GetNoOfNodesToRead returns NoOfNodesToRead (property field)
	GetNoOfNodesToRead() int32
	// GetNodesToRead returns NodesToRead (property field)
	GetNodesToRead() []ExtensionObjectDefinition
}

// HistoryReadRequestExactly can be used when we want exactly this type and not a type which fulfills HistoryReadRequest.
// This is useful for switch cases.
type HistoryReadRequestExactly interface {
	HistoryReadRequest
	isHistoryReadRequest() bool
}

// _HistoryReadRequest is the data-structure of this message
type _HistoryReadRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader             ExtensionObjectDefinition
	HistoryReadDetails        ExtensionObject
	TimestampsToReturn        TimestampsToReturn
	ReleaseContinuationPoints bool
	NoOfNodesToRead           int32
	NodesToRead               []ExtensionObjectDefinition
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadRequest) GetIdentifier() string {
	return "664"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryReadRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_HistoryReadRequest) GetHistoryReadDetails() ExtensionObject {
	return m.HistoryReadDetails
}

func (m *_HistoryReadRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_HistoryReadRequest) GetReleaseContinuationPoints() bool {
	return m.ReleaseContinuationPoints
}

func (m *_HistoryReadRequest) GetNoOfNodesToRead() int32 {
	return m.NoOfNodesToRead
}

func (m *_HistoryReadRequest) GetNodesToRead() []ExtensionObjectDefinition {
	return m.NodesToRead
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryReadRequest factory function for _HistoryReadRequest
func NewHistoryReadRequest(requestHeader ExtensionObjectDefinition, historyReadDetails ExtensionObject, timestampsToReturn TimestampsToReturn, releaseContinuationPoints bool, noOfNodesToRead int32, nodesToRead []ExtensionObjectDefinition) *_HistoryReadRequest {
	_result := &_HistoryReadRequest{
		RequestHeader:              requestHeader,
		HistoryReadDetails:         historyReadDetails,
		TimestampsToReturn:         timestampsToReturn,
		ReleaseContinuationPoints:  releaseContinuationPoints,
		NoOfNodesToRead:            noOfNodesToRead,
		NodesToRead:                nodesToRead,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryReadRequest(structType any) HistoryReadRequest {
	if casted, ok := structType.(HistoryReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadRequest) GetTypeName() string {
	return "HistoryReadRequest"
}

func (m *_HistoryReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (historyReadDetails)
	lengthInBits += m.HistoryReadDetails.GetLengthInBits(ctx)

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (releaseContinuationPoints)
	lengthInBits += 1

	// Simple field (noOfNodesToRead)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRead) > 0 {
		for _curItem, element := range m.NodesToRead {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRead), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryReadRequestParse(ctx context.Context, theBytes []byte, identifier string) (HistoryReadRequest, error) {
	return HistoryReadRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryReadRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryReadRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of HistoryReadRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (historyReadDetails)
	if pullErr := readBuffer.PullContext("historyReadDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for historyReadDetails")
	}
	_historyReadDetails, _historyReadDetailsErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _historyReadDetailsErr != nil {
		return nil, errors.Wrap(_historyReadDetailsErr, "Error parsing 'historyReadDetails' field of HistoryReadRequest")
	}
	historyReadDetails := _historyReadDetails.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("historyReadDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for historyReadDetails")
	}

	// Simple Field (timestampsToReturn)
	if pullErr := readBuffer.PullContext("timestampsToReturn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestampsToReturn")
	}
	_timestampsToReturn, _timestampsToReturnErr := TimestampsToReturnParseWithBuffer(ctx, readBuffer)
	if _timestampsToReturnErr != nil {
		return nil, errors.Wrap(_timestampsToReturnErr, "Error parsing 'timestampsToReturn' field of HistoryReadRequest")
	}
	timestampsToReturn := _timestampsToReturn
	if closeErr := readBuffer.CloseContext("timestampsToReturn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestampsToReturn")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of HistoryReadRequest")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (releaseContinuationPoints)
	_releaseContinuationPoints, _releaseContinuationPointsErr := readBuffer.ReadBit("releaseContinuationPoints")
	if _releaseContinuationPointsErr != nil {
		return nil, errors.Wrap(_releaseContinuationPointsErr, "Error parsing 'releaseContinuationPoints' field of HistoryReadRequest")
	}
	releaseContinuationPoints := _releaseContinuationPoints

	// Simple Field (noOfNodesToRead)
	_noOfNodesToRead, _noOfNodesToReadErr := readBuffer.ReadInt32("noOfNodesToRead", 32)
	if _noOfNodesToReadErr != nil {
		return nil, errors.Wrap(_noOfNodesToReadErr, "Error parsing 'noOfNodesToRead' field of HistoryReadRequest")
	}
	noOfNodesToRead := _noOfNodesToRead

	// Array field (nodesToRead)
	if pullErr := readBuffer.PullContext("nodesToRead", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodesToRead")
	}
	// Count array
	nodesToRead := make([]ExtensionObjectDefinition, utils.Max(noOfNodesToRead, 0))
	// This happens when the size is set conditional to 0
	if len(nodesToRead) == 0 {
		nodesToRead = nil
	}
	{
		_numItems := uint16(utils.Max(noOfNodesToRead, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "637")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'nodesToRead' field of HistoryReadRequest")
			}
			nodesToRead[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("nodesToRead", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodesToRead")
	}

	if closeErr := readBuffer.CloseContext("HistoryReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadRequest")
	}

	// Create a partially initialized instance
	_child := &_HistoryReadRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		HistoryReadDetails:         historyReadDetails,
		TimestampsToReturn:         timestampsToReturn,
		ReleaseContinuationPoints:  releaseContinuationPoints,
		NoOfNodesToRead:            noOfNodesToRead,
		NodesToRead:                nodesToRead,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (historyReadDetails)
		if pushErr := writeBuffer.PushContext("historyReadDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for historyReadDetails")
		}
		_historyReadDetailsErr := writeBuffer.WriteSerializable(ctx, m.GetHistoryReadDetails())
		if popErr := writeBuffer.PopContext("historyReadDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for historyReadDetails")
		}
		if _historyReadDetailsErr != nil {
			return errors.Wrap(_historyReadDetailsErr, "Error serializing 'historyReadDetails' field")
		}

		// Simple Field (timestampsToReturn)
		if pushErr := writeBuffer.PushContext("timestampsToReturn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestampsToReturn")
		}
		_timestampsToReturnErr := writeBuffer.WriteSerializable(ctx, m.GetTimestampsToReturn())
		if popErr := writeBuffer.PopContext("timestampsToReturn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestampsToReturn")
		}
		if _timestampsToReturnErr != nil {
			return errors.Wrap(_timestampsToReturnErr, "Error serializing 'timestampsToReturn' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (releaseContinuationPoints)
		releaseContinuationPoints := bool(m.GetReleaseContinuationPoints())
		_releaseContinuationPointsErr := writeBuffer.WriteBit("releaseContinuationPoints", (releaseContinuationPoints))
		if _releaseContinuationPointsErr != nil {
			return errors.Wrap(_releaseContinuationPointsErr, "Error serializing 'releaseContinuationPoints' field")
		}

		// Simple Field (noOfNodesToRead)
		noOfNodesToRead := int32(m.GetNoOfNodesToRead())
		_noOfNodesToReadErr := writeBuffer.WriteInt32("noOfNodesToRead", 32, (noOfNodesToRead))
		if _noOfNodesToReadErr != nil {
			return errors.Wrap(_noOfNodesToReadErr, "Error serializing 'noOfNodesToRead' field")
		}

		// Array Field (nodesToRead)
		if pushErr := writeBuffer.PushContext("nodesToRead", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodesToRead")
		}
		for _curItem, _element := range m.GetNodesToRead() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodesToRead()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodesToRead' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodesToRead", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodesToRead")
		}

		if popErr := writeBuffer.PopContext("HistoryReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadRequest) isHistoryReadRequest() bool {
	return true
}

func (m *_HistoryReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
