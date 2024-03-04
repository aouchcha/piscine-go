package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := 0; j < len(podium)-1; j++ {
			podium[j], podium[len(podium)-j-1] = podium[len(podium)-j-1], podium[j]
		}
	}
	return podium
}
