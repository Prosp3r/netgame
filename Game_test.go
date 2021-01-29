package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

//TestCreate - Tests creating a game with game struct features
func TestCreate(t *testing.T) {

	exp := [4]int{8, 2, 30, 7}
	G := Create()
	got := [4]int{len(G.ID), len(strings.Split(G.Name, "-")), len(G.RoundScores), len(G.Status)} //Waiting

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to create a game with proper game struct parameters as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Create added new games with expected Game features  -> %s \n", succeed)
	}
}

func TestStop(t *testing.T) {

	exp := "Ended"
	//gameCount := 1
	//InitiateGame(gameCount)
	G := Create()
	G.Stop()
	got := G.Status

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to flag game status as ended as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Stop flaged game status as expected  -> %s \n", succeed)
	}
}

func TestPlay(t *testing.T) {
	exp := 0
	G := Create()
	G.Status = "Live"
	for i := 0; i < len(G.RoundScores); i++ {
		G.Play(i)
	}

	got := 0
	for _, v := range G.RoundScores {
		if v == 0 {
			got++
		}
	}

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to populate some game plays as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success populated all game plays as expected  -> %s \n", succeed)
	}
}

//TestInitiateGame - test to see how much of the game slice will be populated.
func TestInitiateGame(t *testing.T) {
	gameCount := 1
	exp := 1
	go InitiateGame(gameCount)
	Delay(3, "s")
	got := len(Games)

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to the number of games as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success initiated and returned the number of games as expected  -> %s \n", succeed)
	}
}

func Test_fakeUsersPlay(t *testing.T) {
	//ol := len(Players)
	userCount := 3
	exp := 3
	fakeUsersPlay(userCount)
	got := len(Plays)

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to create the number of fake game plays as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success created and returned the number of fake game plays as expected  -> %s \n", succeed)
	}
}

func TestAddUser(t *testing.T) {
	userNames := []string{"John", "Jame", "Joe"}
	for _, v := range userNames {
		AddUser(v)
	}
	exp := len(userNames)
	got := len(Players)
	if got == exp {
		t.Fatalf("Failed to create the number of players as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success created and returned the number of players as expected  -> %s \n", succeed)
	}
}

//Test_createRandomPlays -
func Test_createRandomPlays(t *testing.T) {
	userNames := []string{"Joffrey", "Jamie", "Jekhil"}
	for _, v := range userNames {
		AddUser(v)
	}

	exp := len(Players)
	PurgePlays() //to prevent missmatch error
	createRandomPlays()
	got := len(Plays)
	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to create the number of fake game plays as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success created and returned the number of fake game plays as expected  -> %s \n", succeed)
	}
}

//SetStartScores -
func TestSetStartScores(t *testing.T) {
	G := Create()
	ex := 30
	G.SetStartScores()

	var got int
	for _, v := range G.RoundScores {
		if v == 0 {
			got++
		}
	}

	if got != ex {
		t.Fatalf("Failed! did not adequately zero out Game Scoreboard. Got: %v  -> %s\n", got, failed)
	} else {
		t.Logf("Success! zero out Game Scoreboard as expected  -> %s \n", succeed)
	}
}

//TestSetPlays - Check to see if the plays in queue were consumed and queue cleared
func TestSetPlays(t *testing.T) {
	/*/MinimumGameEntry = 1
	playedChan := make(chan int)
	var wg sync.WaitGroup

	//Plays = []Play{{"QWERTUSD", "John", [3]int{1, 5, 0}}, {"QWREDSOD", "James", [3]int{3, 9, 0}}} //set some fake plays
	userNames := []string{"Joffrey", "Jamie", "Jekhil"}
	for _, v := range userNames {
		AddUser(v)
	}
	createRandomPlays()

	exp := 30
	G := Create()
	wg.Add(1)

	G.SetPlays(playedChan, &wg)
	fmt.Printf("Plays ===> %v \n", Plays)
	wg.Wait()
	//got := <-playedChan //checks the number of rounds played
	got := <-playedChan
	//wg.Wait()
	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to create a game with proper game struct parameters as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Create added new games with expected Game features  -> %s \n", succeed)
	}*/
}

func TestPurgePlays(t *testing.T) {
	userNames := []string{"Joffrey", "Jamie", "Jekhil"}
	for _, v := range userNames {
		AddUser(v)
	}

	createRandomPlays()
	exp := 0
	PurgePlays() //to prevent missmatch error
	got := len(Plays)
	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to purge the game plays as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success PurgePlays Clear the plays list as expected  -> %s \n", succeed)
	}
}

