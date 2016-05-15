package main;

import (
	"github.com/ngraves95/util"
	"github.com/ngraves95/run-temp/weather"
	"github.com/ngraves95/run-temp/tokens"
	"github.com/strava/go.strava"
)

// My athlete id.
const athleteId int64 = 10081791;

func main() {
	// Loop forever in background.
	tokenManager := tokens.NewTokenManager("/home/ngraves3/gocode/src/github.com/ngraves95/run-temp/tokens/");
	client := strava.NewClient(tokenManager.GetToken("strava_access_token"));

	// Gives a list of activity summaries with most recent first.
	activities, err := strava.NewAthletesService(client).
		ListActivities(athleteId).
		Do();
	util.Check(err);

	// This is the most recent activity.
	// We need to Get the activity Id in order to get the ActivityDetailed
	// struct, which contains the description.
	targetActivityId := activities[0].Id;

	// Get the ActivityDetailed struct so we can access the Description.
	activityService := strava.NewActivitiesService(client);
	activity, err := activityService.
		Get(targetActivityId).
		Do();
	util.Check(err);

	// Check if the activity already contains a temperature field to
	// prevent writing it many times.
	description := activity.Description;

	if !weather.ContainsWeatherData(description) {
		// Get start location to use for weather querying.
		// array of 2 float64s.
		var location [2]float64 = activities[0].StartLocation;
		lat := location[0];
		long := location[1];

		// update the activity description to include the temperature.
		strava.NewActivitiesService(client).
			Update(targetActivityId).
			Description(weather.AddWeatherData(lat, long, description)).
			Do();
	}
}
