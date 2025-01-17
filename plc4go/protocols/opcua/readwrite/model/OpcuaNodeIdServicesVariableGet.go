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

// OpcuaNodeIdServicesVariableGet is an enum
type OpcuaNodeIdServicesVariableGet int32

type IOpcuaNodeIdServicesVariableGet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_InputArguments    OpcuaNodeIdServicesVariableGet = 11496
	OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_OutputArguments   OpcuaNodeIdServicesVariableGet = 11497
	OpcuaNodeIdServicesVariableGet_GetPositionMethodType_InputArguments          OpcuaNodeIdServicesVariableGet = 11749
	OpcuaNodeIdServicesVariableGet_GetPositionMethodType_OutputArguments         OpcuaNodeIdServicesVariableGet = 11750
	OpcuaNodeIdServicesVariableGet_GetRejectedListMethodType_OutputArguments     OpcuaNodeIdServicesVariableGet = 12774
	OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_InputArguments      OpcuaNodeIdServicesVariableGet = 15219
	OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_OutputArguments     OpcuaNodeIdServicesVariableGet = 15220
	OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_InputArguments     OpcuaNodeIdServicesVariableGet = 15450
	OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_OutputArguments    OpcuaNodeIdServicesVariableGet = 15451
	OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_InputArguments     OpcuaNodeIdServicesVariableGet = 17532
	OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_OutputArguments    OpcuaNodeIdServicesVariableGet = 17533
	OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_InputArguments        OpcuaNodeIdServicesVariableGet = 23727
	OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_OutputArguments       OpcuaNodeIdServicesVariableGet = 23728
	OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_InputArguments       OpcuaNodeIdServicesVariableGet = 23746
	OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_OutputArguments      OpcuaNodeIdServicesVariableGet = 23747
	OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_InputArguments       OpcuaNodeIdServicesVariableGet = 23768
	OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_OutputArguments      OpcuaNodeIdServicesVariableGet = 23769
	OpcuaNodeIdServicesVariableGet_GetDataSetWriterMethodType_OutputArguments    OpcuaNodeIdServicesVariableGet = 23780
	OpcuaNodeIdServicesVariableGet_GetDataSetReaderMethodType_OutputArguments    OpcuaNodeIdServicesVariableGet = 23791
	OpcuaNodeIdServicesVariableGet_GetGroupMembershipsMethodType_OutputArguments OpcuaNodeIdServicesVariableGet = 25155
	OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_InputArguments      OpcuaNodeIdServicesVariableGet = 32283
	OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_OutputArguments     OpcuaNodeIdServicesVariableGet = 32284
)

