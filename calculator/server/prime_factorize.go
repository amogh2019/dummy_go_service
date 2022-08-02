package main

import (
	"fmt"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

type Seive struct {
	isPrime []bool
	len     int64
}

var sv = Seive{
	isPrime: []bool{false, false, true, true},
	len:     3,
}

func (sv *Seive) IsInputPrime(val int64) bool {
	if val <= sv.len {
		return sv.isPrime[val]
	}

	toBeAppened := make([]bool, val-sv.len)
	for i := range toBeAppened {
		toBeAppened[i] = true
	}
	sv.isPrime = append(sv.isPrime, toBeAppened...)
	sv.len += int64(len(toBeAppened))

	for i := int64(2); (i * i) <= val; i++ {

		if sv.isPrime[i] {
			for j := 2 * i; j <= val; j += i {
				sv.isPrime[j] = false
			}
		}
	}
	return sv.isPrime[val]

}

func (sv *Seive) GetPrimeFactors(val int64) (factor int64, remaining int64) {

	if sv.IsInputPrime(val) {
		fmt.Println(val, " = ", val, " * ", 1)
		return val, 1
	}
	for i := int64(2); i*i <= val; i++ {
		if sv.IsInputPrime(i) && (val%i == 0) {
			fmt.Println(val, " = ", i, " * ", val/i)
			return i, val / i
		}
	}
	return 1, 1
}

func (serverImpl *CalculatorServiceServerImpl) PrimeFactorize(request *pb.PrimeFactorizationRequeset, stream pb.CalculatorService_PrimeFactorizeServer) error {

	fmt.Println("operation received to factorize : ", request)

	rem := int64(request.Number)
	for {
		factor, remaining := sv.GetPrimeFactors(rem)
		stream.Send(&pb.PrimeFactorizationResponse{
			Factor: int32(factor),
		})
		if remaining == 1 {
			break
		}
		rem = remaining

	}

	return nil

}
