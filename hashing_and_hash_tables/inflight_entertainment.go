package hashing_and_hash_tables

func canFindMoviesForFlightDuration(flightLength int, movieLengths []int) bool {
	movieLengthMap := map[int]struct{}{}
	for _, movieLength := range movieLengths {
		movieLengthMap[movieLength] = struct{}{}
		//Is there another movie that equals the flight time?
		if _, ok := movieLengthMap[flightLength-movieLength]; ok {
			return true
		}
	}
	return false
}

func canFindMoviesForFlightDurationWithin20Mins(flightLength int, movieLengths []int) bool {
	lowerBound := flightLength - 20
	upperBound := flightLength + 20
	for i, movieLength := range movieLengths {
		for j := i + 1; j <= len(movieLengths); j++ {
			combinedMovieLength := movieLength + movieLengths[j]
			if combinedMovieLength > lowerBound && combinedMovieLength < upperBound {
				return true
			}
		}
	}
	return false
}
