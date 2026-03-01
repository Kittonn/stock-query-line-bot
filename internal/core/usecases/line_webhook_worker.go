package usecases

import (
	"context"
	"sync"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/Kittonn/stock-query-line-bot/pkg/logger"
)

type LineWebhookWorker struct {
	queue         chan *domain.LineEvent
	lineWebhookUC ports.LineWebhookUsecase
	wg            sync.WaitGroup
	once          sync.Once
	log           logger.Logger
}

func NewLineWebhookWorker(lineWebhookUC ports.LineWebhookUsecase, log logger.Logger) ports.LineWebhookWorker {
	return &LineWebhookWorker{
		queue:         make(chan *domain.LineEvent, 100),
		lineWebhookUC: lineWebhookUC,
		log:           log,
	}
}

func (l *LineWebhookWorker) Enqueue(event *domain.LineEvent) {
	select {
	case l.queue <- event:
	default:
		l.log.Warn("Line webhook event queue is full! Dropping event: ", event.Type)
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
			l.log.Error("recovered from panic workerID: ", workerID, " panic: ", r)
		}
	}()

	l.log.Info("worker handling event workerID: ", workerID, " type: ", event.Type)
	l.lineWebhookUC.HandleEvent(ctx, event)
}

func (l *LineWebhookWorker) Stop() {
	l.once.Do(func() {
		close(l.queue)
	})

	l.wg.Wait()
}
