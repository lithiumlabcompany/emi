# Variable for filename for store running procees id
PID_FILE = /tmp/emi.pid

# Start task performs "go run main.go" command and writes it's process id to PID_FILE.
test:
	go test -v ./... & echo $$! > $(PID_FILE)
# You can also use go build command for start task
# start:
#   go build -o /bin/emi . && \
#   /bin/emi & echo $$! > $(PID_FILE)

# Stop task will kill process by ID stored in PID_FILE (and all child processes by pstree).
stop:
	-kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

# Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
before:
	@echo "STOPED emi" && printf '%*s\n' "40" '' | tr ' ' -

# Restart task will execute stop, before and start tasks in strict order and prints message.
restart: stop before test
	@echo "STARTED emi" && printf '%*s\n' "40" '' | tr ' ' -

# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
tdd: test
	fswatch -or --event=Updated ./ | \
	xargs -n1 -I {} make restart

# .PHONY is used for reserving tasks words
.PHONY: test before stop restart tdd
