// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func OpenAndReadCSV(filename string) ([][]string, error) {
	/*
			    This method reads the CSV file and returns a matrix containing all the CSV info
		        Input:
		            filename: string containing the path+filename of the CSV file
		        Output:
		            rows: contains all the rows info of the csv file
	*/
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	rows, err := ReadCSV(file)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ReadCSV(file io.Reader) ([][]string, error) {

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func ReadDomains(rows [][]string) []string {
	/*
			    This method receives all the CSV info and returns a slice containing all the domains of the file (including duplicates)
		        Input:
		            rows: matrix containing all the CSV info
		        Output:
		            allDomains: slice containing all the valid domains extracted
	*/
	var allDomains []string
	for _, row := range rows {
		email := row[2]
		domain, err := GetDomain(email)
		if err == nil {
			allDomains = append(allDomains, domain)
		} else {
			log.Print("Error parsing ", email, ": ", err)
		}
	}
	return allDomains
}
func GetDomain(email string) (string, error) {
	/*
			    This method receives an email returns its domain
		        Input:
		            email: string containing a specific email
		        Output:
		            parts[1]: part of the email which contains the domain
	*/
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1], nil
	} else {
		return "", errors.New("Incorrect email readed. It will not be added to the domain list")
	}
}

func GetDomainListInfo(domains []string) map[string]int {
	/*
			    This method is used to provide a dictionary containing the domain names and number of clients from input domains list
			    Input:
		            domains: contains the list of all the domains (duplicated)
		        Output:
		            dict: dictionary containing domain name(string) as key values and number of clients (integer) as values
	*/
	dict := make(map[string]int)
	for _, dom := range domains {
		dict[dom] = dict[dom] + 1
	}
	return dict
}
func SortDomainList(list map[string]int, sortMethod int) []string {
	/*
			    This method is used to Sort the domains and customers in with a specific sort order
			    Input:
		            list: dictionary containing domain name(string) as key values and number of clients (integer) as values
		            sortMethod: Ascending/Descending sort: order by Domain Name/Number of customers
	*/
	keys := make([]string, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	switch sortMethod {
	case 1:
		sort.Strings(keys)
	case 2:
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	case 3:
		sort.Slice(keys, func(i, j int) bool {
			return list[keys[i]] < list[keys[j]]
		})
	case 4:
		sort.Slice(keys, func(i, j int) bool {
			return list[keys[i]] > list[keys[j]]
		})
	}

	return keys
}
