package main

import (
	"fmt"
	"log"
	"teamworkinterview/customerimporter"
)

func main() {
	var opt int
MENU:
	fmt.Println("File customers.csv will be loaded. In which order you want to sort?")
	fmt.Println("   1: Sort by Domain name (Ascending order)")
	fmt.Println("   2: Sort by Domain name (Descending order)")
	fmt.Println("   3: Sort by number of Customers in each domain (Ascending order)")
	fmt.Println("   4: Sort by number of Customers in each domain (Descending order)")
	fmt.Println("Choose an option (1-4):")

	if _, err := fmt.Scanln(&opt); err != nil {
		log.Println("Error scanning line: ", err)
		goto MENU
	}
	if opt < 1 || opt > 4 {
		log.Println("Error: Introduce an integer between 1-4")
		goto MENU
	}

	lines, err := customerimporter.OpenAndReadCSV("customers.csv")
	if err != nil {
		panic(err)
	}

	domains := customerimporter.ReadDomains(lines)
	domainsDict := customerimporter.GetDomainListInfo(domains)
	sortedDict := customerimporter.SortDomainList(domainsDict, opt)

	for _, key := range sortedDict {
		fmt.Println("Domain", key, "contains", domainsDict[key], "customers")
	}
}
