package main

import "fmt"

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
	daiIkkyo[0] = "De Ashi Harai \t Advanced Foot Sweep"
	daiIkkyo[1] = "Hiza Guruma \t Knee Whirl"
	daiIkkyo[2] = "Sasae Tsuri Komi Ashi \t Supporting Lift-Pull Foot"
	daiIkkyo[3] = "Uki Goshi \t Floating Hip"
	daiIkkyo[4] = "Osoto Gari \t Major Outer Reap"
	daiIkkyo[5] = "Ogoshi \t Major Hip"
	daiIkkyo[6] = "Ouchi Gari \t Major Inner Reap"
	daiIkkyo[7] = "Morote Seoi Nage \t Two Hand Shoulder Throw"

	//throws for Dai Nikkyo
	daiNikyo[0] = "Kosoto Gari \t Minor Outer Reap"
	daiNikyo[1] = "Kouchi Gari \t Minor Inner Reap"
	daiNikyo[2] = "Koshi Guruma \t Hip Whirl"
	daiNikyo[3] = "Tsuri Komi Goshi \t Lift Pull Hip"
	daiNikyo[4] = "Okuri Ashi Harai \t Assisting Foot Sweep"
	daiNikyo[5] = "Tai Otoshi \t Body Drop"
	daiNikyo[6] = "Harai Goshi \t Sweeping Hip"
	daiNikyo[7] = "Uchi Mata \t Inner Thigh"

	//throws for Dai Sankyo
	daiSankyo[0] = "Kosoto Gake \t Minor Outer Hook"
	daiSankyo[1] = "Tsuri Goshi \t Lifting Hip"
	daiSankyo[2] = "Yoko Otoshi \t Side Drop"
	daiSankyo[3] = "Ashi Guruma \t Foot Whirl"
	daiSankyo[4] = "Hane Goshi \t Spring Hip"
	daiSankyo[5] = "Harai Tsuri Komi Ashi \t Sweeping Lift Pull Foot"
	daiSankyo[6] = "Tomoe Nage \t Circle Throw"
	daiSankyo[7] = "Kata Guruma \t Shoulder Whirl"

	//throws for Dai Yonkyo
	daiYonkyo[0] = "Sumi Gaeshi \t Corner Reversal"
	daiYonkyo[1] = "Tani Otoshi \t Valley Drop"
	daiYonkyo[2] = "Hane Maki Komi \t Springing Body Wrap"
	daiYonkyo[3] = "Sukui Nage \t Scooping Throw"
	daiYonkyo[4] = "Utsuri Goshi \t Transfer Hip"
	daiYonkyo[5] = "O Guruma \t Major Whirl"
	daiYonkyo[6] = "Soto Maki Komi \t Outside Body Wrap"
	daiYonkyo[7] = "Uki Otoshi \t Floating Drop"

	//throws for Dai Gokyo
	daiGokyo[0] = "Osoto Guruma \t Major Outer Whirl"
	daiGokyo[1] = "Uki Waza \t Floating Technique"
	daiGokyo[2] = "Yoko Wakare \t Side Separation"
	daiGokyo[3] = "Yoko Guruma \t Side Whirl"
	daiGokyo[4] = "Ushiro Goshi \t Rear Hip"
	daiGokyo[5] = "Ura Nage \t Back Throw"
	daiGokyo[6] = "Sumi Otoshi \t Corner Drop"
	daiGokyo[7] = "Yoko Gake \t Side Dash"

	//throws from the Shinmesho-no-waza (new techniques)
	shinMesho[0] = "Morote Gari \t Two Hand Reap"
	shinMesho[1] = "Kuchiki Taoshi \t Dead Tree Drop"
	shinMesho[2] = "Kibisu Gaeshi \t Heel Trip"
	shinMesho[3] = "Uchi Mata Sukashi \t Inner Thigh Avoidance"
	shinMesho[4] = "Kouchi Gaeshi \t Minor Inner Reap Counter"
	shinMesho[5] = "Daki Age \t High Lift"
	shinMesho[6] = "Tsubame Gaeshi \t Minor Inner Reap Counter"
	shinMesho[7] = "Osoto Gaeshi \t Major Outer Counter"
	shinMesho[8] = "Ouchi Gaeshi \t Major Inner Counter"
	shinMesho[9] = "Hane Goshi Gaeshi \t Springing Hip Counter"
	shinMesho[10] = "Harai Goshi Gaeshi \t Sweeping Hip Counter"
	shinMesho[11] = "Uchi Mata Gaeshi \t Inner Thigh Counter"
	shinMesho[12] = "Kani Basami \t Crab Claw"
	shinMesho[13] = "Kawazu Gake \t One-leg Entanglement"
	shinMesho[14] = "Osoto Maki Komi \t Major Outer Wrap Around"
	shinMesho[15] = "Uchi Mata Maki Komi \t Inner Thigh Wrap Around"
	shinMesho[16] = "Harai Maki Komi \t Sweeping Wrap Around"
	shinMesho[17] = "Ippon Seoi Nage \t One Arm Shoulder Throw"
	shinMesho[18] = "Sode Tsuri Komi Goshi \t Sleeve Lift-Pull Hip"
	shinMesho[19] = "Obi Otoshi \t Belt Drop"
	shinMesho[20] = "Daki Wakare \t High Seperation"
	shinMesho[21] = "Hikikomi Gaeshi \t Back Fall Reversal"
	shinMesho[22] = "Osoto Otoshi \t Major Outer Drop"
	shinMesho[23] = "Tawara Gaeshi \t Rice Bag Reversal"
	shinMesho[24] = "Uchi Maki Komi \t Inner Wrap Around"
	shinMesho[25] = "Seoi Otoshi \t Shoulder Drop"
	shinMesho[25] = "Yama Arashi \t Mountain Storm"

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

	//print entry to test
	fmt.Println(resultPrint[2])
}
