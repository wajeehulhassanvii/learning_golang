package main

import("math"
"math/rand"
"time"
"fmt")

func main(){
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	c := rand.Intn(10)
	b := int(math.Max(float64(a), float64(c)))

	fmt.Println(a,b,c)
}