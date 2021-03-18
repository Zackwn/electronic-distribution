package core

func Run(electrons Electrons) string {
	sblvlDstn := runSublevelDistribution(electrons)
	return sblvlDstn
}