//Test_randomString - checks to make sure that the right size of string is being created.
func Test_randomString(t *testing.T) {
	exp := 20
	got := len(randomString(exp))

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to create the right size of random string as expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success created create the right size of random string as expected  -> %s \n", succeed)
	}
}

func Test_randInt(t *testing.T) {
	//exp := 10
	min := 1
	max := 5
	omit := 4
	got := randInt(min, max, omit)

	if got == omit {
		t.Fatalf("Failed! returned figure to be ommitted. Got: %v  -> %s\n", got, failed)
	} else if got < min {
		t.Fatalf("Failed! returned figure less than minimum. Got: %v  -> %s\n", got, failed)
	} else if got > max {
		t.Fatalf("Failed! returned figure higher than maximum. Got: %v  -> %s\n", got, failed)
	} else {
		t.Logf("Success created create the right size of random string as expected  -> %s \n", succeed)
	}
}

func Test_isDuplicateGID(t *testing.T) {
	G := Create()
	Games = append(Games, *G)
	ex := G.ID
	got := isDuplicateGID(ex)
	if got == false {
		t.Fatalf("Failed! recognise duplicate iD Got: %v  -> %s\n", got, failed)
	} else {
		t.Logf("Success could track duplicate string as expected  -> %s \n", succeed)
	}
}

//TestDelay - will test if the time delay function does actually delay for the required time
func TestDelay(t *testing.T) {
	tm := time.Now()
	sSec := tm.Unix() //starting seconds
	var fSec int64
	ex := 3
	err := Delay(ex, "s")
	if err == nil {
		fSec = tm.Unix()
	}

	diff := fSec - sSec
	got := diff

	if got != int64(0) {
		t.Fatalf("Failed! time delay is not accurate Expected: %v, Got: %v  -> %s\n", ex, fSec, failed)
	} else {
		t.Logf("Success time delay to the second is accurate as expected  -> %s \n", succeed)
	}
}

//TestTally - Tally's a finished game into the leaderboard. Pass in a sample finihsed game with expected leaderboard result for check.
func TestTally(t *testing.T) {
	var gGame = new(Game)
	input := []byte(`{"id":"UQMRURYE","name":"cool-rattler","starttime":"2021-01-26T23:08:51.270619+01:00","endtime":"2021-01-26T23:09:02.304909+01:00","roundscores":[4,8,3,2,5,9,7,7,2,3,7,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"plays":[{"playerid":"XGDGSPQY","playername":"tapir","entries":[1,3,8]},{"playerid":"NIDWCPOS","playername":"chamois","entries":[4,7,21]},{"playerid":"USMCPKIO","playername":"tuna","entries":[4,9,6]},{"playerid":"AAXGJYXZ","playername":"chicken","entries":[1,7,7]},{"playerid":"JHRGBFUN","playername":"kangaroo","entries":[1,3,9]}],"status":"Ended"}`)

	if err := json.Unmarshal(input, &gGame); err != nil {
		panic(err)
	}

	got, err := json.Marshal(gGame.Tally(gGame.ID))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(got)
	exp := `{"gameid":"UQMRURYE","gamename":"cool-rattler","gamestatus":"Ended","starttime":"2021-01-26T23:08:51.270619+01:00","endtime":"2021-01-26T23:09:02.304909+01:00","TopScorer":[{"playerid":"NIDWCPOS","playername":"chamois","entries":[4,7,21]}],"LowScorer":[{"playerid":"USMCPKIO","playername":"tuna","entries":[4,9,6]}],"Winner":[{"playerid":"NIDWCPOS","playername":"chamois","entries":[4,7,21]}]}`

	//exp := "{"gameid":"UQMRURYE","gamename":"cool-rattler","gamestatus":"Ended","starttime":"2021-01-26T23:08:51.270619+01:00","endtime":"2021-01-26T23:09:02.304909+01:00","TopScorer":[{"playerid":"NIDWCPOS","playername":"chamois","entries":[4,7,21]}],"LowScorer":[{"playerid":"USMCPKIO","playername":"tuna","entries":[4,9,6]}],"Winner":[{"playerid":"NIDWCPOS","playername":"chamois","entries":[4,7,21]}]}"
	if string(got) != exp {
		t.Fatalf("Failed! tally was not accurate and as Expected: %v, Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success tally is accurate as expected  -> %s \n", succeed)
	}
}
