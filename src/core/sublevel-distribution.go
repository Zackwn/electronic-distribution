package core

func runSublevelDistribution(electrons Electrons) string {
	sublevelDistribution := ""

	index := 0

	for electrons > 0 {
		sublevel := sublevelsSequence[index]

		if sublevel.maxElectrons >= electrons {
			sublevelDistribution += formatDistributionSublevel(DistributionSublevel{
				value:          electrons,
				layer:          sublevel.layer,
				representation: sublevel.representation,
			})
			break
		}

		electrons = electrons - sublevel.maxElectrons
		sublevelDistribution += formatDistributionSublevel(DistributionSublevel{
			value:          sublevel.maxElectrons,
			layer:          sublevel.layer,
			representation: sublevel.representation,
		}) + " "
		index++
	}

	return sublevelDistribution
}
