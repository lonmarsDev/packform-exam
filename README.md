# Exam backend

Exam backend is writin in golang, it serve api and import data from csv to mongodb and postgres database.

## Requiments
golang version 1.13
mongodb
postgres Sql

## Installation

```bash
git clone git@github.com:lonmarsDev/packform-exam.git
```

How To change database config file
config files are located on 
db importer - packform-exam/db-importer/config.json
service API- packform-exam/service-api/config.json


## Usage

How import data from csv
```bash
git clone git@github.com:lonmarsDev/packform-exam.git
cd packform-exam/db-importer
go run *.go
```
All csv file are located to packform-exam/db-importer/test_data

How to run api services
```bash
cd packform-exam/service-api
go run *.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
