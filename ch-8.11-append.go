package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	maxDay := costs[0].day // int
	for i := 0; i < len(costs); i++ {
		// Getting max day
		if costs[i].day > maxDay {
			maxDay = costs[i].day
		}
	}

	type itemStr struct {
		day   int
		value float64
	}

	var result []itemStr
	for i := 0; i <= maxDay; i++ {

		var filteredByDay []itemStr
		for x := 0; x < len(costs); x++ {
			if i == costs[x].day {
				filteredByDay = append(filteredByDay, itemStr{day: costs[x].day, value: costs[x].value})
			}
		}

		filteredValueCounter := 0.0
		var addedCost itemStr
		for y := 0; y < len(filteredByDay); y++ {
			filteredValueCounter += filteredByDay[y].value
			addedCost = itemStr{day: filteredByDay[y].day, value: filteredValueCounter}
		}
		result = append(result, addedCost)
	}
	var cleanedResult []float64
	for i := 0; i < len(result); i++ {
		cleanedResult = append(cleanedResult, result[i].value)
	}
	return cleanedResult
}
