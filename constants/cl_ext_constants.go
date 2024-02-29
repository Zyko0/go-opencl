package constants

const (
	CL_LUID_SIZE_KHR                                       = 8
	CL_DEVICE_NODE_MASK_KHR                                = 0x106E
	CL_DEVICE_COMPUTE_CAPABILITY_MINOR_NV                  = 0x4001
	CL_PROGRAM_IL_KHR                                      = 0x1169
	CL_VERSION_PATCH_BITS_KHR                              = 12
	CL_IMPORT_TYPE_ANDROID_HARDWARE_BUFFER_ARM             = 0x41E2
	CL_DEVICE_PARENT_DEVICE_EXT                            = 0x4054
	CL_CHAR_MIN                                            = CL_SCHAR_MIN
	CL_COMMAND_SVM_MEMFILL_ARM                             = 0x40BC
	CL_DEVICE_TERMINATE_CAPABILITY_KHR                     = 0x2031
	CL_DEVICE_AFFINITY_DOMAINS_EXT                         = 0x4056
	CL_MEM_HOST_UNCACHED_QCOM                              = 0x40A4
	CL_MEM_HOST_IOCOHERENT_QCOM                            = 0x40A9
	CL_DEVICE_AVAILABLE_ASYNC_QUEUES_AMD                   = 0x404C
	CL_PLATFORM_EXTENSIONS_WITH_VERSION_KHR                = 0x0907
	CL_COMMAND_SVM_UNMAP_ARM                               = 0x40BE
	CL_DEVICE_WARP_SIZE_NV                                 = 0x4003
	CL_COMMAND_MIGRATE_MEM_OBJECT_EXT                      = 0x4040
	CL_CONTEXT_TERMINATE_KHR                               = 0x2032
	CL_DEVICE_INTEGRATED_MEMORY_NV                         = 0x4006
	CL_COMMAND_RELEASE_GRALLOC_OBJECTS_IMG                 = 0x40D3
	CL_DEVICE_THREAD_TRACE_SUPPORTED_AMD                   = 0x4049
	CL_MEM_USES_SVM_POINTER_ARM                            = 0x40B7
	CL_IMPORT_DMA_BUF_DATA_CONSISTENCY_WITH_HOST_ARM       = 0x41E3
	CL_DEVICE_LOCAL_MEM_BANKS_AMD                          = 0x4048
	CL_MEM_SVM_ATOMICS_ARM                                 = 1
	CL_DEVICE_REFERENCE_COUNT_EXT                          = 0x4057
	CL_AFFINITY_DOMAIN_NEXT_FISSIONABLE_EXT                = 0x100
	CL_DEVICE_MAX_WORK_GROUP_SIZE_AMD                      = 0x4031
	CL_MEM_USE_GRALLOC_PTR_IMG                             = 1
	CL_PLATFORM_NUMERIC_VERSION_KHR                        = 0x0906
	CL_COMMAND_ACQUIRE_GRALLOC_OBJECTS_IMG                 = 0x40D2
	CL_DEVICE_TOPOLOGY_AMD                                 = 0x4037
	CL_QUEUE_THROTTLE_MED_KHR                              = (1 << 1)
	CL_MEM_USE_UNCACHED_CPU_MEMORY_IMG                     = 1
	CL_MEM_SVM_FINE_GRAIN_BUFFER_ARM                       = 1
	CL_DEVICE_ILS_WITH_VERSION_KHR                         = 0x1061
	CL_AFFINITY_DOMAIN_L3_CACHE_EXT                        = 0x3
	CL_COMMAND_SVM_MAP_ARM                                 = 0x40BD
	CL_PLATFORM_ICD_SUFFIX_KHR                             = 0x0920
	CL_DEVICE_LUID_KHR                                     = 0x106D
	CL_DEVICE_LOCAL_MEM_SIZE_PER_COMPUTE_UNIT_AMD          = 0x4047
	CL_NAME_VERSION_MAX_NAME_SIZE_KHR                      = 64
	CL_DEVICE_LUID_VALID_KHR                               = 0x106C
	CL_DEVICE_SVM_FINE_GRAIN_BUFFER_ARM                    = 1
	CL_QUEUE_THROTTLE_LOW_KHR                              = 1 << 2
	CL_KERNEL_EXEC_INFO_SVM_FINE_GRAIN_SYSTEM_ARM          = 0x40B9
	CL_INVALID_PARTITION_NAME_EXT                          = -1059
	CL_DEVICE_SVM_CAPABILITIES_ARM                         = 0x40B6
	CL_QUEUE_THROTTLE_HIGH_KHR                             = 1 << 0
	CL_QUEUE_PRIORITY_MED_KHR                              = 1 << 1
	CL_IMPORT_TYPE_HOST_ARM                                = 0x40B3
	CL_DEVICE_SVM_FINE_GRAIN_SYSTEM_ARM                    = 1
	CL_CONTEXT_MEMORY_INITIALIZE_KHR                       = 0x2030
	CL_DEVICE_GLOBAL_MEM_CHANNEL_BANK_WIDTH_AMD            = 0x4046
	CL_DEVICE_BUILT_IN_KERNELS_WITH_VERSION_KHR            = 0x1062
	CL_KERNEL_EXEC_INFO_WORKGROUP_BATCH_SIZE_ARM           = 0x41E5
	CL_DEVICE_GLOBAL_MEM_CHANNEL_BANKS_AMD                 = 0x4045
	CL_SAMPLER_MIP_FILTER_MODE_KHR                         = 0x1155
	CL_MEM_ANDROID_NATIVE_BUFFER_HOST_PTR_QCOM             = 0x40C6
	CL_DEVICE_PREFERRED_WORK_GROUP_SIZE_AMD                = 0x4030
	CL_DEVICE_GFXIP_MAJOR_AMD                              = 0x404A
	CL_SAMPLER_LOD_MIN_KHR                                 = 0x1156
	CL_VERSION_MINOR_BITS_KHR                              = (10)
	CL_KERNEL_SUB_GROUP_COUNT_FOR_NDRANGE_KHR              = 0x2034
	CL_AFFINITY_DOMAIN_NUMA_EXT                            = 0x10
	CL_DEVICE_JOB_SLOTS_ARM                                = 0x41E0
	CL_DEVICE_NUMERIC_VERSION_KHR                          = 0x105E
	CL_DEVICE_SCHEDULING_WORKGROUP_BATCH_SIZE_ARM          = 1
	CL_KERNEL_MAX_SUB_GROUP_SIZE_FOR_NDRANGE_KHR           = 0x2033
	CL_DEVICE_COMPUTE_CAPABILITY_MAJOR_NV                  = 0x4000
	CL_DEVICE_OPENCL_C_NUMERIC_VERSION_KHR                 = 0x105F
	CL_DEVICE_COMPUTE_UNITS_BITFIELD_ARM                   = 0x40BF
	CL_CONTEXT_TERMINATED_KHR                              = -1121
	CL_DEVICE_EXT_MEM_PADDING_IN_BYTES_QCOM                = 0x40A0
	CL_MEM_HOST_WRITE_COMBINING_QCOM                       = 0x40A7
	CL_AFFINITY_DOMAIN_L2_CACHE_EXT                        = 0x2
	CL_PROGRAM_BINARY_TYPE_INTERMEDIATE                    = 0x40E1
	CL_COMMAND_SVM_MEMCPY_ARM                              = 0x40BB
	CL_DEVICE_SCHEDULING_WORKGROUP_BATCH_SIZE_MODIFIER_ARM = 1
	CL_DEVICE_SIMD_WIDTH_AMD                               = 0x4041
	CL_DEVICE_GLOBAL_MEM_CHANNELS_AMD                      = 0x4044
	CL_DEVICE_SPIR_VERSIONS                                = 0x40E0
	CL_DEVICE_PROFILING_TIMER_OFFSET_AMD                   = 0x4036
	CL_DEVICE_SVM_ATOMICS_ARM                              = 1
	CL_DEVICE_PARTITION_BY_NAMES_EXT                       = 0x4052
	CL_VERSION_MINOR_MASK_KHR                              = 1
	CL_DEVICE_PARTITION_TYPES_EXT                          = 0x4055
	CL_DEVICE_GPU_OVERLAP_NV                               = 0x4004
	CL_KERNEL_EXEC_INFO_SVM_PTRS_ARM                       = 0x40B8
	CL_QUEUE_PRIORITY_KHR                                  = 0x1096
	CL_IMPORT_TYPE_ARM                                     = 0x40B2
	CL_DEVICE_SIMD_PER_COMPUTE_UNIT_AMD                    = 0x4040
	CL_IMAGE_SLICE_ALIGNMENT_QCOM                          = 0x40A3
	CL_VERSION_1_0                                         = 1
	CL_DEVICE_CXX_FOR_OPENCL_NUMERIC_VERSION_EXT           = 0x4230
	CL_VERSION_MAJOR_MASK_KHR                              = 1
	CL_MEM_ION_HOST_PTR_QCOM                               = 0x40A8
	CL_DEVICE_IL_VERSION_KHR                               = 0x105B
	CL_MEM_EXT_HOST_PTR_QCOM                               = 1
	CL_QUEUE_KERNEL_BATCHING_ARM                           = 0x41E7
	CL_IMPORT_TYPE_PROTECTED_ARM                           = 0x40B5
	CL_MEM_HOST_WRITEBACK_QCOM                             = 0x40A5
	CL_DEVICE_PARTITION_BY_COUNTS_EXT                      = 0x4051
	CL_PLATFORM_NOT_FOUND_KHR                              = -1001
	CL_KERNEL_EXEC_INFO_WORKGROUP_BATCH_SIZE_MODIFIER_ARM  = 0x41E6
	CL_DEVICE_SCHEDULING_KERNEL_BATCHING_ARM               = 1
	CL_GRALLOC_RESOURCE_NOT_ACQUIRED_IMG                   = 0x40D4
	CL_YV12_IMG                                            = 0x40D1
	CL_DEVICE_PAGE_SIZE_QCOM                               = 0x40A1
	CL_DEVICE_PREFERRED_CONSTANT_BUFFER_SIZE_AMD           = 0x4033
	CL_NV21_IMG                                            = 0x40D0
	CL_AFFINITY_DOMAIN_L1_CACHE_EXT                        = 0x1
	CL_VERSION_PATCH_MASK_KHR                              = 1
	CL_INVALID_PARTITION_COUNT_EXT                         = -1058
	CL_DEVICE_PARTITION_FAILED_EXT                         = -1057
	CL_DEVICE_IMAGE_PITCH_ALIGNMENT_KHR                    = 0x104A
	CL_DEVICE_SVM_COARSE_GRAIN_BUFFER_ARM                  = 1
	CL_IMPORT_TYPE_DMA_BUF_ARM                             = 0x40B4
	CL_IMAGE_ROW_ALIGNMENT_QCOM                            = 0x40A2
	CL_QUEUE_PRIORITY_HIGH_KHR                             = (1 << 0)
	CL_PRINTF_CALLBACK_ARM                                 = 0x40B0
	CL_QUEUE_JOB_SLOT_ARM                                  = 0x41E1
	CL_MEM_HOST_WRITETHROUGH_QCOM                          = 0x40A6
	CL_PRINTF_BUFFERSIZE_ARM                               = 0x40B1
	CL_DEVICE_IMAGE_BASE_ADDRESS_ALIGNMENT_KHR             = 0x104B
	CL_DEVICE_WAVEFRONT_WIDTH_AMD                          = 0x4043
	CL_DEVICE_SIMD_INSTRUCTION_WIDTH_AMD                   = 0x4042
	CL_DEVICE_PARTITION_STYLE_EXT                          = 0x4058
	CL_MEM_USE_CACHED_CPU_MEMORY_IMG                       = 1
	CL_MIGRATE_MEM_OBJECT_HOST_EXT                         = 0x1
	CL_DEVICE_PCIE_ID_AMD                                  = 0x4034
	CL_DEVICE_GFXIP_MINOR_AMD                              = 0x404B
	CL_COMMAND_SVM_FREE_ARM                                = 0x40BA
	CL_QUEUE_PRIORITY_LOW_KHR                              = (1 << 2)
	CL_SAMPLER_LOD_MAX_KHR                                 = 0x1157
	CL_DEVICE_HALF_FP_CONFIG                               = 0x1033
	CL_DEVICE_GLOBAL_FREE_MEMORY_AMD                       = 0x4039
	CL_VERSION_2_0                                         = 1
	CL_VERSION_2_1                                         = 1
	CL_VERSION_2_2                                         = 1
	CL_DEVICE_PARTITION_EQUALLY_EXT                        = 0x4050
	CL_AFFINITY_DOMAIN_L4_CACHE_EXT                        = 0x4
	CL_DRIVER_UUID_KHR                                     = 0x106B
	CL_DEVICE_KERNEL_EXEC_TIMEOUT_NV                       = 0x4005
	CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN_EXT             = 0x4053
	CL_DEVICE_UUID_KHR                                     = 0x106A
	CL_DEVICE_BOARD_NAME_AMD                               = 0x4038
	CL_DEVICE_SCHEDULING_CONTROLS_CAPABILITIES_ARM         = 0x41E4
	CL_VERSION_MAJOR_BITS_KHR                              = (10)
	CL_DEVICE_MAX_NAMED_BARRIER_COUNT_KHR                  = 0x2035
	CL_DEVICE_REGISTERS_PER_BLOCK_NV                       = 0x4002
	CL_DEVICE_EXTENSIONS_WITH_VERSION_KHR                  = 0x1060
)