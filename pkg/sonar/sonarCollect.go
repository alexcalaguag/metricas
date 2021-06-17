package sonar

import (
	"sync"
	"log"
	"time"
	projects "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/projects"
    model "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/model"
	storage "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/storage"
)

func CollectMetric() {
	start := time.Now()
	defer func() {
		log.Println("Execution Time: ", time.Since(start))
	}()
	projectsNeeded := projects.ReadingFile()
	log.Println("Lista de Projetos: ", projectsNeeded)
	sonarMap := make(map[string]*model.SonarMetric)
	sonarMetric := &model.SonarMetric{}
	wg := sync.WaitGroup{}
	
		for _, id := range projectsNeeded {
			wg.Add(1)
			go func(project string) {
				metric, err := sonarMetric.Get(project)
				if err != nil {
					return
				}
				sonarMap[project] = metric
				storage.SaveMetricSonar(sonarMap[project])
				log.Println("Fetched project : ", project)
				wg.Done()
			}(id)
		}
	wg.Wait()
}
