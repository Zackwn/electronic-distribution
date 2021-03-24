package core

import "fmt"

var LayersRepresentation = map[int]string{
	1: "K",
	2: "L",
	3: "M",
	4: "N",
	5: "O",
	6: "P",
	7: "Q",
}

/*
	main array cap: 7
	layer array cap: 4
*/

type Layer struct {
	representation string
	sublevels      string
	electrons      Electrons
}

func runLayerDistribution(electrons Electrons) string {
	distribution := make([]Layer, 0, 7)

	index := 0

ElectronsLoop:
	for electrons > 0 {
		sublevel := sublevelsSequence[index]
		index++

		dSublevelElectrons := sublevel.maxElectrons
		if sublevel.maxElectrons >= electrons {
			dSublevelElectrons = electrons
			electrons = 0
		} else {
			electrons -= sublevel.maxElectrons
		}

		for dIndex := 0; dIndex < len(distribution); dIndex++ {
			if distribution[dIndex].representation == LayersRepresentation[sublevel.layer] {
				// layer alredy exists
				distribution[dIndex].sublevels += " " + formatDistributionSublevel(DistributionSublevel{
					value:          dSublevelElectrons,
					representation: sublevel.representation,
					layer:          sublevel.layer,
				})
				distribution[dIndex].electrons += dSublevelElectrons
				continue ElectronsLoop
			}
		}

		distribution = append(distribution, Layer{
			electrons:      dSublevelElectrons,
			representation: LayersRepresentation[sublevel.layer],
			sublevels: formatDistributionSublevel(DistributionSublevel{
				value:          dSublevelElectrons,
				representation: sublevel.representation,
				layer:          sublevel.layer,
			}),
		})
	}

	result := ""

	for _, layer := range distribution {
		result += fmt.Sprintf("[%v] %v%v\n", layer.sublevels, layer.representation, layer.electrons)
	}

	return result
}
