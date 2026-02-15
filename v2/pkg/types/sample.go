package types

type SlgParams struct {
	Layers     []int32
	LayerCount int64
	LayerStart float32
	LayerEnd   float32
	Scale      float32
}

type GuidanceParams struct {
	TxtCfg            float32
	ImgCfg            float32
	DistilledGuidance float32
	Slg               SlgParams
}

type SampleParams struct {
	Guidance          GuidanceParams
	Scheduler         int32
	SampleMethod      int32
	SampleSteps       int32
	Eta               float32
	ShiftedTimestep   int32
	CustomSigmas      []float32
	CustomSigmasCount int32
}