package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type race_status struct {
	RaceHubID string
	Locale string
	CreatedAt string
	UpdatedAt string
	FomRaceID string
	BrandColourHexadecimal string
	CircuitSmallImage struct {
		Title string
		Path string
		Url string
	}
	Links []struct {
		Text string
		Url string
	}
	SeasonContext struct {
		Id string
		ContentType string
		CreatedAt time.Time
		UpdatedAt time.Time
		Locale string
		SeasonYear string
		CurrentOrNextMeetingKey string
		State string
		EventState string
		LiveEventId string
		LiveTimingsSource string
		LiveBlog struct {
			ContentType string
			Title string
			ScribbleEventId string
		}
		SeasonState string
		RaceListingOverride int
		DriverAndTeamListingOverride int
		Timetables []struct {
			State string
			Session string
			GmtOffset string
			Description string
			EndTime string //time.Time
			StartTime string //time.Time
		}
		ReplayBaseUrl string
		SeasonContextUIState int
	}
	RaceResults []struct{
	}
	Race struct {
		MeetingCountryName string
		MeetingStartDate time.Time
		MeetingOfficialName string // "FORMULA 1 PIRELLI GRAN PREMIO D’ITALIA 2022",
		MeetingEndDate time.Time
	}
	SeasonYearImage string
}

func getResults(URL string) race_status{
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Printf("reponse error: %s", err)
	}
	req.Header.Add("apiKey", "qPgPPRJyGCIPxFT3el4MF7thXHyJCzAP")
	req.Header.Add("locale", "en")

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("reponse error: %s", err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	res := race_status{}
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return res
}

func main() {
	// branch = 'main'
	// g = Git.open(__dir__)
	// g.checkout(branch)
	// # g.pull

	race := getResults("https://api.formula1.com/v1/event-tracker")
	fmt.Println(race.Race.MeetingCountryName, "will begin on", race.Race.MeetingStartDate)

	// resBytes := []byte(res)
	// var results map[string]interface{}
	// _ = json.Unmarshal(resBytes, &results)
	// fmt.Println(results["seasonContext"].(string))

	// race_status = JSON.parse(response.body)
	// raceinfo = race_status['seasonContext']['timetables'].select { |hash| hash['description'] == 'Race' }.first
	// raceStart = Time.parse("#{raceinfo['startTime']} #{raceinfo['gmtOffset']}").utc

	// config = YAML.load_file("#{__dir__}/_config.yml")
	// config['description'] = "#{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}"
	// File.write("#{__dir__}/_config.yml", config.to_yaml)

	// case raceinfo['state']
	// when 'upcoming'
	// 	answer = "The race hasn't started yet"
	// when 'started'
	// 	answer = "They're racing now"
	// when 'completed'
	// 	winner = race_status['raceResults'].select { |hash| hash['positionNumber'] == '1' }.first
	// 	answer = winner['driverLastName'].downcase == 'hamilton' ? 'YES' : 'NO'
	// end
	// File.write("#{__dir__}/index.md", "# #{answer}")
	// File.write("#{__dir__}/api_results/#{Time.now.strftime('%Y%m%d-%H%M%S')}.json", JSON.pretty_generate(race_status))

	// if g.status.changed.count.positive?
	// 	g.add(all: true)
	// 	g.commit_all("updating for #{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}; race #{raceinfo['state']}")
	// 	g.push(g.remote('origin'), branch)
	// end
}
