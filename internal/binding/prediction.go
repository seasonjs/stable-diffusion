package binding

import (
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/pkg/types"
	"unsafe"
)

var (
	// SD_API const char* sd_prediction_name(enum prediction_t prediction);
	predictionName ffi.Fun

	// SD_API enum prediction_t str_to_prediction(const char* str);
	strToPrediction ffi.Fun
)


func LoadPredictionFuns(lib ffi.Lib) error {
	var err error

	// SD_API const char* sd_prediction_name(enum prediction_t prediction);
	predictionName, err = lib.Prep("sd_prediction_name", &ffi.TypePointer, &ffi.TypeSint32)
	if err != nil {
		return err
	}

	// SD_API enum prediction_t str_to_prediction(const char* str);
	strToPrediction, err = lib.Prep("str_to_prediction", &ffi.TypeSint32, &ffi.TypePointer)
	if err != nil {
		return err
	}

	return nil
}

func PredictionName(prediction types.Prediction) *byte {
	var result *byte
	predictionName.Call(unsafe.Pointer(&result), unsafe.Pointer(&prediction))
	return result
}

func StrToPrediction(str *byte) types.Prediction {
	var result int32
	strToPrediction.Call(unsafe.Pointer(&result), unsafe.Pointer(&str))
	return types.Prediction(result)
}