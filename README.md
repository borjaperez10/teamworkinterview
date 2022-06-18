# TeamworkInterview

  Package customerimporter reads from the given customers.csv file and returns a sorted (data structure of your choice) of email domains along with the number of customers with e-mail addresses for each domain.  Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could* be 1m lines or run on a small machine).
  

## Getting Started

### Dependencies

* This project has been done with Golang 1.16.2.

### Installing
Download source code using:
```
git clone https://github.com/borjaperez10/teamworkinterview.git

```
### Executing program

In order to run, use:
```
go run main.go

```
### Testing program

Some of the functions have been tested. In total, 72.5% of statements have been covered
```
 go test -timeout 30s -run ^TestReadDomains$ teamworkinterview/customerimporter
 go test -timeout 30s -run ^TestGetDomain$ teamworkinterview/customerimporter
 go test -timeout 30s -run ^TestGetDomainListInfo$ teamworkinterview/customerimporter
 go test -timeout 30s -run ^TestReadCSV$ teamworkinterview/customerimporter
 go test -timeout 30s -run ^TestSortDomainList$ teamworkinterview/customerimporter

```
## Author

Borja Perez
