package main

import (
	"log"
	"sync"

	"github.com/wolfmagnate/daily_report/collector/infra"
	"github.com/wolfmagnate/daily_report/collector/integration"
)

type WorkerPool struct {
	numWorkers  int
	jobQueue    chan Job
	lockManager *InMemoryLock
	wg          *sync.WaitGroup

	txProvider infra.TxProvider

	githubIntegrator integration.GithubIntegrator
	slackIntegrator  integration.SlackIntegrator

	githubActivityRepo infra.GithubActivityRepository
	slackActivityRepo  infra.SlackActivityRepository
}

func NewWorkerPool(numWorkers int, queueCapacity int, jobQueue chan Job, lockManager *InMemoryLock, txProvider infra.TxProvider, githubIntegrator integration.GithubIntegrator, slackIntegrator integration.SlackIntegrator, githubActivityRepo infra.GithubActivityRepository, slackActivityRepo infra.SlackActivityRepository) *WorkerPool {
	return &WorkerPool{
		numWorkers:         numWorkers,
		jobQueue:           jobQueue,
		lockManager:        lockManager,
		wg:                 new(sync.WaitGroup),
		txProvider:         txProvider,
		githubIntegrator:   githubIntegrator,
		slackIntegrator:    slackIntegrator,
		githubActivityRepo: githubActivityRepo,
		slackActivityRepo:  slackActivityRepo,
	}
}

func (p *WorkerPool) Run() {
	log.Printf("[WorkerPool] %d人のワーカーでプールを開始します。", p.numWorkers)
	for i := 1; i <= p.numWorkers; i++ {
		worker := NewWorker(i, p.jobQueue, p.lockManager, p.wg, p.txProvider, p.githubIntegrator, p.slackIntegrator, p.githubActivityRepo, p.slackActivityRepo)
		go worker.start()
	}
}

func (p *WorkerPool) Stop() {
	log.Println("[WorkerPool] プールの停止処理を開始します...")
	close(p.jobQueue)
	p.wg.Wait()
	log.Println("[WorkerPool] 全てのワーカーが停止しました。")
}
