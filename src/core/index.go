package core

func Run(electrons Electrons) string {
	sblvlDstn := runSublevelDistribution(electrons)
	layerDstn := runLayerDistribution(electrons)

	return sblvlDstn + "\n\n" + layerDstn
}
