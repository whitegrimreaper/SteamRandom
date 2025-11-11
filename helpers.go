package main

func shouldCheckGame(game GameEntry)(resp bool) {
	// resp = false; <-- add this as default
	// TODO: Implement this
	return false
}

func getUnplayedGames()(resp []GameEntry) {
	list := getAllGameEntries()
	for _, game := range list {
		// gonna assume playtime is in minutes since that's what steam usually uses
		// TODO: Make this configurable
		if game.HasPlayedYet == false || game.Playtime < 15 {
			resp = append(resp, game)
		}
	}
	return resp;
}