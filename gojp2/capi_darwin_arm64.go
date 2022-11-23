// Code generated by 'ccgo -pkgname gojp2 -trace-translation-units -export-externs X -export-defines D -export-fields F -export-structs S -export-typedefs T -I../src/lib/openjp2/ -I../src/bin/common/ ../src/bin/common/color.c ../src/lib/openjp2/bio.c ../src/lib/openjp2/cio.c ../src/lib/openjp2/dwt.c ../src/lib/openjp2/event.c ../src/lib/openjp2/function_list.c ../src/lib/openjp2/helper.c ../src/lib/openjp2/ht_dec.c ../src/lib/openjp2/image.c ../src/lib/openjp2/invert.c ../src/lib/openjp2/j2k.c ../src/lib/openjp2/jp2.c ../src/lib/openjp2/mct.c ../src/lib/openjp2/mqc.c ../src/lib/openjp2/openjpeg.c ../src/lib/openjp2/opj_malloc.c ../src/lib/openjp2/pi.c ../src/lib/openjp2/sparse_array.c ../src/lib/openjp2/t1.c ../src/lib/openjp2/t2.c ../src/lib/openjp2/tcd.c ../src/lib/openjp2/tgt.c ../src/lib/openjp2/thread.c', DO NOT EDIT.

package gojp2

