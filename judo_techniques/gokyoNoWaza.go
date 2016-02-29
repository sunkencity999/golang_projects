package main

import (
	"fmt"
	"golang_projects/judo_techniques/clear"
)

func main() {
	//this application is a searchable library of Judo Techniques

	//multi-dimensional slice that contains the "sets" of techniques.
	gokyo := make([][]string, 100)

	//initializing variables for slices to represent the "sets"
	daiIkkyo := make([]string, 8, 20)
	daiNikyo := make([]string, 8, 20)
	daiSankyo := make([]string, 8, 20)
	daiYonkyo := make([]string, 8, 20)
	daiGokyo := make([]string, 8, 20)
	shinMesho := make([]string, 27, 50)

	//adding values at appropriate locations within the sets.

	//throws for Dai Ikkyo
	daiIkkyo[0] = "De Ashi Harai \t Advanced Foot Sweep \n"
	daiIkkyo[1] = "Hiza Guruma \t Knee Whirl \n"
	daiIkkyo[2] = "Sasae Tsuri Komi Ashi \t Supporting Lift-Pull Foot \n"
	daiIkkyo[3] = "Uki Goshi \t Floating Hip \n"
	daiIkkyo[4] = "Osoto Gari \t Major Outer Reap \n"
	daiIkkyo[5] = "Ogoshi \t Major Hip \n"
	daiIkkyo[6] = "Ouchi Gari \t Major Inner Reap \n"
	daiIkkyo[7] = "Morote Seoi Nage \t Two Hand Shoulder Throw \n"

	//throws for Dai Nikkyo
	daiNikyo[0] = "Kosoto Gari \t Minor Outer Reap\n"
	daiNikyo[1] = "Kouchi Gari \t Minor Inner Reap\n"
	daiNikyo[2] = "Koshi Guruma \t Hip Whirl\n"
	daiNikyo[3] = "Tsuri Komi Goshi \t Lift Pull Hip\n"
	daiNikyo[4] = "Okuri Ashi Harai \t Assisting Foot Sweep\n"
	daiNikyo[5] = "Tai Otoshi \t Body Drop\n"
	daiNikyo[6] = "Harai Goshi \t Sweeping Hip\n"
	daiNikyo[7] = "Uchi Mata \t Inner Thigh\n"

	//throws for Dai Sankyo
	daiSankyo[0] = "Kosoto Gake \t Minor Outer Hook\n"
	daiSankyo[1] = "Tsuri Goshi \t Lifting Hip\n"
	daiSankyo[2] = "Yoko Otoshi \t Side Drop\n"
	daiSankyo[3] = "Ashi Guruma \t Foot Whirl\n"
	daiSankyo[4] = "Hane Goshi \t Spring Hip\n"
	daiSankyo[5] = "Harai Tsuri Komi Ashi \t Sweeping Lift Pull Foot\n"
	daiSankyo[6] = "Tomoe Nage \t Circle Throw\n"
	daiSankyo[7] = "Kata Guruma \t Shoulder Whirl\n"

	//throws for Dai Yonkyo
	daiYonkyo[0] = "Sumi Gaeshi \t Corner Reversal\n"
	daiYonkyo[1] = "Tani Otoshi \t Valley Drop\n"
	daiYonkyo[2] = "Hane Maki Komi \t Springing Body Wrap\n"
	daiYonkyo[3] = "Sukui Nage \t Scooping Throw\n"
	daiYonkyo[4] = "Utsuri Goshi \t Transfer Hip\n"
	daiYonkyo[5] = "O Guruma \t Major Whirl\n"
	daiYonkyo[6] = "Soto Maki Komi \t Outside Body Wrap\n"
	daiYonkyo[7] = "Uki Otoshi \t Floating Drop\n"

	//throws for Dai Gokyo
	daiGokyo[0] = "Osoto Guruma \t Major Outer Whirl\n"
	daiGokyo[1] = "Uki Waza \t Floating Technique\n"
	daiGokyo[2] = "Yoko Wakare \t Side Separation\n"
	daiGokyo[3] = "Yoko Guruma \t Side Whirl\n"
	daiGokyo[4] = "Ushiro Goshi \t Rear Hip\n"
	daiGokyo[5] = "Ura Nage \t Back Throw\n"
	daiGokyo[6] = "Sumi Otoshi \t Corner Drop\n"
	daiGokyo[7] = "Yoko Gake \t Side Dash\n"

	//throws from the Shinmesho-no-waza (new techniques)
	shinMesho[0] = "Morote Gari \t Two Hand Reap\n"
	shinMesho[1] = "Kuchiki Taoshi \t Dead Tree Drop\n"
	shinMesho[2] = "Kibisu Gaeshi \t Heel Trip\n"
	shinMesho[3] = "Uchi Mata Sukashi \t Inner Thigh Avoidance\n"
	shinMesho[4] = "Kouchi Gaeshi \t Minor Inner Reap Counter\n"
	shinMesho[5] = "Daki Age \t High Lift\n"
	shinMesho[6] = "Tsubame Gaeshi \t Minor Inner Reap Counter\n"
	shinMesho[7] = "Osoto Gaeshi \t Major Outer Counter\n"
	shinMesho[8] = "Ouchi Gaeshi \t Major Inner Counter\n"
	shinMesho[9] = "Hane Goshi Gaeshi \t Springing Hip Counter\n"
	shinMesho[10] = "Harai Goshi Gaeshi \t Sweeping Hip Counter\n"
	shinMesho[11] = "Uchi Mata Gaeshi \t Inner Thigh Counter\n"
	shinMesho[12] = "Kani Basami \t Crab Claw\n"
	shinMesho[13] = "Kawazu Gake \t One-leg Entanglement\n"
	shinMesho[14] = "Osoto Maki Komi \t Major Outer Wrap Around\n"
	shinMesho[15] = "Uchi Mata Maki Komi \t Inner Thigh Wrap Around\n"
	shinMesho[16] = "Harai Maki Komi \t Sweeping Wrap Around\n"
	shinMesho[17] = "Ippon Seoi Nage \t One Arm Shoulder Throw\n"
	shinMesho[18] = "Sode Tsuri Komi Goshi \t Sleeve Lift-Pull Hip\n"
	shinMesho[19] = "Obi Otoshi \t Belt Drop\n"
	shinMesho[20] = "Daki Wakare \t High Seperation\n"
	shinMesho[21] = "Hikikomi Gaeshi \t Back Fall Reversal\n"
	shinMesho[22] = "Osoto Otoshi \t Major Outer Drop\n"
	shinMesho[23] = "Tawara Gaeshi \t Rice Bag Reversal\n"
	shinMesho[24] = "Uchi Maki Komi \t Inner Wrap Around\n"
	shinMesho[25] = "Seoi Otoshi \t Shoulder Drop\n"
	shinMesho[25] = "Yama Arashi \t Mountain Storm\n"

	gokyo[0] = daiIkkyo
	gokyo[1] = daiNikyo
	gokyo[2] = daiSankyo
	gokyo[3] = daiYonkyo
	gokyo[4] = daiGokyo
	gokyo[5] = shinMesho

	//assign variables to reference slices within multi-dimensional slice
	daiIkkyoResult := gokyo[0]
	daiNikyoResult := gokyo[1]
	daiSankyoResult := gokyo[2]
	daiYonkyoResult := gokyo[3]
	daiGokyoResult := gokyo[4]
	shinMeshoResult := gokyo[5]

	//prompt for user input and store it, build menu, return values
	var usrInput int
	fmt.Print("Please select the list from the Gokyo No Waza you wish to see: \n 1.Dai Ikkyo\n 2.Dai Nikyo\n 3.Dai Sankyo\n 4.Dai Yonkyo\n 5.Dai Gokyo\n 6.Shin Meshyo\n")
	fmt.Scan(&usrInput)

	inpValue := usrInput
	i := inpValue
	//
	for i = 0; i < 7; i++ {

		if inpValue == 0 {
			fmt.Println(daiIkkyoResult)
		}
		if inpValue == 1 {
			fmt.Println(daiNikyoResult)
		}
		if inpValue == 2 {
			fmt.Println(daiSankyoResult)
		}
		if inpValue == 3 {
			fmt.Println(daiYonkyoResult)
		}
		if inpValue == 4 {
			fmt.Println(daiGokyoResult)
		}
		if inpValue == 5 {
			fmt.Println(shinMeshoResult)
		}
		fmt.Print("Please select the list from the Gokyo No Waza you wish to see: \n 1.Dai Ikkyo\n 2.Dai Nikyo\n 3.Dai Sankyo\n 4.Dai Yonkyo\n 5.Dai Gokyo\n 6.Shin Meshyo\n")
		fmt.Scan(&usrInput)
		clear.CallClear()
		continue

	}

}
