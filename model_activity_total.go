/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A roll-up of metrics pertaining to a set of activities. Values are in seconds and meters.
type ActivityTotal struct {
	// The number of activities considered in this total.
	Count int32 `json:"count,omitempty"`
	// The total distance covered by the considered activities.
	Distance float32 `json:"distance,omitempty"`
	// The total moving time of the considered activities.
	MovingTime int32 `json:"moving_time,omitempty"`
	// The total elapsed time of the considered activities.
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The total elevation gain of the considered activities.
	ElevationGain float32 `json:"elevation_gain,omitempty"`
	// The total number of achievements of the considered activities.
	AchievementCount int32 `json:"achievement_count,omitempty"`
}