package main

import "context"
import "sync"
import "time"
import "Spacetraders/src/Task"

//
// <-----------------------------> Task Controller <------------------------------>
//

type TaskController struct {
	mu           sync.Mutex
	manualActive bool
	cond         *sync.Cond
}

func NewTaskController() *TaskController {
	tc := &TaskController{}
	tc.cond = sync.NewCond(&tc.mu)
	return tc
}

func (tc *TaskController) RunManualTask(fn func() error) error {
	tc.mu.Lock()
	tc.manualActive = true
	tc.mu.Unlock()

	defer func() {
		tc.mu.Lock()
		tc.manualActive = false
		tc.cond.Broadcast()
		tc.mu.Unlock()
	}()

	return fn()
}

func (tc *TaskController) WaitIfManualActive() {
	tc.mu.Lock()
	for tc.manualActive {
		tc.cond.Wait()
	}
	tc.mu.Unlock()
}

//
// <--------------------------------> Scheduler <--------------------------------->
//

type Job struct {
	Name     string
	Interval time.Duration
	Function func(context.Context) error
	nextRun  time.Time
	running  bool
}

type Scheduler struct {
	tc   *TaskController
	jobs []*Job
}

func NewScheduler(tc *TaskController) *Scheduler {
	return &Scheduler{tc: tc}
}

func (s *Scheduler) Register(job *Job) {
	job.nextRun = time.Now().Add(job.Interval)
	s.jobs = append(s.jobs, job)
}

func (s *Scheduler) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			for _, job := range s.jobs {
				if now.After(job.nextRun) && !job.running {
					job.running = true
					job.nextRun = now.Add(job.Interval)
					go func(j *Job) {
						defer func() { j.running = false }()
						s.tc.WaitIfManualActive()
						if err := j.Function(ctx); err != nil {
							Task.LogErr("Job " + j.Name + " failed: " + err.Error())
						}
					}(job)
				}
			}
		}
	}
}

//
// <-----------------------------------> Main <----------------------------------->
//

func main() {
	// TASK Check DB
	if err := Task.CheckDB(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	// TASK Init DB
	if err := Task.DbLite(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	tc := NewTaskController()
	scheduler := NewScheduler(tc)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//
	// <------------------------------> Schedule Jobs <------------------------------->
	//
	
	scheduler.Register(&Job{
		Name:     "sync-server-state",
		Interval: 30 * time.Second,
		Function: func(ctx context.Context) error {
			Task.LogActivity("Running scheduled server sync task")
			Task.UpdateGameServerState()
			return nil 
		},
	})

	scheduler.Register(&Job{
		Name:     "sync-agent-state",
		Interval: 30 * time.Second,
		Function: func(ctx context.Context) error {
			Task.LogActivity("Running scheduled agent sync task")
			Task.UpdateAgentState(nil)
			return nil
		},
	})

	scheduler.Register(&Job{
		Name:     "sync-fleet-state",
		Interval: 30 * time.Second,
		Function: func(ctx context.Context) error {
			Task.LogActivity("Running scheduled fleet sync task")
			Task.UpdateShipState(nil)
			return nil
		},
	})

	//
	// <------------------------------> Startup Tasks <------------------------------->
	//

	// Manual task at startup (blocks scheduler)
	if err := tc.RunManualTask(func() error {
		CFG, _ := Task.GetConfig()

		if CFG.API.AgentToken == "" {
			Task.LogActivity("No agent token found, registering a new agent")
			if err := Task.RegisterNewAgent(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	//
	// <-----------------------------> Start Scheduler <------------------------------>
	//

	go scheduler.Run(ctx)
	Task.LogActivity("Service started, scheduler running")

	select {}
}

