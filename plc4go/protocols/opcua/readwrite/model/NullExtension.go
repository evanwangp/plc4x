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

// NullExtension is the corresponding interface of NullExtension
type NullExtension interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// NullExtensionExactly can be used when we want exactly this type and not a type which fulfills NullExtension.
// This is useful for switch cases.
type NullExtensionExactly interface {
	NullExtension
	isNullExtension() bool
}

// _NullExtension is the data-structure of this message
type _NullExtension struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullExtension) GetIdentifier() string {
	return "0"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullExtension) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_NullExtension) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewNullExtension factory function for _NullExtension
func NewNullExtension() *_NullExtension {
	_result := &_NullExtension{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullExtension(structType any) NullExtension {
	if casted, ok := structType.(NullExtension); ok {
		return casted
	}
	if casted, ok := structType.(*NullExtension); ok {
		return *casted
	}
	return nil
}

func (m *_NullExtension) GetTypeName() string {
	return "NullExtension"
}

func (m *_NullExtension) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullExtension) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NullExtensionParse(ctx context.Context, theBytes []byte, identifier string) (NullExtension, error) {
	return NullExtensionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NullExtensionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NullExtension, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NullExtension"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullExtension")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullExtension"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullExtension")
	}

	// Create a partially initialized instance
	_child := &_NullExtension{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NullExtension) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullExtension) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullExtension"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullExtension")
		}

		if popErr := writeBuffer.PopContext("NullExtension"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullExtension")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullExtension) isNullExtension() bool {
	return true
}

func (m *_NullExtension) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}