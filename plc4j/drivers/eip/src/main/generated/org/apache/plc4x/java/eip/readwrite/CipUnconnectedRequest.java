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
package org.apache.plc4x.java.eip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class CipUnconnectedRequest extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0x52;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Boolean getConnected() {
    return (boolean) false;
  }

  // Constant values.
  public static final Integer ROUTE = 0x0001;

  // Properties.
  protected final PathSegment classSegment;
  protected final PathSegment instanceSegment;
  protected final CipService unconnectedService;
  protected final byte backPlane;
  protected final byte slot;

  // Arguments.
  protected final Integer serviceLen;
  protected final IntegerEncoding order;

  public CipUnconnectedRequest(
      PathSegment classSegment,
      PathSegment instanceSegment,
      CipService unconnectedService,
      byte backPlane,
      byte slot,
      Integer serviceLen,
      IntegerEncoding order) {
    super(serviceLen, order);
    this.classSegment = classSegment;
    this.instanceSegment = instanceSegment;
    this.unconnectedService = unconnectedService;
    this.backPlane = backPlane;
    this.slot = slot;
    this.serviceLen = serviceLen;
    this.order = order;
  }

  public PathSegment getClassSegment() {
    return classSegment;
  }

  public PathSegment getInstanceSegment() {
    return instanceSegment;
  }

  public CipService getUnconnectedService() {
    return unconnectedService;
  }

  public byte getBackPlane() {
    return backPlane;
  }

  public byte getSlot() {
    return slot;
  }

  public int getRoute() {
    return ROUTE;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CipUnconnectedRequest");

    // Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    byte requestPathSize =
        (byte)
            ((((getClassSegment().getLengthInBytes()) + (getInstanceSegment().getLengthInBytes())))
                / (2));
    writeImplicitField(
        "requestPathSize",
        requestPathSize,
        writeSignedByte(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (classSegment)
    writeSimpleField(
        "classSegment",
        classSegment,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (instanceSegment)
    writeSimpleField(
        "instanceSegment",
        instanceSegment,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        (int) 0x9D05,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int messageSize = (int) (((getLengthInBytes()) - (10)) - (4));
    writeImplicitField(
        "messageSize",
        messageSize,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (unconnectedService)
    writeSimpleField(
        "unconnectedService",
        unconnectedService,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (route)
    writeConstField(
        "route",
        ROUTE,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (backPlane)
    writeSimpleField(
        "backPlane",
        backPlane,
        writeSignedByte(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (slot)
    writeSimpleField(
        "slot",
        slot,
        writeSignedByte(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    writeBuffer.popContext("CipUnconnectedRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipUnconnectedRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (requestPathSize)
    lengthInBits += 8;

    // Simple field (classSegment)
    lengthInBits += classSegment.getLengthInBits();

    // Simple field (instanceSegment)
    lengthInBits += instanceSegment.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 16;

    // Implicit Field (messageSize)
    lengthInBits += 16;

    // Simple field (unconnectedService)
    lengthInBits += unconnectedService.getLengthInBits();

    // Const Field (route)
    lengthInBits += 16;

    // Simple field (backPlane)
    lengthInBits += 8;

    // Simple field (slot)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen, IntegerEncoding order)
      throws ParseException {
    readBuffer.pullContext("CipUnconnectedRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte requestPathSize =
        readImplicitField(
            "requestPathSize",
            readSignedByte(readBuffer, 8),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    PathSegment classSegment =
        readSimpleField(
            "classSegment",
            new DataReaderComplexDefault<>(
                () -> PathSegment.staticParse(readBuffer, (IntegerEncoding) (order)), readBuffer),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    PathSegment instanceSegment =
        readSimpleField(
            "instanceSegment",
            new DataReaderComplexDefault<>(
                () -> PathSegment.staticParse(readBuffer, (IntegerEncoding) (order)), readBuffer),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    Integer reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedInt(readBuffer, 16),
            (int) 0x9D05,
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int messageSize =
        readImplicitField(
            "messageSize",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    CipService unconnectedService =
        readSimpleField(
            "unconnectedService",
            new DataReaderComplexDefault<>(
                () ->
                    CipService.staticParse(
                        readBuffer,
                        (boolean) (false),
                        (int) (messageSize),
                        (IntegerEncoding) (order)),
                readBuffer),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    int route =
        readConstField(
            "route",
            readUnsignedInt(readBuffer, 16),
            CipUnconnectedRequest.ROUTE,
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    byte backPlane =
        readSimpleField(
            "backPlane",
            readSignedByte(readBuffer, 8),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    byte slot =
        readSimpleField(
            "slot",
            readSignedByte(readBuffer, 8),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    readBuffer.closeContext("CipUnconnectedRequest");
    // Create the instance
    return new CipUnconnectedRequestBuilderImpl(
        classSegment, instanceSegment, unconnectedService, backPlane, slot, serviceLen, order);
  }

  public static class CipUnconnectedRequestBuilderImpl implements CipService.CipServiceBuilder {
    private final PathSegment classSegment;
    private final PathSegment instanceSegment;
    private final CipService unconnectedService;
    private final byte backPlane;
    private final byte slot;
    private final Integer serviceLen;
    private final IntegerEncoding order;

    public CipUnconnectedRequestBuilderImpl(
        PathSegment classSegment,
        PathSegment instanceSegment,
        CipService unconnectedService,
        byte backPlane,
        byte slot,
        Integer serviceLen,
        IntegerEncoding order) {
      this.classSegment = classSegment;
      this.instanceSegment = instanceSegment;
      this.unconnectedService = unconnectedService;
      this.backPlane = backPlane;
      this.slot = slot;
      this.serviceLen = serviceLen;
      this.order = order;
    }

    public CipUnconnectedRequest build(Integer serviceLen, IntegerEncoding order) {

      CipUnconnectedRequest cipUnconnectedRequest =
          new CipUnconnectedRequest(
              classSegment,
              instanceSegment,
              unconnectedService,
              backPlane,
              slot,
              serviceLen,
              order);
      return cipUnconnectedRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipUnconnectedRequest)) {
      return false;
    }
    CipUnconnectedRequest that = (CipUnconnectedRequest) o;
    return (getClassSegment() == that.getClassSegment())
        && (getInstanceSegment() == that.getInstanceSegment())
        && (getUnconnectedService() == that.getUnconnectedService())
        && (getBackPlane() == that.getBackPlane())
        && (getSlot() == that.getSlot())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getClassSegment(),
        getInstanceSegment(),
        getUnconnectedService(),
        getBackPlane(),
        getSlot());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}