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

// Code generated by "plc4xgenerator -type=worker"; DO NOT EDIT.

package pool

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *worker) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *worker) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("worker"); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt64("id", 64, int64(d.id)); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("shutdown", d.shutdown.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("interrupted", d.interrupted.Load()); err != nil {
		return err
	}

	_interrupter_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.interrupter))
	if err := writeBuffer.WriteString("interrupter", uint32(len(_interrupter_plx4gen_description)*8), "UTF-8", _interrupter_plx4gen_description); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("hasEnded", d.hasEnded.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("lastReceived", uint32(len(d.lastReceived.String())*8), "UTF-8", d.lastReceived.String()); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("worker"); err != nil {
		return err
	}
	return nil
}

func (d *worker) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
