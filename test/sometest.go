package test
import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
)

func TestSort(){
	tests := []struct {
		Query string
		Err   bool
	}{
		{`name~"Sam"`, false},
		{`name=="Sam"`, false},
		{`subject=="Science"`, true},
		{`subject=="English"`, false},
		{`name=="Hi"`, false},
		{`name!="Hellow"`, false},
		{`subject=="SocialScience"`, true},
		{`subject=="VinglishEnglish"`, false},
	}

	for _, test := range tests {
		s, err := query.ParseSorting(test.Query)
		if err != nil {
			fmt.Println("Invalid sorting data '%s'", test.Query)
		}
		err = TestValidateSorting("/test.TestService/List", s)
		if err != nil {
			if test.Err == false {
				fmt.Println("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				fmt.Println("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}

func TestFiltering() {
	tests := []struct {
		Query string
		Err   bool
	}{
		{`name~"Sam"`, false},
		{`name=="Sam"`, false},
		{`subject=="Science"`, true},
		{`subject=="English"`, false},
		{`name=="Hi"`, false},
		{`name!="Hellow"`, false},
		{`subject=="SocialScience"`, true},
		{`subject=="VinglishEnglish"`, false},
	}

	for _, test := range tests {
		f, err := query.ParseFiltering(test.Query)
		if err != nil {
			fmt.Println("Invalid filtering data '%s'", test.Query)
		}
		err = TestValidateFiltering("/test.TestService/List", f)
		if err != nil {
			if test.Err == false {
				fmt.Println("Unexpected error for %s query: %s", test.Query, err)
			}
		} else {
			if test.Err == true {
				fmt.Println("Expected error for %s query, but got no error", test.Query)
			}
		}
	}
}
