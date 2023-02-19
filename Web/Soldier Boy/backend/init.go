package main

import "fmt"


func createFunction() {
	db := GetDb()
	q := `
	CREATE FUNCTION validate(flag VARCHAR(255)) 
	RETURNS INTEGER AS $$
	BEGIN 
	  IF flag LIKE 'Securinets{%}' THEN 
		RETURN 1; 
	  ELSE 
	  	RAISE EXCEPTION 'Flag should start with Securinets{ and end with }';
	  END IF; 
	END; $$
	LANGUAGE plpgsql;
	`
	db.Exec(q)
}


func makeDBReadOnly() {
	db := GetDb()
	db.Exec("ALTER DATABASE previet SET default_transaction_read_only TO on;")
}


func Check() {
	db := GetDb()


	// task := Task{}
// h4tte_SturmFront_s3in_k0nneN
	// db.Where("id = ?", 1).First(&task)
	// fmt.Println(task)

	c := ""
	db.Raw("select validate('Securinets{}') union select case when substr(flag,1,1) = 'q' then 1 else 0 end from tasks -- -").Scan(&c)
	fmt.Println(c)
}

func Init() {

	db := GetDb()

	db.Exec("DROP TABLE tasks")
	db.Exec("DROP FUNCTION validate")

	db.AutoMigrate(&Task{})

	db.Create(&Task{
		Title: "NginXO", 
		Description: "Can you find my secret XO route? flag format: Securinets{[a-zA-Z0-9_]+}", 
		Url: "https://nginxo.vercel.app", 
		Flag: "Securinets{h4tte_SturmFront_s3in_k0nneN}",
	})



	createFunction()

	makeDBReadOnly()
}