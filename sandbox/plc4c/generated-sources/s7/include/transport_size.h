/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#ifndef PLC4C_S7_READ_WRITE_TRANSPORT_SIZE_H_
#define PLC4C_S7_READ_WRITE_TRANSPORT_SIZE_H_

#include <stdbool.h>
#include <stdint.h>
#include "data_transport_size.h"
#include "transport_size.h"


#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_s7_read_write_transport_size {
  plc4c_s7_read_write_transport_size_BOOL = 0x01,
  plc4c_s7_read_write_transport_size_BYTE = 0x02,
  plc4c_s7_read_write_transport_size_WORD = 0x03,
  plc4c_s7_read_write_transport_size_DWORD = 0x04,
  plc4c_s7_read_write_transport_size_LWORD = 0x05,
  plc4c_s7_read_write_transport_size_INT = 0x06,
  plc4c_s7_read_write_transport_size_UINT = 0x07,
  plc4c_s7_read_write_transport_size_SINT = 0x08,
  plc4c_s7_read_write_transport_size_USINT = 0x09,
  plc4c_s7_read_write_transport_size_DINT = 0x0A,
  plc4c_s7_read_write_transport_size_UDINT = 0x0B,
  plc4c_s7_read_write_transport_size_LINT = 0x0C,
  plc4c_s7_read_write_transport_size_ULINT = 0x0D,
  plc4c_s7_read_write_transport_size_REAL = 0x0E,
  plc4c_s7_read_write_transport_size_LREAL = 0x0F,
  plc4c_s7_read_write_transport_size_CHAR = 0x10,
  plc4c_s7_read_write_transport_size_WCHAR = 0x11,
  plc4c_s7_read_write_transport_size_STRING = 0x12,
  plc4c_s7_read_write_transport_size_WSTRING = 0x13,
  plc4c_s7_read_write_transport_size_TIME = 0x14,
  plc4c_s7_read_write_transport_size_LTIME = 0x16,
  plc4c_s7_read_write_transport_size_DATE = 0x17,
  plc4c_s7_read_write_transport_size_TIME_OF_DAY = 0x18,
  plc4c_s7_read_write_transport_size_TOD = 0x19,
  plc4c_s7_read_write_transport_size_DATE_AND_TIME = 0x1A,
  plc4c_s7_read_write_transport_size_DT = 0x1B
};
typedef enum plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size;

// Get an empty NULL-struct
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_null();

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_value_of(char* value_string);

int plc4c_s7_read_write_transport_size_num_values();

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_value_for_index(int index);

bool plc4c_s7_read_write_transport_size_get_supported__s7_300(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_300(bool value);

bool plc4c_s7_read_write_transport_size_get_supported__logo(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__logo(bool value);

uint8_t plc4c_s7_read_write_transport_size_get_code(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_code(uint8_t value);

uint8_t plc4c_s7_read_write_transport_size_get_size_in_bytes(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_size_in_bytes(uint8_t value);

bool plc4c_s7_read_write_transport_size_get_supported__s7_400(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_400(bool value);

bool plc4c_s7_read_write_transport_size_get_supported__s7_1200(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_1200(bool value);

uint8_t plc4c_s7_read_write_transport_size_get_short_name(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_short_name(uint8_t value);

bool plc4c_s7_read_write_transport_size_get_supported__s7_1500(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_supported__s7_1500(bool value);

plc4c_s7_read_write_data_transport_size plc4c_s7_read_write_transport_size_get_data_transport_size(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_data_transport_size(plc4c_s7_read_write_data_transport_size value);

plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_base_type(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_base_type(plc4c_s7_read_write_transport_size value);

char* plc4c_s7_read_write_transport_size_get_data_protocol_id(plc4c_s7_read_write_transport_size value);
plc4c_s7_read_write_transport_size plc4c_s7_read_write_transport_size_get_first_enum_for_field_data_protocol_id(char* value);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_S7_READ_WRITE_TRANSPORT_SIZE_H_
