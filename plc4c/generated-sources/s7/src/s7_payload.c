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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "s7_payload.h"

// Code generated by code-generation. DO NOT EDIT.

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants so we can use the
// enum constant to directly access a given types discriminator values)
const plc4c_s7_read_write_s7_payload_discriminator plc4c_s7_read_write_s7_payload_discriminators[] = {
  {/* plc4c_s7_read_write_s7_payload_read_var_response */
   .parameterParameterType = 0x04, .messageType = 0x03},
  {/* plc4c_s7_read_write_s7_payload_write_var_request */
   .parameterParameterType = 0x05, .messageType = 0x01},
  {/* plc4c_s7_read_write_s7_payload_write_var_response */
   .parameterParameterType = 0x05, .messageType = 0x03},
  {/* plc4c_s7_read_write_s7_payload_user_data */
   .parameterParameterType = 0x00, .messageType = 0x07}

};

// Function returning the discriminator values for a given type constant.
plc4c_s7_read_write_s7_payload_discriminator plc4c_s7_read_write_s7_payload_get_discriminator(plc4c_s7_read_write_s7_payload_type type) {
  return plc4c_s7_read_write_s7_payload_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_s7_read_write_s7_payload plc4c_s7_read_write_s7_payload_null_const;

plc4c_s7_read_write_s7_payload plc4c_s7_read_write_s7_payload_null() {
  return plc4c_s7_read_write_s7_payload_null_const;
}


// Parse function.
plc4c_return_code plc4c_s7_read_write_s7_payload_parse(plc4c_spi_read_buffer* readBuffer, uint8_t messageType, plc4c_s7_read_write_s7_parameter* parameter, plc4c_s7_read_write_s7_payload** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_s7_payload));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
  if((plc4c_s7_read_write_s7_parameter_get_discriminator(parameter->_type).parameterType == 0x04) && (messageType == 0x03)) { /* S7PayloadReadVarResponse */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_read_var_response;
                    
    // Array field (items)
    plc4c_list* items = NULL;
    plc4c_utils_list_create(&items);
    if(items == NULL) {
      return NO_MEMORY;
    }
    {
      // Count array
      uint16_t itemCount = (uint16_t) ((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_read_var_response_num_items;
      for(int curItem = 0; curItem < itemCount; curItem++) {
        bool lastItem = curItem == (itemCount - 1);
        plc4c_s7_read_write_s7_var_payload_data_item* _value = NULL;
        _res = plc4c_s7_read_write_s7_var_payload_data_item_parse(readBuffer, lastItem, (void*) &_value);
        if(_res != OK) {
          return _res;
        }
        plc4c_utils_list_insert_head_value(items, _value);
      }
    }
    (*_message)->s7_payload_read_var_response_items = items;

  } else 
  if((plc4c_s7_read_write_s7_parameter_get_discriminator(parameter->_type).parameterType == 0x05) && (messageType == 0x01)) { /* S7PayloadWriteVarRequest */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_request;
                    
    // Array field (items)
    plc4c_list* items = NULL;
    plc4c_utils_list_create(&items);
    if(items == NULL) {
      return NO_MEMORY;
    }
    {
      // Count array
      uint16_t itemCount = (uint16_t) plc4c_spi_evaluation_helper_count(((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_write_var_request_items);
      for(int curItem = 0; curItem < itemCount; curItem++) {
        bool lastItem = curItem == (itemCount - 1);
        plc4c_s7_read_write_s7_var_payload_data_item* _value = NULL;
        _res = plc4c_s7_read_write_s7_var_payload_data_item_parse(readBuffer, lastItem, (void*) &_value);
        if(_res != OK) {
          return _res;
        }
        plc4c_utils_list_insert_head_value(items, _value);
      }
    }
    (*_message)->s7_payload_write_var_request_items = items;

  } else 
  if((plc4c_s7_read_write_s7_parameter_get_discriminator(parameter->_type).parameterType == 0x05) && (messageType == 0x03)) { /* S7PayloadWriteVarResponse */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_response;
                    
    // Array field (items)
    plc4c_list* items = NULL;
    plc4c_utils_list_create(&items);
    if(items == NULL) {
      return NO_MEMORY;
    }
    {
      // Count array
      uint16_t itemCount = (uint16_t) ((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_write_var_response_num_items;
      for(int curItem = 0; curItem < itemCount; curItem++) {
        bool lastItem = curItem == (itemCount - 1);
        plc4c_s7_read_write_s7_var_payload_status_item* _value = NULL;
        _res = plc4c_s7_read_write_s7_var_payload_status_item_parse(readBuffer, (void*) &_value);
        if(_res != OK) {
          return _res;
        }
        plc4c_utils_list_insert_head_value(items, _value);
      }
    }
    (*_message)->s7_payload_write_var_response_items = items;

  } else 
  if((plc4c_s7_read_write_s7_parameter_get_discriminator(parameter->_type).parameterType == 0x00) && (messageType == 0x07)) { /* S7PayloadUserData */
    (*_message)->_type = plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_user_data;
                    
    // Array field (items)
    plc4c_list* items = NULL;
    plc4c_utils_list_create(&items);
    if(items == NULL) {
      return NO_MEMORY;
    }
    {
      // Count array
      uint16_t itemCount = (uint16_t) plc4c_spi_evaluation_helper_count(((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_user_data_items);
      for(int curItem = 0; curItem < itemCount; curItem++) {
        bool lastItem = curItem == (itemCount - 1);
        plc4c_s7_read_write_s7_payload_user_data_item* _value = NULL;
        _res = plc4c_s7_read_write_s7_payload_user_data_item_parse(readBuffer, ((plc4c_s7_read_write_s7_parameter_user_data_item*) (plc4c_utils_list_get_value(((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_user_data_items, 0)))->s7_parameter_user_data_item_cpu_functions_cpu_function_type, ((plc4c_s7_read_write_s7_parameter_user_data_item*) (plc4c_utils_list_get_value(((plc4c_s7_read_write_s7_parameter*) (parameter))->s7_parameter_user_data_items, 0)))->s7_parameter_user_data_item_cpu_functions_cpu_subfunction, (void*) &_value);
        if(_res != OK) {
          return _res;
        }
        plc4c_utils_list_insert_head_value(items, _value);
      }
    }
    (*_message)->s7_payload_user_data_items = items;

  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_s7_payload_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_payload* _message) {
  plc4c_return_code _res = OK;

  // Switch Field (Depending of the current type, serialize the sub-type elements)
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_read_var_response: {

      // Array field (items)
      {
        uint8_t itemCount = plc4c_utils_list_size(_message->s7_payload_read_var_response_items);
        for(int curItem = 0; curItem < itemCount; curItem++) {
          bool lastItem = curItem == (itemCount - 1);
          plc4c_s7_read_write_s7_var_payload_data_item* _value = (plc4c_s7_read_write_s7_var_payload_data_item*) plc4c_utils_list_get_value(_message->s7_payload_read_var_response_items, curItem);
          _res = plc4c_s7_read_write_s7_var_payload_data_item_serialize(writeBuffer, (void*) _value, lastItem);
          if(_res != OK) {
            return _res;
          }
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_request: {

      // Array field (items)
      {
        uint8_t itemCount = plc4c_utils_list_size(_message->s7_payload_write_var_request_items);
        for(int curItem = 0; curItem < itemCount; curItem++) {
          bool lastItem = curItem == (itemCount - 1);
          plc4c_s7_read_write_s7_var_payload_data_item* _value = (plc4c_s7_read_write_s7_var_payload_data_item*) plc4c_utils_list_get_value(_message->s7_payload_write_var_request_items, curItem);
          _res = plc4c_s7_read_write_s7_var_payload_data_item_serialize(writeBuffer, (void*) _value, lastItem);
          if(_res != OK) {
            return _res;
          }
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_response: {

      // Array field (items)
      {
        uint8_t itemCount = plc4c_utils_list_size(_message->s7_payload_write_var_response_items);
        for(int curItem = 0; curItem < itemCount; curItem++) {
          bool lastItem = curItem == (itemCount - 1);
          plc4c_s7_read_write_s7_var_payload_status_item* _value = (plc4c_s7_read_write_s7_var_payload_status_item*) plc4c_utils_list_get_value(_message->s7_payload_write_var_response_items, curItem);
          _res = plc4c_s7_read_write_s7_var_payload_status_item_serialize(writeBuffer, (void*) _value);
          if(_res != OK) {
            return _res;
          }
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_user_data: {

      // Array field (items)
      {
        uint8_t itemCount = plc4c_utils_list_size(_message->s7_payload_user_data_items);
        for(int curItem = 0; curItem < itemCount; curItem++) {
          bool lastItem = curItem == (itemCount - 1);
          plc4c_s7_read_write_s7_payload_user_data_item* _value = (plc4c_s7_read_write_s7_payload_user_data_item*) plc4c_utils_list_get_value(_message->s7_payload_user_data_items, curItem);
          _res = plc4c_s7_read_write_s7_payload_user_data_item_serialize(writeBuffer, (void*) _value);
          if(_res != OK) {
            return _res;
          }
        }
      }

      break;
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_s7_payload_length_in_bytes(plc4c_s7_read_write_s7_payload* _message) {
  return plc4c_s7_read_write_s7_payload_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_s7_payload_length_in_bits(plc4c_s7_read_write_s7_payload* _message) {
  uint16_t lengthInBits = 0;

  // Depending of the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_read_var_response: {

      // Array field
      if(_message->s7_payload_read_var_response_items != NULL) {
        plc4c_list_element* curElement = _message->s7_payload_read_var_response_items->tail;
        while (curElement != NULL) {
          lengthInBits += plc4c_s7_read_write_s7_var_payload_data_item_length_in_bits((plc4c_s7_read_write_s7_var_payload_data_item*) curElement->value);
          curElement = curElement->next;
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_request: {

      // Array field
      if(_message->s7_payload_write_var_request_items != NULL) {
        plc4c_list_element* curElement = _message->s7_payload_write_var_request_items->tail;
        while (curElement != NULL) {
          lengthInBits += plc4c_s7_read_write_s7_var_payload_data_item_length_in_bits((plc4c_s7_read_write_s7_var_payload_data_item*) curElement->value);
          curElement = curElement->next;
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_write_var_response: {

      // Array field
      if(_message->s7_payload_write_var_response_items != NULL) {
        plc4c_list_element* curElement = _message->s7_payload_write_var_response_items->tail;
        while (curElement != NULL) {
          lengthInBits += plc4c_s7_read_write_s7_var_payload_status_item_length_in_bits((plc4c_s7_read_write_s7_var_payload_status_item*) curElement->value);
          curElement = curElement->next;
        }
      }

      break;
    }
    case plc4c_s7_read_write_s7_payload_type_plc4c_s7_read_write_s7_payload_user_data: {

      // Array field
      if(_message->s7_payload_user_data_items != NULL) {
        plc4c_list_element* curElement = _message->s7_payload_user_data_items->tail;
        while (curElement != NULL) {
          lengthInBits += plc4c_s7_read_write_s7_payload_user_data_item_length_in_bits((plc4c_s7_read_write_s7_payload_user_data_item*) curElement->value);
          curElement = curElement->next;
        }
      }

      break;
    }
  }

  return lengthInBits;
}
