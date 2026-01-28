package binding

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/pkg/types"
)

var (
	// SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	schedulerNameFun ffi.Fun

	// SD_API enum scheduler_t str_to_scheduler(const char* str);
	strToSchedulerFun ffi.Fun

	// SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	getDefaultSchedulerFun ffi.Fun
)

func LoadSchedulerFuns(lib ffi.Lib) error {
	var err error

	// SD_API const char* sd_scheduler_name(enum scheduler_t scheduler);
	schedulerNameFun, err = lib.Prep("sd_scheduler_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum scheduler_t str_to_scheduler(const char* str);
	strToSchedulerFun, err = lib.Prep("str_to_scheduler", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	// SD_API enum scheduler_t sd_get_default_scheduler(const sd_ctx_t* sd_ctx, enum sample_method_t sample_method);
	getDefaultSchedulerFun, err = lib.Prep("sd_get_default_scheduler", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}
	return nil
}

// SchedulerName 获取调度器名称
func SchedulerName(scheduler types.Scheduler) *byte {
	var result *byte
	schedulerNameFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&scheduler))
	return result
}

// StrToScheduler 将字符串转换为调度器
func StrToScheduler(str *byte) types.Scheduler {
	var result int32
	strToSchedulerFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.Scheduler(result)
}

// GetDefaultScheduler 获取默认调度器
func GetDefaultScheduler(ctx Context, sampleMethod types.SampleMethod) types.Scheduler {
	var result int32
	getDefaultSchedulerFun.Call(unsafe.Pointer(&result), unsafe.Pointer(&ctx), unsafe.Pointer(&sampleMethod))
	return types.Scheduler(result)
}
