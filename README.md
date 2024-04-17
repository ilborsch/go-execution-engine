# go-execution-engine

## Steps
1. Receive and parse a go script.
2. Implement rate limit 
3. Validate the script (no unsafe libraries used)
4. Secure the code execution by executing the script in docket container
5. Make sure to limit memory resources and set execution timeouts
6. Collect output from Stdout 
7. Pack all results, panics or compilation errors and send them back to user

## Code design

### For step 1:
1. Request receiver (either from kafka or REST haven't decided yet)

### For step 2:
1. It will be implemented in the web-app microservice

### For step 3:
1. Parse the script 
2. Script validator that will validate the script
3. If the script contains compilation errors return the error to user
4. If code contains imports of forbidden packages also return an error

### For step 4:
1. Create a docker client interface
2. Create an interface that will create a new code environment for the script using docker client

### For step 5:
1. Idk yet need research

### For step 6:
1. Intercept the entire data stream from the stdout while executing the script

### For step 7:
1. Send the results to user via kafka or REST haven't decided yet