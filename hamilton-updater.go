package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
	RaceResults []struct {
		DriverTLA string
		DriverFirstName string
		DriverLastName string
		TeamName string
		PositionNumber string
		RaceTime string
		TeamColourCode string
		GapToLeader string
		DriverImage string
		DriverNameFormat string
	}
	Race struct {
		MeetingCountryName string
		MeetingStartDate time.Time
		MeetingOfficialName string // "FORMULA 1 PIRELLI GRAN PREMIO D’ITALIA 2022",
		MeetingEndDate time.Time
	}
	SeasonYearImage string
	SessionLinkSets struct {
    ReplayLinks []struct {
			Session string
			Text string
			Url string
			LinkType string
		}
	}
	CurrentRaceStatus string
	Winner string
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

	// body, _ := ioutil.ReadAll(response.Body)

	body, err := os.ReadFile("../didhamiltonwin/api_results/20220904-132607.json")

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

	// race_status = JSON.parse(response.body)
	// raceinfo = race_status['seasonContext']['timetables'].select { |hash| hash['description'] == 'Race' }.first
	// raceStart = Time.parse("#{raceinfo['startTime']} #{raceinfo['gmtOffset']}").utc

	race := getResults("https://api.formula1.com/v1/event-tracker")

	// config = YAML.load_file("#{__dir__}/_config.yml")
	// config['description'] = "#{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}"
	// File.write("#{__dir__}/_config.yml", config.to_yaml)

	for _, event := range race.SeasonContext.Timetables {
		if event.Description == "Race" {
			switch event.State {
			case "upcoming":
				race.CurrentRaceStatus = "The race hasn't started yet"
			case "started":
				race.CurrentRaceStatus = "They're racing now"
			case "completed":
				race.CurrentRaceStatus = "They're done racing"
				for _, result := range race.RaceResults {
					position, err := strconv.Atoi(result.PositionNumber)
					if err != nil {
						fmt.Println("err:", err)
					}
					if position == 1 {
						race.Winner = result.DriverTLA
					}
				}
			}
		}
	}

	fmt.Println("state:", race.CurrentRaceStatus)
	fmt.Println("winner:", race.Winner)
	if race.Winner == "HAM" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// File.write("#{__dir__}/index.md", "# #{answer}")
	// File.write("#{__dir__}/api_results/#{Time.now.strftime('%Y%m%d-%H%M%S')}.json", JSON.pretty_generate(race_status))

	// if g.status.changed.count.positive?
	// 	g.add(all: true)
	// 	g.commit_all("updating for #{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}; race #{raceinfo['state']}")
	// 	g.push(g.remote('origin'), branch)
	// end
}
