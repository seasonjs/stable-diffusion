package diffusion

import (
	"github.com/jupiterrider/ffi"
	"github.com/seasonjs/stable-diffusion/v2/internal/binding"
	"github.com/seasonjs/stable-diffusion/v2/internal/utils"
)

func loadDiffusionFuns(lib ffi.Lib) error {
	if err := binding.LoadCacheFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadContextFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadImgGenFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadLogFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadLoraFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadPredictionFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadPreprocessingFuncs(lib); err != nil {
		return err
	}

	if err := binding.LoadPreviewFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadSampleFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadSchedulerFuns(lib); err != nil {
		return err
	}

	if err := binding.LoadToosFuncs(lib); err != nil {
		return err
	}

	if err := binding.LoadUpsScalerFuncs(lib); err != nil {
		return err
	}

	if err := binding.LoadVideoGenFuns(lib); err != nil {
		return err
	}

	if err := utils.LoadStdFuns(lib); err != nil {
		return err
	}

	return nil
}
