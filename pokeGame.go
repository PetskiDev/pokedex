package main

import "fmt"



var state = struct{
	lastPrinted	int
	limit		int
}{
	lastPrinted: -MAP_EXPLORE, //NEEDS TO BE SAME AS LIMIT
	limit: MAP_EXPLORE,
}

func getNextLocations() ([]Location, error){
	state.lastPrinted += state.limit
	locations, err := getLocation(state.lastPrinted, state.limit)
	if err != nil{
		return nil, err
	}
	return locations, nil
}


func getPrevLocations() ([]Location, error){
	if state.lastPrinted <= 0{
		return nil, fmt.Errorf("can't go back")
	}
	state.lastPrinted -= state.limit
	locations, err := getLocation(state.lastPrinted, state.limit)
	if err != nil{
		return nil, err
	}
	return locations, nil
}

