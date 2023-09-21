package student

import "fmt"

type Student struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func (s *Student) ShowDetails() {
	fmt.Println("Nama :", s.Name)
	fmt.Println("Alamat :", s.Address)
	fmt.Println("Pekerjaan :", s.Job)
	fmt.Println("Alasan :", s.Reason)
}
