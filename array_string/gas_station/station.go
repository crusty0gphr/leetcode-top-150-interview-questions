package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	total, start, tank := 0, 0, 0

	for i, amount := range gas {
		tank = tank - cost[i] + amount

		if tank < 0 {
			start = i + 1
			total += tank
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return start
}
