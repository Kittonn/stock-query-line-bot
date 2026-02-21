package usecases

import (
	"context"
	"log"
	"sync"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type LineWebhookWorker struct {
	queue         chan *domain.LineEvent
	lineWebhookUC ports.LineWebhookUsecase
	wg            sync.WaitGroup
}

func NewLineWebhookWorker(lineWebhookUC ports.LineWebhookUsecase) ports.LineWebhookWorker {
	return &LineWebhookWorker{
		queue:         make(chan *domain.LineEvent, 100),
		lineWebhookUC: lineWebhookUC,
	}
}

func (l *LineWebhookWorker) Enqueue(event *domain.LineEvent) {
	select {
	case l.queue <- event:
	default:
		log.Printf("[WARNING] Line webhook event queue is full! Dropping event: %s", event.Type)
	}
}

func (l *LineWebhookWorker) Start(ctx context.Context, workerCount int) {
	for i := 0; i < workerCount; i++ {
		l.wg.Add(1)
		go l.run(i)
	}
}

func (l *LineWebhookWorker) run(workerID int) {
	defer l.wg.Done()

	for event := range l.queue {
		ctx := context.Background()
		l.handle(ctx, workerID, event)
	}
}

func (l *LineWebhookWorker) handle(ctx context.Context, workerID int, event *domain.LineEvent) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	log.Printf("[worker %d] handling %s", workerID, event.Type)
	l.lineWebhookUC.HandleEvent(ctx, *event)
}

func (l *LineWebhookWorker) Stop() {
	close(l.queue)
	l.wg.Wait()
}
