package main

func main() {

}
func cashback(amount int)int {
	const bound =3000
	const persent =5
	result:=0

	if amount >= bound  {
		result = amount * persent/100


	}
	return result
}