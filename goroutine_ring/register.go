package main

type register struct {
	workers map[int]worker
}

var r = register{
	workers: make(map[int]worker),
}

func (r *register) getWorker(id int) worker {
	value, _ := r.workers[id]

	return value
}

func (r *register) exists(id int) bool {
	_, ok := r.workers[id]
	return ok
}

func (r *register) register(id int, item worker) {
	r.workers[id] = item
}

func (r *register) unregister(id int) {
	delete(r.workers, id)
}

func (r *register) size() int {
	return len(r.workers)
}
