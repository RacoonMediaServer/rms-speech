package speech

func (s *Service) processReadyJobs() {
	readyChan := s.workers.DoneChannel()
	for {
		select {
		case <-readyChan:
		case <-s.ctx.Done():
			return
		}
	}
}
