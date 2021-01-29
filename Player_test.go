package main

import (
	"reflect"
	"testing"
)

//MessageQ -
//type MessageQ struct {
//	slice []string
//}

//TestAddPlayer - Validates generated user ids
func TestAddPlayer(t *testing.T) {
	in := []string{"John", "Mary", "Jane", "James", "Willow"}
	exp := []int{8, 8, 8, 8, 8}
	got := make([]int, len(in))

	for i, v := range in {
		r := AddPlayer(v)
		got[i] = len(r)
	}

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to generate proper user id expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success AddPlayer returned with expected size of string ids -> %s \n", succeed)
	}
}

//TestenQ - Tests adding to the message queue
func Test_enQ(t *testing.T) {

	Q := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}
	exp := []int{8, 8, 9}
	got := make([]int, len(in))
	//input
	for _, v := range in {
		Q.enQ(v)
	}
	//check
	for i, x := range Q.slice {
		got[i] = len(x)
	}

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to load content properly in the queue expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success enQ loaded the expected size of string to the queue  -> %s \n", succeed)
	}
}

//TestdeQ - Tests purging content from the message queue
func Test_deQ(t *testing.T) {

	Q := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}

	//input
	for _, v := range in {
		Q.enQ(v)
	}

	for i := 0; i < 2; i++ {
		Q.deQ()
	}

	exp := 1
	got := len(Q.slice)

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to remove content properly from the queue expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success deQ removed the expected count of items from the queue  -> %s \n", succeed)
	}
}

func TestBroadcast(t *testing.T) {
	Q := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}

	//input
	for _, v := range in {
		Q.enQ(v)
		Q.Broadcast()
	}

	exp := 0
	got := len(Q.slice)

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to trigger remove of broadcast content properly from the queue expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Broadcast deQ removed the expected broadcast items count of items from the queue  -> %s \n", succeed)
	}
}

//Test_broadCaster - test the broadcaster by cheking the result in the difference between the message queue content after some time delays
func Test_broadCaster(t *testing.T) {
	//Q := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}
	exp := MQ.slice[0]
	//input
	for _, v := range in {
		MQ.enQ(v)
	}
	go broadCaster()
	//check for first items difference

	Delay(2, "s")

	//check for first items difference
	var got string
	if len(MQ.slice) > 0 {
		got = MQ.slice[0]
	}

	if got == exp {
		t.Fatalf("Failed to trigger braodcast or remove of broadcast content properly from the queue expected less than original size or less than: %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Broadcast worked and deQ removed the expected broadcast items count of items from the queue  -> %s \n", succeed)
	}
}

func Test_scoreMQbroadCaster(t *testing.T) {
	//ScoreMQ := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}
	var exp string
	var got string
	//input
	for _, v := range in {
		ScoreMQ.enQ(v)
	}
	exp = ScoreMQ.slice[0]

	//broadcast messages
	go scoreMQbroadCaster()
	Delay(2, "s")

	if len(ScoreMQ.slice) > 0 {
		got = ScoreMQ.slice[0]
	}

	if got == exp {
		t.Fatalf("Failed to trigger braodcast or remove of broadcast content properly from the ScoreMQ msg queue expected less than original size or less than: %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Broadcast worked and deQ removed the expected broadcast items count of items from the ScoreMQ queue  -> %s \n", succeed)
	}
}

func Test_leadBoardMQbroadCaster(t *testing.T) {
	//ScoreMQ := new(MessageQ)
	in := []string{"Message1", "Message2", "Messages3"}
	var exp string
	var got string
	//input
	for _, v := range in {
		LeadBoardMQ.enQ(v)
	}
	exp = LeadBoardMQ.slice[0]

	//broadcast messages
	go leadBoardMQbroadCaster()
	Delay(2, "s")

	if len(LeadBoardMQ.slice) > 0 {
		got = LeadBoardMQ.slice[0]
	}

	if got == exp {
		t.Fatalf("Failed to trigger braodcast or remove of broadcast content properly from the LeadBoardMQ msg queue expected less than original size or less than: %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success leadBoardMQbroadCaster worked and deQ removed the expected broadcast items count of items from the LeadBoardMQ queue  -> %s \n", succeed)
	}
}

//TestAddPlay -
func TestAddPlay(t *testing.T) {
	//PlayerID, PlayerName string, entries [3]int
	ol := len(Plays)
	var entries [3]int
	entries[0] = 1
	entries[1] = 5
	entries[2] = 0

	PlayerName := "John"
	PlayerID := "WERTFDSE"

	exp := 1
	AddPlay(PlayerID, PlayerName, entries)
	got := len(Plays) - ol

	if reflect.DeepEqual(exp, got) == false {
		t.Fatalf("Failed to add a user play properly expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success AddPlay added the expected user play items count of Plays matches  -> %s \n", succeed)
	}
}

//TestSetID -
func TestSetID(t *testing.T) {
	G := Create()
	exp := G.ID
	G.SetID()

	got := G.ID

	if reflect.DeepEqual(exp, got) == true {
		t.Fatalf("Failed to game ID expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Set new game ID as expected  -> %s \n", succeed)
	}
}

//TestSetPetName
func TestSetPetName(t *testing.T) {
	G := Create()
	exp := G.Name
	G.SetPetName()

	got := G.Name

	if reflect.DeepEqual(exp, got) == true {
		t.Fatalf("Failed to game name expected : %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Set new game name as expected  -> %s \n", succeed)
	}
}

//TestWrap -
func TestWrap(t *testing.T) {
	msg := new(Message)
	exp := "{\"channel\":\"LiveScoreMQ\",\"body\":\"Starting Round 15\n\"}"
	//exp := string({"channel":"LiveScoreMQ","body":"Starting Round 15\n"})
	got := msg.Wrap("LiveScoreMQ", "Starting Round 15")

	if got == exp {
		t.Fatalf("Failed to format message for json queue channel %v Got: %v  -> %s\n", exp, got, failed)
	} else {
		t.Logf("Success Wrap formated the message in Json format as expected  -> %s \n", succeed)
	}
}
