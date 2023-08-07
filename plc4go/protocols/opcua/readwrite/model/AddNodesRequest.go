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

// AddNodesRequest is the corresponding interface of AddNodesRequest
type AddNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToAdd returns NoOfNodesToAdd (property field)
	GetNoOfNodesToAdd() int32
	// GetNodesToAdd returns NodesToAdd (property field)
	GetNodesToAdd() []ExtensionObjectDefinition
}

// AddNodesRequestExactly can be used when we want exactly this type and not a type which fulfills AddNodesRequest.
// This is useful for switch cases.
type AddNodesRequestExactly interface {
	AddNodesRequest
	isAddNodesRequest() bool
}

// _AddNodesRequest is the data-structure of this message
type _AddNodesRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader  ExtensionObjectDefinition
	NoOfNodesToAdd int32
	NodesToAdd     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddNodesRequest) GetIdentifier() string {
	return "488"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddNodesRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_AddNodesRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_AddNodesRequest) GetNoOfNodesToAdd() int32 {
	return m.NoOfNodesToAdd
}

func (m *_AddNodesRequest) GetNodesToAdd() []ExtensionObjectDefinition {
	return m.NodesToAdd
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAddNodesRequest factory function for _AddNodesRequest
func NewAddNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToAdd int32, nodesToAdd []ExtensionObjectDefinition) *_AddNodesRequest {
	_result := &_AddNodesRequest{
		RequestHeader:              requestHeader,
		NoOfNodesToAdd:             noOfNodesToAdd,
		NodesToAdd:                 nodesToAdd,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAddNodesRequest(structType any) AddNodesRequest {
	if casted, ok := structType.(AddNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AddNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AddNodesRequest) GetTypeName() string {
	return "AddNodesRequest"
}

func (m *_AddNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToAdd)
	lengthInBits += 32

	// Array field
	if len(m.NodesToAdd) > 0 {
		for _curItem, element := range m.NodesToAdd {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToAdd), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AddNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AddNodesRequestParse(ctx context.Context, theBytes []byte, identifier string) (AddNodesRequest, error) {
	return AddNodesRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func AddNodesRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (AddNodesRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AddNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of AddNodesRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfNodesToAdd)
	_noOfNodesToAdd, _noOfNodesToAddErr := readBuffer.ReadInt32("noOfNodesToAdd", 32)
	if _noOfNodesToAddErr != nil {
		return nil, errors.Wrap(_noOfNodesToAddErr, "Error parsing 'noOfNodesToAdd' field of AddNodesRequest")
	}
	noOfNodesToAdd := _noOfNodesToAdd

	// Array field (nodesToAdd)
	if pullErr := readBuffer.PullContext("nodesToAdd", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodesToAdd")
	}
	// Count array
	nodesToAdd := make([]ExtensionObjectDefinition, utils.Max(noOfNodesToAdd, 0))
	// This happens when the size is set conditional to 0
	if len(nodesToAdd) == 0 {
		nodesToAdd = nil
	}
	{
		_numItems := uint16(utils.Max(noOfNodesToAdd, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "378")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'nodesToAdd' field of AddNodesRequest")
			}
			nodesToAdd[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("nodesToAdd", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodesToAdd")
	}

	if closeErr := readBuffer.CloseContext("AddNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddNodesRequest")
	}

	// Create a partially initialized instance
	_child := &_AddNodesRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfNodesToAdd:             noOfNodesToAdd,
		NodesToAdd:                 nodesToAdd,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_AddNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddNodesRequest")
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

		// Simple Field (noOfNodesToAdd)
		noOfNodesToAdd := int32(m.GetNoOfNodesToAdd())
		_noOfNodesToAddErr := writeBuffer.WriteInt32("noOfNodesToAdd", 32, (noOfNodesToAdd))
		if _noOfNodesToAddErr != nil {
			return errors.Wrap(_noOfNodesToAddErr, "Error serializing 'noOfNodesToAdd' field")
		}

		// Array Field (nodesToAdd)
		if pushErr := writeBuffer.PushContext("nodesToAdd", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodesToAdd")
		}
		for _curItem, _element := range m.GetNodesToAdd() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodesToAdd()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodesToAdd' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodesToAdd", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodesToAdd")
		}

		if popErr := writeBuffer.PopContext("AddNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddNodesRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddNodesRequest) isAddNodesRequest() bool {
	return true
}

func (m *_AddNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