var OpcuaNodeIdServicesVariableGetValues []OpcuaNodeIdServicesVariableGet

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableGetValues = []OpcuaNodeIdServicesVariableGet{
		OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetPositionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetPositionMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetRejectedListMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetDataSetWriterMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetDataSetReaderMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetGroupMembershipsMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableGetByValue(value int32) (enum OpcuaNodeIdServicesVariableGet, ok bool) {
	switch value {
	case 11496:
		return OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_InputArguments, true
	case 11497:
		return OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_OutputArguments, true
	case 11749:
		return OpcuaNodeIdServicesVariableGet_GetPositionMethodType_InputArguments, true
	case 11750:
		return OpcuaNodeIdServicesVariableGet_GetPositionMethodType_OutputArguments, true
	case 12774:
		return OpcuaNodeIdServicesVariableGet_GetRejectedListMethodType_OutputArguments, true
	case 15219:
		return OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_InputArguments, true
	case 15220:
		return OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_OutputArguments, true
	case 15450:
		return OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_InputArguments, true
	case 15451:
		return OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_OutputArguments, true
	case 17532:
		return OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_InputArguments, true
	case 17533:
		return OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_OutputArguments, true
	case 23727:
		return OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_InputArguments, true
	case 23728:
		return OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_OutputArguments, true
	case 23746:
		return OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_InputArguments, true
	case 23747:
		return OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_OutputArguments, true
	case 23768:
		return OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_InputArguments, true
	case 23769:
		return OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_OutputArguments, true
	case 23780:
		return OpcuaNodeIdServicesVariableGet_GetDataSetWriterMethodType_OutputArguments, true
	case 23791:
		return OpcuaNodeIdServicesVariableGet_GetDataSetReaderMethodType_OutputArguments, true
	case 25155:
		return OpcuaNodeIdServicesVariableGet_GetGroupMembershipsMethodType_OutputArguments, true
	case 32283:
		return OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_InputArguments, true
	case 32284:
		return OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGetByName(value string) (enum OpcuaNodeIdServicesVariableGet, ok bool) {
	switch value {
	case "GetMonitoredItemsMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_InputArguments, true
	case "GetMonitoredItemsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_OutputArguments, true
	case "GetPositionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetPositionMethodType_InputArguments, true
	case "GetPositionMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetPositionMethodType_OutputArguments, true
	case "GetRejectedListMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetRejectedListMethodType_OutputArguments, true
	case "GetSecurityKeysMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_InputArguments, true
	case "GetSecurityKeysMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_OutputArguments, true
	case "GetSecurityGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_InputArguments, true
	case "GetSecurityGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_OutputArguments, true
	case "GetEncryptingKeyMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_InputArguments, true
	case "GetEncryptingKeyMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_OutputArguments, true
	case "GetConnectionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_InputArguments, true
	case "GetConnectionMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_OutputArguments, true
	case "GetWriterGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_InputArguments, true
	case "GetWriterGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_OutputArguments, true
	case "GetReaderGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_InputArguments, true
	case "GetReaderGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_OutputArguments, true
	case "GetDataSetWriterMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetDataSetWriterMethodType_OutputArguments, true
	case "GetDataSetReaderMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetDataSetReaderMethodType_OutputArguments, true
	case "GetGroupMembershipsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetGroupMembershipsMethodType_OutputArguments, true
	case "GetCertificatesMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_InputArguments, true
	case "GetCertificatesMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGetKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableGetValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableGet(structType any) OpcuaNodeIdServicesVariableGet {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableGet {
		if sOpcuaNodeIdServicesVariableGet, ok := typ.(OpcuaNodeIdServicesVariableGet); ok {
			return sOpcuaNodeIdServicesVariableGet
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableGet) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableGet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableGetParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableGet, error) {
	return OpcuaNodeIdServicesVariableGetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableGetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableGet, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableGet", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableGet")
	}
	if enum, ok := OpcuaNodeIdServicesVariableGetByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableGet")
		return OpcuaNodeIdServicesVariableGet(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableGet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableGet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableGet", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableGet) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_InputArguments:
		return "GetMonitoredItemsMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetMonitoredItemsMethodType_OutputArguments:
		return "GetMonitoredItemsMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetPositionMethodType_InputArguments:
		return "GetPositionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetPositionMethodType_OutputArguments:
		return "GetPositionMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetRejectedListMethodType_OutputArguments:
		return "GetRejectedListMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_InputArguments:
		return "GetSecurityKeysMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetSecurityKeysMethodType_OutputArguments:
		return "GetSecurityKeysMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_InputArguments:
		return "GetSecurityGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetSecurityGroupMethodType_OutputArguments:
		return "GetSecurityGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_InputArguments:
		return "GetEncryptingKeyMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetEncryptingKeyMethodType_OutputArguments:
		return "GetEncryptingKeyMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_InputArguments:
		return "GetConnectionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetConnectionMethodType_OutputArguments:
		return "GetConnectionMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_InputArguments:
		return "GetWriterGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetWriterGroupMethodType_OutputArguments:
		return "GetWriterGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_InputArguments:
		return "GetReaderGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetReaderGroupMethodType_OutputArguments:
		return "GetReaderGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetDataSetWriterMethodType_OutputArguments:
		return "GetDataSetWriterMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetDataSetReaderMethodType_OutputArguments:
		return "GetDataSetReaderMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetGroupMembershipsMethodType_OutputArguments:
		return "GetGroupMembershipsMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_InputArguments:
		return "GetCertificatesMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGet_GetCertificatesMethodType_OutputArguments:
		return "GetCertificatesMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableGet) String() string {
	return e.PLC4XEnumName()
}
