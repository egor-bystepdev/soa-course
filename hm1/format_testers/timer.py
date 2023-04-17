import time

class TimerNs:
    def __init__(self):
        self.tm = time.perf_counter_ns()
    
    def start(self):
        self.tm = time.perf_counter_ns()
    
    def finish(self):
        return time.perf_counter_ns() - self.tm
    
