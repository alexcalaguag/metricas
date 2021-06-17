package storage

import (
	model "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "metric_user"
	password = "metric_password"
	dbname   = "metrics"
)


func SaveMetricSonar(metric *model.SonarMetric) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	sqlStatement := `
	INSERT INTO metrics_sonar (project, metric, value)
	VALUES ($1, $2, $3)
	RETURNING id`

	for key, element := range metric.Component.Measures {

		id := 0
		err = db.QueryRow(sqlStatement, metric.Component.Key, element.Metric, element.Value).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("New record ID is:", id)

		fmt.Println("Key:", key, "=>", "Element:", element.Metric)
		fmt.Println("Key:", key, "=>", "Element:", element.Value)
	}

}