var CAPI = map[string]struct{}{
	"__darwin_check_fd_set_overflow":          {},
	"__inline_isfinited":                      {},
	"__inline_isfinitef":                      {},
	"__inline_isfinitel":                      {},
	"__inline_isinfd":                         {},
	"__inline_isinff":                         {},
	"__inline_isinfl":                         {},
	"__inline_isnand":                         {},
	"__inline_isnanf":                         {},
	"__inline_isnanl":                         {},
	"__inline_isnormald":                      {},
	"__inline_isnormalf":                      {},
	"__inline_isnormall":                      {},
	"__inline_signbitd":                       {},
	"__inline_signbitf":                       {},
	"__inline_signbitl":                       {},
	"__isctype":                               {},
	"__istype":                                {},
	"__sincos":                                {},
	"__sincosf":                               {},
	"__sincospi":                              {},
	"__sincospif":                             {},
	"__wcwidth":                               {},
	"color_cmyk_to_rgb":                       {},
	"color_esycc_to_rgb":                      {},
	"color_sycc_to_rgb":                       {},
	"decode_j2k":                              {},
	"isalnum":                                 {},
	"isalpha":                                 {},
	"isascii":                                 {},
	"isblank":                                 {},
	"iscntrl":                                 {},
	"isdigit":                                 {},
	"isgraph":                                 {},
	"islower":                                 {},
	"isprint":                                 {},
	"ispunct":                                 {},
	"isspace":                                 {},
	"isupper":                                 {},
	"isxdigit":                                {},
	"j2k_destroy_cstr_index":                  {},
	"j2k_dump":                                {},
	"j2k_dump_image_comp_header":              {},
	"j2k_dump_image_header":                   {},
	"j2k_get_cstr_index":                      {},
	"j2k_get_cstr_info":                       {},
	"jp2_dump":                                {},
	"jp2_get_cstr_index":                      {},
	"jp2_get_cstr_info":                       {},
	"opj_aligned_32_malloc":                   {},
	"opj_aligned_32_realloc":                  {},
	"opj_aligned_free":                        {},
	"opj_aligned_malloc":                      {},
	"opj_aligned_realloc":                     {},
	"opj_alloc_tile_component_data":           {},
	"opj_bio_create":                          {},
	"opj_bio_destroy":                         {},
	"opj_bio_flush":                           {},
	"opj_bio_inalign":                         {},
	"opj_bio_init_dec":                        {},
	"opj_bio_init_enc":                        {},
	"opj_bio_numbytes":                        {},
	"opj_bio_read":                            {},
	"opj_bio_write":                           {},
	"opj_calculate_norms":                     {},
	"opj_calloc":                              {},
	"opj_codec_set_threads":                   {},
	"opj_cond_create":                         {},
	"opj_cond_destroy":                        {},
	"opj_cond_signal":                         {},
	"opj_cond_wait":                           {},
	"opj_copy_image_header":                   {},
	"opj_create_compress":                     {},
	"opj_create_decompress":                   {},
	"opj_decode":                              {},
	"opj_decode_tile_data":                    {},
	"opj_decoder_set_strict_mode":             {},
	"opj_decompress":                          {},
	"opj_destroy_codec":                       {},
	"opj_destroy_cstr_index":                  {},
	"opj_destroy_cstr_info":                   {},
	"opj_dump_codec":                          {},
	"opj_dwt_calc_explicit_stepsizes":         {},
	"opj_dwt_decode":                          {},
	"opj_dwt_decode_real":                     {},
	"opj_dwt_encode":                          {},
	"opj_dwt_encode_real":                     {},
	"opj_dwt_getnorm":                         {},
	"opj_dwt_getnorm_real":                    {},
	"opj_encode":                              {},
	"opj_encoder_set_extra_options":           {},
	"opj_end_compress":                        {},
	"opj_end_decompress":                      {},
	"opj_event_msg":                           {},
	"opj_free":                                {},
	"opj_get_cstr_index":                      {},
	"opj_get_cstr_info":                       {},
	"opj_get_decoded_tile":                    {},
	"opj_get_encoding_packet_count":           {},
	"opj_get_num_cpus":                        {},
	"opj_has_thread_support":                  {},
	"opj_image_comp_header_update":            {},
	"opj_image_create":                        {},
	"opj_image_create0":                       {},
	"opj_image_data_alloc":                    {},
	"opj_image_data_free":                     {},
	"opj_image_destroy":                       {},
	"opj_image_tile_create":                   {},
	"opj_j2k_convert_progression_order":       {},
	"opj_j2k_create_compress":                 {},
	"opj_j2k_create_decompress":               {},
	"opj_j2k_decode":                          {},
	"opj_j2k_decode_tile":                     {},
	"opj_j2k_decoder_set_strict_mode":         {},
	"opj_j2k_destroy":                         {},
	"opj_j2k_encode":                          {},
	"opj_j2k_encoder_set_extra_options":       {},
	"opj_j2k_end_compress":                    {},
	"opj_j2k_end_decompress":                  {},
	"opj_j2k_get_tile":                        {},
	"opj_j2k_read_header":                     {},
	"opj_j2k_read_tile_header":                {},
	"opj_j2k_set_decode_area":                 {},
	"opj_j2k_set_decoded_components":          {},
	"opj_j2k_set_decoded_resolution_factor":   {},
	"opj_j2k_set_threads":                     {},
	"opj_j2k_setup_decoder":                   {},
	"opj_j2k_setup_encoder":                   {},
	"opj_j2k_setup_mct_encoding":              {},
	"opj_j2k_start_compress":                  {},
	"opj_j2k_write_tile":                      {},
	"opj_jp2_create":                          {},
	"opj_jp2_decode":                          {},
	"opj_jp2_decode_tile":                     {},
	"opj_jp2_decoder_set_strict_mode":         {},
	"opj_jp2_destroy":                         {},
	"opj_jp2_encode":                          {},
	"opj_jp2_encoder_set_extra_options":       {},
	"opj_jp2_end_compress":                    {},
	"opj_jp2_end_decompress":                  {},
	"opj_jp2_get_tile":                        {},
	"opj_jp2_read_header":                     {},
	"opj_jp2_read_tile_header":                {},
	"opj_jp2_set_decode_area":                 {},
	"opj_jp2_set_decoded_components":          {},
	"opj_jp2_set_decoded_resolution_factor":   {},
	"opj_jp2_set_threads":                     {},
	"opj_jp2_setup_decoder":                   {},
	"opj_jp2_setup_encoder":                   {},
	"opj_jp2_start_compress":                  {},
	"opj_jp2_write_tile":                      {},
	"opj_malloc":                              {},
	"opj_matrix_inversion_f":                  {},
	"opj_mct_decode":                          {},
	"opj_mct_decode_custom":                   {},
	"opj_mct_decode_real":                     {},
	"opj_mct_encode":                          {},
	"opj_mct_encode_custom":                   {},
	"opj_mct_encode_real":                     {},
	"opj_mct_get_mct_norms":                   {},
	"opj_mct_get_mct_norms_real":              {},
	"opj_mct_getnorm":                         {},
	"opj_mct_getnorm_real":                    {},
	"opj_mqc_bypass_enc":                      {},
	"opj_mqc_bypass_flush_enc":                {},
	"opj_mqc_bypass_get_extra_bytes":          {},
	"opj_mqc_bypass_init_enc":                 {},
	"opj_mqc_byteout":                         {},
	"opj_mqc_erterm_enc":                      {},
	"opj_mqc_flush":                           {},
	"opj_mqc_init_dec":                        {},
	"opj_mqc_init_enc":                        {},
	"opj_mqc_numbytes":                        {},
	"opj_mqc_raw_init_dec":                    {},
	"opj_mqc_reset_enc":                       {},
	"opj_mqc_resetstates":                     {},
	"opj_mqc_restart_init_enc":                {},
	"opj_mqc_segmark_enc":                     {},
	"opj_mqc_setstate":                        {},
	"opj_mutex_create":                        {},
	"opj_mutex_destroy":                       {},
	"opj_mutex_lock":                          {},
	"opj_mutex_unlock":                        {},
	"opj_pi_create_decode":                    {},
	"opj_pi_create_encode":                    {},
	"opj_pi_destroy":                          {},
	"opj_pi_initialise_encode":                {},
	"opj_pi_next":                             {},
	"opj_pi_update_encoding_parameters":       {},
	"opj_procedure_list_add_procedure":        {},
	"opj_procedure_list_clear":                {},
	"opj_procedure_list_create":               {},
	"opj_procedure_list_destroy":              {},
	"opj_procedure_list_get_first_procedure":  {},
	"opj_procedure_list_get_nb_procedures":    {},
	"opj_read_bytes_BE":                       {},
	"opj_read_bytes_LE":                       {},
	"opj_read_double_BE":                      {},
	"opj_read_double_LE":                      {},
	"opj_read_float_BE":                       {},
	"opj_read_float_LE":                       {},
	"opj_read_header":                         {},
	"opj_read_tile_header":                    {},
	"opj_realloc":                             {},
	"opj_set_MCT":                             {},
	"opj_set_decode_area":                     {},
	"opj_set_decoded_components":              {},
	"opj_set_decoded_resolution_factor":       {},
	"opj_set_default_decoder_parameters":      {},
	"opj_set_default_encoder_parameters":      {},
	"opj_set_default_event_handler":           {},
	"opj_set_error_handler":                   {},
	"opj_set_info_handler":                    {},
	"opj_set_warning_handler":                 {},
	"opj_setup_decoder":                       {},
	"opj_setup_encoder":                       {},
	"opj_sparse_array_int32_create":           {},
	"opj_sparse_array_int32_free":             {},
	"opj_sparse_array_int32_read":             {},
	"opj_sparse_array_int32_write":            {},
	"opj_sparse_array_is_region_valid":        {},
	"opj_start_compress":                      {},
	"opj_stream_create":                       {},
	"opj_stream_create_buffer_stream":         {},
	"opj_stream_create_default_buffer_stream": {},
	"opj_stream_create_default_file_stream":   {},
	"opj_stream_create_file_stream":           {},
	"opj_stream_default_create":               {},
	"opj_stream_default_read":                 {},
	"opj_stream_default_seek":                 {},
	"opj_stream_default_skip":                 {},
	"opj_stream_default_write":                {},
	"opj_stream_destroy":                      {},
	"opj_stream_flush":                        {},
	"opj_stream_get_number_byte_left":         {},
	"opj_stream_has_seek":                     {},
	"opj_stream_read_data":                    {},
	"opj_stream_read_seek":                    {},
	"opj_stream_read_skip":                    {},
	"opj_stream_seek":                         {},
	"opj_stream_set_current_data":             {},
	"opj_stream_set_read_function":            {},
	"opj_stream_set_seek_function":            {},
	"opj_stream_set_skip_function":            {},
	"opj_stream_set_user_data":                {},
	"opj_stream_set_user_data_length":         {},
	"opj_stream_set_write_function":           {},
	"opj_stream_skip":                         {},
	"opj_stream_tell":                         {},
	"opj_stream_write_data":                   {},
	"opj_stream_write_seek":                   {},
	"opj_stream_write_skip":                   {},
	"opj_t1_create":                           {},
	"opj_t1_decode_cblks":                     {},
	"opj_t1_destroy":                          {},
	"opj_t1_encode_cblks":                     {},
	"opj_t1_ht_decode_cblk":                   {},
	"opj_t2_create":                           {},
	"opj_t2_decode_packets":                   {},
	"opj_t2_destroy":                          {},
	"opj_t2_encode_packets":                   {},
	"opj_tcd_copy_tile_data":                  {},
	"opj_tcd_create":                          {},
	"opj_tcd_decode_tile":                     {},
	"opj_tcd_destroy":                         {},
	"opj_tcd_encode_tile":                     {},
	"opj_tcd_get_decoded_tile_size":           {},
	"opj_tcd_get_encoder_input_buffer_size":   {},
	"opj_tcd_init":                            {},
	"opj_tcd_init_decode_tile":                {},
	"opj_tcd_init_encode_tile":                {},
	"opj_tcd_is_band_empty":                   {},
	"opj_tcd_is_subband_area_of_interest":     {},
	"opj_tcd_makelayer":                       {},
	"opj_tcd_makelayer_fixed":                 {},
	"opj_tcd_marker_info_create":              {},
	"opj_tcd_marker_info_destroy":             {},
	"opj_tcd_rateallocate":                    {},
	"opj_tcd_rateallocate_fixed":              {},
	"opj_tcd_reinit_segment":                  {},
	"opj_tcd_update_tile_data":                {},
	"opj_tgt_create":                          {},
	"opj_tgt_decode":                          {},
	"opj_tgt_destroy":                         {},
	"opj_tgt_encode":                          {},
	"opj_tgt_init":                            {},
	"opj_tgt_reset":                           {},
	"opj_tgt_setvalue":                        {},
	"opj_thread_create":                       {},
	"opj_thread_join":                         {},
	"opj_thread_pool_create":                  {},
	"opj_thread_pool_destroy":                 {},
	"opj_thread_pool_get_thread_count":        {},
	"opj_thread_pool_submit_job":              {},
	"opj_thread_pool_wait_completion":         {},
	"opj_tls_get":                             {},
	"opj_tls_set":                             {},
	"opj_version":                             {},
	"opj_write_bytes_BE":                      {},
	"opj_write_bytes_LE":                      {},
	"opj_write_double_BE":                     {},
	"opj_write_double_LE":                     {},
	"opj_write_float_BE":                      {},
	"opj_write_float_LE":                      {},
	"opj_write_tile":                          {},
	"opq_mqc_finish_dec":                      {},
	"toascii":                                 {},
	"tolower":                                 {},
	"toupper":                                 {},
}
