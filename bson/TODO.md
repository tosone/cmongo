### Implementation Progress

##### bson_t

- [x] bson_t
- [ ] bson_append_array()
- [ ] bson_append_array_begin()
- [ ] bson_append_array_end()
- [ ] bson_append_binary()
- [ ] bson_append_bool()
- [ ] bson_append_code()
- [ ] bson_append_code_with_scope()
- [ ] bson_append_date_time()
- [ ] bson_append_dbpointer()
- [ ] bson_append_decimal128()
- [ ] bson_append_document()
- [ ] bson_append_document_begin()
- [ ] bson_append_document_end()
- [ ] bson_append_double()
- [x] bson_append_int32()
- [x] bson_append_int64()
- [ ] bson_append_iter()
- [x] bson_append_maxkey()
- [x] bson_append_minkey()
- [ ] bson_append_now_utc()
- [ ] bson_append_null()
- [x] bson_append_oid()
- [x] bson_append_regex()
- [ ] bson_append_regex_w_len()
- [ ] bson_append_symbol()
- [ ] bson_append_time_t()
- [ ] bson_append_timestamp()
- [ ] bson_append_timeval()
- [ ] bson_append_undefined()
- [ ] bson_append_utf8()
- [ ] bson_append_value()
- [ ] bson_array_as_json()
- [ ] bson_as_canonical_extended_json()
- [ ] bson_as_json()
- [ ] bson_as_relaxed_extended_json()
- [ ] bson_compare()
- [ ] bson_concat()
- [ ] bson_copy()
- [ ] bson_copy_to()
- [ ] bson_copy_to_excluding()
- [ ] bson_copy_to_excluding_noinit()
- [ ] bson_count_keys()
- [ ] bson_destroy()
- [ ] bson_destroy_with_steal()
- [ ] bson_equal()
- [ ] bson_get_data()
- [ ] bson_has_field()
- [ ] bson_init()
- [ ] bson_init_from_json()
- [ ] bson_init_static()
- [ ] bson_new()
- [ ] bson_new_from_buffer()
- [ ] bson_new_from_data()
- [ ] bson_new_from_json()
- [ ] bson_reinit()
- [ ] bson_reserve_buffer()
- [ ] bson_sized_new()
- [ ] bson_steal()
- [ ] bson_validate()
- [ ] bson_validate_with_error()

##### bson_context_t

- [x] bson_context_flags_t
- [x] BSON_CONTEXT_NONE
- [x] BSON_CONTEXT_THREAD_SAFE
- [x] BSON_CONTEXT_DISABLE_HOST_CACHE
- [x] BSON_CONTEXT_DISABLE_PID_CACHE
- [ ] BSON_CONTEXT_USE_TASK_ID
- [x] bson_context_t
- [x] bson_context_destroy()
- [x] bson_context_get_default()
- [x] bson_context_new()

##### bson_decimal128_t

- [x] bson_decimal128_t
- [x] bson_decimal128_from_string()
- [x] bson_decimal128_from_string_w_len()
- [x] bson_decimal128_to_string()

##### bson_error_t

- bson_error_t
- bson_set_error()
- bson_strerror_r()

##### bson_iter_t

- [ ] bson_iter_t
- [ ] bson_iter_array()
- [ ] bson_iter_as_bool()
- [ ] bson_iter_as_double()
- [ ] bson_iter_as_int64()
- [ ] bson_iter_binary()
- [ ] bson_iter_bool()
- [ ] bson_iter_code()
- [ ] bson_iter_codewscope()
- [ ] bson_iter_date_time()
- [ ] bson_iter_dbpointer()
- [ ] bson_iter_decimal128()
- [ ] bson_iter_document()
- [ ] bson_iter_double()
- [ ] bson_iter_dup_utf8()
- [ ] bson_iter_find()
- [ ] bson_iter_find_case()
- [ ] bson_iter_find_descendant()
- [ ] bson_iter_find_w_len()
- [ ] bson_iter_init()
- [ ] bson_iter_init_find()
- [ ] bson_iter_init_find_case()
- [ ] bson_iter_init_find_w_len()
- [ ] bson_iter_init_from_data()
- [ ] bson_iter_int32()
- [ ] bson_iter_int64()
- [ ] bson_iter_key()
- [ ] bson_iter_next()
- [ ] bson_iter_oid()
- [ ] bson_iter_overwrite_bool()
- [ ] bson_iter_overwrite_decimal128()
- [ ] bson_iter_overwrite_double()
- [ ] bson_iter_overwrite_int32()
- [ ] bson_iter_overwrite_int64()
- [ ] bson_iter_recurse()
- [ ] bson_iter_regex()
- [ ] bson_iter_symbol()
- [ ] bson_iter_time_t()
- [ ] bson_iter_timestamp()
- [ ] bson_iter_timeval()
- [ ] bson_iter_type()
- [ ] bson_iter_utf8()
- [ ] bson_iter_value()
- [ ] bson_iter_visit_all()

##### bson_json_reader_t

- [ ] bson_json_reader_t
- [ ] bson_json_error_code_t
- [ ] BSON_JSON_ERROR_READ_CORRUPT_JS
- [ ] BSON_JSON_ERROR_READ_INVALID_PARAM
- [ ] BSON_JSON_ERROR_READ_CB_FAILURE
- [ ] bson_json_data_reader_ingest()
- [ ] bson_json_data_reader_new()
- [ ] bson_json_reader_destroy()
- [ ] bson_json_reader_new()
- [ ] bson_json_reader_new_from_fd()
- [ ] bson_json_reader_new_from_file()
- [ ] bson_json_reader_read()

##### bson_md5_t

- [x] bson_md5_t
- [x] bson_md5_append()
- [x] bson_md5_finish()
- [x] bson_md5_init()

##### bson_oid_t

- [x] bson_oid_t
- [ ] bson_oid_compare()
- [ ] bson_oid_copy()
- [ ] bson_oid_equal()
- [ ] bson_oid_get_time_t()
- [ ] bson_oid_hash()
- [ ] bson_oid_init()
- [ ] bson_oid_init_from_data()
- [ ] bson_oid_init_from_string()
- [ ] bson_oid_init_sequence()
- [ ] bson_oid_is_valid()
- [ ] bson_oid_to_string()

### Never Implementation

##### Memory Management

- bson_free()
- bson_malloc()
- bson_malloc0()
- bson_mem_restore_vtable()
- bson_mem_set_vtable()
- bson_realloc()
- bson_realloc_ctx()
- bson_realloc_func
- bson_zero_free()
